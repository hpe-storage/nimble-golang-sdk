// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.fcConfigService = groupService.GetFibreChannelConfigService()
	suite.fcInterfaceService = groupService.GetFibreChannelInterfaceService()
	suite.arrayGroupService = groupService.GetGroupService()
	if !isFCEnabled(suite.arrayGroupService) {
		suite.T().SkipNow()
	}
}

func (suite *FiberChannelWorkflowSuite) TearDownSuite() {
	suite.groupService.LogoutService()
}

func (suite *FiberChannelWorkflowSuite) TestFiberChannelOfflineOnline() {
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
		assert.Nilf(suite.T(), err, "Unable to set interface offline, err: %v", err)
		onlineInterface := &nimbleos.FibreChannelInterface{
			Online: param.NewBool(true),
		}
		_, err = suite.fcInterfaceService.UpdateFibreChannelInterface(*interfaces[0].ID, onlineInterface)
		assert.Nilf(suite.T(), err, "Unable to set interface to online, err: %v", err)
	}
}

func (suite *FiberChannelWorkflowSuite) TestRegenerateOnlineInterfaces() {
	filter := &param.GetParams{}
	fcs, _ := suite.fcConfigService.GetFibreChannelConfigs(filter)
	interfaces, _ := suite.fcInterfaceService.GetFibreChannelInterfaces(filter)
	for i := 0; i < len(interfaces); i++ {
		if *interfaces[i].Online == true {
			_, err := suite.fcConfigService.RegenerateFibreChannelConfig(*fcs[0].ID, "af:32:f1", true)
			assert.NotNil(suite.T(), err, "Regenerating Fiber channel config with online interfaces should have failed")
		}
	}
}

func TestFiberChannelWorkflowSuite(t *testing.T) {
	suite.Run(t, new(FiberChannelWorkflowSuite))
}
