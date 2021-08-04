// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"
	"os"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

func checkEnvironmentVariableExists(){
	if (os.Getenv("SDK_TARGET_HOST") == "" ||
		os.Getenv("SDK_TARGET_USER") == "" ||
		os.Getenv("SDK_TARGET_USER_PASSWORD") == "") {
			fmt.Println("ERROR: Missing one of these environment variables: SDK_TARGET_HOST, SDK_TARGET_USER, SDK_TARGET_USER_PASSWORD");
			os.Exit(1)
		}
}

func main() {
	checkEnvironmentVariableExists()

	groupService, err := service.NewNimbleGroupService(
		service.WithHost(os.Getenv("SDK_TARGET_HOST")),
		service.WithUser(os.Getenv("SDK_TARGET_USER")),
		service.WithPassword(os.Getenv("SDK_TARGET_USER_PASSWORD")))
	if err != nil {
		fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
		os.Exit(1)
	}
	defer groupService.LogoutService()

	// set debug
	groupService.SetDebug()
	volSvc := groupService.GetVolumeService()

	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	// set volume attributes
	newVolume := &nimbleos.Volume{
		Name:              param.NewString("TestDemo"),
		Size:              &sizeField,
		Description:       &descriptionField,
		ThinlyProvisioned: param.NewBool(true),
		Online:            param.NewBool(true),
		LimitIops:         &limitIopsField,
		LimitMbps:         &limitMbpsField,
		MultiInitiator:    param.NewBool(true),
		AgentType:         nimbleos.NsAgentTypeNone,
	}

	// create volume
	volume, err := volSvc.CreateVolume(newVolume)
	if err != nil {
		fmt.Println("Failed to create volume, ", err)
	}
	// init param
	filter := &param.GetParams{}

	// set attribute fields
	var volAttrList = []string{
		nimbleos.VolumeFields.ID,
		nimbleos.VolumeFields.Name,
		nimbleos.VolumeFields.PerfpolicyName,
	}
	// apply attributes
	filter.WithFields(volAttrList)

	// create a filter
	sf := &param.SearchFilter{
		FieldName: &nimbleos.VolumeFields.Name,
		Operator:  param.EQUALS.String(),
		Value:     "TestDemo",
	}

	filter.WithSearchFilter(sf)

	volumes, err := volSvc.GetVolumes(filter)
	fmt.Printf("%+v\n", *volumes[0])

	// delete
	volSvc.DeleteVolume(*volume.ID)
}
