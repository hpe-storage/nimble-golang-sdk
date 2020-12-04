// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"os"
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param/pagination"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type VolumeServiceTestSuite struct {
	suite.Suite
	groupService              *NsGroupService
	volumeService             sdkprovider.VolumeService
	performancePolicyService  *PerformancePolicyService
	volumeCollectionService   *VolumeCollectionService
	snapshotCollectionService *SnapshotCollectionService
	igroupService             *InitiatorGroupService
	aclService                *AccessControlRecordService
}

func (suite *VolumeServiceTestSuite) config() *NsGroupService {

	groupService, err := NewNsGroupService("1.1.1.1", "xxx", "xxx", "v1", true)
	if err != nil {
		suite.T().Errorf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
		return nil
	}
	// set debug
	//groupService.SetDebug()
	return groupService
}

func (suite *VolumeServiceTestSuite) SetupTest() {
	groupService := suite.config()
	if groupService == nil {
		suite.T().Errorf("Failed to initialized suite.")
		os.Exit(1)
		return
	}
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.performancePolicyService = groupService.GetPerformancePolicyService()
	suite.volumeCollectionService = groupService.GetVolumeCollectionService()
	suite.snapshotCollectionService = groupService.GetSnapshotCollectionService()
	suite.igroupService = groupService.GetInitiatorGroupService()
	suite.aclService = groupService.GetAccessControlRecordService()
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

	//Logout Group service
	suite.groupService.LogoutService()

}

func (suite *VolumeServiceTestSuite) getDefaultVolumeOptions() *nimbleos.Volume {
	perfPolicy, _ := suite.performancePolicyService.GetPerformancePolicyByName("default")
	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	newVolume := &nimbleos.Volume{
		Size:           &sizeField,
		Description:    &descriptionField,
		PerfpolicyId:   perfPolicy.ID,
		Online:         param.NewBool(true),
		LimitIops:      &limitIopsField,
		LimitMbps:      &limitMbpsField,
		MultiInitiator: param.NewBool(true),
		AgentType:      nimbleos.NsAgentTypeNone,
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
		suite.volumeService.DestroyVolume(*volume.ID)
		volume, _ = suite.volumeService.GetVolumeByName(volumeName)
	}

	if volume != nil {
		suite.T().Fatalf("deleteVolume: Failed to delete volume %s", volumeName)
	}
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
	suite.volumeService.DestroyVolume(*volume.ID)
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

func (suite *VolumeServiceTestSuite) TestCloneVolume() {

	volume := suite.createVolume("CloneVolume")
	if volume != nil {
		volcoll, err := suite.volumeCollectionService.GetVolumeCollectionByName("TestVolColl")
		if err != nil {
			suite.T().Fatalf("Failed to get volume collection")
			return
		}
		// add volume to volume collection
		err = suite.volumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to associate CloneVolume to volcoll ")
			return
		}
		// create a snapshot collection
		snapColl, _ := suite.snapshotCollectionService.CreateSnapshotCollection(&nimbleos.SnapshotCollection{
			Name:      param.NewString("CloneSnapColl"),
			VolcollId: volcoll.ID,
		})

		// update the properties of clone volume
		cloneVolume := &nimbleos.Volume{
			Clone:      param.NewBool(true),
			BaseSnapId: snapColl.SnapshotsList[0].SnapId,
			Name:       param.NewString("TestCloneVolume"),
		}
		//restore volume to snapcoll
		_, err = suite.volumeService.CreateVolume(cloneVolume)

		if err != nil {
			suite.T().Fatalf("Failed to clone a volume, id %s", *volume.ID)

		}
		// disassociate volume from volume collection
		err = suite.volumeService.DisassociateVolume(*volume.ID)
		if err != nil {
			suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
			return
		}
		//delete clone first
		suite.deleteVolume("TestCloneVolume")
		suite.deleteVolume("CloneVolume")

	}
}
func (suite *VolumeServiceTestSuite) TestACLVolume() {

	// create igroup with initiators
	initiator := &nimbleos.NsISCSIInitiator{
		Label:     param.NewString("iqn.1998-01.com.vmware:sasi-srm82-pesxi-4b00546a"),
		Iqn:       param.NewString("iqn.1998-01.com.vmware:sasi-srm82-pesxi-4b00546a"),
		IpAddress: param.NewString("*"),
	}

	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, initiator)

	igroup := &nimbleos.InitiatorGroup{
		Name:            param.NewString("sdkigroup"),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}
	igroup, _ = suite.igroupService.CreateInitiatorGroup(igroup)
	assert.NotNil(suite.T(), igroup)

	// create new volume and add ACL to to.
	volume := suite.createVolume("TestAclVolume")
	if volume != nil {
		acl := &nimbleos.AccessControlRecord{
			InitiatorGroupId: igroup.ID,
			VolId:            volume.ID,
		}
		// create acl
		acl, err := suite.aclService.CreateAccessControlRecord(acl)
		if err != nil {
			suite.T().Fatalf("Failed to create access control record, err %v", err)

		}
		suite.deleteVolume("TestAclVolume")
		suite.igroupService.DeleteInitiatorGroup(*igroup.ID)
	}
}

