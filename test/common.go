// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"flag"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

const defaultVolumeName = "DefaultVolumeTest"
const defaultInitiatorGrpName = "DefaultInitiatorgrpTest"
const defaultVolCollName = "DefaultVolCollTest"

var arrayIP = flag.String("arrayIP", "1.1.1.1", "Array IP")

var arrayUsername = flag.String("arrayUsername", "xxx", "Array Username")

var arrayPassword = flag.String("arrayPassword", "xxx", "Array Password")

// Required for group merge test
const sourceArrayIP = "1.1.1.2"
const sourceArrayusername = "xxx"
const sourceArraypassword = "xxx"

func config() (*service.NsGroupService, error) {
	groupService, err := service.NewNsGroupService(*arrayIP, *arrayUsername, *arrayPassword, "v1", true)
	if err != nil {
		return groupService, err
	}
	// set debug
	groupService.SetDebug()
	return groupService, err
}

func createDefaultVolume(volSvc sdkprovider.VolumeService) (*nimbleos.Volume, error) {
	var sizeField int64 = 5120
	newVolume := &nimbleos.Volume{
		Name: param.NewString(defaultVolumeName),
		Size: &sizeField,
	}
	return volSvc.CreateVolume(newVolume)
}

func createDefaultInitiatorGrp(igSvc *service.InitiatorGroupService) (*nimbleos.InitiatorGroup, error) {
	newIG := &nimbleos.InitiatorGroup{
		Name:           param.NewString(defaultInitiatorGrpName),
		Description:    param.NewString("Workflow initiator group"),
		AccessProtocol: nimbleos.NsAccessProtocolIscsi,
	}
	return igSvc.CreateInitiatorGroup(newIG)

}

func createDefaultVolColl(volcollService *service.VolumeCollectionService) (*nimbleos.VolumeCollection, error) {
	newVolColl := &nimbleos.VolumeCollection{
		Name: param.NewString(defaultVolCollName),
	}
	return volcollService.CreateVolumeCollection(newVolColl)
}

func deleteDefaultVolume(volSvc sdkprovider.VolumeService) error {
	volObj, err := volSvc.GetVolumeByName(defaultVolumeName)
	if err != nil {
		return err
	}
	_, err = volSvc.OfflineVolume(*volObj.ID, true)
	if err != nil {
		return err
	}
	return volSvc.DeleteVolume(*volObj.ID)
}

func deleteDefaultInitiatorGroup(igSvc *service.InitiatorGroupService) error {
	ig, err := igSvc.GetInitiatorGroupByName(defaultInitiatorGrpName)
	if err != nil {
		return err
	}
	return igSvc.DeleteInitiatorGroup(*ig.ID)
}

func deleteDefaultVolColl(volcollService *service.VolumeCollectionService) error {
	volcoll, err := volcollService.GetVolumeCollectionByName(defaultVolCollName)
	if err != nil {
		return err
	}
	return volcollService.DeleteVolumeCollection(*volcoll.ID)
}
