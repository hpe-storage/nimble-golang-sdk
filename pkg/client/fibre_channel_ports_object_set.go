// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Fibre Channel ports provide data access. This API provides the list of all Fibre Channel ports configured on the arrays.
const (
	fibreChannelPortPath = "fibre_channel_ports"
)

// FibreChannelPortObjectSet
type FibreChannelPortObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new FibreChannelPort object
func (objectSet *FibreChannelPortObjectSet) CreateObject(payload *nimbleos.FibreChannelPort) (*nimbleos.FibreChannelPort, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelPort")
}

// UpdateObject Modify existing FibreChannelPort object
func (objectSet *FibreChannelPortObjectSet) UpdateObject(id string, payload *nimbleos.FibreChannelPort) (*nimbleos.FibreChannelPort, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on FibreChannelPort")
}

// DeleteObject deletes the FibreChannelPort object with the specified ID
func (objectSet *FibreChannelPortObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on FibreChannelPort")
}

// GetObject returns a FibreChannelPort object with the given ID
func (objectSet *FibreChannelPortObjectSet) GetObject(id string) (*nimbleos.FibreChannelPort, error) {
	resp, err := objectSet.Client.Get(fibreChannelPortPath, id, nimbleos.FibreChannelPort{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.FibreChannelPort), err
}

// GetObjectList returns the list of FibreChannelPort objects
func (objectSet *FibreChannelPortObjectSet) GetObjectList() ([]*nimbleos.FibreChannelPort, error) {
	resp, err := objectSet.Client.List(fibreChannelPortPath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelPortObjectSet(resp), err
}

// GetObjectListFromParams returns the list of FibreChannelPort objects using the given params query info
func (objectSet *FibreChannelPortObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.FibreChannelPort, error) {
	fibreChannelPortObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelPortPath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelPortObjectSet(fibreChannelPortObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelPortObjectSet(response interface{}) []*nimbleos.FibreChannelPort {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.FibreChannelPort, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.FibreChannelPort{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
