// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.keyManagerService = groupService.GetKeyManagerService()
	suite.masterKeyService = groupService.GetMasterKeyService()
	suite.deleteMasterKey()
	suite.createMasterKey()
}

func (suite *EncryptionWorkflowSuite) TearDownSuite() {
	suite.deleteMasterKey()
	suite.groupService.LogoutService()
}

func (suite *EncryptionWorkflowSuite) createMasterKey() {
	var passphrase *string = param.NewString(passphrase)
	key := &nimbleos.MasterKey{
		Passphrase: passphrase,
	}
	_, err := suite.masterKeyService.CreateMasterKey(key)
	assert.Nilf(suite.T(), err, "Master key creation failed, err: %v", err)
}

func (suite *EncryptionWorkflowSuite) deleteMasterKey() {
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	if len(key) != 0 {
		err := suite.masterKeyService.DeleteMasterKey(*key[0].ID)
		assert.Nilf(suite.T(), err, "Unable to delete key, err: %v", err)
	}
}

func (suite *EncryptionWorkflowSuite) TestCreateMasterKeyDuplicate() {
	key := &nimbleos.MasterKey{
		Passphrase: param.NewString("passphrase1"),
	}
	_, err := suite.masterKeyService.CreateMasterKey(key)
	assert.NotNil(suite.T(), err, "Master key creation when one exists should have failed")
}

func (suite *EncryptionWorkflowSuite) TestMasterKeyChangePassphrase() {
	var newPassphrase *string = param.NewString(newPassphrase)
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	updateKey := &nimbleos.MasterKey{
		Passphrase:    param.NewString(passphrase),
		NewPassphrase: newPassphrase,
	}
	_, err := suite.masterKeyService.UpdateMasterKey(*key[0].ID, updateKey)
	assert.Nilf(suite.T(), err, "Unable to update master key, err: %v", err)
}

func (suite *EncryptionWorkflowSuite) TestMasterKeyUpdateInactive() {
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	updateKey := &nimbleos.MasterKey{
		Passphrase: param.NewString(newPassphrase),
		Active:     param.NewBool(false),
	}
	_, err := suite.masterKeyService.UpdateMasterKey(*key[0].ID, updateKey)
	assert.Nilf(suite.T(), err, "Unable to set master key to inactive, err: %v", err)
}

func (suite *EncryptionWorkflowSuite) TestPurgeKey() {
	filter := &param.GetParams{}
	key, _ := suite.masterKeyService.GetMasterKeys(filter)
	var age int = 0
	err := suite.masterKeyService.PurgeInactiveMasterKey(*key[0].ID, &age)
	assert.Nilf(suite.T(), err, "Unable to purge inactive key, err: %v", err)
}

func TestEncryptionWorkflowSuite(t *testing.T) {
	suite.Run(t, new(EncryptionWorkflowSuite))
}
