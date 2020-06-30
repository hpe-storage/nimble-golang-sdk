// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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
	response, err := objectSet.Client.Post(networkInterfacePath, payload)
	return response.(*model.NetworkInterface), err
}

// UpdateObject Modify existing NetworkInterface object
func (objectSet *NetworkInterfaceObjectSet) UpdateObject(id string, payload *model.NetworkInterface) (*model.NetworkInterface, error) {
	response, err := objectSet.Client.Put(networkInterfacePath, id, payload)
	return response.(*model.NetworkInterface), err
}

// DeleteObject deletes the NetworkInterface object with the specified ID
func (objectSet *NetworkInterfaceObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(networkInterfacePath, id)
}

// GetObject returns a NetworkInterface object with the given ID
func (objectSet *NetworkInterfaceObjectSet) GetObject(id string) (*model.NetworkInterface, error) {
	response, err := objectSet.Client.Get(networkInterfacePath, id, model.NetworkInterface{})
	if response == nil {
		return nil, err
	}
	return response.(*model.NetworkInterface), err
}

// GetObjectList returns the list of NetworkInterface objects
func (objectSet *NetworkInterfaceObjectSet) GetObjectList() ([]*model.NetworkInterface, error) {
	response, err := objectSet.Client.List(networkInterfacePath)
	return buildNetworkInterfaceObjectSet(response), err
}

// GetObjectListFromParams returns the list of NetworkInterface objects using the given params query info
func (objectSet *NetworkInterfaceObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.NetworkInterface, error) {
	response, err := objectSet.Client.ListFromParams(networkInterfacePath, params)
	return buildNetworkInterfaceObjectSet(response), err
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