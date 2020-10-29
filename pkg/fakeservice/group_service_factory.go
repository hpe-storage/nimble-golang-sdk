// Copyright 2020 Hewlett Packard Enterprise Development LP

package fakeservice

// NsGroupService type
type NsGroupService struct {
	ip string
	// Declare all the supported services
	volumeService *VolumeService
	poolService   *PoolService
}

// NewNsGroupService - initializes NsGroupService
func NewNsGroupService(ip, username, password, apiVersion string, synchronous bool) (gs *NsGroupService, err error) {

	return &NsGroupService{ip: ip}, nil
}

// LogoutService - delete the session token
func (gs *NsGroupService) LogoutService() error {
	return nil
}

// GetVolumeService - returns service of a type *VolumeService
func (gs *NsGroupService) GetVolumeService() (vs *VolumeService) {
	if gs.volumeService == nil {
		gs.volumeService = NewVolumeService(gs)
	}
	return gs.volumeService
}

// GetPoolService - returns service of a type *PoolService
func (gs *NsGroupService) GetPoolService() (vs *PoolService) {
	if gs.poolService == nil {
		gs.poolService = NewPoolService(gs)
	}
	return gs.poolService
}
