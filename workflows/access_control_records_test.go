// Copyright 2020 Hewlett Packard Enterprise Development LP
package workflows

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
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

func (suite *ACRWorkflowSuite) SetupTest() {
	groupService := config()
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.igService = groupService.GetInitiatorGroupService()
	suite.acrService = groupService.GetAccessControlRecordService()
	createDefaultVolume(groupService)
	createDefaultInitiatorGrp(groupService)
}

func (suite *ACRWorkflowSuite) TearDownSuite() {
	err := suite.acrService.DeleteAccessControlRecord(*suite.acr.ID)
	if err != nil {
		suite.T().Errorf("Access Control Record deletion failed: %v", err.Error())
	}
	deleteVolume(suite.groupService, defaultVolumeName)
	deleteInitiatorGroup(suite.groupService, defaultInitiatorGrp)
}

func (suite *ACRWorkflowSuite) TestCreateACR() {
	vol, _ := suite.volumeService.GetVolumeByName(defaultVolumeName)
	ig, _ := suite.igService.GetInitiatorGroupByName(defaultInitiatorGrp)
	newACR := &nimbleos.AccessControlRecord{
		InitiatorGroupId: ig.ID,
		VolId:            vol.ID,
	}
	acr, err := suite.acrService.CreateAccessControlRecord(newACR)
	if err != nil {
		suite.T().Errorf("Access Control Record creation failed: %v", err.Error())
	}
	suite.acr = acr
}

func (suite *ACRWorkflowSuite) TestModifyACR() {
	updateACR := &nimbleos.AccessControlRecord{
		ApplyTo: nimbleos.NsAccessApplyToVolume,
	}
	_, err := suite.acrService.UpdateAccessControlRecord(*suite.acr.ID, updateACR)
	if err == nil {
		suite.T().Error("ACR update should have failed")
	}
}

func (suite *ACRWorkflowSuite) TestDeleteIGAssociated() {
	ig, _ := suite.igService.GetInitiatorGroupByName(defaultInitiatorGrp)
	err := suite.igService.DeleteInitiatorGroup(*ig.ID)
	if err == nil {
		suite.T().Errorf("Deletion of Initiator Group associated with volume should have failed")
	}
}

func TestACRWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ACRWorkflowSuite))
}
