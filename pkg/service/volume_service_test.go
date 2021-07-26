// Copyright 2020 - 2021 Hewlett Packard Enterprise Development LP

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
	nonTenantGroupService              *NsGroupService
	nonTenantVolumeService             sdkprovider.VolumeService
	nonTenantPerformancePolicyService  *PerformancePolicyService
	nonTenantVolumeCollectionService   *VolumeCollectionService
	nonTenantSnapshotCollectionService *SnapshotCollectionService
	nonTenantIgroupService             *InitiatorGroupService
	nonTenantAclService                *AccessControlRecordService
	nonTenantFolderService			   *FolderService
	nonTenantPoolService			   sdkprovider.PoolService
	nonTenantUserService			   *UserService

	tenantGroupService       	        *NsGroupService
	tenantVolumeService      	        sdkprovider.VolumeService
	tenantPerformancePolicyService  	*PerformancePolicyService
	tenantVolumeCollectionService  	 	*VolumeCollectionService
	tenantSnapshotCollectionService		*SnapshotCollectionService
	tenantIgroupService             	*InitiatorGroupService
	tenantAclService                	*AccessControlRecordService
	tenantFolderService			  		*FolderService
}

func (suite *VolumeServiceTestSuite) config() (*NsGroupService) {

	nonTenantGroupService, err := NewNimbleGroupService(WithHost("10.157.82.90"),
		WithUser("admin"), WithPassword("admin"))

	if err != nil {
		suite.T().Errorf("NewGroupService(): Unable to connect to non-tenant group, err: %v", err.Error())
		return nil
	}

	// set debug
	//groupService.SetDebug()
	return nonTenantGroupService
}

