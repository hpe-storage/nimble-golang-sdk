// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// Controller Service - Controller is a redundant collection of hardware capable of running the array software.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ControllerService type
type ControllerService struct {
	objectSet *client.ControllerObjectSet
}

// NewControllerService - method to initialize "ControllerService"
func NewControllerService(gs *NsGroupService) *ControllerService {
	objectSet := gs.client.GetControllerObjectSet()
	return &ControllerService{objectSet: objectSet}
}

// GetControllers - method returns a array of pointers of type "Controllers"
func (svc *ControllerService) GetControllers(params *param.GetParams) ([]*nimbleos.Controller, error) {
	controllerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return controllerResp, nil
}

// CreateController - method creates a "Controller"
func (svc *ControllerService) CreateController(obj *nimbleos.Controller) (*nimbleos.Controller, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateController: invalid parameter specified, %v", obj)
	}

	controllerResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return controllerResp, nil
}

// UpdateController - method modifies  the "Controller"
func (svc *ControllerService) UpdateController(id string, obj *nimbleos.Controller) (*nimbleos.Controller, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateController: invalid parameter specified, %v", obj)
	}

	controllerResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return controllerResp, nil
}

// GetControllerById - method returns a pointer to "Controller"
func (svc *ControllerService) GetControllerById(id string) (*nimbleos.Controller, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetControllerById: invalid parameter specified, %v", id)
	}

	controllerResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return controllerResp, nil
}

// GetControllerByName - method returns a pointer "Controller"
func (svc *ControllerService) GetControllerByName(name string) (*nimbleos.Controller, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: param.NewString(nimbleos.VolumeFields.Name),
			Operator:  param.EQUALS.String(),
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

	return controllerResp[0], nil
}

// DeleteController - deletes the "Controller"
func (svc *ControllerService) DeleteController(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteController: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
