// Copyright 2020 Hewlett Packard Enterprise Development LP

// Group tests
// ------------
//    1. Get group info
//    2. validate group merge
//    3. Do group merge
//    4. modify default_iscsi_target_scope
// Pool tests
// -----------
//   1. Get pool details - get pools name
//   2. Merge two pools
//   3. Edit Pool - remove array from pool
//   4. Create a new pool with available / above removed pool
//   5. Change pool name
//   6. Create volume in new pool
//   7. Delete pool when volumes are present
//       Should fail, validate error msg
//   8. Delete volume from new pool
//   9. Delete pool
//       Should be deleted

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GroupPoolWorkflowSuite struct {
	suite.Suite
	groupService      *service.NsGroupService
	arraygroupService *service.GroupService
	volumeService     sdkprovider.VolumeService
	poolService       sdkprovider.PoolService
	arrayService      *service.ArrayService
}

func (suite *GroupPoolWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Could not connect to array: %v", arrayIP)
	suite.groupService = groupService
	suite.arraygroupService = groupService.GetGroupService()
	suite.poolService = groupService.GetPoolService()
	suite.arrayService = groupService.GetArrayService()
	suite.volumeService = groupService.GetVolumeService()
}

func (suite *GroupPoolWorkflowSuite) TearDownSuite() {
	suite.groupService.LogoutService()
}

func (suite *GroupPoolWorkflowSuite) connectSourceArray() *service.NsGroupService {
	srcgroupService, err := service.NewNsGroupService(*sourceArrayIP, *sourceArrayUsername, *sourceArrayPassword, "v1", true)
	assert.Nilf(suite.T(), err, "Could not connect to array: %v", sourceArrayIP)
	srcgroupService.SetDebug()
	return srcgroupService
}

// Get Leader array pool details
func (suite *GroupPoolWorkflowSuite) TestGetGroupDetailsLeaderArray() {
	filter := &param.GetParams{}
	groupResp, err := suite.arraygroupService.GetGroups(filter)
	groupName := *groupResp[0].Name
	fmt.Printf("source array group name %+v\n", groupName)
	assert.Nil(suite.T(), err, "Failed to get group details")
	_, err = suite.arraygroupService.GetGroupByName(groupName)
	assert.Nilf(suite.T(), err, "Failed to get group by name: %v", groupName)
}

// Validate & Performe Group Merge
func (suite *GroupPoolWorkflowSuite) TestGroupMerge() {
	filter := &param.GetParams{}
	srcgroupService := suite.connectSourceArray()
	sourcearrayGroupService := srcgroupService.GetGroupService()
	sourcegroupResp, _ := sourcearrayGroupService.GetGroups(filter)
	sourcegroupName := *sourcegroupResp[0].Name
	sourcegroupArrayName := *sourcegroupResp[0].MemberList[0]
	groupResp, _ := suite.arraygroupService.GetGroups(filter)
	groupID := *groupResp[0].ID
	fmt.Printf("sourcegroup ID %+v\n", groupID)

	// Validate Group Merge
	_, err := suite.arraygroupService.ValidateMergeGroup(groupID, sourcegroupName, *sourceArrayIP, *sourceArrayUsername, *sourceArrayPassword, nil, nil)
	assert.Nil(suite.T(), err, "Group Merge Validate failed")

	// Perform group merge
	_, err = suite.arraygroupService.MergeGroup(groupID, sourcegroupName, *sourceArrayIP, *sourceArrayUsername, *sourceArrayPassword, nil, nil, nil)
	fmt.Printf("Wait 2 mins for group merge")
	time.Sleep(120 * time.Second)

	// Check if the array is added to leader group
	groupResp, _ = suite.arraygroupService.GetGroups(filter)
	groupMemberList := groupResp[0].MemberList
	var exisitingGroupMemberList []string
	for i := 0; i < len(groupMemberList); i++ {
		exisitingGroupMemberList = append(exisitingGroupMemberList, *groupMemberList[i])
	}
	_, contains := contains(exisitingGroupMemberList, sourcegroupArrayName)
	if !contains {
		suite.T().Errorf("Array %v not found in group after group merge", sourcegroupArrayName)
	}
}

