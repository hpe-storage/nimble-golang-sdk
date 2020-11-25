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

const snapcollName = "SnapshotCollectionTest"

type SnapCollWorkflowSuite struct {
	suite.Suite
	groupService              *service.NsGroupService
	snapcollService           *service.SnapshotCollectionService
	volumeService             *service.VolumeService
	volcollService            *service.VolumeCollectionService
	protectionScheduleService *service.ProtectionScheduleService
}

func (suite *SnapCollWorkflowSuite) SetupTest() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.snapcollService = groupService.GetSnapshotCollectionService()
	suite.volumeService = groupService.GetVolumeService()
	suite.volcollService = groupService.GetVolumeCollectionService()
	suite.protectionScheduleService = groupService.GetProtectionScheduleService()
	_, err = createDefaultVolume(suite.volumeService)
	assert.Nilf(suite.T(), err, "Unable to create default volume, err: %v", err)
	_, err = createDefaultVolColl(suite.volcollService)
	assert.Nilf(suite.T(), err, "Unable to create default volume collection, err: %v", err)
}

func (suite *SnapCollWorkflowSuite) TearDownSuite() {
	err := deleteDefaultVolume(suite.volumeService)
	assert.Nilf(suite.T(), err, "Unable to delete default volume, err: %v", err)
	err = deleteDefaultVolColl(suite.volcollService)
	assert.Nilf(suite.T(), err, "Unable to delete default volume collection, err: %v", err)
	suite.deleteSnapColl(snapcollName)
	suite.groupService.LogoutService()
}

func (suite *SnapCollWorkflowSuite) deleteSnapColl(snapcollName string) {
	snapcoll, _ := suite.snapcollService.GetSnapshotCollectionByName(snapcollName)
	if snapcoll != nil {
		err := suite.snapcollService.DeleteSnapshotCollection(*snapcoll.ID)
		assert.Nilf(suite.T(), err, "Unable to delete snapshot collection, err: %v", err)
	}
}

func (suite *SnapCollWorkflowSuite) TestCreateSnapshotCollection() {
	// Create volume collection & associate volumes
	volcoll, _ := suite.volcollService.GetVolumeCollectionByName(defaultVolCollName)
	vol, _ := suite.volumeService.GetVolumeByName(defaultVolumeName)
	err := suite.volumeService.AssociateVolume(*vol.ID, *volcoll.ID)
	assert.Nilf(suite.T(), err, "Unable to associate volume to volume collection, err: %v", err)
	// Create snapshot schedule and set frequency to 1 min
	var period *int64 = param.NewInt64(1)
	var numRetain *int64 = param.NewInt64(12)
	newPS := &nimbleos.ProtectionSchedule{
		Name:                  param.NewString(protectionScheduleName),
		VolcollOrProttmplType: nimbleos.NsProtectionPolicyTypeVolumeCollection,
		VolcollOrProttmplId:   volcoll.ID,
		Period:                period,
		PeriodUnit:            nimbleos.NsPeriodUnitMinutes,
		NumRetain:             numRetain,
	}
	_, err = suite.protectionScheduleService.CreateProtectionSchedule(newPS)
	assert.Nilf(suite.T(), err, "Unable to create protection schedule", err)

	filter := &param.GetParams{}
	snapcollBefore, _ := suite.snapcollService.GetSnapshotCollections(filter)

	snapcoll := &nimbleos.SnapshotCollection{
		Name:      param.NewString(snapcollName),
		VolcollId: volcoll.ID,
	}
	_, err = suite.snapcollService.CreateSnapshotCollection(snapcoll)
	assert.Nilf(suite.T(), err, "Unable to create snapshot collection, err: %v ", err)
	snapcollCreate, _ := suite.snapcollService.GetSnapshotCollections(filter)
	assert.Greater(suite.T(), len(snapcollCreate), len(snapcollBefore), "Unable to create snapshot collection")
	_ = suite.volumeService.DisassociateVolume(*vol.ID)
	updateSnapcoll := &nimbleos.SnapshotCollection{
		Description: param.NewString("Updated snapshot collection"),
	}
	currSnapcoll, _ := suite.snapcollService.GetSnapshotCollectionByName(snapcollName)
	_, err = suite.snapcollService.UpdateSnapshotCollection(*currSnapcoll.ID, updateSnapcoll)
	assert.Nilf(suite.T(), err, "Unable to update snapshot collection, err: %v")
}

func TestSnapCollWorkflowSuite(t *testing.T) {
	suite.Run(t, new(SnapCollWorkflowSuite))
}
