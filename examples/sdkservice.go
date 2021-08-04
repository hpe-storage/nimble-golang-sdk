package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/fakeservice"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
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

func getFakeService(clientOpts ...service.ServiceOption) (sdkprovider.NsGroupService, error) {
	grpSvc, err := fakeservice.NewNimbleGroupService(clientOpts...)
	return grpSvc, err
}

func getRealService(host string, isTenant bool, username, password string) (*service.NsGroupService, error) {
	if isTenant {
		groupService, err := service.NewNimbleGroupService(
			service.WithHost(host),
			service.WithTenantUser(username),
			service.WithPassword(password))
		if err != nil {
			fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
			os.Exit(1)
		}
		return groupService, err
	} else {
		groupService, err := service.NewNimbleGroupService(
			service.WithHost(host),
			service.WithUser(username),
			service.WithPassword(password))
		if err != nil {
			fmt.Printf("NewGroupService(): Unable to connect to group, err: %v", err.Error())
			os.Exit(1)
		}
		return groupService, err
	}
}

func main() {
	arg := &param.GetParams{}
	host, isTenant, username, password := parseOptionalArgument()
	groupService, _ := getFakeService(
		service.WithHost(host),
		service.WithUser(username),
		service.WithPassword(password))
	defer groupService.LogoutService()
	groupService.SetDebug()

	p, _ := groupService.GetPoolService().GetPools(arg)
	fmt.Printf("Fake Pools %+v \n", p)

	// volume operations
	var sizeField int64 = 5120
	descriptionField := "This volume was created as part of a unit test"
	var limitIopsField int64 = 256
	var limitMbpsField int64 = 1

	newVolume := &nimbleos.Volume{
		Name:           param.NewString("fake-test-vol1"),
		Size:           &sizeField,
		Description:    &descriptionField,
		Online:         param.NewBool(true),
		LimitIops:      &limitIopsField,
		LimitMbps:      &limitMbpsField,
		MultiInitiator: param.NewBool(true),
		AgentType:      nimbleos.NsAgentTypeNone,
	}

	vol, _ := groupService.GetVolumeService().CreateVolume(newVolume)
	fmt.Printf("Fake volume %+v \n", vol)

	// Get real service
	groupService, _ = getRealService(host, isTenant, username, password)

	defer groupService.LogoutService()
	groupService.SetDebug()

	p, _ = groupService.GetPoolService().GetPools(arg)
	fmt.Printf("Real Pools %+v \n", p)

	// reset id
	newVolume.ID = nil
	vol, _ = groupService.GetVolumeService().CreateVolume(newVolume)
	fmt.Printf("Real volume %+v \n", vol)

	groupService.GetVolumeService().DestroyVolume(*vol.ID)
}
