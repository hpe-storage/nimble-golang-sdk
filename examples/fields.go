// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

func main() {
	groupService, err := service.NewNimbleGroupService(
		service.WithHost("10.157.82.90"),
 		service.WithUser("admin"),
 		service.WithPassword("admin"))
	if err != nil {
		fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
		return
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
		FieldName: param.NewString(nimbleos.VolumeFields.Name),
		Operator:  param.EQUALS.String(),
		Value:     "TestDemo",
	}

	filter.WithSearchFilter(sf)

	volumes, err := volSvc.GetVolumes(filter)
	fmt.Printf("%+v\n", *volumes[0])

	// delete
	volSvc.DeleteVolume(*volume.ID)
}
