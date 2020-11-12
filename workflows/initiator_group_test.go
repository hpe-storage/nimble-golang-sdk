package workflow

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/suite"
	"testing"
)

var initiatorGroupName = "InitiatorGroupTest"

type IGWorkflowSuite struct {
	suite.Suite
	groupService        *service.NsGroupService
	initiatorGrpService *service.InitiatorGroupService
	volumeService       *service.VolumeService
}

func (suite *IGWorkflowSuite) SetupTest() {
	groupService, err := config()
	if err != nil {
		suite.T().Errorf("Unable to connect to group: %v", err.Error())
	}
	suite.groupService = groupService
	suite.initiatorGrpService = groupService.GetInitiatorGroupService()
	suite.volumeService = groupService.GetVolumeService()
	createDefaultVolume(suite.volumeService)
}

func (suite *IGWorkflowSuite) TearDownSuite() {
	suite.deleteInitiatorGroup("TestIGIscsi")
	suite.deleteInitiatorGroup(initiatorGroupName)
	deleteDefaultVolume(suite.volumeService)
}

func (suite *IGWorkflowSuite) deleteInitiatorGroup(igName string) {
	ig, _ := suite.initiatorGrpService.GetInitiatorGroupByName(igName)
	if ig != nil {
		err := suite.initiatorGrpService.DeleteInitiatorGroup(*ig.ID)
		if err != nil {
			suite.T().Errorf("Issue deleting Initiator group: %v", igName)
		}
	}
}

func (suite *IGWorkflowSuite) TestCreateIGMissingParam() {
	newIG := &nimbleos.InitiatorGroup{
		Name: param.NewString(initiatorGroupName),
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	if err == nil {
		suite.T().Error("Initiator group creation should have failed with error missing param")
	}
}

func (suite *IGWorkflowSuite) TestCreateIG() {
	var desc *string = param.NewString("Workflow initiator group")
	newIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString(initiatorGroupName),
		Description:    desc,
		AccessProtocol: nimbleos.NsAccessProtocolIscsi,
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	if err != nil {
		suite.T().Errorf("Initiator group creation failed with error: %v", err.Error())
	}
	ig, _ := suite.initiatorGrpService.GetInitiatorGroupByName(initiatorGroupName)
	if *ig.Description != *desc {
		suite.T().Errorf("Initiator group creation failed to have correct description: %v", desc)
	}
}

func (suite *IGWorkflowSuite) TestCreateIGDuplicate() {
	dupIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString(initiatorGroupName),
		Description:    param.NewString("Workflow initiator group"),
		AccessProtocol: nimbleos.NsAccessProtocolIscsi,
	}
	_, err := suite.initiatorGrpService.CreateInitiatorGroup(dupIG)
	if err == nil {
		suite.T().Error("Initiator group creation should have failed")
	}
}

func (suite *IGWorkflowSuite) TestCreateIGIscsiInitiators() {
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
	if err != nil {
		suite.T().Errorf("Initiator group creation failed: %v", err.Error())
	}
}

func (suite *IGWorkflowSuite) TestCreateIGInvalidIscsi() {
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
	if err == nil {
		suite.T().Errorf("Initiator group cretaion with invalid iscsi initiators should have failed")
	}
}

func (suite *IGWorkflowSuite) TestModifyIGDescription() {
	var desc *string = param.NewString("Workflow tests for initiator group")
	updateIG := &nimbleos.InitiatorGroup{
		Description: desc,
	}
	ig, _ := suite.initiatorGrpService.GetInitiatorGroupByName(initiatorGroupName)
	if ig != nil {
		suite.initiatorGrpService.UpdateInitiatorGroup(*ig.ID, updateIG)
	}
	ig, _ = suite.initiatorGrpService.GetInitiatorGroupByName(initiatorGroupName)
	if *ig.Description != *desc {
		suite.T().Errorf("Initiator group description did not update: %v", ig.Description)
	}
}

func (suite *IGWorkflowSuite) TestModifyIGInitiators() {
	var ipAdd *string = param.NewString("10.1.0.0")
	iscsiInitiator := &nimbleos.NsISCSIInitiator{
		IpAddress: ipAdd,
		Label:     param.NewString("NewLabel"),
	}
	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, iscsiInitiator)
	updateIG := &nimbleos.InitiatorGroup{
		IscsiInitiators: initiatorList,
	}
	currentIG, _ := suite.initiatorGrpService.GetInitiatorGroupByName("TestIGIscsi")
	updatedIG, err := suite.initiatorGrpService.UpdateInitiatorGroup(*currentIG.ID, updateIG)
	if err != nil {
		suite.T().Errorf("Modifying IP address of initiator failed: %v", err.Error())
	} else if *updatedIG.IscsiInitiators[0].IpAddress != *ipAdd {
		suite.T().Errorf("Updating iscsi initiatos failed: %v", updatedIG.IscsiInitiators[0].IpAddress)
	}
}

func TestInitiatorGroupSuite(t *testing.T) {
	suite.Run(t, new(IGWorkflowSuite))
}
