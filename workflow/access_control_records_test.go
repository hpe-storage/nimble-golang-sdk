// Copyright 2020 Hewlett Packard Enterprise Development LP
package workflow

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ACRWorkflowSuite struct {
	suite.Suite
	acr           *nimbleos.AccessControlRecord
	groupService  *service.NsGroupService
	volumeService *service.VolumeService
	igService     *service.InitiatorGroupService
	acrService    *service.AccessControlRecordService
}

func (suite *ACRWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.igService = groupService.GetInitiatorGroupService()
	suite.acrService = groupService.GetAccessControlRecordService()
	createDefaultVolume(suite.volumeService)
	createDefaultInitiatorGrp(suite.igService)
}

func (suite *ACRWorkflowSuite) TearDownSuite() {
	if suite.acr != nil {
		err := suite.acrService.DeleteAccessControlRecord(*suite.acr.ID)
		assert.Nilf(suite.T(), err, "Unable to delete Access control record, err: %v", err)
	}
	err := deleteDefaultVolume(suite.volumeService)
	assert.Nilf(suite.T(), err, "Unable to delete default volume, err: %v", err)
	err = deleteDefaultInitiatorGroup(suite.igService)
	assert.Nilf(suite.T(), err, "Unable to delete default initiator group, err: %v", err)
}

func (suite *ACRWorkflowSuite) TestCreateACR() {
	vol, _ := suite.volumeService.GetVolumeByName(defaultVolumeName)
	ig, _ := suite.igService.GetInitiatorGroupByName(defaultInitiatorGrpName)
	newACR := &nimbleos.AccessControlRecord{
		InitiatorGroupId: ig.ID,
		VolId:            vol.ID,
	}
	acr, err := suite.acrService.CreateAccessControlRecord(newACR)
	assert.Nilf(suite.T(), err, "Unable to create Access Control Record, err: %v", err)
	suite.acr = acr
}

func (suite *ACRWorkflowSuite) TestModifyACR() {
	updateACR := &nimbleos.AccessControlRecord{
		ApplyTo: nimbleos.NsAccessApplyToVolume,
	}
	_, err := suite.acrService.UpdateAccessControlRecord(*suite.acr.ID, updateACR)
	assert.NotNilf(suite.T(), err, "Access Control Record update should have failed")
}

func (suite *ACRWorkflowSuite) TestDeleteIGAssociated() {
	ig, _ := suite.igService.GetInitiatorGroupByName(defaultInitiatorGrpName)
	err := suite.igService.DeleteInitiatorGroup(*ig.ID)
	assert.NotNilf(suite.T(), err, "Deletion of Initiator Group associated with volume should have failed")
}

func TestACRWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ACRWorkflowSuite))
}