// Edit Default target Scope
func (suite *GroupPoolWorkflowSuite) TestUpdateGroupProperties() {
	filter := &param.GetParams{}
	groupResp, _ := suite.arraygroupService.GetGroups(filter)
	groupID := *groupResp[0].ID
	var newDefaultIscsiTargetScope *nimbleos.NsTargetScope
	var newDefaultVolumeLimit int64 = 98
	arrayVersion := getArrayVersion(suite.arraygroupService)
	fmt.Printf("arrayVersion: %v\n", arrayVersion)
	var updateDefaultTargetScope *nimbleos.Group
	if arrayVersion < 5.1 {
		updateDefaultTargetScope = &nimbleos.Group{
			DefaultVolumeLimit: param.NewInt64(newDefaultVolumeLimit),
		}
	} else {
		defaultIscsiTargetScope := *groupResp[0].DefaultIscsiTargetScope
		if defaultIscsiTargetScope == "volume" {
			newDefaultIscsiTargetScope = nimbleos.NsTargetScopeGroup
		} else {
			newDefaultIscsiTargetScope = nimbleos.NsTargetScopeVolume
		}
		updateDefaultTargetScope = &nimbleos.Group{
			DefaultIscsiTargetScope: newDefaultIscsiTargetScope,
			DefaultVolumeLimit:      param.NewInt64(newDefaultVolumeLimit),
		}
	}
	_, err := suite.arraygroupService.UpdateGroup(groupID, updateDefaultTargetScope)
	assert.Nil(suite.T(), err, "Default iscsi target update failed")
	groupResponse, _ := suite.arraygroupService.GetGroupById(groupID)
	assert.Equal(suite.T(), newDefaultVolumeLimit, *groupResponse.DefaultVolumeLimit, "Default Volume Limit not updated")
	if arrayVersion >= 5.1 {
		var updatedDefaultIscsiTargetScope *nimbleos.NsTargetScope
		if *groupResponse.DefaultIscsiTargetScope == "volume" {
			updatedDefaultIscsiTargetScope = nimbleos.NsTargetScopeVolume
		} else {
			updatedDefaultIscsiTargetScope = nimbleos.NsTargetScopeGroup
		}
		assert.Equal(suite.T(), newDefaultIscsiTargetScope, updatedDefaultIscsiTargetScope, "Default Iscsi TargetScope not updated")
	}
}

// Create Volume in the specified pool
func (suite *GroupPoolWorkflowSuite) CreateVolumeWithPoolName(volumeName string, poolName string) {
	poolResp := suite.GetPoolDetail(poolName)
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name:   param.NewString(volumeName),
		Size:   &sizeField,
		PoolId: param.NewString(*poolResp.ID),
		Online: param.NewBool(false),
	}
	_, err := suite.volumeService.CreateVolume(newVolume)
	assert.Nilf(suite.T(), err, "Unable to create volume %v", volumeName)
}

// Delete Volume
func (suite *GroupPoolWorkflowSuite) deleteVolume(volumeName string) {
	volObj, _ := suite.volumeService.GetVolumeByName(volumeName)
	err := suite.volumeService.DeleteVolume(*volObj.ID)
	assert.Nilf(suite.T(), err, "Unable to delete volume %v", volumeName)

}

// Get Pool Detail
func (suite *GroupPoolWorkflowSuite) GetPoolDetail(poolName string) *nimbleos.Pool {
	poolResp, err := suite.poolService.GetPoolByName(poolName)
	assert.Nilf(suite.T(), err, "Unable to get %v pool details", poolName)
	return poolResp
}

// Merge non-default Pool to default Pool
func (suite *GroupPoolWorkflowSuite) TestMergePool() {
	// Get Non default pool name from array (after group merge a new pool will be created)
	filter := &param.GetParams{}
	poolResp, _ := suite.poolService.GetPools(filter)
	var poolCount int = len(poolResp)
	var nonDefaulPoolName string = ""
	var nonDefaulPoolID string = ""
	for i := 0; i < poolCount; i++ {
		if *poolResp[i].Name != "default" {
			nonDefaulPoolName = *poolResp[i].Name
			nonDefaulPoolID = *poolResp[i].ID
			fmt.Printf("Non default Pool Name %+v\n", nonDefaulPoolName)
			fmt.Printf("Non default Pool ID %+v\n", nonDefaulPoolID)
		}
	}
	defaultPool := suite.GetPoolDetail("default")
	defaultPoolID := *defaultPool.ID
	fmt.Printf("Default Pool ID %+v\n", defaultPoolID)

	// Pool Merge - Merge pool to 'default' pool
	_, err := suite.poolService.MergePool(nonDefaulPoolID, defaultPoolID, nil)
	assert.Nil(suite.T(), err, "Pool Merge failed")
}

// Edit Pool - remove array from pool
func (suite *GroupPoolWorkflowSuite) TestModifyPoolArray() {
	// defaultPool, _ := suite.poolService.GetPoolByName("default")
	defaultPool := suite.GetPoolDetail("default")
	existingPoolArrayList := defaultPool.ArrayList
	defaultPoolID := *defaultPool.ID
	arrayCount := int(*defaultPool.ArrayCount)
	fmt.Printf("Default Pool Array Count %+v\n", arrayCount)

	// Get one arrayId to retain in pool. Others Arrays will be unassigned from pool
	arrayID := *existingPoolArrayList[0].ID
	arraylist := &nimbleos.NsArrayDetail{
		ID:      param.NewString(arrayID),
		ArrayId: param.NewString(arrayID),
	}

	var updateArrayList []*nimbleos.NsArrayDetail
	updateArrayList = append(updateArrayList, arraylist)

	updatePoolInput := &nimbleos.Pool{
		ArrayList: updateArrayList,
	}
	_, err := suite.poolService.UpdatePool(defaultPoolID, updatePoolInput)
	assert.Nil(suite.T(), err, "Pool Array list update failed")
	fmt.Printf("Updated pool array list\n")
	time.Sleep(5 * time.Second)
}

