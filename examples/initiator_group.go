// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"os"
)

func main() {
	host := os.Getenv("SDK_TARGET_HOST")
	user := os.Getenv("SDK_TARGET_USER")
	password := os.Getenv("SDK_TARGET_PASSWORD")

	if host == "" || user == "" || password == "" {
		fmt.Println("ERROR: Missing one of these environment variables: SDK_TARGET_HOST, SDK_TARGET_USER, SDK_TARGET_PASSWORD")
		fmt.Println(
			`Usage:
			- SDK_TARGET_HOST				// Managment hostname or IP of array
			- SDK_TARGET_USER				// User (non-tenant) username
			- SDK_TARGET_USER_PASSWORD		// User (non-tenant) password`)
		os.Exit(1)
	}

	groupService, err := service.NewNimbleGroupService(
		service.WithHost(host),
		service.WithUser(user),
		service.WithPassword(password))
	if err != nil {
		fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
		os.Exit(1)
	}

	defer groupService.LogoutService()
	// set debug
	groupService.SetDebug()

	// get volume service instance
	volSvc := groupService.GetVolumeService()

	// get initiator service instance
	igroupSvc := groupService.GetInitiatorGroupService()

	// get access control service instance
	aclSvc := groupService.GetAccessControlRecordService()

	// get peformance policy service instance
	perfSvc := groupService.GetPerformancePolicyService()

	perfPolicy, _ := perfSvc.GetPerformancePolicyByName("default")
	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	newVolume := &nimbleos.Volume{
		Name:              param.NewString("Test2020"),
		Size:              &sizeField,
		Description:       &descriptionField,
		PerfpolicyId:      perfPolicy.ID,
		ThinlyProvisioned: param.NewBool(true),
		Online:            param.NewBool(true),
		LimitIops:         &limitIopsField,
		LimitMbps:         &limitMbpsField,
		MultiInitiator:    param.NewBool(true),
		AgentType:         nimbleos.NsAgentTypeNone,
	}

	// initiator init
	initiator := &nimbleos.NsISCSIInitiator{
		Label:     param.NewString("iqn.1998-01.com.vmware:sasi-srm82-pesxi-4b00546a"),
		Iqn:       param.NewString("iqn.1998-01.com.vmware:sasi-srm82-pesxi-4b00546a"),
		IpAddress: param.NewString("*"),
	}

	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, initiator)

	// igroup init
	igroup := &nimbleos.InitiatorGroup{
		Name:            param.NewString("sdkigroup"),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}

	// create initiator group
	igroup, _ = igroupSvc.CreateInitiatorGroup(igroup)

	// create a new volume and assigne acl
	volume, _ := volSvc.CreateVolume(newVolume)
	if volume != nil {
		acl := &nimbleos.AccessControlRecord{
			InitiatorGroupId: igroup.ID,
			VolId:            volume.ID,
		}

		// create acl
		_, err := aclSvc.CreateAccessControlRecord(acl)
		if err != nil {
			fmt.Println(err)
		}
	}

	// get igroup by name
	igroup, err = igroupSvc.GetInitiatorGroupByName("sdkigroup")
	if err != nil {
		fmt.Printf("Failed to get igroup by name, err %v\n", err)
	}
	fmt.Println(igroup)

	// cleaup
	updateVol := &nimbleos.Volume{
		Online: param.NewBool(false),
	}

	// set volume offline
	volume, _ = volSvc.UpdateVolume(*volume.ID, updateVol)

	// delete the volume
	_ = volSvc.DeleteVolume(*volume.ID)

	// delete igroup
	_ = igroupSvc.DeleteInitiatorGroup(*igroup.ID)
}
