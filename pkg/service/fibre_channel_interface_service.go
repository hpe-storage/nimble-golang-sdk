// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// FibreChannelInterface Service - Represent information of specified Fibre Channel interfaces. Fibre Channel interfaces are hosted on Fibre Channel ports to provide data access.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// FibreChannelInterfaceService type
type FibreChannelInterfaceService struct {
	objectSet *client.FibreChannelInterfaceObjectSet
}

// NewFibreChannelInterfaceService - method to initialize "FibreChannelInterfaceService"
func NewFibreChannelInterfaceService(gs *NsGroupService) *FibreChannelInterfaceService {
	objectSet := gs.client.GetFibreChannelInterfaceObjectSet()
	return &FibreChannelInterfaceService{objectSet: objectSet}
}

// GetFibreChannelInterfaces - method returns a array of pointers of type "FibreChannelInterfaces"
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaces(params *param.GetParams) ([]*nimbleos.FibreChannelInterface, error) {
	fibreChannelInterfaceResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return fibreChannelInterfaceResp, nil
}

// CreateFibreChannelInterface - method creates a "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) CreateFibreChannelInterface(obj *nimbleos.FibreChannelInterface) (*nimbleos.FibreChannelInterface, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateFibreChannelInterface: invalid parameter specified, %v", obj)
	}

	fibreChannelInterfaceResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return fibreChannelInterfaceResp, nil
}

// UpdateFibreChannelInterface - method modifies  the "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) UpdateFibreChannelInterface(id string, obj *nimbleos.FibreChannelInterface) (*nimbleos.FibreChannelInterface, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateFibreChannelInterface: invalid parameter specified, %v", obj)
	}

	fibreChannelInterfaceResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return fibreChannelInterfaceResp, nil
}

// GetFibreChannelInterfaceById - method returns a pointer to "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaceById(id string) (*nimbleos.FibreChannelInterface, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetFibreChannelInterfaceById: invalid parameter specified, %v", id)
	}

	fibreChannelInterfaceResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return fibreChannelInterfaceResp, nil
}

// GetFibreChannelInterfaceByName - method returns a pointer "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaceByName(name string) (*nimbleos.FibreChannelInterface, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	fibreChannelInterfaceResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(fibreChannelInterfaceResp) == 0 {
		return nil, nil
	}

	return fibreChannelInterfaceResp[0], nil
}

// DeleteFibreChannelInterface - deletes the "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) DeleteFibreChannelInterface(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteFibreChannelInterface: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