func (suite *VolumeServiceTestSuite) SetupTest() {
	nonTenantGroupService := suite.config()
	if nonTenantGroupService == nil {
		suite.T().Errorf("Failed to initialized non-tenant suite.")
		os.Exit(1)
		return
	}

	suite.nonTenantGroupService = nonTenantGroupService
	suite.nonTenantVolumeService = nonTenantGroupService.GetVolumeService()
	suite.nonTenantPerformancePolicyService = nonTenantGroupService.GetPerformancePolicyService()
	suite.nonTenantVolumeCollectionService = nonTenantGroupService.GetVolumeCollectionService()
	suite.nonTenantSnapshotCollectionService = nonTenantGroupService.GetSnapshotCollectionService()
	suite.nonTenantIgroupService = nonTenantGroupService.GetInitiatorGroupService()
	suite.nonTenantAclService = nonTenantGroupService.GetAccessControlRecordService()
	suite.nonTenantFolderService = nonTenantGroupService.GetFolderService()
	suite.nonTenantPoolService = nonTenantGroupService.GetPoolService()
	suite.nonTenantUserService = nonTenantGroupService.GetUserService()

	/**
	commenting everything out for now
	*/
	// create folder
	// folder, _ := suite.nonTenantFolderService.GetFolderByName("tenant3")
	// pool, _ := suite.nonTenantPoolService.GetPoolByName("default")

	// if folder == nil {
	// 	folderParams := &nimbleos.Folder{
	// 		Name: param.NewString("tenant3"),
	// 		PoolId:pool.ID,
	// 	}
	// 	folder, _ = suite.nonTenantFolderService.CreateFolder(folderParams)
	// }

	// create tenant
	// user, _ := suite.nonTenantUserService.GetUserByName("testtenant")
	// if user == nil {
	// 	user_param := &nimbleos.User{
	// 		Name: param.NewString("testtenant"),
	// 		Password: param.NewString("TestP4ssW0rd"),
	// 		Folders: folder,
	// 	}
	// 	user, err := suite.nonTenantUserService.CreateUser(user_param)
	// 	if err != nil {
	// 		suite.T().Errorf("Error creating a tenant user %v, message: %v", user, err.Error())
	// 	}
	// }

	// // TODO: delete later -- debugging
	// tenantId := *user.ID
	// admin, err := suite.nonTenantUserService.GetUserByName("admin")
	// adminId := *admin.ID
	// fmt.Println(tenantId, adminId, err)
	// tenantOnFolder, _:= suite.nonTenantUserService.GetUserById("2f1687ff771a5ea819000000000000000000000004")
	// fmt.Println((tenantOnFolder))

	// // try 1: assign folder's tenant id to user's tenant id.
	// userObj := &nimbleos.User{
	// 	TenantId: folder.TenantId,
	// }
	// user, _ = suite.nonTenantUserService.UpdateUser(tenantId, userObj)
	// fmt.Println(*user.TenantId)

	// // why is password NIL
	// // how to assign folder to user..?
	// tenantGroupService, err := NewNimbleGroupService(WithHost("10.157.82.90"),
	// 	WithTenantUser(*user.Name), WithPassword(*user.Password))

	// if err != nil {
	// 	suite.T().Errorf("NewGroupService(): Unable to connect to tenant group, err: %v", err.Error())
	// }

	// update folder with tenant id

	tenantGroupService, err := NewNimbleGroupService(WithHost("10.157.82.90"),
		WithTenantUser("raunak"), WithPassword("Nim123Boli"))
	if err != nil {
		suite.T().Errorf("NewGroupService(): Unable to connect to tenant group, err: %v", err.Error())
	}

	suite.tenantGroupService = tenantGroupService
	suite.tenantVolumeService = tenantGroupService.GetVolumeService()
	suite.tenantPerformancePolicyService = tenantGroupService.GetPerformancePolicyService()
	suite.tenantVolumeCollectionService = tenantGroupService.GetVolumeCollectionService()
	suite.tenantSnapshotCollectionService = tenantGroupService.GetSnapshotCollectionService()
	suite.tenantIgroupService = tenantGroupService.GetInitiatorGroupService()
	suite.tenantAclService = tenantGroupService.GetAccessControlRecordService()
	suite.tenantFolderService = tenantGroupService.GetFolderService()

	suite.nonTenantCreateVolume("GetVolume")
	suite.nonTenantCreateVolume("DeleteVolume")
	suite.nonTenantCreateVolume("DestroyVolume")

	suite.tenantCreateVolume("GetVolume")
	suite.tenantCreateVolume("DeleteVolume")
	suite.tenantCreateVolume("DestroyVolume")
	//volcoll
	suite.nonTenantCreateVolColl("TestVolColl")
	suite.tenantCreateVolColl("TestVolColl1")
}

func (suite *VolumeServiceTestSuite) TearDownTest() {

	suite.nonTenantDeleteVolume("DestroyVolume")
	suite.nonTenantDeleteVolume("DeleteVolume")
	suite.nonTenantDeleteVolume("GetVolume")

	suite.tenantDeleteVolume("DestroyVolume")
	suite.tenantDeleteVolume("DeleteVolume")
	suite.tenantDeleteVolume("GetVolume")

	suite.nonTenantDeleteVollColl("TestVolColl")
	suite.tenantDeleteVollColl("TestVolColl1")

	//Logout Group service
	suite.nonTenantGroupService.LogoutService()
	suite.tenantGroupService.LogoutService()

}

func (suite *VolumeServiceTestSuite) getDefaultVolumeOptions() *nimbleos.Volume {
	perfPolicy, _ := suite.nonTenantPerformancePolicyService.GetPerformancePolicyByName("default")
	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1
	folder, _ := suite.nonTenantFolderService.GetFolderByName("tenant1")

	newVolume := &nimbleos.Volume{
		Size:           &sizeField,
		Description:    &descriptionField,
		PerfpolicyId:   perfPolicy.ID,
		Online:         param.NewBool(true),
		LimitIops:      &limitIopsField,
		LimitMbps:      &limitMbpsField,
		MultiInitiator: param.NewBool(true),
		AgentType:      nimbleos.NsAgentTypeNone,
		FolderId:		folder.ID,
	}
	return newVolume
}

