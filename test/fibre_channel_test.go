// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FiberChannelWorkflowSuite struct {
	suite.Suite
	groupService       *service.NsGroupService
	fcConfigService    *service.FibreChannelConfigService
	fcInterfaceService *service.FibreChannelInterfaceService
	arrayGroupService  *service.GroupService
}

func (suite *FiberChannelWorkflowSuite) SetupSuite() {
	groupService, err := config()
	require.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.fcConfigService = groupService.GetFibreChannelConfigService()
	suite.fcInterfaceService = groupService.GetFibreChannelInterfaceService()
	suite.arrayGroupService = groupService.GetGroupService()
	if !isFCEnabled(suite.arrayGroupService) {
		suite.T().SkipNow()
	}
}

func (suite *FiberChannelWorkflowSuite) TearDownSuite() {
	var fibreChannelTestResult string
	if len(testStarted) == len(testCompleted) {
		fibreChannelTestResult = "PASS"
	} else {
		fibreChannelTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(fibreChannelTestResult, "C545228", "Fibre Channel workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.groupService.LogoutService()
}

func (suite *FiberChannelWorkflowSuite) TestFiberChannelOfflineOnline() {
	testStarted = append(testStarted, true)
	filter := &param.GetParams{}
	interfaces, _ := suite.fcInterfaceService.GetFibreChannelInterfaces(filter)
	var interfaceID *string = nil
	for i := 0; i < len(interfaces); i++ {
		if *interfaces[i].Online == true {
			interfaceID = interfaces[i].ID
			break
		}
	}
	if interfaceID != nil {
		offlineInterface := &nimbleos.FibreChannelInterface{
			Online: param.NewBool(false),
		}
		_, err := suite.fcInterfaceService.UpdateFibreChannelInterface(*interfaceID, offlineInterface)
		require.Nilf(suite.T(), err, "Unable to set interface offline, err: %v", err)
		onlineInterface := &nimbleos.FibreChannelInterface{
			Online: param.NewBool(true),
		}
		_, err = suite.fcInterfaceService.UpdateFibreChannelInterface(*interfaces[0].ID, onlineInterface)
		require.Nilf(suite.T(), err, "Unable to set interface to online, err: %v", err)
	}
	testCompleted = append(testCompleted, true)
}

func (suite *FiberChannelWorkflowSuite) TestRegenerateOnlineInterfaces() {
	testStarted = append(testStarted, true)
	filter := &param.GetParams{}
	fcs, _ := suite.fcConfigService.GetFibreChannelConfigs(filter)
	interfaces, _ := suite.fcInterfaceService.GetFibreChannelInterfaces(filter)
	for i := 0; i < len(interfaces); i++ {
		if *interfaces[i].Online == true {
			_, err := suite.fcConfigService.RegenerateFibreChannelConfig(*fcs[0].ID, "af:32:f1", true)
			require.NotNil(suite.T(), err, "Regenerating Fiber channel config with online interfaces should have failed")
		}
	}
	testCompleted = append(testCompleted, true)
}

func TestFiberChannelWorkflowSuite(t *testing.T) {
	suite.Run(t, new(FiberChannelWorkflowSuite))
}
