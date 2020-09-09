// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelPort Service - Fibre Channel ports provide data access. This API provides the list of all Fibre Channel ports configured on the arrays.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// FibreChannelPortService type
type FibreChannelPortService struct {
	objectSet *client.FibreChannelPortObjectSet
}

// NewFibreChannelPortService - method to initialize "FibreChannelPortService"
func NewFibreChannelPortService(gs *NsGroupService) *FibreChannelPortService {
	objectSet := gs.client.GetFibreChannelPortObjectSet()
	return &FibreChannelPortService{objectSet: objectSet}
}

// GetFibreChannelPorts - method returns a array of pointers of type "FibreChannelPorts"
func (svc *FibreChannelPortService) GetFibreChannelPorts(params *param.GetParams) ([]*nimbleos.FibreChannelPort, error) {
	if params == nil {
		return nil, fmt.Errorf("GetFibreChannelPorts: invalid parameter specified, %v", params)
	}

	fibreChannelPortResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return fibreChannelPortResp, nil
}

// CreateFibreChannelPort - method creates a "FibreChannelPort"
func (svc *FibreChannelPortService) CreateFibreChannelPort(obj *nimbleos.FibreChannelPort) (*nimbleos.FibreChannelPort, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateFibreChannelPort: invalid parameter specified, %v", obj)
	}

	fibreChannelPortResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return fibreChannelPortResp, nil
}

// UpdateFibreChannelPort - method modifies  the "FibreChannelPort"
func (svc *FibreChannelPortService) UpdateFibreChannelPort(id string, obj *nimbleos.FibreChannelPort) (*nimbleos.FibreChannelPort, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateFibreChannelPort: invalid parameter specified, %v", obj)
	}

	fibreChannelPortResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return fibreChannelPortResp, nil
}

// GetFibreChannelPortById - method returns a pointer to "FibreChannelPort"
func (svc *FibreChannelPortService) GetFibreChannelPortById(id string) (*nimbleos.FibreChannelPort, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetFibreChannelPortById: invalid parameter specified, %v", id)
	}

	fibreChannelPortResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return fibreChannelPortResp, nil
}

// DeleteFibreChannelPort - deletes the "FibreChannelPort"
func (svc *FibreChannelPortService) DeleteFibreChannelPort(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteFibreChannelPort: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