func (suite *VolumeServiceTestSuite) nonTenantCreateVolume(volumeName string) *nimbleos.Volume {
	// Create a volume for non tenant
	volume, _ := suite.nonTenantVolumeService.GetVolumeByName(volumeName)
	if volume == nil {
		volume = suite.getDefaultVolumeOptions()
		volume.Name = &volumeName
		volume, _ = suite.nonTenantVolumeService.CreateVolume(volume)
	}
	return volume
}

func (suite *VolumeServiceTestSuite) tenantCreateVolume(volumeName string) *nimbleos.Volume {
	// Create a volume for tenant
	volume, _ := suite.tenantVolumeService.GetVolumeByName(volumeName)
	if volume == nil {
		volume = suite.getDefaultVolumeOptions()
		volume.Name = &volumeName
		volume, err := suite.tenantVolumeService.CreateVolume(volume)
		if err != nil {
			suite.T().Errorf("volume=%v, error=%v", volume, err.Error())
		}
	}
	return volume
}

func (suite *VolumeServiceTestSuite) nonTenantDeleteVolume(volumeName string) {
	volume, _ := suite.nonTenantVolumeService.GetVolumeByName(volumeName)
	if volume != nil {
		suite.nonTenantVolumeService.DestroyVolume(*volume.ID)
		volume, _ = suite.nonTenantVolumeService.GetVolumeByName(volumeName)
	}

	if volume != nil {
		suite.T().Fatalf("deleteVolume: Failed to delete volume %s", volumeName)
	}
}

func (suite *VolumeServiceTestSuite) tenantDeleteVolume(volumeName string) {
	volume, _ := suite.tenantVolumeService.GetVolumeByName(volumeName)
	if volume != nil {
		suite.tenantVolumeService.DestroyVolume(*volume.ID)
		volume, _ = suite.tenantVolumeService.GetVolumeByName(volumeName)
	}

	if volume != nil {
		suite.T().Fatalf("deleteVolume: Failed to delete volume %s", volumeName)
	}
}

func (suite *VolumeServiceTestSuite) nonTenantDeleteVollColl(name string) {
	volcoll, _ := suite.nonTenantVolumeCollectionService.GetVolumeCollectionByName(name)
	suite.nonTenantVolumeCollectionService.DeleteVolumeCollection(*volcoll.ID)
}

func (suite *VolumeServiceTestSuite) nonTenantCreateVolColl(name string) {
	volcoll := &nimbleos.VolumeCollection{
		Name: param.NewString(name),
	}
	suite.nonTenantVolumeCollectionService.CreateVolumeCollection(volcoll)
}

func (suite *VolumeServiceTestSuite) tenantDeleteVollColl(name string) {
	volcoll, _ := suite.tenantVolumeCollectionService.GetVolumeCollectionByName(name)
	suite.tenantVolumeCollectionService.DeleteVolumeCollection(*volcoll.ID)
}

func (suite *VolumeServiceTestSuite) tenantCreateVolColl(name string) {
	volcoll := &nimbleos.VolumeCollection{
		Name: param.NewString(name),
	}
	suite.tenantVolumeCollectionService.CreateVolumeCollection(volcoll)
}

func (suite *VolumeServiceTestSuite) TestGetNonExistentVolumeByID() {

	// Test non tenant
	volume, err := suite.nonTenantVolumeService.GetVolumeById("06aaaaaaaaaaaaaaaa000000000000000000000000")
	if err != nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Unable to ge non-existent volume, err: %v", err.Error())
		return
	}
	assert.Nil(suite.T(), volume)

	volume, err = suite.nonTenantVolumeService.GetVolumeById("badf0rmat")
	if err == nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Expected error")
		return
	}

	// Test tenant

	volume, err = suite.tenantVolumeService.GetVolumeById("06aaaaaaaaaaaaaaaa000000000000000000000000")
	if err != nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Unable to ge non-existent volume, err: %v", err.Error())
		return
	}
	/** Ideally, volume should be NIL (i.e. API returns HTTP 404 NOT FOUND)
	 * But for some reason, it's returning HTTP 200, body = {}
	 * So right not we test it by asserting volume.ID = nil.
	 * TODO: Change this to assert.Nil(suite.T(), volume) once the bug is fixed
	 */

	assert.Nil(suite.T(), volume.ID)

	volume, err = suite.tenantVolumeService.GetVolumeById("badf0rmat")
	if err == nil {
		suite.T().Errorf("TestGetNonExistentVolumeByID(): Expected error")
		return
	}
}

