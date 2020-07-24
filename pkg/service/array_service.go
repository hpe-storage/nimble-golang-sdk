// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Array Service - Retrieve information of specified arrays. The array is the management and configuration for the underlying physical hardware array box.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ArrayService type
type ArrayService struct {
	objectSet *client.ArrayObjectSet
}

// NewArrayService - method to initialize "ArrayService"
func NewArrayService(gs *NsGroupService) *ArrayService {
	objectSet := gs.client.GetArrayObjectSet()
	return &ArrayService{objectSet: objectSet}
}

// GetArrays - method returns a array of pointers of type "Arrays"
func (svc *ArrayService) GetArrays(params *param.GetParams) ([]*nimbleos.Array, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	arrayResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return arrayResp, nil
}

// CreateArray - method creates a "Array"
func (svc *ArrayService) CreateArray(obj *nimbleos.Array) (*nimbleos.Array, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	arrayResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return arrayResp, nil
}

// UpdateArray - method modifies  the "Array"
func (svc *ArrayService) UpdateArray(id string, obj *nimbleos.Array) (*nimbleos.Array, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	arrayResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return arrayResp, nil
}

// GetArrayById - method returns a pointer to "Array"
func (svc *ArrayService) GetArrayById(id string) (*nimbleos.Array, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	arrayResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return arrayResp, nil
}

// GetArrayByName - method returns a pointer "Array"
func (svc *ArrayService) GetArrayByName(name string) (*nimbleos.Array, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	arrayResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(arrayResp) == 0 {
		return nil, nil
	}

	return arrayResp[0], nil
}

// DeleteArray - deletes the "Array"
func (svc *ArrayService) DeleteArray(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
