// Copyright 2020 Hewlett Packard Enterprise Development LP

package fakeservice

// Pool Service - Manage pools. Pools are an aggregation of arrays.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// PoolService type
type PoolService struct {
	// create a map
}

// NewPoolService - method to initialize "PoolService"
func NewPoolService(gs *NsGroupService) *PoolService {

	return &PoolService{}
}

// GetPools - method returns a array of pointers of type "Pools"
func (svc *PoolService) GetPools(params *param.GetParams) ([]*nimbleos.Pool, error) {
	return nil, nil
}

// CreatePool - method creates a "Pool"
func (svc *PoolService) CreatePool(obj *nimbleos.Pool) (*nimbleos.Pool, error) {
	return nil, nil
}

// UpdatePool - method modifies  the "Pool"
func (svc *PoolService) UpdatePool(id string, obj *nimbleos.Pool) (*nimbleos.Pool, error) {
	return nil, nil
}

// GetPoolById - method returns a pointer to "Pool"
func (svc *PoolService) GetPoolById(id string) (*nimbleos.Pool, error) {
	return nil, nil
}

// GetPoolByName - method returns a pointer "Pool"
func (svc *PoolService) GetPoolByName(name string) (*nimbleos.Pool, error) {
	return nil, nil
}

// DeletePool - deletes the "Pool"
func (svc *PoolService) DeletePool(id string) error {

	return nil
}

// MergePool - merge the specified pool into the target pool. All volumes on the specified pool are moved to the target pool and the specified pool is then deleted. All the arrays in the pool are assigned to the target pool.
//   Required parameters:
//       id - ID of the specified pool.
//       targetPoolId - ID of the target pool.

//   Optional parameters:
//       force - Forcibly merge the specified pool into target pool.

func (svc *PoolService) MergePool(id string, targetPoolId string, force *bool) (*nimbleos.NsPoolMergeReturn, error) {

	return nil, nil
}
