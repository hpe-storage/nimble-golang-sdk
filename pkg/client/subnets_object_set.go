// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *SubnetObjectSet) CreateObject(payload *model.Subnet) (*model.Subnet, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Subnet")
}

// UpdateObject Modify existing Subnet object
func (objectSet *SubnetObjectSet) UpdateObject(id string, payload *model.Subnet) (*model.Subnet, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Subnet")
}

// DeleteObject deletes the Subnet object with the specified ID
func (objectSet *SubnetObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Subnet")
}

// GetObject returns a Subnet object with the given ID
func (objectSet *SubnetObjectSet) GetObject(id string) (*model.Subnet, error) {
	subnetObjectSetResp, err := objectSet.Client.Get(subnetPath, id, model.Subnet{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if subnetObjectSetResp == nil {
		return nil,nil
	}
	return subnetObjectSetResp.(*model.Subnet), err
}

// GetObjectList returns the list of Subnet objects
func (objectSet *SubnetObjectSet) GetObjectList() ([]*model.Subnet, error) {
	subnetObjectSetResp, err := objectSet.Client.List(subnetPath)
	if err != nil {
		return nil, err
	}
	return buildSubnetObjectSet(subnetObjectSetResp), err
}

// GetObjectListFromParams returns the list of Subnet objects using the given params query info
func (objectSet *SubnetObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Subnet, error) {
	subnetObjectSetResp, err := objectSet.Client.ListFromParams(subnetPath, params)
	if err != nil {
		return nil, err
	}
	return buildSubnetObjectSet(subnetObjectSetResp), err
}
// generated function to build the appropriate response types
func buildSubnetObjectSet(response interface{}) ([]*model.Subnet) {
	values := reflect.ValueOf(response)
	results := make([]*model.Subnet, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Subnet{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
