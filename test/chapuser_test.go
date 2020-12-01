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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const chapUserName = "testchapuser"

type ChapuserWorkflowSuite struct {
	suite.Suite
	groupService        *service.NsGroupService
	userService         *service.UserService
	chapuserService     *service.ChapUserService
	initiatorGrpService *service.InitiatorGroupService
}

func (suite *ChapuserWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Could not connect to array")
	suite.groupService = groupService
	suite.chapuserService = groupService.GetChapUserService()
	suite.initiatorGrpService = groupService.GetInitiatorGroupService()
}

func (suite *ChapuserWorkflowSuite) CreateIGIscsiInitiators(iqnLable string, iqn string, ipAddress string) string {
	iscsiInitiator := &nimbleos.NsISCSIInitiator{
		Label:     param.NewString(iqnLable),
		Iqn:       param.NewString(iqn),
		IpAddress: param.NewString(ipAddress),
	}
	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, iscsiInitiator)
	newIG := &nimbleos.InitiatorGroup{
		Name:            param.NewString("TestChapIGIscsi"),
		Description:     param.NewString("Workflow initiator group"),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}
	igResp, err := suite.initiatorGrpService.CreateInitiatorGroup(newIG)
	assert.Nilf(suite.T(), err, "Unable to create initiator group, err: %v", err)
	return *igResp.ID
}

func (suite *ChapuserWorkflowSuite) TestCreateModifyDeleteChapuser() {
	chapUserPassword := "password_25-24"
	chapUserescription := "Chap User created by SDK module for testing"
	// Create chapuser without initiator
	newChapUser := &nimbleos.ChapUser{
		Name:        param.NewString(chapUserName),
		Password:    param.NewString(chapUserPassword),
		Description: param.NewString(chapUserescription),
	}
	_, err := suite.chapuserService.CreateChapUser(newChapUser)
	assert.Nilf(suite.T(), err, "Chap User creation failed")
	chapuserResp, err := suite.chapuserService.GetChapUserByName(chapUserName)
	assert.Nilf(suite.T(), err, "Failed to get chap user by name")
	assert.Equal(suite.T(), chapUserescription, *chapuserResp.Description, "In-correct Description for Chap user")
	chapuserID := *chapuserResp.ID

	// Update chapuser with initiators
	iqnLable := "chapiqn1"
	iqn := "iqn.1994-05.com.redhat:48e040d53bf6"
	ipAddress := "22.2.2.2"
	igID := suite.CreateIGIscsiInitiators(iqnLable, iqn, ipAddress)
	iscsiInitiator := &nimbleos.NsISCSIIQN{
		Name: param.NewString(iqn),
	}
	var initiatorList []*nimbleos.NsISCSIIQN
	initiatorList = append(initiatorList, iscsiInitiator)
	updateChapUser := &nimbleos.ChapUser{
		InitiatorIqns: initiatorList,
	}
	_, err = suite.chapuserService.UpdateChapUser(chapuserID, updateChapUser)
	assert.Nilf(suite.T(), err, "Chap User creation failed")
	_, err = suite.chapuserService.GetChapUserByName(chapUserName)
	assert.Nilf(suite.T(), err, "Failed to get chap user by name")

	// Delete Chap user with initiators associated
	err = suite.chapuserService.DeleteChapUser(chapuserID)
	assert.NotNilf(suite.T(), err, "Failed to delete chap user")

	// Delete Initiatorgroup
	err = suite.initiatorGrpService.DeleteInitiatorGroup(igID)
	assert.Nilf(suite.T(), err, "Failed to delete initiator group")

	// Delete Chap user
	err = suite.chapuserService.DeleteChapUser(chapuserID)
	assert.Nilf(suite.T(), err, "Failed to delete chap user")

}

func TestChapuserWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ChapuserWorkflowSuite))
}
