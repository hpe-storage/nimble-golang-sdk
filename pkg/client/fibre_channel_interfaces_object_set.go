// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Represent information of specified Fibre Channel interfaces. Fibre Channel interfaces are hosted on Fibre Channel ports to provide data access.
const (
	fibreChannelInterfacePath = "fibre_channel_interfaces"
)

// FibreChannelInterfaceObjectSet
type FibreChannelInterfaceObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new FibreChannelInterface object
func (objectSet *FibreChannelInterfaceObjectSet) CreateObject(payload *nimbleos.FibreChannelInterface) (*nimbleos.FibreChannelInterface, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelInterface")
}

// UpdateObject Modify existing FibreChannelInterface object
func (objectSet *FibreChannelInterfaceObjectSet) UpdateObject(id string, payload *nimbleos.FibreChannelInterface) (*nimbleos.FibreChannelInterface, error) {
	resp, err := objectSet.Client.Put(fibreChannelInterfacePath, id, payload, &nimbleos.FibreChannelInterface{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.FibreChannelInterface), err
}

// DeleteObject deletes the FibreChannelInterface object with the specified ID
func (objectSet *FibreChannelInterfaceObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on FibreChannelInterface")
}

// GetObject returns a FibreChannelInterface object with the given ID
func (objectSet *FibreChannelInterfaceObjectSet) GetObject(id string) (*nimbleos.FibreChannelInterface, error) {
	resp, err := objectSet.Client.Get(fibreChannelInterfacePath, id, &nimbleos.FibreChannelInterface{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.FibreChannelInterface), err
}

// GetObjectList returns the list of FibreChannelInterface objects
func (objectSet *FibreChannelInterfaceObjectSet) GetObjectList() ([]*nimbleos.FibreChannelInterface, error) {
	resp, err := objectSet.Client.List(fibreChannelInterfacePath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelInterfaceObjectSet(resp), err
}

// GetObjectListFromParams returns the list of FibreChannelInterface objects using the given params query info
func (objectSet *FibreChannelInterfaceObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.FibreChannelInterface, error) {
	fibreChannelInterfaceObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelInterfacePath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelInterfaceObjectSet(fibreChannelInterfaceObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelInterfaceObjectSet(response interface{}) []*nimbleos.FibreChannelInterface {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.FibreChannelInterface, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.FibreChannelInterface{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
