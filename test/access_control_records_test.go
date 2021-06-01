// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ACRWorkflowSuite struct {
	suite.Suite
	acr           *nimbleos.AccessControlRecord
	groupService  *service.NsGroupService
	volumeService sdkprovider.VolumeService
	igService     *service.InitiatorGroupService
	acrService    *service.AccessControlRecordService
}

func (suite *ACRWorkflowSuite) SetupSuite() {
	groupService, err := config()
	require.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.igService = groupService.GetInitiatorGroupService()
	suite.acrService = groupService.GetAccessControlRecordService()
	_, err = createDefaultVolume(suite.volumeService)
	require.Nilf(suite.T(), err, "Unable to create default volume, err: %v", err)
	_, err = createDefaultInitiatorGrp(suite.igService)
	require.Nilf(suite.T(), err, "Unable to create default initiator group, err: %v", err)
	suite.createACR()
}

func (suite *ACRWorkflowSuite) TearDownSuite() {
	var accessControlTestResult string
	if len(testStarted) == len(testCompleted) {
		accessControlTestResult = "PASS"
	} else {
		accessControlTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(accessControlTestResult, "C545074", "Access Control Records workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	if suite.acr != nil {
		err := suite.acrService.DeleteAccessControlRecord(*suite.acr.ID)
		require.Nilf(suite.T(), err, "Unable to delete Access control record, err: %v", err)
	}
	err := deleteDefaultVolume(suite.volumeService)
	require.Nilf(suite.T(), err, "Unable to delete default volume, err: %v", err)
	err = deleteDefaultInitiatorGroup(suite.igService)
	require.Nilf(suite.T(), err, "Unable to delete default initiator group, err: %v", err)
	suite.groupService.LogoutService()
}

func (suite *ACRWorkflowSuite) createACR() {
	vol, _ := suite.volumeService.GetVolumeByName(defaultVolumeName)
	ig, _ := suite.igService.GetInitiatorGroupByName(defaultInitiatorGrpName)
	newACR := &nimbleos.AccessControlRecord{
		InitiatorGroupId: ig.ID,
		VolId:            vol.ID,
	}
	acr, err := suite.acrService.CreateAccessControlRecord(newACR)
	require.Nilf(suite.T(), err, "Unable to create Access Control Record, err: %v", err)
	suite.acr = acr
}

func (suite *ACRWorkflowSuite) TestModifyACR() {
	testStarted = append(testStarted, true)
	updateACR := &nimbleos.AccessControlRecord{
		ApplyTo: nimbleos.NsAccessApplyToVolume,
	}
	_, err := suite.acrService.UpdateAccessControlRecord(*suite.acr.ID, updateACR)
	require.NotNilf(suite.T(), err, "Access Control Record update should have failed")
	testCompleted = append(testCompleted, true)
}

func (suite *ACRWorkflowSuite) TestDeleteIGAssociated() {
	testStarted = append(testStarted, true)
	ig, _ := suite.igService.GetInitiatorGroupByName(defaultInitiatorGrpName)
	err := suite.igService.DeleteInitiatorGroup(*ig.ID)
	require.NotNilf(suite.T(), err, "Deletion of Initiator Group associated with volume should have failed")
	testCompleted = append(testCompleted, true)
}

func TestACRWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ACRWorkflowSuite))
}
