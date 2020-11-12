package workflow

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

var defaultVolumeName = "DefaultVolumeTest"
var defaultInitiatorGrpName = "DefaultInitiatorgrpTest"
var defaultVolCollName = "DefaultVolCollTest"

func config() (*service.NsGroupService, error) {
	groupService, err := service.NewNsGroupService("", "admin", "admin", "v1", true)
	if err != nil {
		fmt.Errorf("NewGroupService(): unable to connect to group, err: %v", err.Error())
	}
	defer groupService.LogoutService()
	// set debug
	groupService.SetDebug()
	return groupService, err
}

func createDefaultVolume(volSvc *service.VolumeService) (*nimbleos.Volume, error) {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(defaultVolumeName),
		Size: &sizeField,
	}
	return volSvc.CreateVolume(newVolume)
}

func createDefaultInitiatorGrp(igSvc *service.InitiatorGroupService) {
	newIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString(defaultInitiatorGrpName),
		Description:    param.NewString("Workflow initiator group"),
		AccessProtocol: nimbleos.NsAccessProtocolIscsi,
	}
	_, err := igSvc.CreateInitiatorGroup(newIG)
	if err != nil {
		fmt.Errorf("Initiator group creation failed with error: %v", err.Error())
	}
}

func createDefaultVolColl(volcollService *service.VolumeCollectionService) {
	newVolColl := &nimbleos.VolumeCollection{
		Name: param.NewString(defaultVolCollName),
	}
	_, err := volcollService.CreateVolumeCollection(newVolColl)
	if err != nil {
		fmt.Errorf("Volume Collection creation failed: %v", err.Error())
	}
}

func deleteDefaultVolume(volSvc *service.VolumeService) {
	volObj, err := volSvc.GetVolumeByName(defaultVolumeName)
	if err != nil {
		fmt.Errorf("Unable to volume, err: %v", err.Error())
	} else {
		err := volSvc.DeleteVolume(*volObj.ID)
		if err != nil {
			fmt.Errorf("Unable to delete volume: %v", defaultVolumeName)
		}
	}
}

func deleteDefaultInitiatorGroup(igSvc *service.InitiatorGroupService) {
	ig, err := igSvc.GetInitiatorGroupByName(defaultInitiatorGrpName)
	if err != nil {
		fmt.Errorf("Unable to get initiator group, err: %v", err.Error())
	}
	if ig != nil {
		err := igSvc.DeleteInitiatorGroup(*ig.ID)
		if err != nil {
			fmt.Errorf("Unable to delete Initiator group: %v", defaultInitiatorGrpName)
		}
	}
}

func deleteDefaultVolColl(volcollService *service.VolumeCollectionService) {
	volcoll, err := volcollService.GetVolumeCollectionByName(defaultVolCollName)
	if err != nil {
		fmt.Errorf("Unable to get volume collection, err: %v", err.Error())
	}
	if volcoll != nil {
		err := volcollService.DeleteVolumeCollection(*volcoll.ID)
		if err != nil {
			fmt.Errorf("Unable to delete volume collection: %v", defaultVolCollName)
		}
	}
}
