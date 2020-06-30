// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Manage group wide Fibre Channel configuration.
const (
    fibreChannelConfigPath = "fibre_channel_configs"
)

// FibreChannelConfigObjectSet
type FibreChannelConfigObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new FibreChannelConfig object
func (objectSet *FibreChannelConfigObjectSet) CreateObject(payload *model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	response, err := objectSet.Client.Post(fibreChannelConfigPath, payload)
	return response.(*model.FibreChannelConfig), err
}

// UpdateObject Modify existing FibreChannelConfig object
func (objectSet *FibreChannelConfigObjectSet) UpdateObject(id string, payload *model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	response, err := objectSet.Client.Put(fibreChannelConfigPath, id, payload)
	return response.(*model.FibreChannelConfig), err
}

// DeleteObject deletes the FibreChannelConfig object with the specified ID
func (objectSet *FibreChannelConfigObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(fibreChannelConfigPath, id)
}

// GetObject returns a FibreChannelConfig object with the given ID
func (objectSet *FibreChannelConfigObjectSet) GetObject(id string) (*model.FibreChannelConfig, error) {
	response, err := objectSet.Client.Get(fibreChannelConfigPath, id, model.FibreChannelConfig{})
	if response == nil {
		return nil, err
	}
	return response.(*model.FibreChannelConfig), err
}

// GetObjectList returns the list of FibreChannelConfig objects
func (objectSet *FibreChannelConfigObjectSet) GetObjectList() ([]*model.FibreChannelConfig, error) {
	response, err := objectSet.Client.List(fibreChannelConfigPath)
	return buildFibreChannelConfigObjectSet(response), err
}

// GetObjectListFromParams returns the list of FibreChannelConfig objects using the given params query info
func (objectSet *FibreChannelConfigObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelConfig, error) {
	response, err := objectSet.Client.ListFromParams(fibreChannelConfigPath, params)
	return buildFibreChannelConfigObjectSet(response), err
}

// generated function to build the appropriate response types
func buildFibreChannelConfigObjectSet(response interface{}) ([]*model.FibreChannelConfig) {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelConfig, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelConfig{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}