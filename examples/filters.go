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

	// get volume service instance
	volSvc := groupService.GetVolumeService()

	// init param
	sfilter := &param.GetParams{}

	// set search filter
	f := &param.SearchFilter{
		FieldName: &nimbleos.VolumeFields.Name,
		Operator:  param.EQUALS.String(),
		Value:     "GetVolume",
	}

	// apply the filter
	sfilter.WithSearchFilter(f)

	//GetVolumes
	volumes, _ := volSvc.GetVolumes(sfilter)
	fmt.Printf("%v", volumes)
}
