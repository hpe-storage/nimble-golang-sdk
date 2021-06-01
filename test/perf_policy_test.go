// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
	require.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.perfpolicyService = groupService.GetPerformancePolicyService()
	suite.volumeService = groupService.GetVolumeService()
	suite.createPerfPolicy(perfPolicyName)
}

func (suite *PerformancePolicyWorkflowSuite) TearDownSuite() {
	var perfPolicyTestResult string
	if len(testStarted) == len(testCompleted) {
		perfPolicyTestResult = "PASS"
	} else {
		perfPolicyTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(perfPolicyTestResult, "C560652", "Performance Policy workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	err := deleteDefaultVolume(suite.volumeService)
	require.Nilf(suite.T(), err, "Unable to delete default volume, err: %v", err)
	suite.deletePerfPolicy(perfPolicyName)
	suite.groupService.LogoutService()
}

func (suite *PerformancePolicyWorkflowSuite) createPerfPolicy(perfPolicyName string) {
	perfPolicy := &nimbleos.PerformancePolicy{
		Name: param.NewString(perfPolicyName),
	}
	_, err := suite.perfpolicyService.CreatePerformancePolicy(perfPolicy)
	require.Nilf(suite.T(), err, "Unable to create Performance Policy, err: %v", err)
}

func (suite *PerformancePolicyWorkflowSuite) deletePerfPolicy(perfPolicyName string) {
	perfPolicy, _ := suite.perfpolicyService.GetPerformancePolicyByName(perfPolicyName)
	if perfPolicy != nil {
		err := suite.perfpolicyService.DeletePerformancePolicy(*perfPolicy.ID)
		require.Nilf(suite.T(), err, "Unable to delete performance policy, err: %v", err)
	}
}

func (suite *PerformancePolicyWorkflowSuite) TestAssociateVolumePerfPolicy() {
	testStarted = append(testStarted, true)
	perfPolicy, err := suite.perfpolicyService.GetPerformancePolicyByName(perfPolicyName)
	require.Nilf(suite.T(), err, "Unable to find performance policy, err: %v", err)
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name:          param.NewString(defaultVolumeName),
		Size:          &sizeField,
		DedupeEnabled: param.NewBool(false),
	}
	vol, err := suite.volumeService.CreateVolume(newVolume)
	require.Nilf(suite.T(), err, "Unable to create volume, err: %v", err)
	updateVol := &nimbleos.Volume{
		PerfpolicyId: perfPolicy.ID,
	}
	_, err = suite.volumeService.UpdateVolume(*vol.ID, updateVol)
	require.Nilf(suite.T(), err, "Unable to update volume, err: %v", err)
	// Verify performance policy cannot be deleted when volume is associated
	err = suite.perfpolicyService.DeletePerformancePolicy(*perfPolicy.ID)
	require.NotNil(suite.T(), err, "Deleting performance policy should have failed when volume is associated")
	testCompleted = append(testCompleted, true)
}

func (suite *PerformancePolicyWorkflowSuite) TestPreDefinedPolicy() {
	testStarted = append(testStarted, true)
	perfPolicy, err := suite.perfpolicyService.GetPerformancePolicyByName("default")
	require.Nilf(suite.T(), err, "Unable to find performance policy, err: %v", err)
	err = suite.perfpolicyService.DeletePerformancePolicy(*perfPolicy.ID)
	require.NotNil(suite.T(), err, "Deleting pre defined performance policy should have failed")
	testCompleted = append(testCompleted, true)
}

func TestPerfPolicyWorkflowSuite(t *testing.T) {
	suite.Run(t, new(PerformancePolicyWorkflowSuite))
}