// Check if the element is present in the provided list/array/slice
func contains(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Create, Modify, Delete new pool with unassigned array
func (suite *GroupPoolWorkflowSuite) TestNewPoolCreateModifyDelete() {
	// Get default pool details to get assigned arrays details
	fmt.Printf("Starting new pool creation")
	filter := &param.GetParams{}
	groupResp, _ := suite.arraygroupService.GetGroups(filter)
	defaultPool := suite.GetPoolDetail("default")
	fmt.Printf("Pool array list %+v\n", defaultPool.ArrayList)
	groupMemberList := groupResp[0].MemberList
	fmt.Printf("Group member array list %+v\n", groupMemberList)
	arrayCount := int(*defaultPool.ArrayCount)
	var existingPoolArrayList []string
	for i := 0; i < arrayCount; i++ {
		existingPoolArrayList = append(existingPoolArrayList, *defaultPool.ArrayList[i].Name)
	}
	fmt.Printf("Array(s) Assigned to default pool %+v\n", existingPoolArrayList)

	// Get an Array name which is not assigned to default Pool
	var unassignedArrayName string
	fmt.Printf("initiated Unassigned array name %+v\n", unassignedArrayName)
	for i := 0; i < len(groupMemberList); i++ {
		_, contains := contains(existingPoolArrayList, *groupMemberList[i])
		if !contains {
			unassignedArrayName = *groupMemberList[i]
			break
		}
	}
	fmt.Printf("Unassigned array name %+v\n", unassignedArrayName)
	respArray, _ := suite.arrayService.GetArrayByName(unassignedArrayName)
	unassignedArrayID := *respArray.ID
	fmt.Printf("Unassigned arrayID %+v\n", unassignedArrayID)
	var newPoolName = "testnondefaultpool"
	// Create new pool with unassigned array
	arraylist := &nimbleos.NsArrayDetail{
		ID:      param.NewString(unassignedArrayID),
		ArrayId: param.NewString(unassignedArrayID),
	}
	var poolArrayList []*nimbleos.NsArrayDetail
	poolArrayList = append(poolArrayList, arraylist)
	newPooolDesription := "Creating a new for GoLang SDK testing"
	createPoolInput := &nimbleos.Pool{
		ArrayList:   poolArrayList,
		Name:        param.NewString(newPoolName),
		Description: param.NewString(newPooolDesription),
	}
	createpoolResp, err := suite.poolService.CreatePool(createPoolInput)
	assert.Nilf(suite.T(), err, "Failed to create Pool: %v", newPoolName)
	newPoolID := *createpoolResp.ID
	assert.Equal(suite.T(), newPooolDesription, *createpoolResp.Description, "Pool description does not match")

	// Modify pool name & description
	newPoolName = "newpooltest"
	newDescription := "Updated pool description for GoLang SDK testing"
	updatePoolInput := &nimbleos.Pool{
		Name:        param.NewString(newPoolName),
		Description: param.NewString(newDescription),
	}
	_, err = suite.poolService.UpdatePool(newPoolID, updatePoolInput)
	assert.Nilf(suite.T(), err, "Failed to update Pool: %v", newPoolName)
	poolResp := suite.GetPoolDetail(newPoolName)
	assert.Equal(suite.T(), newPoolName, *poolResp.Name, "Pool name nnot updated")
	assert.Equal(suite.T(), newDescription, *poolResp.Description, "Pool description not updated")

	// Create Volume in New Pool
	volumeName := "testvolume"
	suite.CreateVolumeWithPoolName(volumeName, newPoolName)

	// Delete pool when volumes are present, expected to fail
	err = suite.poolService.DeletePool(newPoolID)
	assert.NotNil(suite.T(), err, "Pool Deletion expected to be failed")

	// Delete Volume
	suite.deleteVolume(volumeName)
	time.Sleep(5 * time.Second)

	// Delete Pool
	err = suite.poolService.DeletePool(newPoolID)
	assert.Nilf(suite.T(), err, "Failed to delete Pool: %v", newPoolName)
}

func TestGroupPoolWorkflowSuite(t *testing.T) {
	suite.Run(t, new(GroupPoolWorkflowSuite))
}
