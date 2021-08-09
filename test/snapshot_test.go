// Copyright 2020 Hewlett Packard Enterprise Development LP

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
 9. Delete offline snapshot from the volume
 10. Delete the volume
*/

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

type SnapshotWorkflowSuite struct {
	suite.Suite
	groupService    *service.NsGroupService
	snapshotService *service.SnapshotService
	volumeService   sdkprovider.VolumeService
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

func (suite *SnapshotWorkflowSuite) TearDownSuite() {
	var snapshotTestResult string
	if len(testStarted) == len(testCompleted) {
		snapshotTestResult = "PASS"
	} else {
		snapshotTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(snapshotTestResult, "C560653", "Snapshot workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.groupService.LogoutService()
}

// Create a volume/clone volume
func (suite *SnapshotWorkflowSuite) createVolume(volumeName string, baseSnapshotID string) string {

	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name:   &volumeName,
		Size:   &sizeField,
		Online: param.NewBool(false),
	}
	if baseSnapshotID != "" {
		newVolume.BaseSnapId = &baseSnapshotID
		newVolume.Clone = param.NewBool(true)
	}
	volResp, err := suite.volumeService.CreateVolume(newVolume)
	require.Nilf(suite.T(), err, "Unable to create volume: %v", volumeName)
	return *volResp.ID
}

func (suite *SnapshotWorkflowSuite) deleteVolume(volumeID string) error {
	delVol := suite.volumeService.DeleteVolume(volumeID)
	return delVol
}

// Creat an offline/online snapshot for the volume. Delete volume
func (suite *SnapshotWorkflowSuite) TestCreateUpdateSnapshot() {
	testStarted = append(testStarted, true)
	// Create, Get, Delete Offline Snapshot
	volumeName := "testvolumesnap1"
	volumeID := suite.createVolume(volumeName, "")
	offlineSnapshotName := "snapoffline"
	newOfflineSnapshot := &nimbleos.Snapshot{
		Name:  &offlineSnapshotName,
		VolId: &volumeID,
	}
	_, err := suite.snapshotService.CreateSnapshot(newOfflineSnapshot)
	require.Nilf(suite.T(), err, "Failed to create snapshot: %v", offlineSnapshotName)

	getSnapResp, err := suite.snapshotService.GetSnapshotByName(volumeName)
	require.Nilf(suite.T(), err, "Failed to get snapshot(s) by volume name: %v", volumeName)

	err = suite.snapshotService.DeleteSnapshot(*getSnapResp[0].ID)
	require.Nilf(suite.T(), err, "Failed to delete snapshot: %v", offlineSnapshotName)

	// Create Online Snapshot
	onlineSnapshotName := "snaponline"
	newOnlineSnapshot := &nimbleos.Snapshot{
		Name:     &onlineSnapshotName,
		VolId:    &volumeID,
		Online:   param.NewBool(true),
		Writable: param.NewBool(true),
	}
	snapResp, err := suite.snapshotService.CreateSnapshot(newOnlineSnapshot)
	require.Nilf(suite.T(), err, "Failed to create snapshot: %v", newOnlineSnapshot)
	onlineSnapshotID := *snapResp.ID

	// Delete Volume with online snapshot
	err = suite.deleteVolume(volumeID)
	require.NotNil(suite.T(), err, "Volume deletion expected to throw error, as there are online snapshots")

	// Offline the Snapshot & update name, description
	newOnlineSnapshotName := "updatesnaponline"
	newSnapshotDescription := "Offline the snapshot"
	updateSnapshot := &nimbleos.Snapshot{
		Name:        &newOnlineSnapshotName,
		Online:      param.NewBool(false),
		Description: &newSnapshotDescription,
	}
	_, err = suite.snapshotService.UpdateSnapshot(onlineSnapshotID, updateSnapshot)
	require.Nil(suite.T(), err, "Failed to update snapshot")

	getSnapIDResp, err := suite.snapshotService.GetSnapshotById(onlineSnapshotID)
	require.Nil(suite.T(), err, "Failed to get snapshot by ID")
	require.Equal(suite.T(), newSnapshotDescription, *getSnapIDResp.Description, "In-correct Description for Snapshot")
	require.Equal(suite.T(), newOnlineSnapshotName, *getSnapIDResp.Name, "Name not updated")
	require.Equal(suite.T(), false, *getSnapIDResp.Online, "Name not updated")

	err = suite.deleteVolume(volumeID)
	require.Nilf(suite.T(), err, "Failed to delete volume: %v", volumeName)
	testCompleted = append(testCompleted, true)
}

// Snapshot Clone Volume test
func (suite *SnapshotWorkflowSuite) TestCloneVolumeSnapshot() {
	testStarted = append(testStarted, true)
	volumeName := "testvolumesnap2"
	parentVolumeID := suite.createVolume(volumeName, "")
	baseSnapshotName := "snapoffline"
	newOfflineSnapshot := &nimbleos.Snapshot{
		Name:   &baseSnapshotName,
		VolId:  &parentVolumeID,
		Online: param.NewBool(true),
	}
	snapResp, err := suite.snapshotService.CreateSnapshot(newOfflineSnapshot)
	require.Nil(suite.T(), err, "Failed to create snapshot")
	baseSnapshotID := *snapResp.ID

	// Create clone volume
	cloneVolumeName := "clonevolumesnap2"
	cloneVolumeID := suite.createVolume(cloneVolumeName, baseSnapshotID)

	// Delete Snapshot with Clone Volume
	err = suite.snapshotService.DeleteSnapshot(baseSnapshotID)
	require.NotNil(suite.T(), err, "Snapshot deletion should have thrown error")

	// Delete clone volume
	err = suite.deleteVolume(cloneVolumeID)
	require.Nil(suite.T(), err, "Failed to delete volume")

	// Delete Snapshot after deleting Clone Volume
	err = suite.snapshotService.DeleteSnapshot(baseSnapshotID)
	require.Nil(suite.T(), err, "Failed to delete snapshot")

	// Delete parent volume
	err = suite.deleteVolume(parentVolumeID)
	require.Nil(suite.T(), err, "Failed to delete volume")
	testCompleted = append(testCompleted, true)

}

func TestSnapshotWorkflowSuite(t *testing.T) {
	suite.Run(t, new(SnapshotWorkflowSuite))
}
