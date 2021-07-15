// Copyright 2020 Hewlett Packard Enterprise Development LP

// Disk tests
// -----------
//   1. Get Disk info from array
//   2. Get the shelf location & slot number of the last in-use slot ("state": "in use")
//   3. Remove the disk slot from shelf
//           Validate the state
//   4. Add the removed disk slot to the shelf again
//           Validate the state

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DiskWorkflowSuite struct {
	suite.Suite
	groupService *service.NsGroupService
	diskService  *service.DiskService
}

func (suite *DiskWorkflowSuite) SetupSuite() {
	groupService, err := config()
	require.Nilf(suite.T(), err, "Could not connect to array: %v", arrayIP)
	suite.groupService = groupService
	suite.diskService = groupService.GetDiskService()
}

func (suite *DiskWorkflowSuite) TearDownSuite() {
	var diskTestResult string
	if len(testStarted) == len(testCompleted) {
		diskTestResult = "PASS"
	} else {
		diskTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(diskTestResult, "C545227", "Disk Workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.groupService.LogoutService()
}

func (suite *DiskWorkflowSuite) TestAddRemoveDisk() {
	testStarted = append(testStarted, true)
	filter := &param.GetParams{}
	diskResp, err := suite.diskService.GetDisks(filter)
	require.Nil(suite.T(), err, "Failed to get disk details")
	fmt.Print(*diskResp[0].State)

	// Get a disk which is in use
	var diskID string
	for i := 0; i < len(diskResp); i++ {
		if *diskResp[i].State == "in use" {
			diskID = *diskResp[i].ID
			break
		}
	}
	require.NotEmpty(suite.T(), diskID, "No Disk found which are 'in-use'")

	// Remove Disk
	removeDisk := &nimbleos.Disk{
		DiskOp: nimbleos.NsDiskOpRemove,
	}
	_, err = suite.diskService.UpdateDisk(diskID, removeDisk)
	require.Nil(suite.T(), err, "Failed to Remove Disk")
	getDiskResp, err := suite.diskService.GetDiskById(diskID)
	require.Nil(suite.T(), err, "Failed to Get Disk details")
	require.Equal(suite.T(), nimbleos.NsDiskState("removed"), *getDiskResp.State, "Failed to Remove Disk")

	// Add Disk
	addDisk := &nimbleos.Disk{
		DiskOp: nimbleos.NsDiskOpAdd,
	}
	_, err = suite.diskService.UpdateDisk(diskID, addDisk)
	require.Nil(suite.T(), err, "Failed to Add Disk")

	// Wait till the state changes from 'valid' to 'in use'
	time.Sleep(2 * time.Second)
	getDiskResp, err = suite.diskService.GetDiskById(diskID)
	require.Nil(suite.T(), err, "Failed to Get Disk details")
	require.Equal(suite.T(), nimbleos.NsDiskState("in use"), *getDiskResp.State, "Failed to Add Disk")
	testCompleted = append(testCompleted, true)

}

func TestDiskWorkflowSuite(t *testing.T) {
	suite.Run(t, new(DiskWorkflowSuite))
}
