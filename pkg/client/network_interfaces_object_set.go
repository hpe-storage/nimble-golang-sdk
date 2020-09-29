// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *NetworkInterfaceObjectSet) CreateObject(payload *nimbleos.NetworkInterface) (*nimbleos.NetworkInterface, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on NetworkInterface")
}

// UpdateObject Modify existing NetworkInterface object
func (objectSet *NetworkInterfaceObjectSet) UpdateObject(id string, payload *nimbleos.NetworkInterface) (*nimbleos.NetworkInterface, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on NetworkInterface")
}

// DeleteObject deletes the NetworkInterface object with the specified ID
func (objectSet *NetworkInterfaceObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on NetworkInterface")
}

// GetObject returns a NetworkInterface object with the given ID
func (objectSet *NetworkInterfaceObjectSet) GetObject(id string) (*nimbleos.NetworkInterface, error) {
	resp, err := objectSet.Client.Get(networkInterfacePath, id, &nimbleos.NetworkInterface{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.NetworkInterface), err
}

// GetObjectList returns the list of NetworkInterface objects
func (objectSet *NetworkInterfaceObjectSet) GetObjectList() ([]*nimbleos.NetworkInterface, error) {
	resp, err := objectSet.Client.List(networkInterfacePath)
	if err != nil {
		return nil, err
	}
	return buildNetworkInterfaceObjectSet(resp), err
}

// GetObjectListFromParams returns the list of NetworkInterface objects using the given params query info
func (objectSet *NetworkInterfaceObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.NetworkInterface, error) {
	networkInterfaceObjectSetResp, err := objectSet.Client.ListFromParams(networkInterfacePath, params)
	if err != nil {
		return nil, err
	}
	return buildNetworkInterfaceObjectSet(networkInterfaceObjectSetResp), err
}

// generated function to build the appropriate response types
func buildNetworkInterfaceObjectSet(response interface{}) []*nimbleos.NetworkInterface {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.NetworkInterface, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.NetworkInterface{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
