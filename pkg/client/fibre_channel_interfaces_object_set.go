// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *FibreChannelInterfaceObjectSet) CreateObject(payload *model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelInterface")
}

// UpdateObject Modify existing FibreChannelInterface object
func (objectSet *FibreChannelInterfaceObjectSet) UpdateObject(id string, payload *model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	newPayload, err := model.EncodeFibreChannelInterface(payload)
	resp, err := objectSet.Client.Put(fibreChannelInterfacePath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return model.DecodeFibreChannelInterface(resp)
}

// DeleteObject deletes the FibreChannelInterface object with the specified ID
func (objectSet *FibreChannelInterfaceObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on FibreChannelInterface")
}

// GetObject returns a FibreChannelInterface object with the given ID
func (objectSet *FibreChannelInterfaceObjectSet) GetObject(id string) (*model.FibreChannelInterface, error) {
	fibreChannelInterfaceObjectSetResp, err := objectSet.Client.Get(fibreChannelInterfacePath, id, model.FibreChannelInterface{})
	if err != nil {
		return nil, err
	}

	// null check
	if fibreChannelInterfaceObjectSetResp == nil {
		return nil, nil
	}
	return fibreChannelInterfaceObjectSetResp.(*model.FibreChannelInterface), err
}

// GetObjectList returns the list of FibreChannelInterface objects
func (objectSet *FibreChannelInterfaceObjectSet) GetObjectList() ([]*model.FibreChannelInterface, error) {
	fibreChannelInterfaceObjectSetResp, err := objectSet.Client.List(fibreChannelInterfacePath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelInterfaceObjectSet(fibreChannelInterfaceObjectSetResp), err
}

// GetObjectListFromParams returns the list of FibreChannelInterface objects using the given params query info
func (objectSet *FibreChannelInterfaceObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelInterface, error) {
	fibreChannelInterfaceObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelInterfacePath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelInterfaceObjectSet(fibreChannelInterfaceObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelInterfaceObjectSet(response interface{}) []*model.FibreChannelInterface {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelInterface, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelInterface{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
