// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *FibreChannelPortObjectSet) CreateObject(payload *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelPort")
}

// UpdateObject Modify existing FibreChannelPort object
func (objectSet *FibreChannelPortObjectSet) UpdateObject(id string, payload *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on FibreChannelPort")
}

// DeleteObject deletes the FibreChannelPort object with the specified ID
func (objectSet *FibreChannelPortObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on FibreChannelPort")
}

// GetObject returns a FibreChannelPort object with the given ID
func (objectSet *FibreChannelPortObjectSet) GetObject(id string) (*model.FibreChannelPort, error) {
	fibreChannelPortObjectSetResp, err := objectSet.Client.Get(fibreChannelPortPath, id, model.FibreChannelPort{})
	if err != nil {
		return nil, err
	}

	// null check
	if fibreChannelPortObjectSetResp == nil {
		return nil, nil
	}
	return fibreChannelPortObjectSetResp.(*model.FibreChannelPort), err
}

// GetObjectList returns the list of FibreChannelPort objects
func (objectSet *FibreChannelPortObjectSet) GetObjectList() ([]*model.FibreChannelPort, error) {
	fibreChannelPortObjectSetResp, err := objectSet.Client.List(fibreChannelPortPath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelPortObjectSet(fibreChannelPortObjectSetResp), err
}

// GetObjectListFromParams returns the list of FibreChannelPort objects using the given params query info
func (objectSet *FibreChannelPortObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelPort, error) {
	fibreChannelPortObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelPortPath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelPortObjectSet(fibreChannelPortObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelPortObjectSet(response interface{}) []*model.FibreChannelPort {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelPort, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelPort{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
