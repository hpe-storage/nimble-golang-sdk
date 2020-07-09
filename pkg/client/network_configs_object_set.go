// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
)

// Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.
const (
	networkConfigPath = "network_configs"
)

// NetworkConfigObjectSet
type NetworkConfigObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new NetworkConfig object
func (objectSet *NetworkConfigObjectSet) CreateObject(payload *model.NetworkConfig) (*model.NetworkConfig, error) {
	newPayload, err := model.EncodeNetworkConfig(payload)
	resp, err := objectSet.Client.Post(networkConfigPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return model.DecodeNetworkConfig(resp)
}

// UpdateObject Modify existing NetworkConfig object
func (objectSet *NetworkConfigObjectSet) UpdateObject(id string, payload *model.NetworkConfig) (*model.NetworkConfig, error) {
	newPayload, err := model.EncodeNetworkConfig(payload)
	resp, err := objectSet.Client.Put(networkConfigPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return model.DecodeNetworkConfig(resp)
}

// DeleteObject deletes the NetworkConfig object with the specified ID
func (objectSet *NetworkConfigObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(networkConfigPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a NetworkConfig object with the given ID
func (objectSet *NetworkConfigObjectSet) GetObject(id string) (*model.NetworkConfig, error) {
	networkConfigObjectSetResp, err := objectSet.Client.Get(networkConfigPath, id, model.NetworkConfig{})
	if err != nil {
		return nil, err
	}

	// null check
	if networkConfigObjectSetResp == nil {
		return nil, nil
	}
	return networkConfigObjectSetResp.(*model.NetworkConfig), err
}

// GetObjectList returns the list of NetworkConfig objects
func (objectSet *NetworkConfigObjectSet) GetObjectList() ([]*model.NetworkConfig, error) {
	networkConfigObjectSetResp, err := objectSet.Client.List(networkConfigPath)
	if err != nil {
		return nil, err
	}
	return buildNetworkConfigObjectSet(networkConfigObjectSetResp), err
}

// GetObjectListFromParams returns the list of NetworkConfig objects using the given params query info
func (objectSet *NetworkConfigObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.NetworkConfig, error) {
	networkConfigObjectSetResp, err := objectSet.Client.ListFromParams(networkConfigPath, params)
	if err != nil {
		return nil, err
	}
	return buildNetworkConfigObjectSet(networkConfigObjectSetResp), err
}

// generated function to build the appropriate response types
func buildNetworkConfigObjectSet(response interface{}) []*model.NetworkConfig {
	values := reflect.ValueOf(response)
	results := make([]*model.NetworkConfig, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.NetworkConfig{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
