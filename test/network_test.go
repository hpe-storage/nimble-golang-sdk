// Copyright 2020 Hewlett Packard Enterprise Development LP

/*
Network Config test
--------------
  1. Get array's network info
  2. Delete 'Draft' network config ("role": "draft") if it is already present
  3. Create new network config
      Note: Use the name 'draft' when creating a draft configuration.
            Possible values: 'active', 'backup', 'draft'.
  4. Update draft config
  6. validate 'draft'
  7. delete draft
*/

package test

import (
	"fmt"
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type NetworkWorkflowSuite struct {
	suite.Suite
	groupService         *service.NsGroupService
	networkConfigService *service.NetworkConfigService
}

func (suite *NetworkWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Could not connect to array %v", arrayIP)
	suite.groupService = groupService
	suite.networkConfigService = groupService.GetNetworkConfigService()
}

func (suite *NetworkWorkflowSuite) TearDownSuite() {
	suite.groupService.LogoutService()
}

func (suite *NetworkWorkflowSuite) TestCreateUpdateDeleteDraft() {
	filter := &param.GetParams{}
	getNetworkResp, err := suite.networkConfigService.GetNetworkConfigs(filter)
	assert.Nil(suite.T(), err, "Failed to get Network Config Details")
	var draftConfigID string = ""
	var activeConfigID string = ""

	// Check if draft config is present, if present delete it
	for i, item := range getNetworkResp {
		if *item.Name == *nimbleos.NsNetConfigNameDraft {
			draftConfigID = *item.ID
			fmt.Printf("%v config is present in array", *getNetworkResp[i].Name)
		}
		if *item.Name == *nimbleos.NsNetConfigNameActive {
			activeConfigID = *item.ID
		}
	}
	if draftConfigID != "" {
		err := suite.networkConfigService.DeleteNetworkConfig(draftConfigID)
		assert.Nil(suite.T(), err, "Failed to Delete Draft Config")
	}

	// Create Draft config (Use Active Config to create Draft config)
	getActiceConfigResp, err := suite.networkConfigService.GetNetworkConfigById(activeConfigID)
	activeConfigSubnetList := getActiceConfigResp.SubnetList
	activeConfigRouteList := getActiceConfigResp.RouteList
	activeConfigArrayList := getActiceConfigResp.ArrayList
	var draftSubnetList []*nimbleos.NsSubnet
	draftSubnet := &nimbleos.NsSubnet{}
	for i, item := range activeConfigSubnetList {
		fmt.Printf("Adding subnet label: %v to subnet list", *activeConfigSubnetList[i].Label)
		draftSubnet = &nimbleos.NsSubnet{
			Label:              item.Label,
			Network:            item.Network,
			Netmask:            item.Netmask,
			NetzoneType:        item.NetzoneType,
			Type:               item.Type,
			AllowIscsi:         item.AllowIscsi,
			AllowGroup:         item.AllowGroup,
			DiscoveryIp:        item.DiscoveryIp,
			Mtu:                item.Mtu,
			VlanId:             item.VlanId,
			Failover:           item.Failover,
			FailoverEnableTime: item.FailoverEnableTime,
		}
		draftSubnetList = append(draftSubnetList, draftSubnet)
	}

	draftRoute := &nimbleos.NsRoute{
		TgtNetwork: activeConfigRouteList[0].TgtNetwork,
		TgtNetmask: activeConfigRouteList[0].TgtNetmask,
		Gateway:    activeConfigRouteList[0].Gateway,
	}
	var draftRouteList []*nimbleos.NsRoute
	draftRouteList = append(draftRouteList, draftRoute)

	// Create NIC List and add to ArrayList
	var draftNICList []*nimbleos.NsNIC
	draftNIC := &nimbleos.NsNIC{}
	for index, item := range activeConfigArrayList[0].NicList {
		fmt.Printf("Adding NIC: %v to NIC List", *activeConfigArrayList[0].NicList[index].Name)
		draftNIC = &nimbleos.NsNIC{
			Name:        item.Name,
			SubnetLabel: item.SubnetLabel,
			DataIp:      item.DataIp,
			Tagged:      item.Tagged,
		}
		draftNICList = append(draftNICList, draftNIC)
	}
	draftArray := &nimbleos.NsArrayNet{
		Name:            activeConfigArrayList[0].Name,
		CtrlrASupportIp: activeConfigArrayList[0].CtrlrASupportIp,
		CtrlrBSupportIp: activeConfigArrayList[0].CtrlrBSupportIp,
		NicList:         draftNICList,
	}
	var draftArrayList []*nimbleos.NsArrayNet
	draftArrayList = append(draftArrayList, draftArray)

	draftNetworkConfig := &nimbleos.NetworkConfig{
		RouteList:                      draftRouteList,
		SubnetList:                     draftSubnetList,
		ArrayList:                      draftArrayList,
		Name:                           nimbleos.NsNetConfigNameDraft,
		MgmtIp:                         getActiceConfigResp.MgmtIp,
		IscsiAutomaticConnectionMethod: param.NewBool(false),
		IscsiConnectionRebalancing:     param.NewBool(false),
	}

	// Create Draft Network Config
	draftConfigResp, err := suite.networkConfigService.CreateNetworkConfig(draftNetworkConfig)
	assert.Nil(suite.T(), err, "Failed to Create Draft Network Config")
	draftConfigID = *draftConfigResp.ID

	// Update Draft Network Config - Array API expects all the mandatory params to be passed for update
	updateDraftNetworkConfig := &nimbleos.NetworkConfig{
		RouteList:                      draftRouteList,
		SubnetList:                     draftSubnetList,
		ArrayList:                      draftArrayList,
		Name:                           nimbleos.NsNetConfigNameDraft,
		MgmtIp:                         getActiceConfigResp.MgmtIp,
		IscsiAutomaticConnectionMethod: param.NewBool(true),
		IscsiConnectionRebalancing:     param.NewBool(true),
	}

	draftConfigResp, err = suite.networkConfigService.UpdateNetworkConfig(draftConfigID, updateDraftNetworkConfig)
	assert.Nil(suite.T(), err, "Failed to update drfat network config")
	assert.Equal(suite.T(), *draftConfigResp.IscsiConnectionRebalancing, true, "Failed to update IscsiConnectionRebalancing value in draft network config")
	assert.Equal(suite.T(), *draftConfigResp.IscsiAutomaticConnectionMethod, true, "Failed to update IscsiAutomaticConnectionMethod value in draft network config")

	// Validate Draft Config - Expected to fail as we did not add draft with accurate data
	var ignoreValidationMask uint64 = 0
	err = suite.networkConfigService.ValidateNetconfigNetworkConfig(draftConfigID, ignoreValidationMask)
	assert.Nil(suite.T(), err, "Validation failed for Draft Config")

	// Delete Draft Config
	err = suite.networkConfigService.DeleteNetworkConfig(draftConfigID)
	assert.Nil(suite.T(), err, "Failed to Delete Draft Config")

}

func TestNetworkWorkflowSuite(t *testing.T) {
	suite.Run(t, new(NetworkWorkflowSuite))
}
