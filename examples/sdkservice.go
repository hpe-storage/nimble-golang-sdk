package main

import (
	"fmt"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/fakeservice"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

func getFakeService(ip, username, password, apiVersion string, synchronous bool) (sdkprovider.NsGroupService, error) {

	grpSvc, err := fakeservice.NewNsGroupService(ip, username, password, apiVersion, synchronous)
	return grpSvc, err
}

func getRealService(ip, username, password, apiVersion string, synchronous, tenantAware bool) (sdkprovider.NsGroupService, error) {
	grpSvc, err := service.NewNsGroupService(ip, username, password, apiVersion, synchronous, tenantAware)
	return grpSvc, err
}

func main() {
	arg := &param.GetParams{}
	groupService, _ := getFakeService("1.1.1.1", "xxx", "xxx", "v1", true)
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
	groupService, _ = getRealService("1.1.1.1", "xxxx", "xxxx", "v1", true, true)

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
