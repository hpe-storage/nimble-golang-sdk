// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// NetworkInterface Service - Manage per array network interface configuration.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (svc *NetworkInterfaceService) GetNetworkInterfaces(params *util.GetParams) ([]*model.NetworkInterface, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	networkInterfaceResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// CreateNetworkInterface - method creates a "NetworkInterface"
func (svc *NetworkInterfaceService) CreateNetworkInterface(obj *model.NetworkInterface) (*model.NetworkInterface, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	networkInterfaceResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// UpdateNetworkInterface - method modifies  the "NetworkInterface"
func (svc *NetworkInterfaceService) UpdateNetworkInterface(id string, obj *model.NetworkInterface) (*model.NetworkInterface, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	networkInterfaceResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// GetNetworkInterfaceById - method returns a pointer to "NetworkInterface"
func (svc *NetworkInterfaceService) GetNetworkInterfaceById(id string) (*model.NetworkInterface, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	networkInterfaceResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return networkInterfaceResp, nil
}

// GetNetworkInterfaceByName - method returns a pointer "NetworkInterface"
func (svc *NetworkInterfaceService) GetNetworkInterfaceByName(name string) (*model.NetworkInterface, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
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
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
