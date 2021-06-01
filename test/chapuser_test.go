// Copyright 2020 Hewlett Packard Enterprise Development LP

/*
Chap User tests
---------------
  1. Create a Chap user
  2. Update Chap user name, description
  3. Create Initiator group & Update Chap user iqns
  4. Delete Chap user when iqn is associated
      It should fail, validate error message
  5. Remove iqns from chap user & delete chap user
  6. Delete initiator group
*/

package test

import (
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const chapUserName string = "testchapuser"
const chapUserPassword string = "password_25-24"
const initiatorGroupNameChap string = "TestChapIGIscsi"

type ChapuserWorkflowSuite struct {
	suite.Suite
	groupService        *service.NsGroupService
	userService         *service.UserService
	chapuserService     *service.ChapUserService
	initiatorGrpService *service.InitiatorGroupService
	arrayGroupService   *service.GroupService
}

func (suite *ChapuserWorkflowSuite) SetupSuite() {
	groupService, err := config()
	require.Nilf(suite.T(), err, "Could not connect to array %v", arrayIP)
	suite.groupService = groupService
	suite.chapuserService = groupService.GetChapUserService()
	suite.initiatorGrpService = groupService.GetInitiatorGroupService()
	suite.arrayGroupService = groupService.GetGroupService()
}

func (suite *ChapuserWorkflowSuite) TearDownSuite() {
	var chapTestResult string
	if len(testStarted) == len(testCompleted) {
		chapTestResult = "PASS"
	} else {
		chapTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(chapTestResult, "C545226", "Chapuser workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.groupService.LogoutService()
}

func (suite *ChapuserWorkflowSuite) CreateIGIscsiInitiators(iqnLabel string, iqn string, ipAddress string) string {
	iscsiInitiator := &nimbleos.NsISCSIInitiator{
		Label:     &iqnLabel,
		Iqn:       &iqn,
		IpAddress: &ipAddress,
	}
	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, iscsiInitiator)
	newIG := &nimbleos.InitiatorGroup{
		Name:            param.NewString(initiatorGroupNameChap),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}
	igResp, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	require.Nilf(suite.T(), err, "Initiator group creation failed: %v", initiatorGroupNameChap)
	return *igResp.ID
}

func (suite *ChapuserWorkflowSuite) TestCreateModifyDeleteChapuser() {
	testStarted = append(testStarted, true)
	// Skipping for FC array, test with iscsi initiators
	if isFCEnabled(suite.arrayGroupService) {
		suite.T().Skip()
	}
	chapUserDescription := "Chap User created by SDK module for testing"
	// Create chapuser without initiator
	newChapUser := &nimbleos.ChapUser{
		Name:        param.NewString(chapUserName),
		Password:    param.NewString(chapUserPassword),
		Description: &chapUserDescription,
	}
	_, err := suite.chapuserService.CreateChapUser(newChapUser)
	require.Nilf(suite.T(), err, "Chap User %v creation failed", chapUserName)
	chapuserResp, err := suite.chapuserService.GetChapUserByName(chapUserName)
	require.Nilf(suite.T(), err, "Failed to get chap user by name %v", chapUserName)
	require.Equal(suite.T(), chapUserDescription, *chapuserResp.Description, "In-correct Description for Chap user")
	chapuserID := *chapuserResp.ID

	arrayVersion := getArrayVersion(suite.arrayGroupService)
	if arrayVersion > 5.1 {
		// Update chapuser with initiators
		iqnLabel := "chapiqn1"
		iqn := "iqn.1994-05.com.redhat:48e040d53bf6"
		ipAddress := "22.2.2.2"
		igID := suite.CreateIGIscsiInitiators(iqnLabel, iqn, ipAddress)
		iscsiInitiator := &nimbleos.NsISCSIIQN{
			Name: &iqn,
		}
		var initiatorList []*nimbleos.NsISCSIIQN
		initiatorList = append(initiatorList, iscsiInitiator)
		updateChapUser := &nimbleos.ChapUser{
			InitiatorIqns: initiatorList,
		}
		_, err = suite.chapuserService.UpdateChapUser(chapuserID, updateChapUser)
		require.Nilf(suite.T(), err, "Failed to update chap user: %v", chapUserName)
		_, err = suite.chapuserService.GetChapUserByName(chapUserName)
		require.Nilf(suite.T(), err, "Failed to get chap user by name: %v", chapUserName)

		// Delete Chap user with initiators associated
		err = suite.chapuserService.DeleteChapUser(chapuserID)
		require.NotNilf(suite.T(), err, "Deleted chap user %v with initiators", chapUserName)

		// Delete Initiatorgroup
		err = suite.initiatorGrpService.DeleteInitiatorGroup(igID)
		require.Nilf(suite.T(), err, "Failed to delete initiator group: %v", initiatorGroupNameChap)
	}

	// Delete Chap user
	err = suite.chapuserService.DeleteChapUser(chapuserID)
	require.Nilf(suite.T(), err, "Failed to delete chap user: %v", chapUserName)
	testCompleted = append(testCompleted, true)

}

func TestChapuserWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ChapuserWorkflowSuite))
}