func (suite *VolumeServiceTestSuite) TestGetVolumesPagination() {

	arg := new(param.GetParams)
	admin_pagination := new(pagination.Page)
	// set batch size
	admin_pagination.SetPageSize(2)

	arg.Page = admin_pagination

	// Test non tenant
	// fetch only 2 volumes
	for hasNextPage := true; hasNextPage; hasNextPage = arg.Page.NextPage() {
		volumes, err := suite.nonTenantVolumeService.GetVolumes(arg)
		if err != nil {
			suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch volumes, err: %v", err.Error())
			return
		}

		if len(volumes) > 2 {
			suite.T().Errorf("TestGetVolumesPagination(): Returned volumes are more than the requested page size")
			return
		}
	}

	// Test tenant

	/** Right now this test is commented out because there's an internal bug
	 * that can not process filters in params
	 * TODO: Uncomment this once the bug is fixed
	 */

	// arg_tenant := new(param.GetParams)
	// pagination_tenant := new(pagination.Page)
	// // set batch size
	// pagination_tenant.SetPageSize(2)

	// arg_tenant.Page = pagination_tenant
	// // fetch only 2 volumes
	// for hasNextPage := true; hasNextPage; hasNextPage = arg_tenant.Page.NextPage() {
	// 	volumes, err := suite.tenantVolumeService.GetVolumes(arg_tenant)
	// 	if err != nil {
	// 		suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch volumes, err: %v", err.Error())
	// 		return
	// 	}

	// 	if len(volumes) > 2 {
	// 		suite.T().Errorf("TestGetVolumesPagination(): Returned volumes are more than the requested page size")
	// 		return
	// 	}
	// }

	// dont set batch size, should return all the volumes
	narg := new(param.GetParams)

	// Test non tenant
	// fetch all volumes
	vols, err := suite.nonTenantVolumeService.GetVolumes(narg)
	if err != nil {
		suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch all volumes, err: %v", err.Error())
		return
	}
	if len(vols) < 3 {
		suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch all the volumes")
		return
	}

	// Test tenant
	vols, err = suite.tenantVolumeService.GetVolumes(narg)
	if err != nil {
		suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch all volumes, err: %v", err.Error())
		return
	}
	if len(vols) < 3 {
		suite.T().Errorf("TestGetVolumesPagination(): Unable to fetch all the volumes, len: %v", len(vols))
		return
	}

}
func (suite *VolumeServiceTestSuite) TestOnlineBulkVolumes() {

	// Test non tenant
	volume, _ := suite.nonTenantVolumeService.GetVolumeByName("GetVolume")
	if volume != nil {
		var volList [1]*string
		volList[0] = volume.ID
		err := suite.nonTenantVolumeService.BulkSetOnlineAndOfflineVolumes(volList[:], false)
		if err != nil {
			suite.T().Fatalf("BulkOnlineVolumes: Failed to set volumes %v online", volList)
		}

	}

	/**
	* Right now we comment this test out because the API
	* v1/volumes/actions/bulk_set_online_and_offline
	* doesn't work for tenant
	* TODO: remove comments once the bug is fixed
	*/

	// Test tenant
	// volume, _ = suite.tenantVolumeService.GetVolumeByName("GetVolume")
	// if volume != nil {
	// 	var volList [1]*string
	// 	volList[0] = volume.ID
	// 	err := suite.tenantVolumeService.BulkSetOnlineAndOfflineVolumes(volList[:], false)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		suite.T().Fatalf("BulkOnlineVolumes: Failed to set volumes %v online", volList)
	// 	}

	// }
}

