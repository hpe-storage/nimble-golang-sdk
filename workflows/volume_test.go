package workflows

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/suite"
	"testing"
)

var volumeName = "VolumeTest"

type VolumeWorkflowSuite struct {
	suite.Suite
	groupService        *service.NsGroupService
	volumeService       *service.VolumeService
	snapshotService     *service.SnapshotService
	volcolService       *service.VolumeCollectionService
	initiatorGrpService *service.InitiatorGroupService
}

func (suite *VolumeWorkflowSuite) SetupTest() {
	groupService := config()
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.snapshotService = groupService.GetSnapshotService()
	suite.volcolService = groupService.GetVolumeCollectionService()
	suite.initiatorGrpService = groupService.GetInitiatorGroupService()
	createDefaultInitiatorGrp(groupService)
	createDefaultVolumeCollection(groupService)
}

func (suite *VolumeWorkflowSuite) TearDownSuite() {
	deleteVolume(suite.groupService, volumeName)
	deleteInitiatorGroup(suite.groupService, defaultInitiatorGrp)
	deleteVolCol(suite.groupService, defaultVolColName)
}

func (suite *VolumeWorkflowSuite) TestVolumeCreateWithMissParams() {
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	if err == nil {
		suite.T().Error("Volume creation should have failed")
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeCreate() {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
		Size: &sizeField,
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	if err != nil {
		suite.T().Errorf("Volume creation failed with error: %v", err.Error())
	}
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if *vol.Size == sizeField {
		fmt.Printf("Volume created successfully: %v", *vol.Name)
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeCreateDuplicate() {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
		Size: &sizeField,
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	if err == nil {
		suite.T().Error("Creating duplicate volume should have failed")
	}
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
		if err != nil {
			suite.T().Error("Volume clone failed")
		}
		cloneVol, _ := suite.volumeService.GetVolumeByName("TestClone")
		if cloneVol == nil {
			suite.T().Error("TestClone should have been present")
		} else {
			suite.volumeService.DeleteVolume(*cloneVol.ID)
		}
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
		if err != nil {
			suite.T().Errorf("Volume modify didnt go well %v:", err.Error())
		}
	}
	updateVol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if *updateVol.LimitMbps != limitMbpsField {
		suite.T().Error("Volume modify failed to update LimitMbps ")
	}
	if *updateVol.Size != sizeIncField {
		suite.T().Error("Volume size increase failed.")
	}
	var sizeDecField int64 = 4000
	newVolume = &nimbleos.Volume{
		Size: &sizeDecField,
	}
	if vol != nil {
		_, err := suite.volumeService.UpdateVolume(*vol.ID, newVolume)
		if err == nil {
			suite.T().Errorf("Decreasing volume size without force should have failed")
		}
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateOfflineVolume() {
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if vol != nil {
		_, err := suite.volumeService.OfflineVolume(*vol.ID, false)
		if err != nil {
			suite.T().Errorf("Setting volume offline failed %v", err.Error())
		}
		vol, _ = suite.volumeService.GetVolumeByName(volumeName)
		if *vol.VolState != *nimbleos.NsVolStatusOffline {
			suite.T().Errorf("Volume did not go offline")
		}

	} else {
		suite.T().Errorf("Could not find the volume")
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateOnlineVolume() {
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if vol != nil {
		_, err := suite.volumeService.OnlineVolume(*vol.ID, false)
		if err != nil {
			suite.T().Errorf("Setting volume online failed %v", err.Error())
		}
		vol, _ = suite.volumeService.GetVolumeByName(volumeName)
		if *vol.VolState != *nimbleos.NsVolStatusOnline {
			suite.T().Errorf("Volume did not come online")
		}
	} else {
		suite.T().Error("Could not find volume")
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeVolColAssociate() {
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	volcol, _ := suite.volcolService.GetVolumeCollectionByName(defaultVolColName)
	err := suite.volumeService.AssociateVolume(*vol.ID, *volcol.ID)
	if err != nil {
		suite.T().Errorf("Associating volume to a volume collection failed: %v", err.Error())
	}
	err = suite.volumeService.DisassociateVolume(*vol.ID)
	if err != nil {
		suite.T().Errorf("Disassociating volume to a volume collection failed: %v", err.Error())
	}
}

func TestVolumeWorkflowSuite(t *testing.T) {
	suite.Run(t, new(VolumeWorkflowSuite))
}
