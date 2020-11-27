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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DiskWorkflowSuite struct {
	suite.Suite
	groupService *service.NsGroupService
	diskService  *service.DiskService
}

func (suite *DiskWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Could not connect to array")
	suite.groupService = groupService
	suite.diskService = groupService.GetDiskService()
}

func (suite *DiskWorkflowSuite) TestAddRemoveDisk() {
	filter := &param.GetParams{}
	diskResp, err := suite.diskService.GetDisks(filter)
	assert.Nilf(suite.T(), err, "Failed to get disk details")
	fmt.Print(*diskResp[0].State)

	// Get a disk which is in use
	var diskID string
	for i := 0; i < len(diskResp); i++ {
		if *diskResp[i].State == "in use" {
			diskID = *diskResp[i].ID
			break
		}
	}
	assert.NotEmpty(suite.T(), diskID, "No Disk found which are 'in-use'")

	// Remove Disk
	removeDisk := &nimbleos.Disk{
		DiskOp: nimbleos.NsDiskOpRemove,
	}
	_, err = suite.diskService.UpdateDisk(diskID, removeDisk)
	assert.Nilf(suite.T(), err, "Failed to Remove Disk")
	getDiskResp, err := suite.diskService.GetDiskById(diskID)
	assert.Nilf(suite.T(), err, "Failed to Get Disk details")
	assert.Equal(suite.T(), nimbleos.NsDiskState("removed"), *getDiskResp.State, "Failed to Remove Disk")

	// Add Disk
	addDisk := &nimbleos.Disk{
		DiskOp: nimbleos.NsDiskOpAdd,
	}
	_, err = suite.diskService.UpdateDisk(diskID, addDisk)
	assert.Nilf(suite.T(), err, "Failed to Add Disk")
	getDiskResp, err = suite.diskService.GetDiskById(diskID)

	// Wait till the state changes from 'valid' to 'in use'
	time.Sleep(2 * time.Second)
	assert.Nilf(suite.T(), err, "Failed to Get Disk details")
	assert.Equal(suite.T(), nimbleos.NsDiskState("in use"), *getDiskResp.State, "Failed to Add Disk")

}

func TestDiskWorkflowSuite(t *testing.T) {
	suite.Run(t, new(DiskWorkflowSuite))
}