func (suite *VolumeServiceTestSuite) TestAddVolumeVolcoll() {

	volume := suite.nonTenantCreateVolume("TestAddVolVolcoll")

	// Test non tenant
	if volume != nil {
		volcoll, _ := suite.nonTenantVolumeCollectionService.GetVolumeCollectionByName("TestVolColl")
		// add volume to volcoll
		err := suite.nonTenantVolumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to add a %s volume to %s volume collection", *volume.ID, *volcoll.ID)
			return
		}

	}
	// remove volume from volcoll
	err := suite.nonTenantVolumeService.DisassociateVolume(*volume.ID)
	if err != nil {
		suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
		return
	}

	volume = suite.tenantCreateVolume("TestAddVolVolcoll1")
	// Test tenant
	if volume != nil {
		volcoll, _ := suite.tenantVolumeCollectionService.GetVolumeCollectionByName("TestVolColl1")
		// add volume to volcoll
		suite.tenantVolumeService.DisassociateVolume(*volume.ID)
		err := suite.tenantVolumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to add a %s volume to %s volume collection", *volume.ID, *volcoll.ID)
			return
		}

	}
	// remove volume from volcoll
	err = suite.tenantVolumeService.DisassociateVolume(*volume.ID)
	if err != nil {
		suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
		return
	}
	suite.nonTenantDeleteVolume("TestAddVolVolcoll")
	suite.tenantDeleteVolume("TestAddVolVolcoll1")
}

func (suite *VolumeServiceTestSuite) TestRestoreVolume() {

	volume := suite.nonTenantCreateVolume("RestoreVolume")

	if volume != nil {
		volcoll, err := suite.nonTenantVolumeCollectionService.GetVolumeCollectionByName("TestVolColl")
		if err != nil {
			suite.T().Fatalf("Failed to get volume collection")
			return
		}
		// Test non-tenant: add volume to volume collection
		err = suite.nonTenantVolumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to associate RestoreVolume to volcoll ")
			return
		}
		// Test non-tenant: create a snapshot collection
		snapColl, _ := suite.nonTenantSnapshotCollectionService.CreateSnapshotCollection(&nimbleos.SnapshotCollection{
			Name:      param.NewString("RestoreSnapColl"),
			VolcollId: volcoll.ID,
		})
		// Test non-tenant: set the volume offline before restore
		suite.nonTenantVolumeService.OfflineVolume(*volume.ID, true)
		// Test non-tenant: restore volume to snapcoll
		err = suite.nonTenantVolumeService.RestoreVolume(*volume.ID, *snapColl.SnapshotsList[0].SnapId)

		if err != nil {
			suite.T().Fatalf("Failed to restore volume")

		}
		// Test non-tenant: disassociate volume from volume collection
		err = suite.nonTenantVolumeService.DisassociateVolume(*volume.ID)
		if err != nil {
			suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
			return
		}
		suite.nonTenantDeleteVolume("RestoreVolume")
	}

	// Tenant is not authorized to restore volume, no need to test.
}

