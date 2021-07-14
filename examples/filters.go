// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

func main() {

	// login to Array, get groupService instance
	groupService, err := service.NewNsGroupService("1.1.1.1", "xxxx", "xxxx", "v1", true, true)
	if err != nil {
		fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
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
		FieldName: nimbleos.VolumeFields.Name,
		Operator:  param.EQUALS.String(),
		Value:     "GetVolume",
	}
	// apply the filter
	sfilter.WithSearchFilter(f)

	//GetVolumes
	volumes, err := volSvc.GetVolumes(sfilter)
	fmt.Printf("%v", volumes)
}
