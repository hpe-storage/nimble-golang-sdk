// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

func handleOptionError(target string){
	fmt.Printf("ERROR: Environment variable SDK_TARGET_%s is not defined.\n", strings.ToUpper(target))
	os.Exit(1)
}

/*
	This function returns: host, isTenant (indicates whether we're using regular user or tenantUser), username, and password
*/
func parseOptionalArgument()(string, bool, string, string){
	// define flags
	host := flag.String("host", "", "SDK Target Host")
	user := flag.String("user", "", "SDK Target User")
	password := flag.String("user-password", "", "SDK Target User Password")
	tenant := flag.String("tenant", "", "SDK Target Tenant")
	tenantPassword := flag.String("tenant-password", "", "SDK Target Tenant Password")

	flag.Parse()
	fmt.Println(*host, *user, *password, *tenant, *tenantPassword);

	if *host == "" {
		*host = os.Getenv("SDK_TARGET_HOST")
	}

	if *user == "" {
		*user = os.Getenv("SDK_TARGET_USER")
	}

	if *password == "" {
		*password = os.Getenv("SDK_TARGET_USER_PASSWORD")
	}

	if *tenant == "" {
		*tenant = os.Getenv("SDK_TARGET_TENANT_USER")
	}

	if *tenantPassword == "" {
		*tenantPassword = os.Getenv("SDK_TARGET_TENANT_PASSWORD")
	}

	// Host has to be defined
	if *host == "" {
		handleOptionError("host")
	}

	// Either (user & password) or (tenant & tenantPassword) has to be defined
	isTenant := false

	if *user == "" && *password== "" && *tenant == "" && *tenantPassword == "" {
		fmt.Println("ERROR: Must define either a user or a tenant.")
		os.Exit(1)
	}

	// Initalize with user instead of tenant
	if *tenant == "" && *tenantPassword == "" {
		// Error handling
		if (*user != "" || *password != "") {
			if *user != "" && *password == "" {
				handleOptionError("user_password")
			}
			if *user == "" && *password != "" {
				handleOptionError("user")
			}
		}
		return *host, isTenant, *user, *password
	}

	// Initialize with a tenant
	if *user == "" && *password == "" {
		isTenant = true
		// Error handling
		if (*tenant != "" || *tenantPassword != ""){
			if *tenant != "" && *tenantPassword == "" {
				handleOptionError("tenant_password")
			}
			if *tenant == "" && *tenantPassword != "" {
				handleOptionError("tenant_user")
			}
		}
		return *host, isTenant, *tenant, *tenantPassword
	}

	// if all user, password, tenant, and tenantPassword are initialized
	// default to regular user
	fmt.Println("Both tenant's password & username and user's password & username are defined")
	fmt.Println("User's username and password are used for service initialization")
	return *host, isTenant, *user, *password
}

func initializeGroupService(host string, isTenant bool, username, password string) (*service.NsGroupService) {
	if isTenant {
		groupService, err := service.NewNimbleGroupService(
			service.WithHost(host),
			service.WithTenantUser(username),
			service.WithPassword(password))
		if err != nil {
			fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
			os.Exit(1)
		}
		return groupService
	} else {
		groupService, err := service.NewNimbleGroupService(
			service.WithHost(host),
			service.WithUser(username),
			service.WithPassword(password))
		if err != nil {
			fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
			os.Exit(1)
		}
		return groupService
	}
}

func main() {
	host, isTenant, username, password := parseOptionalArgument()

	groupService := initializeGroupService(host, isTenant, username, password)
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
