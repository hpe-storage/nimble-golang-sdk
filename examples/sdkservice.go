package main

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/fakeservice"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"os"
)

func getFakeService(clientOpts ...service.ServiceOption) (sdkprovider.NsGroupService, error) {
	grpSvc, err := fakeservice.NewNimbleGroupService(clientOpts...)
	return grpSvc, err
}

func getRealService(clientOpts ...service.ServiceOption) (sdkprovider.NsGroupService, error) {
	grpSvc, err := service.NewNimbleGroupService(clientOpts...)
	return grpSvc, err
}

func main() {
	arg := &param.GetParams{}
	host := os.Getenv("SDK_TARGET_HOST")
	user := os.Getenv("SDK_TARGET_USER")
	password := os.Getenv("SDK_TARGET_PASSWORD")

	if host == "" || user == "" || password == "" {
		fmt.Println("ERROR: Missing one of these environment variables: SDK_TARGET_HOST, SDK_TARGET_USER, SDK_TARGET_PASSWORD")
		fmt.Println("Usage:")
		fmt.Println("SDK_TARGET_HOST - Management hostname or IP of array")
		fmt.Println("SDK_TARGET_USER - User (non-tenant) username")
		fmt.Println("SDK_TARGET_USER_PASSWORD - User (non-tenant) password")
		os.Exit(1)
	}

	groupService, _ := getFakeService(
		service.WithHost(host),
		service.WithUser(user),
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
	groupService, _ = getRealService(
		service.WithHost(host),
		service.WithUser(user),
		service.WithPassword(password))

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
