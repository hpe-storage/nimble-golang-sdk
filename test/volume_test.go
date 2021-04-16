// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"	
	"testing"
)

const volumeName = "VolumeTest"
var testResult = "PASS"
var result ="PASS"

type VolumeWorkflowSuite struct {
	suite.Suite
	groupService        *service.NsGroupService
	volumeService       sdkprovider.VolumeService
	snapshotService     *service.SnapshotService
	volcollService      *service.VolumeCollectionService
	initiatorGrpService *service.InitiatorGroupService
}
func (suite *VolumeWorkflowSuite) SetupTest() {
    fmt.Println("SetupTest")
	if testResult == "PASS" && result == "PASS" {
	testResult = "PASS"	
	} else {
		testResult = "FAIL"
	}
	fmt.Println(testResult)    
}
func (suite *VolumeWorkflowSuite) SetupSuite() {
	fmt.Printf("setting up suite\n")
	groupService, err := config()
	require.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.snapshotService = groupService.GetSnapshotService()
	suite.volcollService = groupService.GetVolumeCollectionService()
	suite.initiatorGrpService = groupService.GetInitiatorGroupService()
	_, err = createDefaultInitiatorGrp(suite.initiatorGrpService)
	require.Nilf(suite.T(), err, "Unable to create default initiator group, err: %v", err)
	_, err = createDefaultVolColl(suite.volcollService)
	require.Nilf(suite.T(), err, "Unable to create default volume collection, err: %v", err)
	suite.createVolume()
	
}


func (suite *VolumeWorkflowSuite) TearDownSuite() {
	
	if testResult == "PASS" && result == "PASS" {
		testResult = "PASS"	
	} else {
			testResult = "FAIL"
	}
	fmt.Println("teardo testResult")
	fmt.Println(testResult)
	if pushResultToDB == true {
	pushResultToDashboard(testResult,"C545072","Volume workflow")
	}
	fmt.Printf("tear down up suite\n")
	suite.deleteVolume(volumeName)
	err := deleteDefaultInitiatorGroup(suite.initiatorGrpService)
	require.Nilf(suite.T(), err, "Unable to delete default initiator group, err: %v", err)
	err = deleteDefaultVolColl(suite.volcollService)
	require.Nilf(suite.T(), err, "Unable to delete default volume collection, err: %v", err)
	suite.groupService.LogoutService()
}

func (suite *VolumeWorkflowSuite) deleteVolume(volumeName string) {
	volObj, _ := suite.volumeService.GetVolumeByName(volumeName)
	_, err := suite.volumeService.OfflineVolume(*volObj.ID, true)
	require.Nilf(suite.T(), err, "Unable to offline volume, err: %v", err)
	if volObj != nil {
		err := suite.volumeService.DeleteVolume(*volObj.ID)
		require.Nilf(suite.T(), err, "Unable to delete volume, err: %v", volumeName)
	}
}

func (suite *VolumeWorkflowSuite) createVolume() {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
		Size: &sizeField,
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	require.Nilf(suite.T(), err, "Volume creation failed with error: %v", err)
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	require.Equal(suite.T(), sizeField, *vol.Size, "Size was not set correctly")
}

func (suite *VolumeWorkflowSuite) TestVolumeCreateWithMissParams() {
	result ="FAIL"
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	require.NotNil(suite.T(), err, "Volume creation should have failed")
	result = "PASS"
}

func (suite *VolumeWorkflowSuite) TestVolumeCreateDuplicate() {
	result ="FAIL"
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(volumeName),
		Size: &sizeField,
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	require.NotNil(suite.T(), err, "Creating duplicate volume should have failed")
	result = "PASS"
}

func (suite *VolumeWorkflowSuite) TestCreateVolumeByCloning() {
	result ="FAIL"
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
		fmt.Printf("tests executed\n")
		require.Nilf(suite.T(), err, "Unable to clone volume, err: %v", err)
		cloneVol, _ := suite.volumeService.GetVolumeByName("TestClone")
		require.NotNil(suite.T(), cloneVol, "TestClone should have been present")
		_, err = suite.volumeService.OfflineVolume(*cloneVol.ID, true)
		require.Nilf(suite.T(), err, "Unable to offline volume, err: %v", err)
		err = suite.volumeService.DeleteVolume(*cloneVol.ID)
		require.Nilf(suite.T(), err, "Unable to delete volume, err: %v", err)
		result = "PASS"
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateSize() {
	result ="FAIL"
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
		require.Nilf(suite.T(), err, "Unable to update volume size, err: %v", err)
	}
	updateVol, _ := suite.volumeService.GetVolumeByName(volumeName)
	require.Equal(suite.T(), *updateVol.LimitMbps, limitMbpsField, "Volume modify failed to update LimitMbps")
	require.Equal(suite.T(), *updateVol.Size, sizeIncField, "Volume size increase failed.")
	var sizeDecField int64 = 4000
	newVolume = &nimbleos.Volume{
		Size: &sizeDecField,
	}
	if vol != nil {
		_, err := suite.volumeService.UpdateVolume(*vol.ID, newVolume)
		require.NotNil(suite.T(), err, "Decreasing volume size without force should have failed")
		result = "PASS"
	}
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateOfflineVolume() {		
	result = "FAIL"
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)	
	if vol != nil {
		_, err := suite.volumeService.OfflineVolume(*vol.ID, false)		
		require.Nilf(suite.T(), err, "Unable to set volume offline, err: %v", err)
		vol, _ = suite.volumeService.GetVolumeByName(volumeName)
		require.Equal(suite.T(),*vol.VolState, *nimbleos.NsVolStatusOffline, "Volume did not go offline")		
		result = "PASS" 
	} else {
		suite.T().Errorf("Could not find the volume")		
	}	
	
}

func (suite *VolumeWorkflowSuite) TestVolumeUpdateOnlineVolume() {
	
	result = "FAIL"
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	if vol != nil {
		_, err := suite.volumeService.OnlineVolume(*vol.ID, false)
		fmt.Println(err)
		require.Nilf(suite.T(), err, "Unable to set volume online, error: %v", err)
		vol, _ = suite.volumeService.GetVolumeByName(volumeName)
		fmt.Println(*vol.VolState)
		require.Equal(suite.T(),*vol.VolState, *nimbleos.NsVolStatusOnline, "Volume did not come online")
		result = "PASS"
		
	} else {		
		suite.T().Error("Could not find volume")
		
	}	
}

func (suite *VolumeWorkflowSuite) TestVolumeVolCollAssociate() {
	result ="FAIL"
	vol, _ := suite.volumeService.GetVolumeByName(volumeName)
	volcoll, _ := suite.volcollService.GetVolumeCollectionByName(defaultVolCollName)
	err := suite.volumeService.AssociateVolume(*vol.ID, *volcoll.ID)
	require.Nilf(suite.T(), err, "Associating volume to a volume collection failed: %v", err)
	err = suite.volumeService.DisassociateVolume(*vol.ID)
	require.Nilf(suite.T(), err, "Disassociating volume to a volume collection failed: %v", err)
	result = "PASS"
}

func TestVolumeWorkflowSuite(t *testing.T) {
	suite.Run(t, new(VolumeWorkflowSuite))
}

