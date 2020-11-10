package workflows

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

var defaultVolumeName = "DefaultVolumeTest"
var defaultInitiatorGrp = "DefaultInitiatorgrpTest"
var defaultVolColName = "DefaultVolColTest"

func config() *service.NsGroupService {
	groupService, err := service.NewNsGroupService("10.21.1.216", "admin", "admin")
	if err != nil {
		fmt.Errorf("NewGroupService(): unable to connect to group, err: %v", err.Error())
	}

	// set debug
	groupService.SetDebug()
	// suite.groupService = groupService
	// suite.volumeService = groupService.GetVolumeService()
	return groupService
}

func createDefaultVolume(groupService *service.NsGroupService) (*nimbleos.Volume, error) {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(defaultVolumeName),
		Size: &sizeField,
	}
	volSvc := groupService.GetVolumeService()
	return volSvc.CreateVolume(newVolume)
}

func createDefaultInitiatorGrp(groupService *service.NsGroupService) {
	newIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString(defaultInitiatorGrp),
		Description:    param.NewString("Workflow initiator group"),
		AccessProtocol: nimbleos.NsAccessProtocolIscsi,
	}
	igSvc := groupService.GetInitiatorGroupService()
	_, err := igSvc.CreateInitiatorGroup(newIG)
	if err != nil {
		fmt.Errorf("Initiator group creation failed with error: %v", err.Error())
	}
}

func createDefaultVolumeCollection(groupService *service.NsGroupService) {
	newVolCol := &nimbleos.VolumeCollection{
		Name: param.NewString(defaultVolColName),
	}
	volcolService := groupService.GetVolumeCollectionService()
	_, err := volcolService.CreateVolumeCollection(newVolCol)
	if err != nil {
		fmt.Errorf("Volume Collection creation failed: %v", err.Error())
	}
}

func deleteVolume(groupService *service.NsGroupService, volName string) {
	volSvc := groupService.GetVolumeService()
	volObj, _ := volSvc.GetVolumeByName(volName)
	if volObj != nil {
		err := volSvc.DeleteVolume(*volObj.ID)
		if err != nil {
			fmt.Errorf("Issue deleting volume: %v", volName)
		}
	}
}

func deleteInitiatorGroup(groupService *service.NsGroupService, igName string) {
	igSvc := groupService.GetInitiatorGroupService()
	ig, _ := igSvc.GetInitiatorGroupByName(igName)
	if ig != nil {
		err := igSvc.DeleteInitiatorGroup(*ig.ID)
		if err != nil {
			fmt.Errorf("Issue deleting Initiator group: %v", igName)
		}
	}
}

func deleteVolCol(groupService *service.NsGroupService, volcolName string) {
	volcolSvc := groupService.GetVolumeCollectionService()
	volcol, _ := volcolSvc.GetVolumeCollectionByName(volcolName)
	if volcol != nil {
		err := volcolSvc.DeleteVolumeCollection(*volcol.ID)
		if err != nil {
			fmt.Errorf("Issue deleting volume collection: %v", volcolName)
		}
	}
}
