package main

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/fakeservice"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

func getFakeService(clientOpts ...service.ServiceOptions) (sdkprovider.NsGroupService, error) {
	grpSvc, err := fakeservice.NewNimbleGroupService(clientOpts...)
	return grpSvc, err
}

func getRealService(clientOpts ...service.ServiceOptions) (sdkprovider.NsGroupService, error) {
	grpSvc, err := service.NewNimbleGroupService(clientOpts...)
	return grpSvc, err
}

func main() {
	arg := &param.GetParams{}
	groupService, _ := getFakeService(service.WithHost("1.1.1.1"),
		service.WithUser("xxx"), service.WithPassword("xxx"),
		service.WithoutWaitForAsyncJobs())
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
	groupService, _ = getRealService(service.WithHost("1.1.1.1"),
		service.WithUser("xxx"), service.WithPassword("xxx"),
		service.WithoutWaitForAsyncJobs())

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