func (suite *VolumeServiceTestSuite) TestExpiredToken() {

	volume := suite.createVolume("TestExpiredToken")
	if volume != nil {
		//Set invalid session Token 9d255b36c700ec8b56e2064e67f01c45
		suite.groupService.client.SessionToken = "9d255b36c700ec8b56e2064e67f01c45"
		suite.deleteVolume("TestExpiredToken")
	}
}
func (suite *VolumeServiceTestSuite) TestGetVolumes() {

	volumes, err := suite.volumeService.GetVolumes(nil)
	if err != nil || len(volumes) == 0 {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volumes, err: %v", err.Error())
		return
	}

}
func (suite *VolumeServiceTestSuite) TestVolumeSortFilters() {
	// sort unit test
	sfilter := new(param.GetParams)
	var sortOrderList []param.SortOrder
	sortOrderList = append(sortOrderList, param.SortOrder{
		Field:     *nimbleos.VolumeFields.Name,
		Ascending: true,
	})

	sfilter.WithSortBy(sortOrderList)
	volumes, err := suite.volumeService.GetVolumes(sfilter)
	if err != nil || len(volumes) == 0 {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volumes, err: %v", err.Error())
		return
	}

}

func (suite *VolumeServiceTestSuite) TestVolumeSearchFilters() {
	// search filter unit test
	sfilter := &param.GetParams{}
	sf := &param.SearchFilter{
		FieldName: nimbleos.VolumeFields.Name,
		Operator:  param.EQUALS.String(),
		Value:     "GetVolume",
	}

	sfilter.WithSearchFilter(sf)
	volumes, _ := suite.volumeService.GetVolumes(sfilter)
	if len(volumes) == 0 {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volumes")
		return
	}

}
func (suite *VolumeServiceTestSuite) TestVolumeMultiSearchFilters() {
	// compound search filter unit test
	filters := &param.GetParams{}

	var searchList []*param.SearchFilter
	sf1 := &param.SearchFilter{
		FieldName: nimbleos.VolumeFields.Name,
		Operator:  param.EQUALS.String(),
		Value:     "GetVolume",
	}
	sf2 := &param.SearchFilter{
		FieldName: nimbleos.VolumeFields.SearchName,
		Operator:  param.EQUALS.String(),
		Value:     "Volume",
	}
	searchList = append(searchList, sf1)
	searchList = append(searchList, sf2)

	// Logical AND on multiple filters
	compoundFilter := &param.SearchFilter{
		Operator: param.OR.String(),
		Criteria: searchList,
	}
	filters.WithSearchFilter(compoundFilter)

	_, err := suite.volumeService.GetVolumes(filters)
	if err != nil {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volume, err: %v", err.Error())
		return
	}
}

func (suite *VolumeServiceTestSuite) TestVolumeFields() {
	// fields unit test
	filter := &param.GetParams{}
	var volAttrList = []string{
		*nimbleos.VolumeFields.ID,
		*nimbleos.VolumeFields.Name,
		*nimbleos.VolumeFields.PerfpolicyName,
	}

	filter.WithFields(volAttrList)

	sf := &param.SearchFilter{
		FieldName: nimbleos.VolumeFields.Name,
		Operator:  param.EQUALS.String(),
		Value:     "GetVolume",
	}
	filter.WithSearchFilter(sf)

	_, err := suite.volumeService.GetVolumes(filter)
	if err != nil {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volume")
		return
	}
}

// Runs all test via go test
func TestVolumeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(VolumeServiceTestSuite))
}
