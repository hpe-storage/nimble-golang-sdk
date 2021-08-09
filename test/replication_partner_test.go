// Copyright 2020 Hewlett Packard Enterprise Development LP
package test

import (
	"testing"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const repPartner = "ReplcationPartnerTest"
const repPassphrase = "Passphrase123"

type ReplicationPartnerWorkflowSuite struct {
	suite.Suite
	groupService                *service.NsGroupService
	downstreamGroupService      *service.NsGroupService
	upstreamRepPartnerService   *service.ReplicationPartnerService
	downstreamRepPartnerService *service.ReplicationPartnerService
	upstreamArrayGrpService     *service.GroupService
	downstreamArrayGrpService   *service.GroupService
	upstreamGroupName           string
	downstreamGroupName         string
	volumeService               sdkprovider.VolumeService
	volcollService              *service.VolumeCollectionService
	protectionScheduleService   *service.ProtectionScheduleService
}

func (suite *ReplicationPartnerWorkflowSuite) SetupSuite() {
	groupService, err := config()
	require.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.volumeService = groupService.GetVolumeService()
	suite.upstreamRepPartnerService = groupService.GetReplicationPartnerService()
	suite.upstreamArrayGrpService = groupService.GetGroupService()
	filter := &param.GetParams{}
	groupResp, _ := suite.upstreamArrayGrpService.GetGroups(filter)
	suite.upstreamGroupName = *groupResp[0].Name
	// Skip test run if no downstream array is passed
	if *downstreamArrayIP == "" {
		suite.T().SkipNow()
	}
	// Downstream Array
	groupService2, err := service.NewNsGroupService(*downstreamArrayIP, "admin", "admin", "v1", true)
	require.Nilf(suite.T(), err, "Unable to connect to the downstream partner, err: %v", err)
	suite.downstreamGroupService = groupService2
	suite.downstreamRepPartnerService = groupService2.GetReplicationPartnerService()
	suite.downstreamArrayGrpService = groupService2.GetGroupService()
	groupResp, _ = suite.downstreamArrayGrpService.GetGroups(filter)
	suite.downstreamGroupName = *groupResp[0].Name
	suite.createReplicationPartner()
}

func (suite *ReplicationPartnerWorkflowSuite) TearDownSuite() {
	var replicationPartnerTestResult string
	if len(testStarted) == len(testCompleted) {
		replicationPartnerTestResult = "PASS"
	} else {
		replicationPartnerTestResult = "FAIL"
	}
	var postResult = *postResultToDashboard
	if postResult == "true" {
		pushResultToDashboard(replicationPartnerTestResult, "C545231", "Replication Partner workflow")
	}
	//cleanup test result
	testStarted = nil
	testCompleted = nil
	suite.deleteReplicationPartner()
	suite.groupService.LogoutService()
	suite.downstreamGroupService.LogoutService()
}

func (suite *ReplicationPartnerWorkflowSuite) createReplicationPartner() {
	// Upstream
	upstream := &nimbleos.ReplicationPartner{
		Name:        param.NewString(suite.downstreamGroupName),
		Secret:      param.NewString(repPassphrase),
		Hostname:    downstreamArrayIP,
		SubnetLabel: param.NewString("mgmt-data"),
	}
	upstreamRP, err := suite.upstreamRepPartnerService.CreateReplicationPartner(upstream)
	require.Nilf(suite.T(), err, "Unable to set replication partner, err: %v", err)
	// Downstream
	downstream := &nimbleos.ReplicationPartner{
		Name:        param.NewString(suite.upstreamGroupName),
		Secret:      param.NewString(repPassphrase),
		Hostname:    arrayIP,
		SubnetLabel: param.NewString("mgmt-data"),
	}
	downstreamRP, err := suite.downstreamRepPartnerService.CreateReplicationPartner(downstream)
	require.Nilf(suite.T(), err, "Unable to set replication partner, err: %v", err)

	// Test replication partner
	err = suite.upstreamRepPartnerService.TestReplicationPartner(*upstreamRP.ID)
	require.Nilf(suite.T(), err, "Unable to test replication partners, err: %v", err)
	err = suite.downstreamRepPartnerService.TestReplicationPartner(*downstreamRP.ID)
	require.Nilf(suite.T(), err, "Unable to test replication partners, err: %v", err)
}

func (suite *ReplicationPartnerWorkflowSuite) deleteReplicationPartner() {
	upstreamRP, err := suite.upstreamRepPartnerService.GetReplicationPartnerByName(suite.downstreamGroupName)
	require.Nilf(suite.T(), err, "Unable to find the replication partner, err: %v", err)
	err = suite.upstreamRepPartnerService.DeleteReplicationPartner(*upstreamRP.ID)
	require.Nilf(suite.T(), err, "Unable to delete upstream replication partner, err: %v", err)
	downstreamRP, err := suite.downstreamRepPartnerService.GetReplicationPartnerByName(suite.upstreamGroupName)
	require.Nilf(suite.T(), err, "Unable to find the replication partner, err: %v", err)
	err = suite.downstreamRepPartnerService.DeleteReplicationPartner(*downstreamRP.ID)
	require.Nilf(suite.T(), err, "Unable to delete downstram replication partner, err: %v", err)
}

func (suite *ReplicationPartnerWorkflowSuite) TestPauseResumeRP() {
	testStarted = append(testStarted, true)
	upstreamRP, err := suite.upstreamRepPartnerService.GetReplicationPartnerByName(suite.downstreamGroupName)
	require.Nilf(suite.T(), err, "Unable to find replication partner, err: %v", err)

	// Pause Replication Partner
	err = suite.upstreamRepPartnerService.PauseReplicationPartner(*upstreamRP.ID)
	require.Nilf(suite.T(), err, "Unable to pause replication parner, err: %v", err)

	// Resume replication partner
	err = suite.upstreamRepPartnerService.ResumeReplicationPartner(*upstreamRP.ID)
	require.Nilf(suite.T(), err, "Unable to resume replication partner")
	testCompleted = append(testCompleted, true)
}

func (suite *ReplicationPartnerWorkflowSuite) TestUpdateRP() {
	testStarted = append(testStarted, true)
	upstreamRP, err := suite.upstreamRepPartnerService.GetReplicationPartnerByName(suite.downstreamGroupName)
	require.Nilf(suite.T(), err, "Unable to find replication partner, err: %v", err)
	updateRP := &nimbleos.ReplicationPartner{
		Description: param.NewString("Testing replication partner"),
	}
	_, err = suite.upstreamRepPartnerService.UpdateReplicationPartner(*upstreamRP.ID, updateRP)
	require.Nilf(suite.T(), err, "Unable to update replication partner, err: %v", err)
	testCompleted = append(testCompleted, true)
}

func TestReplicationPartnerWorkflowSuite(t *testing.T) {
	suite.Run(t, new(ReplicationPartnerWorkflowSuite))
}