func (suite *VolumeServiceTestSuite) TestCloneVolume() {
	volume := suite.nonTenantCreateVolume("CloneVolume")
	if volume != nil {
		volcoll, err := suite.nonTenantVolumeCollectionService.GetVolumeCollectionByName("TestVolColl")
		if err != nil {
			suite.T().Fatalf("Failed to get volume collection")
			return
		}
		// Test non-tenant: add volume to volume collection
		suite.nonTenantVolumeService.DisassociateVolume(*volume.ID)
		err = suite.nonTenantVolumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to associate CloneVolume to volcoll ")
			return
		}
		// Test non-tenant: create a snapshot collection
		snapColl, _ := suite.nonTenantSnapshotCollectionService.GetSnapshotCollectionByName("CloneSnapColl")
		if snapColl == nil {
			snapColl, _ = suite.nonTenantSnapshotCollectionService.CreateSnapshotCollection(&nimbleos.SnapshotCollection{
				Name:      param.NewString("CloneSnapColl"),
				VolcollId: volcoll.ID,
			})
		}

		// Test non-tenant: update the properties of clone volume
		cloneVolume := &nimbleos.Volume{
			Clone:      param.NewBool(true),
			BaseSnapId: snapColl.SnapshotsList[0].SnapId,
			Name:       param.NewString("TestCloneVolume"),
		}
		// Test non-tenant: restore volume to snapcoll
		_, err = suite.nonTenantVolumeService.CreateVolume(cloneVolume)

		if err != nil {
			suite.T().Fatalf("Failed to clone a volume, id %s", *volume.ID)

		}
		// Test non-tenant: disassociate volume from volume collection
		err = suite.nonTenantVolumeService.DisassociateVolume(*volume.ID)
		if err != nil {
			suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
			return
		}

		//delete clone first
		suite.nonTenantDeleteVolume("TestCloneVolume")
		suite.nonTenantDeleteVolume("CloneVolume")
	}

	volume = suite.nonTenantCreateVolume("CloneVolume")
	if volume != nil {
		volcoll, err := suite.tenantVolumeCollectionService.GetVolumeCollectionByName("TestVolColl1")
		// Test tenant: add volume to volume collection
		err = suite.tenantVolumeService.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			suite.T().Fatalf("Failed to associate CloneVolume to volcoll ")
			return
		}
		// Test tenant: create a snapshot collection
		snapColl, _ := suite.nonTenantSnapshotCollectionService.GetSnapshotCollectionByName("CloneSnapColl1")
		if snapColl == nil {
			snapColl, _ = suite.tenantSnapshotCollectionService.CreateSnapshotCollection(&nimbleos.SnapshotCollection{
				Name:      param.NewString("CloneSnapColl"),
				VolcollId: volcoll.ID,
			})
		}

		// Test tenant: update the properties of clone volume
		cloneVolume := &nimbleos.Volume{
			Clone:      param.NewBool(true),
			BaseSnapId: snapColl.SnapshotsList[0].SnapId,
			Name:       param.NewString("TestCloneVolume"),
			FolderId: 	volume.FolderId,
		}
		// Test tenant: restore volume to snapcoll
		_, err = suite.tenantVolumeService.CreateVolume(cloneVolume)

		if err != nil {
			suite.T().Fatalf("Failed to clone a volume, id %s. err=%s", *volume.ID, err.Error())
		}
		// Test tenant: disassociate volume from volume collection
		err = suite.tenantVolumeService.DisassociateVolume(*volume.ID)
		if err != nil {
			suite.T().Fatalf("Failed to remove %s volume from volume collection", *volume.ID)
			return
		}
		//delete clone first
		suite.tenantDeleteVolume("TestCloneVolume")
		suite.tenantDeleteVolume("CloneVolume")
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

	ig, _ := suite.nonTenantIgroupService.GetInitiatorGroupByName("sdkigroup")
	if ig == nil {
		ig, _ = suite.nonTenantIgroupService.CreateInitiatorGroup(igroup)
	}
	assert.NotNil(suite.T(), ig)

	// create new volume and add ACL to to.
	volume := suite.nonTenantCreateVolume("TestAclVolume")
	if volume != nil {
		acl := &nimbleos.AccessControlRecord{
			InitiatorGroupId: ig.ID,
			VolId:            volume.ID,
		}
		// NonTenant: create acl
		acl, err := suite.nonTenantAclService.CreateAccessControlRecord(acl)
		if err != nil {
			suite.T().Fatalf("Failed to create access control record, err %v", err)
		}
		suite.nonTenantDeleteVolume("TestAclVolume")
		suite.nonTenantIgroupService.DeleteInitiatorGroup(*ig.ID)
	}

	igroup_tenant := &nimbleos.InitiatorGroup{
		Name:            param.NewString("igrouptenant1"),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}

	ig_tenant, _ := suite.tenantIgroupService.GetInitiatorGroupByName("igrouptenant1")
	if ig_tenant == nil {
		ig_tenant, _ = suite.tenantIgroupService.CreateInitiatorGroup(igroup_tenant)
	}
	assert.NotNil(suite.T(), ig_tenant)

	volume = suite.tenantCreateVolume("TestAclVolume1")
	if volume != nil {
		acl := &nimbleos.AccessControlRecord{
			InitiatorGroupId: ig_tenant.ID,
			VolId:            volume.ID,
		}
		// Tenant: create acl
		acl, err := suite.tenantAclService.CreateAccessControlRecord(acl)
		if err != nil {
			suite.T().Fatalf("Failed to create access control record, err %v", err)

		}

		suite.tenantDeleteVolume("TestAclVolume1")
		suite.tenantIgroupService.DeleteInitiatorGroup(*ig_tenant.ID)
	}
}

