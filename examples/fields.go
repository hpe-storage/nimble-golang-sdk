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
