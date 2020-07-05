// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Controller Service - Controller is a redundant collection of hardware capable of running the array software.

import (
	"fmt"
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
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	controllerResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return controllerResp,nil
}

// CreateController - method creates a "Controller"
func (svc *ControllerService) CreateController(obj *model.Controller) (*model.Controller, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	controllerResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return controllerResp,nil
}

// UpdateController - method modifies  the "Controller" 
func (svc *ControllerService) UpdateController(id string, obj *model.Controller) (*model.Controller, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	controllerResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return controllerResp,nil
}

// GetControllerById - method returns a pointer to "Controller"
func (svc *ControllerService) GetControllerById(id string) (*model.Controller, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	controllerResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return controllerResp,nil
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
	controllerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(controllerResp) == 0 {
    	return nil, nil
    }
    
	return controllerResp[0],nil
}	

// DeleteController - deletes the "Controller"
func (svc *ControllerService) DeleteController(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
