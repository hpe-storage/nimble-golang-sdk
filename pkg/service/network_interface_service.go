// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// NetworkInterface Service - Manage per array network interface configuration.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// NetworkInterfaceService type 
type NetworkInterfaceService struct {
	objectSet *client.NetworkInterfaceObjectSet
}

// NewNetworkInterfaceService - method to initialize "NetworkInterfaceService" 
func NewNetworkInterfaceService(gs *NsGroupService) (*NetworkInterfaceService) {
	objectSet := gs.client.GetNetworkInterfaceObjectSet()
	return &NetworkInterfaceService{objectSet: objectSet}
}

// GetNetworkInterfaces - method returns a array of pointers of type "NetworkInterfaces"
func (svc *NetworkInterfaceService) GetNetworkInterfaces(params *util.GetParams) ([]*model.NetworkInterface, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateNetworkInterface - method creates a "NetworkInterface"
func (svc *NetworkInterfaceService) CreateNetworkInterface(obj *model.NetworkInterface) (*model.NetworkInterface, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateNetworkInterface - method modifies  the "NetworkInterface" 
func (svc *NetworkInterfaceService) UpdateNetworkInterface(id string, obj *model.NetworkInterface) (*model.NetworkInterface, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetNetworkInterfaceById - method returns a pointer to "NetworkInterface"
func (svc *NetworkInterfaceService) GetNetworkInterfaceById(id string) (*model.NetworkInterface, error) {
	return svc.objectSet.GetObject(id)
}

// GetNetworkInterfaceByName - method returns a pointer "NetworkInterface" 
func (svc *NetworkInterfaceService) GetNetworkInterfaceByName(name string) (*model.NetworkInterface, error) {
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

// DeleteNetworkInterface - deletes the "NetworkInterface"
func (svc *NetworkInterfaceService) DeleteNetworkInterface(id string) error {
	return svc.objectSet.DeleteObject(id)
}
