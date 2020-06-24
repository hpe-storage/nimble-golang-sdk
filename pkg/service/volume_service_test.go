// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

type VolumeServiceTestSuite struct {
	suite.Suite
	groupService             *GroupService
	volumeService            *VolumeService
	performancePolicyService *PerformancePolicyService
}

func (suite *VolumeServiceTestSuite) SetupTest() {
	groupService, err := NewGroupService("10.18.174.8", "admin", "admin")
	if err != nil {
		suite.T().Errorf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
		return
	}
	// groupService.client.EnableDebug()

	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.performancePolicyService = groupService.GetPerformancePolicyService()

	suite.createVolume("GetVolume")
	suite.createVolume("DeleteVolume")
	suite.createVolume("DestroyVolume")
}

func (suite *VolumeServiceTestSuite) TearDownTest() {
	suite.deleteVolume("DestroyVolume")
	suite.deleteVolume("DeleteVolume")
	suite.deleteVolume("GetVolume")
}

func (suite *VolumeServiceTestSuite) getDefaultVolumeOptions() *model.Volume {

	perfPolicy, _ := suite.performancePolicyService.GetPerformancePolicyByName("default")
	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	// Enum check
	var agentType model.NsAgentType
	agentType = model.NSAGENTTYPE_NONE

	newVolume := &model.Volume{
		Size:              &sizeField,
		Description:       &descriptionField,
		PerfpolicyID:      perfPolicy.ID,
		ThinlyProvisioned: util.NewBool(true),
		Online:            util.NewBool(true),
		LimitIops:         &limitIopsField,
		LimitMbps:         &limitMbpsField,
		MultiInitiator:    util.NewBool(true),
		AgentType:		   &agentType,

	}
	return newVolume
}

func (suite *VolumeServiceTestSuite) createVolume(volumeName string) *model.Volume {
	volume, _ := suite.volumeService.GetVolumeByName(volumeName)
	if volume == nil {
		volume = suite.getDefaultVolumeOptions()
		volume.Name = &volumeName
		volume, _ = suite.volumeService.CreateVolume(volume)
	}
	return volume
}

func (suite *VolumeServiceTestSuite) deleteVolume(volumeName string) {
	volume, _ := suite.volumeService.GetVolumeByName(volumeName)
	if volume != nil {
		suite.volumeService.DestroyVolume(*volume.ID)
		volume, _ = suite.volumeService.GetVolumeByName(volumeName)
	}

	if volume != nil {
		suite.T().Fatalf("deleteVolume: Failed to delete volume %s", volumeName)
	}
}

func (suite *VolumeServiceTestSuite) TestVolumeService() {
	assert.NotNil(suite.T(), suite.volumeService)

	volumes, err := suite.volumeService.GetVolumesWithFields([]string{*model.VolumeFields.Name})
	if err != nil {
		suite.T().Errorf("TestVolumeService(): Unable to get volumes, err: %v", err.Error())
		return
	}
	assert.NotEmpty(suite.T(), volumes)
}

func (suite *VolumeServiceTestSuite) TestGetNonExistentVolumeByID() {
	volume, err := suite.volumeService.GetVolumeById("06aaaaaaaaaaaaaaaa000000000000000000000000")
	if err != nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Unable to get non-existent volume, err: %v", err.Error())
		return
	}
	assert.Nil(suite.T(), volume)

	volume, err = suite.volumeService.GetVolumeById("badf0rmat")
	if err == nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Expected error")
		return
	}
}

// Runs all test via go test
func TestVolumeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(VolumeServiceTestSuite))
}
