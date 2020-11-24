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

const protectionTemplateName = "ProtectionTemplateTest"
const protectionScheduleName = "ProtectionScheduleTest"

type ProtectionTemplateWorkflowSuite struct {
	suite.Suite
	groupService              *service.NsGroupService
	protectionTemplateService *service.ProtectionTemplateService
	protectionScheduleService *service.ProtectionScheduleService
}

func (suite *ProtectionTemplateWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Unable to connect to group: %v", err)
	suite.protectionTemplateService = groupService.GetProtectionTemplateService()
	suite.protectionScheduleService = groupService.GetProtectionScheduleService()
}

func (suite *ProtectionTemplateWorkflowSuite) TearDownSuite() {
	suite.deleteProtectionTemplate(protectionTemplateName)
	suite.deleteProtectionSchedule(protectionScheduleName)
}

func (suite *ProtectionTemplateWorkflowSuite) deleteProtectionTemplate(ptName string) {
	pt, _ := suite.protectionTemplateService.GetProtectionTemplateByName(ptName)
	if pt != nil {
		err := suite.protectionTemplateService.DeleteProtectionTemplate(*pt.ID)
		assert.Nilf(suite.T(), err, "Unable to delete Protection Template, err: %v", err)
	}
}

func (suite *ProtectionTemplateWorkflowSuite) deleteProtectionSchedule(psName string) {
	ps, _ := suite.protectionScheduleService.GetProtectionScheduleByName(psName)
	if ps != nil {
		err := suite.protectionScheduleService.DeleteProtectionSchedule(*ps.ID)
		assert.Nilf(suite.T(), err, "Unable to delete Protection Schedule, err: %v", err)
	}
}

func (suite *ProtectionTemplateWorkflowSuite) TestCreateProtectionTemplate() {
	newPT := &nimbleos.ProtectionTemplate{
		Name: param.NewString(protectionTemplateName),
	}
	pt, err := suite.protectionTemplateService.CreateProtectionTemplate(newPT)
	assert.Nilf(suite.T(), err, "Unable to create Protection template, err: %v", err)
	var numRetain *int64 = param.NewInt64(1234)
	newPS := &nimbleos.ProtectionSchedule{
		Name:                  param.NewString(protectionScheduleName),
		VolcollOrProttmplType: nimbleos.NsProtectionPolicyTypeProtectionTemplate,
		VolcollOrProttmplId:   pt.ID,
		NumRetain:             numRetain,
	}
	_, err = suite.protectionScheduleService.CreateProtectionSchedule(newPS)
	assert.Nilf(suite.T(), err, "Unable to create protection schedule, err: %v", err)
	// Update protecion template app sync
	updatePT := &nimbleos.ProtectionTemplate{
		AppSync:   nimbleos.NsAppSyncTypeVss,
		AppServer: param.NewString("10.1.1.1"),
		AppId:     nimbleos.NsAppIdTypeExchangeDag,
	}
	updatedPT, err := suite.protectionTemplateService.UpdateProtectionTemplate(*pt.ID, updatePT)
	assert.Nilf(suite.T(), err, "Unable to update app syn of protection template, err: %v", err)
	assert.Equal(suite.T(), *nimbleos.NsAppSyncTypeVss, *updatedPT.AppSync, "Update failed")
}

func (suite *ProtectionTemplateWorkflowSuite) TestCreateProtectionTemplateDup() {
	newPT := &nimbleos.ProtectionTemplate{
		Name: param.NewString(protectionTemplateName),
	}
	_, err := suite.protectionTemplateService.CreateProtectionTemplate(newPT)
	assert.NotNil(suite.T(), err, "Creating duplicate protection template should have failed")
}

func TestProtectionTemplateWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ProtectionTemplateWorkflowSuite))
}
