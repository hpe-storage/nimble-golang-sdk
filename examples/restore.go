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

	// get  volume service instance
	volSvc := groupService.GetVolumeService()

	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	// set volume attributes
	newVolume := &nimbleos.Volume{
		Name:              param.NewString("RestoreVolume"),
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
	if volume != nil {
		volcollName := &nimbleos.VolumeCollection{
			Name: param.NewString("TestVolcoll"),
		}
		volcoll, err := groupService.GetVolumeCollectionService().CreateVolumeCollection(volcollName)
		if err != nil {
			fmt.Println("Failed to create volume collection")
			return
		}

		// add volume to volume collection
		err = volSvc.AssociateVolume(*volume.ID, *volcoll.ID)
		if err != nil {
			fmt.Println("Failed to associate RestoreVolume to volcoll ")
			return
		}

		// create a snapshot collection
		snapColl, _ := groupService.GetSnapshotCollectionService().CreateSnapshotCollection(&nimbleos.SnapshotCollection{
			Name:      param.NewString("RestoreSnapColl"),
			VolcollId: volcoll.ID,
		})

		// Get snapshot collection by name
		snapColl, err = groupService.GetSnapshotCollectionService().GetSnapshotCollectionByName("RestoreSnapColl")
		if err != nil {
			fmt.Printf("Failed to get snapshot collection by name, err: %v\n", err)
		}
		fmt.Println(snapColl)

		// set the volume offline before restore
		volSvc.OfflineVolume(*volume.ID, true)
		//restore volume to snapcoll
		err = volSvc.RestoreVolume(*volume.ID, *snapColl.SnapshotsList[0].SnapId)

		if err != nil {
			fmt.Println("Failed to restore volume")

		}

		// cleanup
		// disassociate volume from volume collection
		err = volSvc.DisassociateVolume(*volume.ID)

		if err != nil {
			fmt.Printf("Failed to remove %s volume from volume collection\n", *volume.ID)
			return
		}

		// delete volcoll
		groupService.GetVolumeCollectionService().DeleteVolumeCollection(*volcoll.ID)
		volSvc.DeleteVolume(*volume.ID)

	}

}
