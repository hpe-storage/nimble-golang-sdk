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
	fmt.Println("Using user's username and password are used for service initialization")
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

	// get volume service instance
	volSvc := groupService.GetVolumeService()

	// get initiator service instance
	igroupSvc := groupService.GetInitiatorGroupService()

	// get access control service instance
	aclSvc := groupService.GetAccessControlRecordService()

	// get peformance policy service instance
	perfSvc := groupService.GetPerformancePolicyService()

	perfPolicy, _ := perfSvc.GetPerformancePolicyByName("default")
	// Initialize volume attributes
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	newVolume := &nimbleos.Volume{
		Name:              param.NewString("Test2020"),
		Size:              &sizeField,
		Description:       &descriptionField,
		PerfpolicyId:      perfPolicy.ID,
		ThinlyProvisioned: param.NewBool(true),
		Online:            param.NewBool(true),
		LimitIops:         &limitIopsField,
		LimitMbps:         &limitMbpsField,
		MultiInitiator:    param.NewBool(true),
		AgentType:         nimbleos.NsAgentTypeNone,
	}

	// initiator init
	initiator := &nimbleos.NsISCSIInitiator{
		Label:     param.NewString("iqn.1998-01.com.vmware:sasi-srm82-pesxi-4b00546a"),
		Iqn:       param.NewString("iqn.1998-01.com.vmware:sasi-srm82-pesxi-4b00546a"),
		IpAddress: param.NewString("*"),
	}

	var initiatorList []*nimbleos.NsISCSIInitiator
	initiatorList = append(initiatorList, initiator)

	// igroup init
	igroup := &nimbleos.InitiatorGroup{
		Name:            param.NewString("sdkigroup"),
		AccessProtocol:  nimbleos.NsAccessProtocolIscsi,
		IscsiInitiators: initiatorList,
	}

	// create initiator group
	igroup, _ = igroupSvc.CreateInitiatorGroup(igroup)

	// create a new volume and assigne acl
	volume, _ := volSvc.CreateVolume(newVolume)
	if volume != nil {
		acl := &nimbleos.AccessControlRecord{
			InitiatorGroupId: igroup.ID,
			VolId:            volume.ID,
		}

		// create acl
		_, err := aclSvc.CreateAccessControlRecord(acl)
		if err != nil {
			fmt.Println(err)
		}
	}

	// get igroup by name
	igroup, err := igroupSvc.GetInitiatorGroupByName("sdkigroup")
	if err != nil {
		fmt.Printf("Failed to get igroup by name, err %v\n", err)
	}
	fmt.Println(igroup)

	// cleaup
	updateVol := &nimbleos.Volume{
		Online: param.NewBool(false),
	}

	// set volume offline
	volume, _ = volSvc.UpdateVolume(*volume.ID, updateVol)

	// delete the volume
	_ = volSvc.DeleteVolume(*volume.ID)

	// delete igroup
	_ = igroupSvc.DeleteInitiatorGroup(*igroup.ID)
}
