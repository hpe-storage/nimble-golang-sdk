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
	require.Nilf(suite.T(), err, "Unable to connect to group: %v", err)
	suite.groupService = groupService
	suite.protectionTemplateService = groupService.GetProtectionTemplateService()
	suite.protectionScheduleService = groupService.GetProtectionScheduleService()
	suite.createProtectionTemplate(protectionTemplateName)
}

func (suite *ProtectionTemplateWorkflowSuite) TearDownSuite() {
	var protectionTemplateTestResult string
	if len(testStarted) == len(testCompleted) {
		protectionTemplateTestResult = "PASS"
	} else {
		protectionTemplateTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(protectionTemplateTestResult, "C545077", "Protection Template workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.deleteProtectionTemplate(protectionTemplateName)
	suite.deleteProtectionSchedule(protectionScheduleName)
	suite.groupService.LogoutService()
}

func (suite *ProtectionTemplateWorkflowSuite) createProtectionTemplate(ptName string) {
	newPT := &nimbleos.ProtectionTemplate{
		Name: param.NewString(ptName),
	}
	_, err := suite.protectionTemplateService.CreateProtectionTemplate(newPT)
	require.Nilf(suite.T(), err, "Unable to create Protection template, err: %v", err)
}

func (suite *ProtectionTemplateWorkflowSuite) deleteProtectionTemplate(ptName string) {
	pt, _ := suite.protectionTemplateService.GetProtectionTemplateByName(ptName)
	if pt != nil {
		err := suite.protectionTemplateService.DeleteProtectionTemplate(*pt.ID)
		require.Nilf(suite.T(), err, "Unable to delete Protection Template, err: %v", err)
	}
}

func (suite *ProtectionTemplateWorkflowSuite) deleteProtectionSchedule(psName string) {
	ps, _ := suite.protectionScheduleService.GetProtectionScheduleByName(psName)
	if ps != nil {
		err := suite.protectionScheduleService.DeleteProtectionSchedule(*ps.ID)
		require.Nilf(suite.T(), err, "Unable to delete Protection Schedule, err: %v", err)
	}
}

func (suite *ProtectionTemplateWorkflowSuite) TestCreateProtectionTemplate() {
	testStarted = append(testStarted, true)
	pt, err := suite.protectionTemplateService.GetProtectionTemplateByName(protectionTemplateName)
	require.Nilf(suite.T(), err, "Unable to get protection template, err: %v", err)
	var numRetain *int64 = param.NewInt64(1234)
	newPS := &nimbleos.ProtectionSchedule{
		Name:                  param.NewString(protectionScheduleName),
		VolcollOrProttmplType: nimbleos.NsProtectionPolicyTypeProtectionTemplate,
		VolcollOrProttmplId:   pt.ID,
		NumRetain:             numRetain,
	}
	_, err = suite.protectionScheduleService.CreateProtectionSchedule(newPS)
	require.Nilf(suite.T(), err, "Unable to create protection schedule, err: %v", err)
	// Update protecion template app sync
	updatePT := &nimbleos.ProtectionTemplate{
		AppSync:   nimbleos.NsAppSyncTypeVss,
		AppServer: param.NewString("10.1.1.1"),
		AppId:     nimbleos.NsAppIdTypeExchangeDag,
	}
	updatedPT, err := suite.protectionTemplateService.UpdateProtectionTemplate(*pt.ID, updatePT)
	require.Nilf(suite.T(), err, "Unable to update app syn of protection template, err: %v", err)
	require.Equal(suite.T(), *nimbleos.NsAppSyncTypeVss, *updatedPT.AppSync, "Update failed")
	testCompleted = append(testCompleted, true)
}

func (suite *ProtectionTemplateWorkflowSuite) TestCreateProtectionTemplateDup() {
	testStarted = append(testStarted, true)
	newPT := &nimbleos.ProtectionTemplate{
		Name: param.NewString(protectionTemplateName),
	}
	_, err := suite.protectionTemplateService.CreateProtectionTemplate(newPT)
	require.NotNil(suite.T(), err, "Creating duplicate protection template should have failed")
	testCompleted = append(testCompleted, true)
}

func TestProtectionTemplateWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ProtectionTemplateWorkflowSuite))
}
