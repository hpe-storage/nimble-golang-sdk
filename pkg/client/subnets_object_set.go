// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets
// let you create logical addressing for selective routing.
const (
	subnetPath = "subnets"
)

// SubnetObjectSet
type SubnetObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Subnet object
func (objectSet *SubnetObjectSet) CreateObject(payload *nimbleos.Subnet) (*nimbleos.Subnet, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Subnet")
}

// UpdateObject Modify existing Subnet object
func (objectSet *SubnetObjectSet) UpdateObject(id string, payload *nimbleos.Subnet) (*nimbleos.Subnet, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Subnet")
}

// DeleteObject deletes the Subnet object with the specified ID
func (objectSet *SubnetObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Subnet")
}

// GetObject returns a Subnet object with the given ID
func (objectSet *SubnetObjectSet) GetObject(id string) (*nimbleos.Subnet, error) {
	resp, err := objectSet.Client.Get(subnetPath, id, &nimbleos.Subnet{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Subnet), err
}

// GetObjectList returns the list of Subnet objects
func (objectSet *SubnetObjectSet) GetObjectList() ([]*nimbleos.Subnet, error) {
	resp, err := objectSet.Client.List(subnetPath)
	if err != nil {
		return nil, err
	}
	return buildSubnetObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Subnet objects using the given params query info
func (objectSet *SubnetObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Subnet, error) {
	subnetObjectSetResp, err := objectSet.Client.ListFromParams(subnetPath, params)
	if err != nil {
		return nil, err
	}
	return buildSubnetObjectSet(subnetObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSubnetObjectSet(response interface{}) []*nimbleos.Subnet {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Subnet, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Subnet{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
