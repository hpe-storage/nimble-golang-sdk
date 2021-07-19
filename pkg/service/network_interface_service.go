// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// NetworkInterface Service - Manage per array network interface configuration.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// NetworkInterfaceService type
type NetworkInterfaceService struct {
	objectSet *client.NetworkInterfaceObjectSet
}

// NewNetworkInterfaceService - method to initialize "NetworkInterfaceService"
func NewNetworkInterfaceService(gs *NsGroupService) *NetworkInterfaceService {
	objectSet := gs.client.GetNetworkInterfaceObjectSet()
	return &NetworkInterfaceService{objectSet: objectSet}
}

// GetNetworkInterfaces - method returns a array of pointers of type "NetworkInterfaces"
func (svc *NetworkInterfaceService) GetNetworkInterfaces(params *param.GetParams) ([]*nimbleos.NetworkInterface, error) {
	networkInterfaceResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// CreateNetworkInterface - method creates a "NetworkInterface"
func (svc *NetworkInterfaceService) CreateNetworkInterface(obj *nimbleos.NetworkInterface) (*nimbleos.NetworkInterface, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateNetworkInterface: invalid parameter specified, %v", obj)
	}

	networkInterfaceResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// UpdateNetworkInterface - method modifies  the "NetworkInterface"
func (svc *NetworkInterfaceService) UpdateNetworkInterface(id string, obj *nimbleos.NetworkInterface) (*nimbleos.NetworkInterface, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateNetworkInterface: invalid parameter specified, %v", obj)
	}

	networkInterfaceResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// GetNetworkInterfaceById - method returns a pointer to "NetworkInterface"
func (svc *NetworkInterfaceService) GetNetworkInterfaceById(id string) (*nimbleos.NetworkInterface, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetNetworkInterfaceById: invalid parameter specified, %v", id)
	}

	networkInterfaceResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// GetNetworkInterfaceByName - method returns a pointer "NetworkInterface"
func (svc *NetworkInterfaceService) GetNetworkInterfaceByName(name string) (*nimbleos.NetworkInterface, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	networkInterfaceResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(networkInterfaceResp) == 0 {
		return nil, nil
	}

	return networkInterfaceResp[0], nil
}

// DeleteNetworkInterface - deletes the "NetworkInterface"
func (svc *NetworkInterfaceService) DeleteNetworkInterface(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteNetworkInterface: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
