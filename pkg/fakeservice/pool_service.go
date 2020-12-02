// Copyright 2020 Hewlett Packard Enterprise Development LP

package fakeservice

// Pool Service - Manage pools. Pools are an aggregation of arrays.

import (
	"fmt"

	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// PoolService type
type PoolService struct {
	// create a map
	PoolMemory map[string]*nimbleos.Pool
}

// NewPoolService - method to initialize "PoolService"
func NewPoolService(gs *NsGroupService) *PoolService {

	return &PoolService{}
}

// GetPools - method returns a array of pointers of type "Pools"
func (svc *PoolService) GetPools(params *param.GetParams) ([]*nimbleos.Pool, error) {
	var pools []*nimbleos.Pool
	for _, v := range svc.PoolMemory {
		pools = append(pools, v)
	}
	return pools, nil
}

// CreatePool - method creates a "Pool"
func (svc *PoolService) CreatePool(obj *nimbleos.Pool) (*nimbleos.Pool, error) {
	if svc.PoolMemory == nil {
		svc.PoolMemory = make(map[string]*nimbleos.Pool)
	}
	if _, ok := svc.PoolMemory[*obj.ID]; !ok {
		id := getFakeNimbleID(*obj.Name)
		*obj.ID = id
		svc.PoolMemory[*obj.ID] = obj
	} else {
		return nil, fmt.Errorf("object already exists, duplicate pool %v", obj.Name)
	}
	return obj, nil
}

// UpdatePool - method modifies  the "Pool"
func (svc *PoolService) UpdatePool(id string, obj *nimbleos.Pool) (*nimbleos.Pool, error) {
	if _, ok := svc.PoolMemory[id]; ok {
		svc.PoolMemory[id] = obj
	} else {
		return nil, fmt.Errorf("object doesn't exist, failed to update pool %v", obj.Name)
	}
	return obj, nil
}

// GetPoolById - method returns a pointer to "Pool"
func (svc *PoolService) GetPoolById(id string) (*nimbleos.Pool, error) {
	if _, ok := svc.PoolMemory[id]; ok {
		return svc.PoolMemory[id], nil
	}
	return nil, nil
}

// GetPoolByName - method returns a pointer "Pool"
func (svc *PoolService) GetPoolByName(name string) (*nimbleos.Pool, error) {
	for _, v := range svc.PoolMemory {
		if *v.Name == name {
			return v, nil
		}
	}
	return nil, nil
}

// DeletePool - deletes the "Pool"
func (svc *PoolService) DeletePool(id string) error {
	if _, ok := svc.PoolMemory[id]; ok {
		delete(svc.PoolMemory, id)
	}
	return nil
}

// MergePool - merge the specified pool into the target pool. All volumes on the specified pool are moved to the target pool and the specified pool is then deleted. All the arrays in the pool are assigned to the target pool.
//   Required parameters:
//       id - ID of the specified pool.
//       targetPoolId - ID of the target pool.

//   Optional parameters:
//       force - Forcibly merge the specified pool into target pool.

func (svc *PoolService) MergePool(id string, targetPoolId string, force *bool) (*nimbleos.NsPoolMergeReturn, error) {
	if _, ok := svc.PoolMemory[id]; ok {
		delete(svc.PoolMemory, id)
	}
	return nil, nil
}
