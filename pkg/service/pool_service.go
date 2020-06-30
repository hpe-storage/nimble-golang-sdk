// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Pool Service - Manage pools. Pools are an aggregation of arrays.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetPoolsWithFields - method returns a array of pointers of type "Pool" 
func (svc *PoolService) GetPoolsWithFields(fields []string) ([]*model.Pool, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreatePool - method creates a "Pool"
func (svc *PoolService) CreatePool(obj *model.Pool) (*model.Pool, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditPool - method modifies  the "Pool" 
func (svc *PoolService) EditPool(id string, obj *model.Pool) (*model.Pool, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyPool - private method for more than one element check. 
func onlyPool(objs []*model.Pool) (*model.Pool, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Pool found with the given filter")
	}

	return objs[0], nil
}

 
// GetPoolsByID - method returns associative a array of pointers of type "Pool", filter by Id
func (svc *PoolService) GetPoolsByID(pool *model.Pool, fields []string) (map[string]*model.Pool, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetPools(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Pool)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetPoolById - method returns a pointer to "Pool"
func (svc *PoolService) GetPoolById(id string) (*model.Pool, error) {
	return svc.objectSet.GetObject(id)
}

// GetPoolsByName - method returns a associative array of pointers of type "Pool", filter by name 
func (svc *PoolService) GetPoolsByName(pool *model.Pool, fields []string) (map[string]*model.Pool, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetPools(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Pool)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlyPool(objs)
}	

