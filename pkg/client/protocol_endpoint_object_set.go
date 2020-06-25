// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Protocol endpoints are administrative logical units (LUs) in an LU conglomerate to be used with VMware Virtual Volumes.
const (
    protocolEndpointPath = "protocol_endpoints"
)


// ProtocolEndpointObjectSet
type ProtocolEndpointObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ProtocolEndpoint object
func (objectSet *ProtocolEndpointObjectSet) CreateObject(payload *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	response, err := objectSet.Client.Post(protocolEndpointPath, payload)
	return response.(*model.ProtocolEndpoint), err
}

// UpdateObject Modify existing ProtocolEndpoint object
func (objectSet *ProtocolEndpointObjectSet) UpdateObject(id string, payload *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	response, err := objectSet.Client.Put(protocolEndpointPath, id, payload)
	return response.(*model.ProtocolEndpoint), err
}

// DeleteObject deletes the ProtocolEndpoint object with the specified ID
func (objectSet *ProtocolEndpointObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(protocolEndpointPath, id)
}

// GetObject returns a ProtocolEndpoint object with the given ID
func (objectSet *ProtocolEndpointObjectSet) GetObject(id string) (*model.ProtocolEndpoint, error) {
	response, err := objectSet.Client.Get(protocolEndpointPath, id, model.ProtocolEndpoint{})
	if response == nil {
		return nil, err
	}
	return response.(*model.ProtocolEndpoint), err
}

// GetObjectList returns the list of ProtocolEndpoint objects
func (objectSet *ProtocolEndpointObjectSet) GetObjectList() ([]*model.ProtocolEndpoint, error) {
	response, err := objectSet.Client.List(protocolEndpointPath)
	return buildProtocolEndpointObjectSet(response), err
}

// GetObjectListFromParams returns the list of ProtocolEndpoint objects using the given params query info
func (objectSet *ProtocolEndpointObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ProtocolEndpoint, error) {
	response, err := objectSet.Client.ListFromParams(protocolEndpointPath, params)
	return buildProtocolEndpointObjectSet(response), err
}

// generated function to build the appropriate response types
func buildProtocolEndpointObjectSet(response interface{}) ([]*model.ProtocolEndpoint) {
	values := reflect.ValueOf(response)
	results := make([]*model.ProtocolEndpoint, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ProtocolEndpoint{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}