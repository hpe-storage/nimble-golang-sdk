// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
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
	return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelConfig")
}

// UpdateObject Modify existing FibreChannelConfig object
func (objectSet *FibreChannelConfigObjectSet) UpdateObject(id string, payload *model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on FibreChannelConfig")
}

// DeleteObject deletes the FibreChannelConfig object with the specified ID
func (objectSet *FibreChannelConfigObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on FibreChannelConfig")
}

// GetObject returns a FibreChannelConfig object with the given ID
func (objectSet *FibreChannelConfigObjectSet) GetObject(id string) (*model.FibreChannelConfig, error) {
	fibreChannelConfigObjectSetResp, err := objectSet.Client.Get(fibreChannelConfigPath, id, model.FibreChannelConfig{})
	if err != nil {
		return nil, err
	}

	// null check
	if fibreChannelConfigObjectSetResp == nil {
		return nil, nil
	}
	return fibreChannelConfigObjectSetResp.(*model.FibreChannelConfig), err
}

// GetObjectList returns the list of FibreChannelConfig objects
func (objectSet *FibreChannelConfigObjectSet) GetObjectList() ([]*model.FibreChannelConfig, error) {
	fibreChannelConfigObjectSetResp, err := objectSet.Client.List(fibreChannelConfigPath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelConfigObjectSet(fibreChannelConfigObjectSetResp), err
}

// GetObjectListFromParams returns the list of FibreChannelConfig objects using the given params query info
func (objectSet *FibreChannelConfigObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelConfig, error) {
	fibreChannelConfigObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelConfigPath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelConfigObjectSet(fibreChannelConfigObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelConfigObjectSet(response interface{}) []*model.FibreChannelConfig {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelConfig, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelConfig{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
