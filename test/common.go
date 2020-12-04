// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"flag"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"strconv"
	"strings"
)

const defaultVolumeName = "DefaultVolumeTest"
const defaultInitiatorGrpName = "DefaultInitiatorgrpTest"
const defaultVolCollName = "DefaultVolCollTest"

var arrayIP = flag.String("arrayIP", "1.1.1.1", "Array IP")

var arrayUsername = flag.String("arrayUsername", "xxx", "Array Username")

var arrayPassword = flag.String("arrayPassword", "xxx", "Array Password")

var downstreamArrayIP = flag.String("downstream", "", "Array IP to be used as downstream")

// Required for group merge test
var sourceArrayIP = flag.String("sourceArrayIP", "1.1.1.2", "Source array IP for group tests")

var sourceArrayusername = flag.String("sourceArrayusername", "xxx", "Source array username")

var sourceArraypassword = flag.String("sourceArraypassword", "xxx", "Source array password")

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

func isFCEnabled(arrayGroupService *service.GroupService) bool {
	filter := &param.GetParams{}
	groups, _ := arrayGroupService.GetGroups(filter)
	group := groups[0]
	return *group.FcEnabled
}

func isIscsiEnabled(arrayGroupService *service.GroupService) bool {
	filter := &param.GetParams{}
	groups, _ := arrayGroupService.GetGroups(filter)
	group := groups[0]
	return *group.IscsiEnabled
}

func getArrayVersion(groupService *service.GroupService) float64 {
	filter := &param.GetParams{}
	groupResp, _ := groupService.GetGroups(filter)
	currentVersion := *groupResp[0].VersionCurrent
	currentVersion = strings.Split(strings.TrimSpace(currentVersion), "-")[0]
	currentVersionSplitlist := strings.Split(currentVersion, ".")[:2]
	arrayVersion, _ := strconv.ParseFloat(strings.Join(currentVersionSplitlist, "."), 8)
	return arrayVersion
}
