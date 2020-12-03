// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const perfPolicyName = "PerformancePolicyTest"

type PerformancePolicyWorkflowSuite struct {
	suite.Suite
	groupService      *service.NsGroupService
	perfpolicyService *service.PerformancePolicyService
	volumeService     sdkprovider.VolumeService
}

func (suite *PerformancePolicyWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.perfpolicyService = groupService.GetPerformancePolicyService()
	suite.volumeService = groupService.GetVolumeService()
	suite.createPerfPolicy(perfPolicyName)
}

func (suite *PerformancePolicyWorkflowSuite) TearDownSuite() {
	err := deleteDefaultVolume(suite.volumeService)
	assert.Nilf(suite.T(), err, "Unable to delete default volume, err: %v", err)
	suite.deletePerfPolicy(perfPolicyName)
	suite.groupService.LogoutService()
}

func (suite *PerformancePolicyWorkflowSuite) createPerfPolicy(perfPolicyName string) {
	perfPolicy := &nimbleos.PerformancePolicy{
		Name: param.NewString(perfPolicyName),
	}
	_, err := suite.perfpolicyService.CreatePerformancePolicy(perfPolicy)
	assert.Nilf(suite.T(), err, "Unable to create Performance Policy, err: %v", err)
}

func (suite *PerformancePolicyWorkflowSuite) deletePerfPolicy(perfPolicyName string) {
	perfPolicy, _ := suite.perfpolicyService.GetPerformancePolicyByName(perfPolicyName)
	if perfPolicy != nil {
		err := suite.perfpolicyService.DeletePerformancePolicy(*perfPolicy.ID)
		assert.Nilf(suite.T(), err, "Unable to delete performance policy, err: %v", err)
	}
}

func (suite *PerformancePolicyWorkflowSuite) TestAssociateVolumePerfPolicy() {
	perfPolicy, err := suite.perfpolicyService.GetPerformancePolicyByName(perfPolicyName)
	assert.Nilf(suite.T(), err, "Unable to find performance policy, err: %v", err)
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name:          param.NewString(defaultVolumeName),
		Size:          &sizeField,
		DedupeEnabled: param.NewBool(false),
	}
	vol, err := suite.volumeService.CreateVolume(newVolume)
	assert.Nilf(suite.T(), err, "Unable to create volume, err: %v", err)
	updateVol := &nimbleos.Volume{
		PerfpolicyId: perfPolicy.ID,
	}
	_, err = suite.volumeService.UpdateVolume(*vol.ID, updateVol)
	assert.Nilf(suite.T(), err, "Unable to update volume, err: %v", err)
	// Verify performance policy cannot be deleted when volume is associated
	err = suite.perfpolicyService.DeletePerformancePolicy(*perfPolicy.ID)
	assert.NotNil(suite.T(), err, "Deleting performance policy should have failed when volume is associated")
}

func (suite *PerformancePolicyWorkflowSuite) TestPreDefinedPolicy() {
	perfPolicy, err := suite.perfpolicyService.GetPerformancePolicyByName("default")
	assert.Nilf(suite.T(), err, "Unable to find performance policy, err: %v", err)
	err = suite.perfpolicyService.DeletePerformancePolicy(*perfPolicy.ID)
	assert.NotNil(suite.T(), err, "Deleting pre defined performance policy should have failed")
}

func TestPerfPolicyWorkflowSuite(t *testing.T) {
	suite.Run(t, new(PerformancePolicyWorkflowSuite))
}
