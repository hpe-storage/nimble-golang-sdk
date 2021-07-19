// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// Pool Service - Manage pools. Pools are an aggregation of arrays.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// PoolService type
type PoolService struct {
	objectSet *client.PoolObjectSet
}

// NewPoolService - method to initialize "PoolService"
func NewPoolService(gs *NsGroupService) *PoolService {
	objectSet := gs.client.GetPoolObjectSet()
	return &PoolService{objectSet: objectSet}
}

// GetPools - method returns a array of pointers of type "Pools"
func (svc *PoolService) GetPools(params *param.GetParams) ([]*nimbleos.Pool, error) {
	poolResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// CreatePool - method creates a "Pool"
func (svc *PoolService) CreatePool(obj *nimbleos.Pool) (*nimbleos.Pool, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreatePool: invalid parameter specified, %v", obj)
	}

	poolResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// UpdatePool - method modifies  the "Pool"
func (svc *PoolService) UpdatePool(id string, obj *nimbleos.Pool) (*nimbleos.Pool, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdatePool: invalid parameter specified, %v", obj)
	}

	poolResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// GetPoolById - method returns a pointer to "Pool"
func (svc *PoolService) GetPoolById(id string) (*nimbleos.Pool, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetPoolById: invalid parameter specified, %v", id)
	}

	poolResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return poolResp, nil
}

// GetPoolByName - method returns a pointer "Pool"
func (svc *PoolService) GetPoolByName(name string) (*nimbleos.Pool, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	poolResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(poolResp) == 0 {
		return nil, nil
	}

	return poolResp[0], nil
}

// DeletePool - deletes the "Pool"
func (svc *PoolService) DeletePool(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeletePool: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
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

	if len(id) == 0 || len(targetPoolId) == 0 {
		return nil, fmt.Errorf("MergePool: invalid parameter specified id: %v, targetPoolId: %v ", id, targetPoolId)
	}

	resp, err := svc.objectSet.Merge(&id, &targetPoolId, force)
	return resp, err
}
