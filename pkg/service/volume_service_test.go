// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param/pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VolumeServiceTestSuite struct {
	suite.Suite
	groupService              *NsGroupService
	volumeService             *VolumeService
	performancePolicyService  *PerformancePolicyService
	volumeCollectionService   *VolumeCollectionService
	snapshotCollectionService *SnapshotCollectionService
}

func (suite *VolumeServiceTestSuite) config() *NsGroupService {
	groupService, err := NewNsGroupService("10.18.174.8", "admin", "admin")
	if err != nil {
		suite.T().Errorf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
		return nil
	}
	// set debug
	groupService.SetDebug()
	return groupService
}

func (suite *VolumeServiceTestSuite) SetupTest() {
	groupService := suite.config()
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.performancePolicyService = groupService.GetPerformancePolicyService()
	suite.volumeCollectionService = groupService.GetVolumeCollectionService()
	suite.snapshotCollectionService = groupService.GetSnapshotCollectionService()

	suite.createVolume("GetVolume")
	suite.createVolume("DeleteVolume")
	suite.createVolume("DestroyVolume")
	//volcoll
	suite.createVolColl("TestVolColl")
}

func (suite *VolumeServiceTestSuite) TearDownTest() {
	suite.deleteVolume("DestroyVolume")
	suite.deleteVolume("DeleteVolume")
	suite.deleteVolume("GetVolume")
	suite.deleteVollColl("TestVolColl")
}

func (suite *VolumeServiceTestSuite) getDefaultVolumeOptions() *nimbleos.Volume {
	perfPolicy, _ := suite.performancePolicyService.GetPerformancePolicyByName("default")
	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	newVolume := &nimbleos.Volume{
		Size:              &sizeField,
		Description:       &descriptionField,
		PerfpolicyId:      perfPolicy.ID,
		ThinlyProvisioned: param.NewBool(true),
		Online:            param.NewBool(true),
		LimitIops:         &limitIopsField,
		LimitMbps:         &limitMbpsField,
		MultiInitiator:    param.NewBool(true),
		AgentType:         nimbleos.NsAgentTypeNone,
	}
	return newVolume
}

func (suite *VolumeServiceTestSuite) createVolume(volumeName string) *nimbleos.Volume {
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
		suite.volumeService.DeleteVolume(*volume.ID)
		volume, _ = suite.volumeService.GetVolumeByName(volumeName)
	}

	if volume != nil {
		suite.T().Fatalf("deleteVolume: Failed to delete volume %s", volumeName)
	}
}

func (suite *VolumeServiceTestSuite) TestGetNonExistentVolumeByID() {
	volume, err := suite.volumeService.GetVolumeById("06aaaaaaaaaaaaaaaa000000000000000000000000")
	if err != nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Unable to ge non-existent volume, err: %v", err.Error())
		return
	}
	assert.Nil(suite.T(), volume)

	volume, err = suite.volumeService.GetVolumeById("badf0rmat")
	if err == nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Expected error")
		return
	}
}

func (suite *VolumeServiceTestSuite) TestGetVolumesPagination() {
	arg := new(param.GetParams)
	//arg.Page = new(pagination.Page)
	pagination := new(pagination.Page)
	// set batch size
	pagination.SetPageSize(2)

	arg.Page = pagination
	// fetch only 2 volumes
	for hasNextPage := true; hasNextPage; hasNextPage = arg.Page.NextPage() {
		volumes, err := suite.volumeService.GetVolumes(arg)
		if err != nil {
			suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch volumes, err: %v", err.Error())
			return
		}

		if len(volumes) > 2 {
			suite.T().Errorf("TestGetVolumesPagination(): Returned volumes are more than the requested page size")
			return
		}
	}

	// dont set batch size, should return all the volumes
	narg := new(param.GetParams)

	// fetch all volumes
	vols, err := suite.volumeService.GetVolumes(narg)
	if err != nil {
		suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch all volumes, err: %v", err.Error())
		return
	}
	if len(vols) < 3 {
		suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch all the volumes")
		return
	}

}
func (suite *VolumeServiceTestSuite) TestOnlineBulkVolumes() {
	volume, _ := suite.volumeService.GetVolumeByName("GetVolume")
	if volume != nil {
		var volList [1]*string
		volList[0] = volume.ID
		err := suite.volumeService.BulkSetOnlineAndOfflineVolumes(volList[:], false)
		if err != nil {
			suite.T().Fatalf("BulkOnlineVolumes: Failed to set volumes %v online", volList)
		}

	}
	suite.volumeService.DeleteVolume(*volume.ID)
}
func (suite *VolumeServiceTestSuite) deleteVollColl(name string) {
	volcoll, _ := suite.volumeCollectionService.GetVolumeCollectionByName("TestVolColl")
	suite.volumeCollectionService.DeleteVolumeCollection(*volcoll.ID)
}

func (suite *VolumeServiceTestSuite) createVolColl(name string) {
	volcoll := &nimbleos.VolumeCollection{
		Name: param.NewString(name),
	}
	suite.volumeCollectionService.CreateVolumeCollection(volcoll)
}
func (suite *VolumeServiceTestSuite) TestAddVolumeVolcoll() {

	volume := suite.createVolume("TestAddVolVolcoll")
	if volume != nil {
		volcoll, _ := suite.volumeCollectionService.GetVolumeCollectionByName("TestVolColl")
		// add volume to volcoll
		err := suite.volumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to add a %s volume to %s volume collection", *volume.ID, *volcoll.ID)
			return
		}

	}
	// remove volume from volcoll
	err := suite.volumeService.DisassociateVolume(*volume.ID)
	if err != nil {
		suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
		return
	}
	suite.deleteVolume("TestAddVolVolcoll")
}

func (suite *VolumeServiceTestSuite) TestRestoreVolume() {
	volume := suite.createVolume("RestoreVolume")
	if volume != nil {
		volcoll, err := suite.volumeCollectionService.GetVolumeCollectionByName("TestVolColl")
		if err != nil {
			suite.T().Fatalf("Failed to get volume collection")
			return
		}
		// add volume to volume collection
		err = suite.volumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to associate RestoreVolume to volcoll ")
			return
		}
		// create a snapshot collection
		snapColl, _ := suite.snapshotCollectionService.CreateSnapshotCollection(&nimbleos.SnapshotCollection{
			Name:      param.NewString("RestoreSnapColl"),
			VolcollId: volcoll.ID,
		})
		// set the volume offline before restore
		suite.volumeService.OfflineVolume(*volume.ID, true)
		//restore volume to snapcoll
		err = suite.volumeService.RestoreVolume(*volume.ID, *snapColl.SnapshotsList[0].SnapId)

		if err != nil {
			suite.T().Fatalf("Failed to restore volume")

		}
		// disassociate volume from volume collection
		err = suite.volumeService.DisassociateVolume(*volume.ID)
		if err != nil {
			suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
			return
		}
		suite.deleteVolume("RestoreVolume")
	}
}

// Runs all test via go test
func TestVolumeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(VolumeServiceTestSuite))
}
