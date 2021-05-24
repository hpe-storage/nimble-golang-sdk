// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"flag"
	"strconv"
	"strings"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

const defaultVolumeName = "DefaultVolumeTest"
const defaultInitiatorGrpName = "DefaultInitiatorgrpTest"
const defaultVolCollName = "DefaultVolCollTest"

var arrayIP = flag.String("arrayIP", "", "Array IP")

var arrayUsername = flag.String("arrayUsername", "xxx", "Array Username")

var arrayPassword = flag.String("arrayPassword", "xxx", "Array Password")

var downstreamArrayIP = flag.String("downstream", "", "IP address of an array to be used as a downstream replica. Pass this IP to execute replication partner test cases.")

// Required for group merge test
var sourceArrayIP = flag.String("sourceArrayIP", "", "IP address of source array, used for group and merge test cases.")

var sourceArrayUsername = flag.String("sourceArrayUsername", "xxx", "Source array username")

var sourceArrayPassword = flag.String("sourceArrayPassword", "xxx", "Source array password")

//  Dashboard
var postResultToDashboard = flag.String("postResultToDashboard", "xxx", "Dashboard flag")
var dashboardUsername = flag.String("dashboardUsername", "xxx", "Dashboard username")
var dashboardPassword = flag.String("dashboardPassword", "xxx", "Dashboard password")
var dashboardURL = flag.String("dashboardURL", "xxx", "Dashboard Url")
var testStarted []bool
var testCompleted []bool

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
	if group != nil {
		return *group.FcEnabled
	}
	return false
}

func isIscsiEnabled(arrayGroupService *service.GroupService) bool {
	filter := &param.GetParams{}
	groups, _ := arrayGroupService.GetGroups(filter)
	group := groups[0]
	if group != nil {
		return *group.IscsiEnabled
	}
	return false
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
