// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

func main() {
	groupService, err := service.NewNimbleGroupService(service.WithHost("1.1.1.1"), service.WithUser("xxx"), service.WithPassword("xxx"), service.WithoutWaitForAsyncJobs())
	if err != nil {
		fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
		return
	}
	// set debug
	groupService.SetDebug()
	defer groupService.LogoutService()
	// get  volume service instance
	volSvc := groupService.GetVolumeService()
	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	// set volume attributes
	newVolume := &nimbleos.Volume{
		Name:              param.NewString("TestDemo1"),
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
		fmt.Printf("Failed to create volume, err ", err)
		return
	}
	fmt.Println(volume)
	// delete volume, cleanup
	volSvc.DeleteVolume(*volume.ID)
}
