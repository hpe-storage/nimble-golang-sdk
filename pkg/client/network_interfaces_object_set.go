// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manage per array network interface configuration.
const (
    networkInterfacePath = "network_interfaces"
)

// NetworkInterfaceObjectSet
type NetworkInterfaceObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new NetworkInterface object
func (objectSet *NetworkInterfaceObjectSet) CreateObject(payload *model.NetworkInterface) (*model.NetworkInterface, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on NetworkInterface")
}

// UpdateObject Modify existing NetworkInterface object
func (objectSet *NetworkInterfaceObjectSet) UpdateObject(id string, payload *model.NetworkInterface) (*model.NetworkInterface, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on NetworkInterface")
}

// DeleteObject deletes the NetworkInterface object with the specified ID
func (objectSet *NetworkInterfaceObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on NetworkInterface")
}

// GetObject returns a NetworkInterface object with the given ID
func (objectSet *NetworkInterfaceObjectSet) GetObject(id string) (*model.NetworkInterface, error) {
	networkInterfaceObjectSetResp, err := objectSet.Client.Get(networkInterfacePath, id, model.NetworkInterface{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if networkInterfaceObjectSetResp == nil {
		return nil,nil
	}
	return networkInterfaceObjectSetResp.(*model.NetworkInterface), err
}

// GetObjectList returns the list of NetworkInterface objects
func (objectSet *NetworkInterfaceObjectSet) GetObjectList() ([]*model.NetworkInterface, error) {
	networkInterfaceObjectSetResp, err := objectSet.Client.List(networkInterfacePath)
	if err != nil {
		return nil, err
	}
	return buildNetworkInterfaceObjectSet(networkInterfaceObjectSetResp), err
}

// GetObjectListFromParams returns the list of NetworkInterface objects using the given params query info
func (objectSet *NetworkInterfaceObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.NetworkInterface, error) {
	networkInterfaceObjectSetResp, err := objectSet.Client.ListFromParams(networkInterfacePath, params)
	if err != nil {
		return nil, err
	}
	return buildNetworkInterfaceObjectSet(networkInterfaceObjectSetResp), err
}
// generated function to build the appropriate response types
func buildNetworkInterfaceObjectSet(response interface{}) ([]*model.NetworkInterface) {
	values := reflect.ValueOf(response)
	results := make([]*model.NetworkInterface, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.NetworkInterface{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
