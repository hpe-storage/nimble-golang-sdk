// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"os"
)

func checkEnvironmentVariableExists(host, user, password string) {
	if host == "" || user == "" || password == "" {
		fmt.Println("ERROR: Missing one of these environment variables: SDK_TARGET_HOST, SDK_TARGET_USER, SDK_TARGET_PASSWORD")
		fmt.Println("See README for usage")
		os.Exit(1)
	}
}

func main() {
	host := os.Getenv("SDK_TARGET_HOST")
	user := os.Getenv("SDK_TARGET_USER")
	password := os.Getenv("SDK_TARGET_PASSWORD")

	checkEnvironmentVariableExists(host, user, password)

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
	volumes, err := volSvc.GetVolumes(sfilter)
	fmt.Printf("%v", volumes)

	if err != nil {
		fmt.Printf("NewGroupService(): Unable to get volumes, err: %v", err.Error())
	}
}
