// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Pool Service - Manage pools. Pools are an aggregation of arrays.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// PoolService type 
type PoolService struct {
	objectSet *client.PoolObjectSet
}

// NewPoolService - method to initialize "PoolService" 
func NewPoolService(gs *NsGroupService) (*PoolService) {
	objectSet := gs.client.GetPoolObjectSet()
	return &PoolService{objectSet: objectSet}
}

// GetPools - method returns a array of pointers of type "Pools"
func (svc *PoolService) GetPools(params *util.GetParams) ([]*model.Pool, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreatePool - method creates a "Pool"
func (svc *PoolService) CreatePool(obj *model.Pool) (*model.Pool, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdatePool - method modifies  the "Pool" 
func (svc *PoolService) UpdatePool(id string, obj *model.Pool) (*model.Pool, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetPoolById - method returns a pointer to "Pool"
func (svc *PoolService) GetPoolById(id string) (*model.Pool, error) {
	return svc.objectSet.GetObject(id)
}

// GetPoolByName - method returns a pointer "Pool" 
func (svc *PoolService) GetPoolByName(name string) (*model.Pool, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeletePool - deletes the "Pool"
func (svc *PoolService) DeletePool(id string) error {
	return svc.objectSet.DeleteObject(id)
}
