// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// UninitializedArray Service - Lists discovered arrays that are not members of any group and are in the same subnet.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// UninitializedArrayService type 
type UninitializedArrayService struct {
	objectSet *client.UninitializedArrayObjectSet
}

// NewUninitializedArrayService - method to initialize "UninitializedArrayService" 
func NewUninitializedArrayService(gs *NsGroupService) (*UninitializedArrayService) {
	objectSet := gs.client.GetUninitializedArrayObjectSet()
	return &UninitializedArrayService{objectSet: objectSet}
}

// GetUninitializedArrays - method returns a array of pointers of type "UninitializedArrays"
func (svc *UninitializedArrayService) GetUninitializedArrays(params *util.GetParams) ([]*model.UninitializedArray, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetUninitializedArraysWithFields - method returns a array of pointers of type "UninitializedArray" 
func (svc *UninitializedArrayService) GetUninitializedArraysWithFields(fields []string) ([]*model.UninitializedArray, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateUninitializedArray - method creates a "UninitializedArray"
func (svc *UninitializedArrayService) CreateUninitializedArray(obj *model.UninitializedArray) (*model.UninitializedArray, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditUninitializedArray - method modifies  the "UninitializedArray" 
func (svc *UninitializedArrayService) EditUninitializedArray(id string, obj *model.UninitializedArray) (*model.UninitializedArray, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyUninitializedArray - private method for more than one element check. 
func onlyUninitializedArray(objs []*model.UninitializedArray) (*model.UninitializedArray, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one UninitializedArray found with the given filter")
	}

	return objs[0], nil
}

 
// GetUninitializedArraysByID - method returns associative a array of pointers of type "UninitializedArray", filter by Id
func (svc *UninitializedArrayService) GetUninitializedArraysByID(pool *model.Pool, fields []string) (map[string]*model.UninitializedArray, error) {
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
	objs, err := svc.GetUninitializedArrays(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.UninitializedArray)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetUninitializedArrayById - method returns a pointer to "UninitializedArray"
func (svc *UninitializedArrayService) GetUninitializedArrayById(id string) (*model.UninitializedArray, error) {
	return svc.objectSet.GetObject(id)
}


