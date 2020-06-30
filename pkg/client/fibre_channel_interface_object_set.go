// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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
	response, err := objectSet.Client.Post(fibreChannelInterfacePath, payload)
	return response.(*model.FibreChannelInterface), err
}

// UpdateObject Modify existing FibreChannelInterface object
func (objectSet *FibreChannelInterfaceObjectSet) UpdateObject(id string, payload *model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	response, err := objectSet.Client.Put(fibreChannelInterfacePath, id, payload)
	return response.(*model.FibreChannelInterface), err
}

// DeleteObject deletes the FibreChannelInterface object with the specified ID
func (objectSet *FibreChannelInterfaceObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(fibreChannelInterfacePath, id)
}

// GetObject returns a FibreChannelInterface object with the given ID
func (objectSet *FibreChannelInterfaceObjectSet) GetObject(id string) (*model.FibreChannelInterface, error) {
	response, err := objectSet.Client.Get(fibreChannelInterfacePath, id, model.FibreChannelInterface{})
	if response == nil {
		return nil, err
	}
	return response.(*model.FibreChannelInterface), err
}

// GetObjectList returns the list of FibreChannelInterface objects
func (objectSet *FibreChannelInterfaceObjectSet) GetObjectList() ([]*model.FibreChannelInterface, error) {
	response, err := objectSet.Client.List(fibreChannelInterfacePath)
	return buildFibreChannelInterfaceObjectSet(response), err
}

// GetObjectListFromParams returns the list of FibreChannelInterface objects using the given params query info
func (objectSet *FibreChannelInterfaceObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelInterface, error) {
	response, err := objectSet.Client.ListFromParams(fibreChannelInterfacePath, params)
	return buildFibreChannelInterfaceObjectSet(response), err
}

// generated function to build the appropriate response types
func buildFibreChannelInterfaceObjectSet(response interface{}) ([]*model.FibreChannelInterface) {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelInterface, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelInterface{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}