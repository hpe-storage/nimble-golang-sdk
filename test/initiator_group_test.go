// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const initiatorGroupName = "InitiatorGroupTest"

type IGWorkflowSuite struct {
	suite.Suite
	groupService        *service.NsGroupService
	initiatorGrpService *service.InitiatorGroupService
	volumeService       sdkprovider.VolumeService
	arrayGroupService   *service.GroupService
}

func (suite *IGWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.initiatorGrpService = groupService.GetInitiatorGroupService()
	suite.volumeService = groupService.GetVolumeService()
	_, err = createDefaultVolume(suite.volumeService)
	assert.Nilf(suite.T(), err, "Unable to create default volume, err: %v", err)
	suite.arrayGroupService = groupService.GetGroupService()
}

func (suite *IGWorkflowSuite) TearDownSuite() {
	err := deleteDefaultVolume(suite.volumeService)
	assert.Nilf(suite.T(), err, "Unable to delete default volume, err: %v", err)
	suite.groupService.LogoutService()
}

func (suite *IGWorkflowSuite) deleteInitiatorGroup(igName string) {
	ig, _ := suite.initiatorGrpService.GetInitiatorGroupByName(igName)
	if ig != nil {
		err := suite.initiatorGrpService.DeleteInitiatorGroup(*ig.ID)
		assert.Nilf(suite.T(), err, "Unable to delete Initiator group, err: %v", err)
	}
}

func (suite *IGWorkflowSuite) TestCreateIGMissingParam() {
	newIG := &nimbleos.InitiatorGroup{
		Name: param.NewString(initiatorGroupName),
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	assert.NotNil(suite.T(), err, "Initiator group creation should have failed with error missing param")
}

func (suite *IGWorkflowSuite) TestCreateIG() {
	if !isIscsiEnabled(suite.arrayGroupService) {
		suite.T().Skip()
	}
	var desc *string = param.NewString("Workflow initiator group")
	newIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString(initiatorGroupName),
		Description:    desc,
		AccessProtocol: nimbleos.NsAccessProtocolIscsi,
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	assert.Nilf(suite.T(), err, "Unable to create initiator group, err: %v", err)
	ig, _ := suite.initiatorGrpService.GetInitiatorGroupByName(initiatorGroupName)
	assert.Equal(suite.T(), *desc, *ig.Description, "Initiator group creation does not have expected description.")
	var newDesc *string = param.NewString("Workflow tests for initiator group")
	updateIG := &nimbleos.InitiatorGroup{
		Description: newDesc,
	}
	_, err = suite.initiatorGrpService.UpdateInitiatorGroup(*ig.ID, updateIG)
	assert.Nilf(suite.T(), err, "Unable to update initiator group, err: %v", err)
	ig, _ = suite.initiatorGrpService.GetInitiatorGroupByName(initiatorGroupName)
	assert.Equal(suite.T(), *newDesc, *ig.Description, "Unable to update Initiator Group")

	// Create duplicate - should fail
	dupIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString(initiatorGroupName),
		Description:    param.NewString("Workflow initiator group"),
		AccessProtocol: nimbleos.NsAccessProtocolIscsi,
	}
	_, err = suite.initiatorGrpService.CreateInitiatorGroup(dupIG)
	assert.NotNil(suite.T(), err, "Initiator group creation should have failed")

	// Clean up
	suite.deleteInitiatorGroup(initiatorGroupName)
}

func (suite *IGWorkflowSuite) TestCreateIGIscsiInitiators() {
	if !isIscsiEnabled(suite.arrayGroupService) {
		suite.T().Skip()
	}
	iscsiInitiator := &nimbleos.NsISCSIInitiator{
		Label:     param.NewString("initiator1"),
		Iqn:       param.NewString("iqn.1994-05.com.redhat:48e040d53bf6"),
		IpAddress: param.NewString("10.0.0.0"),
	}
	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, iscsiInitiator)
	newIG := &nimbleos.InitiatorGroup{
		Name:            param.NewString("TestIGIscsi"),
		Description:     param.NewString("Workflow initiator group"),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	assert.Nilf(suite.T(), err, "Unable to create initiator group, err: %v", err)
	// Update initiator group
	ipAdd := "10.1.0.0"
	updateIscsiInitiator := &nimbleos.NsISCSIInitiator{
		IpAddress: &ipAdd,
		Label:     param.NewString("NewLabel"),
	}
	var updatedInitiatorList []*nimbleos.NsISCSIInitiator
	updatedInitiatorList = append(updatedInitiatorList, updateIscsiInitiator)
	updateIG := &nimbleos.InitiatorGroup{
		IscsiInitiators: updatedInitiatorList,
	}
	currentIG, _ := suite.initiatorGrpService.GetInitiatorGroupByName("TestIGIscsi")
	_, err = suite.initiatorGrpService.UpdateInitiatorGroup(*currentIG.ID, updateIG)
	assert.Nilf(suite.T(), err, "Modifying IP address of initiator failed: %v", err)
	assert.Equal(suite.T(), ipAdd, *updateIG.IscsiInitiators[0].IpAddress, "Updating iscsi initiators failed")
	// Clean up
	suite.deleteInitiatorGroup("TestIGIscsi")
}

func (suite *IGWorkflowSuite) TestCreateFCInitiators() {
	if !isFCEnabled(suite.arrayGroupService) {
		suite.T().Skip()
	}
	fcInitiator := &nimbleos.NsFCInitiator{
		Wwpn: param.NewString("af:32:f1:20:bc:ba:43:1a"),
	}
	var fcList []*nimbleos.NsFCInitiator
	fcList = append(fcList, fcInitiator)
	newIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString("TestIGFC"),
		Description:    param.NewString("Workflow initiator group"),
		AccessProtocol: nimbleos.NsAccessProtocolFc,
		FcInitiators:   fcList,
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	assert.Nilf(suite.T(), err, "Unable to create initiator group, err: %v", err)
	suite.deleteInitiatorGroup("TestIGFC")
}

func (suite *IGWorkflowSuite) TestCreateIGInvalidIscsi() {
	if !isIscsiEnabled(suite.arrayGroupService) {
		suite.T().Skip()
	}
	iscsiInitiator := &nimbleos.NsISCSIInitiator{
		Label:     param.NewString("initiator1"),
		Iqn:       param.NewString("a"),
		IpAddress: param.NewString("10.0.0.0"),
	}
	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, iscsiInitiator)
	newIG := &nimbleos.InitiatorGroup{
		Name:            param.NewString("TestIGIscsiInvalid"),
		Description:     param.NewString("Workflow initiator group"),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	assert.NotNil(suite.T(), err, "Initiator group creation with invalid iscsi initiators should have failed")
}

func TestInitiatorGroupSuite(t *testing.T) {
	suite.Run(t, new(IGWorkflowSuite))
}
