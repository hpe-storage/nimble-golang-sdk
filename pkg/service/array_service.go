// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Array Service - Retrieve information of specified arrays. The array is the management and configuration for the underlying physical hardware array box.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateArray - method creates a "Array"
func (svc *ArrayService) CreateArray(obj *model.Array) (*model.Array, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateArray - method modifies  the "Array" 
func (svc *ArrayService) UpdateArray(id string, obj *model.Array) (*model.Array, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetArrayById - method returns a pointer to "Array"
func (svc *ArrayService) GetArrayById(id string) (*model.Array, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteArray - deletes the "Array"
func (svc *ArrayService) DeleteArray(id string) error {
	return svc.objectSet.DeleteObject(id)
}
