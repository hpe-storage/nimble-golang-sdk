// Copyright 2020 Hewlett Packard Enterprise Development LP
package main

import (
	"fmt"
	"flag"
	"os"
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
