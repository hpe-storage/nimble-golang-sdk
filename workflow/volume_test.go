package workflow

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const volumeName = "VolumeTest"

type VolumeWorkflowSuite struct {
	suite.Suite
	groupService        *service.NsGroupService
	volumeService       *service.VolumeService
	snapshotService     *service.SnapshotService
	volcollService      *service.VolumeCollectionService
	initiatorGrpService *service.InitiatorGroupService
}

func (suite *VolumeWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.snapshotService = groupService.GetSnapshotService()
	suite.volcollService = groupService.GetVolumeCollectionService()
	suite.initiatorGrpService = groupService.GetInitiatorGroupService()
	createDefaultInitiatorGrp(suite.initiatorGrpService)
	createDefaultVolColl(suite.volcollService)
}

func (suite *VolumeWorkflowSuite) TearDownSuite() {
	suite.deleteVolume(volumeName)
	err := deleteDefaultInitiatorGroup(suite.initiatorGrpService)
	assert.Nilf(suite.T(), err, "Unable to delete default initiator group, err: %v", err)
	err = deleteDefaultVolColl(suite.volcollService)
	assert.Nilf(suite.T(), err, "Unable to delete default volume collection, err: %v", err)
}

func (suite *VolumeWorkflowSuite) deleteVolume(volumeName string) {
	volObj, _ := suite.volumeService.GetVolumeByName(volumeName)
	if volObj != nil {
		err := suite.volumeService.DeleteVolume(*volObj.ID)
		assert.Nilf(suite.T(), err, "Unable to delete volume, err: %v", volumeName)
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeCreateWithMissParams() {
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	assert.NotNil(suite.T(), err, "Volume creation should have failed")
}

func (suite *VolumeWorkflowSuite) TestVolumeCreate() {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
		Size: &sizeField,
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	assert.Nilf(suite.T(), err, "Volume creation failed with error: %v", err)
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	assert.Equal(suite.T(), sizeField, *vol.Size, "Size was not set correctly")
}

func (suite *VolumeWorkflowSuite) TestVolumeCreateDuplicate() {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
		Size: &sizeField,
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	assert.NotNil(suite.T(), err, "Creating duplicate volume should have failed")
}

func (suite *VolumeWorkflowSuite) TestCreateVolumeByCloning() {
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if vol != nil {
		newSnapshot := &nimbleos.Snapshot{
			Name:  param.NewString("TestSnap"),
			VolId: vol.ID,
		}
		snap, _ := suite.snapshotService.CreateSnapshot(newSnapshot)
		newVolume := &nimbleos.Volume{
			Name:       param.NewString("TestClone"),
			Clone:      param.NewBool(true),
			BaseSnapId: snap.ID,
		}
		_, err := suite.volumeService.CreateVolume(newVolume)
		assert.Nilf(suite.T(), err, "Unable to clone volume, err: %v", err)
		cloneVol, _ := suite.volumeService.GetVolumeByName("TestClone")
		assert.NotNil(suite.T(), cloneVol, "TestClone should have been present")
		suite.volumeService.DeleteVolume(*cloneVol.ID)
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateSize() {
	var sizeIncField int64 = 10240
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1
	newVolume := &nimbleos.Volume{
		Size:      &sizeIncField,
		LimitIops: &limitIopsField,
		LimitMbps: &limitMbpsField,
	}
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if vol != nil {
		_, err := suite.volumeService.UpdateVolume(*vol.ID, newVolume)
		assert.Nilf(suite.T(), err, "Unable to update volume size, err: %v", err)
	}
	updateVol, _ := suite.volumeService.GetVolumeByName(volumeName)
	assert.Equal(suite.T(), *updateVol.LimitMbps, limitMbpsField, "Volume modify failed to update LimitMbps")
	assert.Equal(suite.T(), *updateVol.Size, sizeIncField, "Volume size increase failed.")
	var sizeDecField int64 = 4000
	newVolume = &nimbleos.Volume{
		Size: &sizeDecField,
	}
	if vol != nil {
		_, err := suite.volumeService.UpdateVolume(*vol.ID, newVolume)
		assert.NotNil(suite.T(), err, "Decreasing volume size without force should have failed")
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateOfflineVolume() {
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if vol != nil {
		_, err := suite.volumeService.OfflineVolume(*vol.ID, false)
		assert.Nilf(suite.T(), err, "Unable to set volume offline, err: %v", err)
		vol, _ = suite.volumeService.GetVolumeByName(volumeName)
		assert.Equal(suite.T(), *vol.VolState, *nimbleos.NsVolStatusOffline, "Volume did not go offline")
	} else {
		suite.T().Errorf("Could not find the volume")
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateOnlineVolume() {
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if vol != nil {
		_, err := suite.volumeService.OnlineVolume(*vol.ID, false)
		assert.Nilf(suite.T(), err, "Unable to set volume online, error: %v", err)
		vol, _ = suite.volumeService.GetVolumeByName(volumeName)
		assert.Equal(suite.T(), *vol.VolState, *nimbleos.NsVolStatusOnline, "Volume did not come online")
	} else {
		suite.T().Error("Could not find volume")
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeVolCollAssociate() {
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	volcoll, _ := suite.volcollService.GetVolumeCollectionByName(defaultVolCollName)
	err := suite.volumeService.AssociateVolume(*vol.ID, *volcoll.ID)
	assert.Nilf(suite.T(), err, "Associating volume to a volume collection failed: %v", err)
	err = suite.volumeService.DisassociateVolume(*vol.ID)
	assert.Nilf(suite.T(), err, "Disassociating volume to a volume collection failed: %v", err)
}

func TestVolumeWorkflowSuite(t *testing.T) {
	suite.Run(t, new(VolumeWorkflowSuite))
}
