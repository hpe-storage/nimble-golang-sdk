/*
Create/Modify/Delete Snapshot tests:
-----------------
 1. Create a volume
 2. Creat an online snapshot snapshot1 for the volume
      Validate if Snapshot is created for the volume
 3. Creat an offline snapshot snapshot2 for the volume
 4. Delete volume with online snapshot
			It should fail
 5. Update following params for the same snapshot
    		online, description, name
         writable value cannot be modified
 6. Create clone volume with snapshot1
 7. Delete snapshot1 which has cloned volume
         It should fail
 8. Delete the cloned volume
 10. Delete offline snapshot from the volume
 11. Delete the volume
*/

package test

import (
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SnapshotWorkflowSuite struct {
	suite.Suite
	groupService    *service.NsGroupService
	snapshotService *service.SnapshotService
	volumeService   *service.VolumeService
}

func (suite *SnapshotWorkflowSuite) SetupSuite() {
	groupService, err := config()
	if err != nil {
		suite.T().Errorf("Unable to connect to group: %v", err.Error())
	}
	suite.groupService = groupService
	suite.snapshotService = groupService.GetSnapshotService()
	suite.volumeService = groupService.GetVolumeService()
}

// Create a volume/clone volume
func (suite *SnapshotWorkflowSuite) createVolume(volumeName string, baseSnapshotID string) string {

	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name:   param.NewString(volumeName),
		Size:   &sizeField,
		Online: param.NewBool(false),
	}
	if baseSnapshotID != "" {
		newVolume.BaseSnapId = param.NewString(baseSnapshotID)
		newVolume.Clone = param.NewBool(true)
	}
	volResp, err := suite.volumeService.CreateVolume(newVolume)
	assert.Nilf(suite.T(), err, "Unable to create volume %v", volumeName)
	return *volResp.ID
}

func (suite *SnapshotWorkflowSuite) deleteVolume(volumeID string) error {
	delVol := suite.volumeService.DeleteVolume(volumeID)
	return delVol
}

// Creat an offline/online snapshot for the volume. Delete volume
func (suite *SnapshotWorkflowSuite) TestCreateUpdateSnapshot() {
	// Create, Get, Delete Offline Snapshot
	volumeName := "testvolumesnap1"
	volumeID := suite.createVolume(volumeName, "")
	offlineSnapshotName := "snapoffline"
	newOfflineSnapshot := &nimbleos.Snapshot{
		Name:  param.NewString(offlineSnapshotName),
		VolId: param.NewString(volumeID),
	}
	_, err := suite.snapshotService.CreateSnapshot(newOfflineSnapshot)
	assert.Nilf(suite.T(), err, "Failed to create snapshot")

	getSnapResp, err := suite.snapshotService.GetSnapshotByName(volumeName)
	assert.Nilf(suite.T(), err, "Failed to get snapshot")

	err = suite.snapshotService.DeleteSnapshot(*getSnapResp[0].ID)
	assert.Nilf(suite.T(), err, "Failed to delete snapshot")

	// Create Online Snapshot
	onlineSnapshotName := "snaponline"
	newOnlineSnapshot := &nimbleos.Snapshot{
		Name:     param.NewString(onlineSnapshotName),
		VolId:    param.NewString(volumeID),
		Online:   param.NewBool(true),
		Writable: param.NewBool(true),
	}
	snapResp, err := suite.snapshotService.CreateSnapshot(newOnlineSnapshot)
	assert.Nilf(suite.T(), err, "Failed to create snapshot")
	onlineSnapshotID := *snapResp.ID

	// Delete Volume with online snapshot
	err = suite.deleteVolume(volumeID)
	assert.NotNilf(suite.T(), err, "Volume deletion expected to throw error, as there are online snapshots")

	// Offline the Snapshot & update description
	updateSnapshot := &nimbleos.Snapshot{
		Name:        param.NewString("snapshot2"),
		Online:      param.NewBool(false),
		Description: param.NewString("Offline the snapshot"),
	}
	_, err = suite.snapshotService.UpdateSnapshot(onlineSnapshotID, updateSnapshot)
	assert.Nilf(suite.T(), err, "Failed to update snapshot")

	err = suite.deleteVolume(volumeID)
	assert.Nilf(suite.T(), err, "Failed to delete volume")
}

// Snapshot Clone Volume test
func (suite *SnapshotWorkflowSuite) TestCloneVolumeSnapshot() {
	volumeName := "testvolumesnap2"
	parentVolumeID := suite.createVolume(volumeName, "")
	baseSnapshotName := "snapoffline"
	newOfflineSnapshot := &nimbleos.Snapshot{
		Name:   param.NewString(baseSnapshotName),
		VolId:  param.NewString(parentVolumeID),
		Online: param.NewBool(true),
	}
	snapResp, err := suite.snapshotService.CreateSnapshot(newOfflineSnapshot)
	assert.Nilf(suite.T(), err, "Failed to create snapshot")
	baseSnapshotID := *snapResp.ID

	// Create clone volume
	cloneVolumeName := "clonevolumesnap2"
	cloneVolumeID := suite.createVolume(cloneVolumeName, baseSnapshotID)

	// Delete Snapshot with Clone Volume
	err = suite.snapshotService.DeleteSnapshot(baseSnapshotID)
	assert.NotNilf(suite.T(), err, "Snapshot deletion should have thrown error")

	// Delete clone volume
	err = suite.deleteVolume(cloneVolumeID)
	assert.Nilf(suite.T(), err, "Failed to delete volume")

	// Delete Snapshot after deleting Clone Volume
	err = suite.snapshotService.DeleteSnapshot(baseSnapshotID)
	assert.Nilf(suite.T(), err, "Failed to delete snapshot")

	// Delete parent volume
	err = suite.deleteVolume(parentVolumeID)
	assert.Nilf(suite.T(), err, "Failed to delete volume")

}

func TestSnapshotWorkflowSuite(t *testing.T) {
	suite.Run(t, new(SnapshotWorkflowSuite))
}
