// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Array Service - Retrieve information of specified arrays. The array is the management and configuration for the underlying physical hardware array box.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// ArrayService type 
type ArrayService struct {
	objectSet *client.ArrayObjectSet
}

// NewArrayService - method to initialize "ArrayService" 
func NewArrayService(gs *NsGroupService) (*ArrayService) {
	objectSet := gs.client.GetArrayObjectSet()
	return &ArrayService{objectSet: objectSet}
}

// GetArrays - method returns a array of pointers of type "Arrays"
func (svc *ArrayService) GetArrays(params *util.GetParams) ([]*model.Array, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetArraysWithFields - method returns a array of pointers of type "Array" 
func (svc *ArrayService) GetArraysWithFields(fields []string) ([]*model.Array, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateArray - method creates a "Array"
func (svc *ArrayService) CreateArray(obj *model.Array) (*model.Array, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditArray - method modifies  the "Array" 
func (svc *ArrayService) EditArray(id string, obj *model.Array) (*model.Array, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyArray - private method for more than one element check. 
func onlyArray(objs []*model.Array) (*model.Array, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Array found with the given filter")
	}

	return objs[0], nil
}

 
// GetArraysByID - method returns associative a array of pointers of type "Array", filter by Id
func (svc *ArrayService) GetArraysByID(pool *model.Pool, fields []string) (map[string]*model.Array, error) {
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
	objs, err := svc.GetArrays(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Array)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetArrayById - method returns a pointer to "Array"
func (svc *ArrayService) GetArrayById(id string) (*model.Array, error) {
	return svc.objectSet.GetObject(id)
}

// GetArraysByName - method returns a associative array of pointers of type "Array", filter by name 
func (svc *ArrayService) GetArraysByName(pool *model.Pool, fields []string) (map[string]*model.Array, error) {
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
	objs, err := svc.GetArrays(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Array)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetArrayByName - method returns a pointer "Array" 
func (svc *ArrayService) GetArrayByName(name string) (*model.Array, error) {
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
	return onlyArray(objs)
}	

