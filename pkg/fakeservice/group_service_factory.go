// Copyright 2020 Hewlett Packard Enterprise Development LP

package fakeservice

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/sdkprovider"
)

func getFakeNimbleID(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// NsGroupService type
type NsGroupService struct {
	ip string
	// Declare all the supported services
	volumeService sdkprovider.VolumeService
	poolService   sdkprovider.PoolService
}

// NewNsGroupService - initializes NsGroupService
// This function is deprecated. Use NewNimbleGroupService() to instantiate a new service
func NewNsGroupService(ip, username, password, apiVersion string, synchronous bool) (gs *NsGroupService, err error) {
	return &NsGroupService{ip: ip}, nil
}

func NewNimbleGroupService(serviceOpts ...service.ServiceOptions) (gs *NsGroupService, err error) {
	fakeGroupServiceOption := &service.GroupServiceOption{}

	for _, opt := range serviceOpts {
		opt(fakeGroupServiceOption)
	}

	return &NsGroupService{ip: fakeGroupServiceOption.Host}, nil
}

// SetDebug - enable debugging
func (gs *NsGroupService) SetDebug() {
	return
}

// LogoutService - delete the session token
func (gs *NsGroupService) LogoutService() error {
	return nil
}

// GetVolumeService - returns service of a type *VolumeService
func (gs *NsGroupService) GetVolumeService() (vs sdkprovider.VolumeService) {
	if gs.volumeService == nil {
		gs.volumeService = NewVolumeService(gs)
	}
	return gs.volumeService
}

// GetPoolService - returns service of a type *PoolService
func (gs *NsGroupService) GetPoolService() (vs sdkprovider.PoolService) {
	if gs.poolService == nil {
		gs.poolService = NewPoolService(gs)
	}
	return gs.poolService
}
