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

const volcollName = "VolumeCollectionTest"

type VolCollWorkflowSuite struct {
	suite.Suite
	groupService   *service.NsGroupService
	volcollService *service.VolumeCollectionService
	volService     sdkprovider.VolumeService
}

func (suite *VolCollWorkflowSuite) SetupSuite() {
	groupService, err := config()
	assert.Nilf(suite.T(), err, "Unable to connect to group, err: %v", err)
	suite.groupService = groupService
	suite.volcollService = groupService.GetVolumeCollectionService()
	suite.volService = groupService.GetVolumeService()
	_, err = createDefaultVolume(suite.volService)
	assert.Nilf(suite.T(), err, "Unable to create default volume, err: %v", err)
	suite.createVolColl(volcollName)
}

func (suite *VolCollWorkflowSuite) TearDownSuite() {
	err := deleteDefaultVolume(suite.volService)
	assert.Nilf(suite.T(), err, "Unable to delete default volume, err: %v", err)
	suite.deleteVolColl(volcollName)
	suite.groupService.LogoutService()
}

func (suite *VolCollWorkflowSuite) createVolColl(volcollName string) {
	newVolColl := &nimbleos.VolumeCollection{
		Name: param.NewString(volcollName),
	}
	volcoll, err := suite.volcollService.CreateVolumeCollection(newVolColl)
	assert.Nilf(suite.T(), err, "Unable to create volume collection, err: %v", err)
	vol, _ := suite.volService.GetVolumeByName(defaultVolumeName)
	err = suite.volService.AssociateVolume(*vol.ID, *volcoll.ID)
	assert.Nilf(suite.T(), err, "Associating volume to volume collection failed, err: %v", err)
}

func (suite *VolCollWorkflowSuite) deleteVolColl(volcollName string) {
	volcoll, _ := suite.volcollService.GetVolumeCollectionByName(volcollName)
	if volcoll != nil {
		err := suite.volcollService.DeleteVolumeCollection(*volcoll.ID)
		assert.Nilf(suite.T(), err, "Unable to delete volume collection, err: %v", defaultVolCollName)
	}
}

func (suite *VolCollWorkflowSuite) TestCreateVolCollDuplicate() {
	newVolColl := &nimbleos.VolumeCollection{
		Name: param.NewString(volcollName),
	}
	_, err := suite.volcollService.CreateVolumeCollection(newVolColl)
	assert.NotNil(suite.T(), err, "Duplicate Volume collection should have failed.")
}

func (suite *VolCollWorkflowSuite) TestCreateGenericVolCol() {
	newVolColl := &nimbleos.VolumeCollection{
		Name:          param.NewString("VolColGeneric"),
		AppSync:       nimbleos.NsAppSyncTypeGeneric,
		AgentHostname: param.NewString("10.1.0.0"),
		AgentUsername: param.NewString("xxx"),
		AgentPassword: param.NewString("xxx"),
	}
	volcoll, err := suite.volcollService.CreateVolumeCollection(newVolColl)
	assert.Nilf(suite.T(), err, "Unable to create VOlColl for generic app sync, err: %v", err)
	assert.Equal(suite.T(), *nimbleos.NsAppSyncTypeGeneric, *volcoll.AppSync, "VolColl app sync type not set correctly.")
	suite.deleteVolColl(*volcoll.Name)
}

func (suite *VolCollWorkflowSuite) TestCreateVssVolCol() {
	newVolColl := &nimbleos.VolumeCollection{
		Name:      param.NewString("VolColVss"),
		AppSync:   nimbleos.NsAppSyncTypeVss,
		AppServer: param.NewString("10.1.1.1"),
		AppId:     nimbleos.NsAppIdTypeExchangeDag,
	}
	volcoll, err := suite.volcollService.CreateVolumeCollection(newVolColl)
	assert.Nilf(suite.T(), err, "Unable to create VolColl for VSS, err: %v", err)
	assert.Equal(suite.T(), *nimbleos.NsAppSyncTypeVss, *volcoll.AppSync, "VolColl app sync type not set correctly.")
	suite.deleteVolColl(*volcoll.Name)
}

func (suite *VolCollWorkflowSuite) TestCreateVmwareVolColl() {
	newVolColl := &nimbleos.VolumeCollection{
		Name:            param.NewString("VolCollVmare"),
		AppSync:         nimbleos.NsAppSyncTypeVmware,
		VcenterHostname: param.NewString("10.0.0.0"),
		VcenterUsername: param.NewString("xxx"),
		VcenterPassword: param.NewString("xxx"),
	}
	volcoll, err := suite.volcollService.CreateVolumeCollection(newVolColl)
	assert.Nilf(suite.T(), err, "Unable to create VolColl with VMWare app sync, err: %v", err)
	assert.Equal(suite.T(), *nimbleos.NsAppSyncTypeVmware, *volcoll.AppSync, "VolColl app sync type not set correctly.")
	suite.deleteVolColl(*volcoll.Name)
}

func (suite *VolCollWorkflowSuite) TestDeleteProtectedVolume() {
	vol, _ := suite.volService.GetVolumeByName(defaultVolumeName)
	if vol != nil {
		err := suite.volService.DeleteVolume(*vol.ID)
		assert.NotNil(suite.T(), err, "Delete volume should have failed")
	}
}

func (suite *VolCollWorkflowSuite) TestDissassociateVol() {
	vol, _ := suite.volService.GetVolumeByName(defaultVolumeName)
	err := suite.volService.DisassociateVolume(*vol.ID)
	assert.Nilf(suite.T(), err, "Disassociating volume from volume collection failed, err: %v", err)

}

func (suite *VolCollWorkflowSuite) TestVolCollSchedule() {
	var numRetain *int64 = param.NewInt64(123)
	volcoll, _ := suite.volcollService.GetVolumeCollectionByName(volcollName)
	protectionScheduleService := suite.groupService.GetProtectionScheduleService()
	ps := &nimbleos.ProtectionSchedule{
		Name:                  param.NewString("Schedule-1"),
		VolcollOrProttmplId:   volcoll.ID,
		VolcollOrProttmplType: nimbleos.NsProtectionPolicyTypeVolumeCollection,
		NumRetain:             numRetain,
	}
	_, err := protectionScheduleService.CreateProtectionSchedule(ps)
	assert.Nilf(suite.T(), err, "Unable to associated protection schedule", err)

}

func TestVolCollWorkflowSuite(t *testing.T) {
	suite.Run(t, new(VolCollWorkflowSuite))
}