func (suite *VolumeServiceTestSuite) TestExpiredToken() {
	volume := suite.nonTenantCreateVolume("TestExpiredToken")
	if volume != nil {
		//Set invalid session Token 9d255b36c700ec8b56e2064e67f01c45
		suite.nonTenantGroupService.client.SessionToken = "9d255b36c700ec8b56e2064e67f01c45"
		suite.nonTenantDeleteVolume("TestExpiredToken")
	}

	volume = suite.tenantCreateVolume("TestExpiredToken1")
	if volume != nil {
		//Set invalid session Token 9d255b36c700ec8b56e2064e67f01c45
		suite.tenantGroupService.client.SessionToken = "9d255b36c700ec8b56e2064e67f01c45"
		suite.tenantDeleteVolume("TestExpiredToken1")
	}
}
func (suite *VolumeServiceTestSuite) TestGetVolumes() {

	// Test non-tenant
	volumes, err := suite.nonTenantVolumeService.GetVolumes(nil)
	if err != nil || len(volumes) == 0 {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volumes, err: %v", err.Error())
		return
	}

	// Test tenant
	volumes, err = suite.tenantVolumeService.GetVolumes(nil)
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

	// Test nonTenant
	volumes, err := suite.nonTenantVolumeService.GetVolumes(sfilter)
	if err != nil || len(volumes) == 0 {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volumes, err: %v", err.Error())
		return
	}

	// Test Tenant
	/** Right now this test is commented out because there's an internal bug
	 * that can not process filters in params
	 * TODO: Uncomment this once the bug is fixed
	 */

	// volumes, err = suite.tenantVolumeService.GetVolumes(sfilter)
	// if err != nil || len(volumes) == 0 {
	// 	suite.T().Errorf("TestGetVolumes(): Unable to fetch volumes, err: %v", err.Error())
	// 	return
	// }

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

	// Test Non-tenant
	volumes, _ := suite.nonTenantVolumeService.GetVolumes(sfilter)
	if len(volumes) == 0 {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volumes")
		return
	}

	// Test tenant
	volumes, _ = suite.tenantVolumeService.GetVolumes(sfilter)
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

	// Test non-tenant
	_, err := suite.nonTenantVolumeService.GetVolumes(filters)
	if err != nil {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volume, err: %v", err.Error())
		return
	}

	// Test tenant
	_, err = suite.tenantVolumeService.GetVolumes(filters)
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

	// Test non tenant
	_, err := suite.nonTenantVolumeService.GetVolumes(filter)
	if err != nil {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volume")
		return
	}

	// Test tenant
	_, err = suite.tenantVolumeService.GetVolumes(filter)
	if err != nil {
		suite.T().Errorf("TestGetVolumes(): Unable to fetch volume")
		return
	}
}

// Runs all test via go test
func TestVolumeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(VolumeServiceTestSuite))
}
