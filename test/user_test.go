// Copyright 2020 Hewlett Packard Enterprise Development LP

// User Create/Delete tests:
// --------------------
// 1. Create a poweruser
// 		Validate power user properties
// 		Create volume using newly created poweruser
// 		Edit poweruser to guest user
// 		Delete volume
// 2. Create Guest USer
// 		Validate guest user properties
// 3. Delete both users & Validate user deleted
// 4. Update User Policy

package test

import (
	"fmt"
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const powerUserName = "testpoweruser"
const guestUserName = "testguestuser"

type UserWorkflowSuite struct {
	suite.Suite
	groupService      *service.NsGroupService
	userService       *service.UserService
	volumeService     sdkprovider.VolumeService
	userPolicyService *service.UserPolicyService
	arrayGroupService *service.GroupService
}

func (suite *UserWorkflowSuite) SetupSuite() {
	groupService, err := config()
	require.Nilf(suite.T(), err, "Could not connect to array")
	suite.groupService = groupService
	suite.userService = groupService.GetUserService()
	suite.userPolicyService = groupService.GetUserPolicyService()
	suite.arrayGroupService = groupService.GetGroupService()
}

func (suite *UserWorkflowSuite) TearDownSuite() {
	var userTestResult string
	if len(testStarted) == len(testCompleted) {
		userTestResult = "PASS"
	} else {
		userTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(userTestResult, "C545232", "User workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.groupService.LogoutService()
}

// Delete User
func (suite *UserWorkflowSuite) deleteUser(userName string) error {
	userObj, err := suite.userService.GetUserByName(userName)
	require.NotNilf(suite.T(), userObj, "Failed to get user details")
	err = suite.userService.DeleteUser(*userObj.ID)
	return err
}

// Connect to array with different user
func (suite *UserWorkflowSuite) connectArray(arrayIP string, username string, password string) *service.NsGroupService {
	srcgroupService, err := service.NewNsGroupService(arrayIP, username, password, "v1", true)
	require.Nilf(suite.T(), err, "Could not connect to array")
	srcgroupService.SetDebug()
	return srcgroupService
}

// Create Power User
func (suite *UserWorkflowSuite) TestCreateModifyDeletePowerUser() {
	testStarted = append(testStarted, true)
	powerUserPassword := "testpowerauth"
	powerUserRole := nimbleos.NsUserRolesPoweruser
	powerUserFullName := "Test Power User"
	powerUserDescription := "User created by Ansible module for testing"
	powerUserEmailAddr := "testuser@test.com"
	var powerUserInactivityTimeout int64 = 1800
	newPowerUser := &nimbleos.User{
		Name:              param.NewString(powerUserName),
		Role:              powerUserRole,
		Password:          param.NewString(powerUserPassword),
		FullName:          param.NewString(powerUserFullName),
		Description:       param.NewString(powerUserDescription),
		EmailAddr:         param.NewString(powerUserEmailAddr),
		InactivityTimeout: param.NewInt64(powerUserInactivityTimeout),
	}
	_, err := suite.userService.CreateUser(newPowerUser)
	require.Nilf(suite.T(), err, "Power User creation failed")
	userResp, err := suite.userService.GetUserByName(powerUserName)
	require.Nilf(suite.T(), err, "Get user detail failed")
	require.Equal(suite.T(), nimbleos.NsUserRoles("poweruser"), *userResp.Role, "User created with in-correct Role")
	require.Equal(suite.T(), powerUserDescription, *userResp.Description, "Description not updated for Power user")
	require.Equal(suite.T(), powerUserEmailAddr, *userResp.EmailAddr, "EmailAddr not updated for Power user")
	require.Equal(suite.T(), powerUserFullName, *userResp.FullName, "FullName not updated for Power user")
	require.Equal(suite.T(), powerUserInactivityTimeout, *userResp.InactivityTimeout, "Inactive timeoue not updated for Power user")
	powerUserID := *userResp.ID

	// Create a volume using power user
	volumeName := "poweruser-testvol"
	powerusergroupService := suite.connectArray(*arrayIP, powerUserName, powerUserPassword)
	poweruserVolumeService := powerusergroupService.GetVolumeService()
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name:   param.NewString(volumeName),
		Size:   &sizeField,
		Online: param.NewBool(false),
	}
	_, err = poweruserVolumeService.CreateVolume(newVolume)
	require.Nilf(suite.T(), err, "Unable to create volume %v using power user", volumeName)

	// Delete volume
	volObj, _ := poweruserVolumeService.GetVolumeByName(volumeName)
	fmt.Printf("%v", volObj)
	err = poweruserVolumeService.DeleteVolume(*volObj.ID)
	require.Nilf(suite.T(), err, "Failed to delete volume")

	// Modify Power user role to Guest user
	guestUserRole := nimbleos.NsUserRolesGuest
	updateGuestUser := &nimbleos.User{
		Role: guestUserRole,
	}
	userResp, err = suite.userService.UpdateUser(powerUserID, updateGuestUser)
	require.Nilf(suite.T(), err, "Unable to modify user role")
	require.Equal(suite.T(), nimbleos.NsUserRoles("guest"), *userResp.Role, "User Role not updated")

	// Delete User
	err = suite.deleteUser(powerUserName)
	require.Nilf(suite.T(), err, "Unable to delete User %v", powerUserName)
	testCompleted = append(testCompleted, true)
}

func (suite *UserWorkflowSuite) TestCreateDeleteGuestUser() {
	testStarted = append(testStarted, true)
	guestUserName := "testguestuser"
	guestUserPassword := "testguestauth"
	guestUserRole := nimbleos.NsUserRolesGuest
	guestUserFullName := "Test Guest User"
	guestUserDescription := "User created by Ansible module for testing"
	guestUserEmailAddr := "testguest@test.com"
	var guestUserInactivityTimeout int64 = 1800
	newguestUser := &nimbleos.User{
		Name:              param.NewString(guestUserName),
		Role:              guestUserRole,
		Password:          param.NewString(guestUserPassword),
		FullName:          param.NewString(guestUserFullName),
		Description:       param.NewString(guestUserDescription),
		EmailAddr:         param.NewString(guestUserEmailAddr),
		InactivityTimeout: param.NewInt64(guestUserInactivityTimeout),
	}
	_, err := suite.userService.CreateUser(newguestUser)
	require.Nilf(suite.T(), err, "Guest User creation failed")
	userResp, err := suite.userService.GetUserByName(guestUserName)
	require.Nilf(suite.T(), err, "Get user detail failed")
	require.Equal(suite.T(), nimbleos.NsUserRoles("guest"), *userResp.Role, "In-correct User Role")
	require.Equal(suite.T(), guestUserDescription, *userResp.Description, "Description not updated for Guest user")
	require.Equal(suite.T(), guestUserEmailAddr, *userResp.EmailAddr, "EmailAddr not updated for Guest user")
	require.Equal(suite.T(), guestUserFullName, *userResp.FullName, "FullName not updated for Guest user")
	require.Equal(suite.T(), guestUserInactivityTimeout, *userResp.InactivityTimeout, "Inactivity timeout not updated for Guest user")

	// Delete User
	err = suite.deleteUser(guestUserName)
	require.Nilf(suite.T(), err, "Unable to Guest delete User %v", guestUserName)
	testCompleted = append(testCompleted, true)
}

func (suite *UserWorkflowSuite) TestUpdateUserPolicy() {
	testStarted = append(testStarted, true)
	arrayVersion := getArrayVersion(suite.arrayGroupService)
	if arrayVersion < 5.1 {
		suite.T().Skip()
	}
	var allowedAttempts int64 = 10
	var minLength int64 = 9
	var upperCaseCount int64 = 3
	var lowerCaseCount int64 = 9
	var specialCharCount int64 = 3
	var digitCount int64 = 4

	updateUserPolicy := &nimbleos.UserPolicy{
		AllowedAttempts: param.NewInt64(allowedAttempts),
		MinLength:       param.NewInt64(minLength),
		Upper:           param.NewInt64(upperCaseCount),
		Lower:           param.NewInt64(lowerCaseCount),
		Special:         param.NewInt64(specialCharCount),
		Digit:           param.NewInt64(digitCount),
	}
	filter := &param.GetParams{}
	userPolicyResp, err := suite.userPolicyService.GetUserPolicies(filter)
	userPolicyID := *userPolicyResp[0].ID
	_, err = suite.userPolicyService.UpdateUserPolicy(userPolicyID, updateUserPolicy)
	require.Nilf(suite.T(), err, "Failed to create userpolicy")
	userPolicyResp, err = suite.userPolicyService.GetUserPolicies(filter)
	require.Nilf(suite.T(), err, "Get user policy detail failed")
	require.Equal(suite.T(), allowedAttempts, *userPolicyResp[0].AllowedAttempts, "Wrong allowedAttempts count")
	require.Equal(suite.T(), minLength, *userPolicyResp[0].MinLength, "Wrong minLength count")
	require.Equal(suite.T(), upperCaseCount, *userPolicyResp[0].Upper, "Wrong upperCaseCount count")
	require.Equal(suite.T(), lowerCaseCount, *userPolicyResp[0].Lower, "Wrong lowerCaseCount count")
	require.Equal(suite.T(), digitCount, *userPolicyResp[0].Digit, "Wrong digitCount count")
	require.Equal(suite.T(), specialCharCount, *userPolicyResp[0].Special, "Wrong specialCharCount count")
	testCompleted = append(testCompleted, true)
}

func TestUserWorkflowSuite(t *testing.T) {
	suite.Run(t, new(UserWorkflowSuite))
}
