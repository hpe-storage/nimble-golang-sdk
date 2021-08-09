// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const passphrase = "passphrase-91"
const newPassphrase = "passphrase-92"

type EncryptionWorkflowSuite struct {
	suite.Suite
	groupService      *service.NsGroupService
	keyManagerService *service.KeyManagerService
	masterKeyService  *service.MasterKeyService
}

func (suite *EncryptionWorkflowSuite) SetupSuite() {
	groupService, err := config()
	require.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.keyManagerService = groupService.GetKeyManagerService()
	suite.masterKeyService = groupService.GetMasterKeyService()
	suite.deleteMasterKey()
	suite.createMasterKey()
}

func (suite *EncryptionWorkflowSuite) TearDownSuite() {
	var encryptionTestResult string
	if len(testStarted) == len(testCompleted) {
		encryptionTestResult = "PASS"
	} else {
		encryptionTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(encryptionTestResult, "C545078", "Encryption workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.deleteMasterKey()
	suite.groupService.LogoutService()
}

func (suite *EncryptionWorkflowSuite) createMasterKey() {
	var passphrase *string = param.NewString(passphrase)
	key := &nimbleos.MasterKey{
		Passphrase: passphrase,
	}
	_, err := suite.masterKeyService.CreateMasterKey(key)
	require.Nilf(suite.T(), err, "Master key creation failed, err: %v", err)
}

func (suite *EncryptionWorkflowSuite) deleteMasterKey() {
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	if len(key) != 0 {
		err := suite.masterKeyService.DeleteMasterKey(*key[0].ID)
		require.Nilf(suite.T(), err, "Unable to delete key, err: %v", err)
	}
}

func (suite *EncryptionWorkflowSuite) TestCreateMasterKeyDuplicate() {
	testStarted = append(testStarted, true)
	key := &nimbleos.MasterKey{
		Passphrase: param.NewString("passphrase1"),
	}
	_, err := suite.masterKeyService.CreateMasterKey(key)
	require.NotNil(suite.T(), err, "Master key creation when one exists should have failed")
	testCompleted = append(testCompleted, true)
}

func (suite *EncryptionWorkflowSuite) TestMasterKeyChangePassphrase() {
	testStarted = append(testStarted, true)
	var newPassphrase *string = param.NewString(newPassphrase)
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	updateKey := &nimbleos.MasterKey{
		Passphrase:    param.NewString(passphrase),
		NewPassphrase: newPassphrase,
	}
	_, err := suite.masterKeyService.UpdateMasterKey(*key[0].ID, updateKey)
	require.Nilf(suite.T(), err, "Unable to update master key, err: %v", err)
	testCompleted = append(testCompleted, true)
}

func (suite *EncryptionWorkflowSuite) TestMasterKeyUpdateInactive() {
	testStarted = append(testStarted, true)
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	updateKey := &nimbleos.MasterKey{
		Passphrase: param.NewString(newPassphrase),
		Active:     param.NewBool(false),
	}
	_, err := suite.masterKeyService.UpdateMasterKey(*key[0].ID, updateKey)
	require.Nilf(suite.T(), err, "Unable to set master key to inactive, err: %v", err)
	testCompleted = append(testCompleted, true)
}

func (suite *EncryptionWorkflowSuite) TestPurgeKey() {
	testStarted = append(testStarted, true)
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	var age int = 0
	err := suite.masterKeyService.PurgeInactiveMasterKey(*key[0].ID, &age)
	require.Nilf(suite.T(), err, "Unable to purge inactive key, err: %v", err)
	testCompleted = append(testCompleted, true)
}

func TestEncryptionWorkflowSuite(t *testing.T) {
	suite.Run(t, new(EncryptionWorkflowSuite))
}
