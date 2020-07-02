// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Controller Service - Controller is a redundant collection of hardware capable of running the array software.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ControllerService type 
type ControllerService struct {
	objectSet *client.ControllerObjectSet
}

// NewControllerService - method to initialize "ControllerService" 
func NewControllerService(gs *NsGroupService) (*ControllerService) {
	objectSet := gs.client.GetControllerObjectSet()
	return &ControllerService{objectSet: objectSet}
}

// GetControllers - method returns a array of pointers of type "Controllers"
func (svc *ControllerService) GetControllers(params *util.GetParams) ([]*model.Controller, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateController - method creates a "Controller"
func (svc *ControllerService) CreateController(obj *model.Controller) (*model.Controller, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateController - method modifies  the "Controller" 
func (svc *ControllerService) UpdateController(id string, obj *model.Controller) (*model.Controller, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetControllerById - method returns a pointer to "Controller"
func (svc *ControllerService) GetControllerById(id string) (*model.Controller, error) {
	return svc.objectSet.GetObject(id)
}

// GetControllerByName - method returns a pointer "Controller" 
func (svc *ControllerService) GetControllerByName(name string) (*model.Controller, error) {
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

// DeleteController - deletes the "Controller"
func (svc *ControllerService) DeleteController(id string) error {
	return svc.objectSet.DeleteObject(id)
}
