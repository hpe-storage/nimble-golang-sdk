// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (objectSet *NetworkConfigObjectSet) CreateObject(payload *nimbleos.NetworkConfig) (*nimbleos.NetworkConfig, error) {
	resp, err := objectSet.Client.Post(networkConfigPath, payload, &nimbleos.NetworkConfig{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NetworkConfig), err
}

// UpdateObject Modify existing NetworkConfig object
func (objectSet *NetworkConfigObjectSet) UpdateObject(id string, payload *nimbleos.NetworkConfig) (*nimbleos.NetworkConfig, error) {
	resp, err := objectSet.Client.Put(networkConfigPath, id, payload, &nimbleos.NetworkConfig{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NetworkConfig), err
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
func (objectSet *NetworkConfigObjectSet) GetObject(id string) (*nimbleos.NetworkConfig, error) {
	resp, err := objectSet.Client.Get(networkConfigPath, id, nimbleos.NetworkConfig{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.NetworkConfig), err
}

// GetObjectList returns the list of NetworkConfig objects
func (objectSet *NetworkConfigObjectSet) GetObjectList() ([]*nimbleos.NetworkConfig, error) {
	resp, err := objectSet.Client.List(networkConfigPath)
	if err != nil {
		return nil, err
	}
	return buildNetworkConfigObjectSet(resp), err
}

// GetObjectListFromParams returns the list of NetworkConfig objects using the given params query info
func (objectSet *NetworkConfigObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.NetworkConfig, error) {
	networkConfigObjectSetResp, err := objectSet.Client.ListFromParams(networkConfigPath, params)
	if err != nil {
		return nil, err
	}
	return buildNetworkConfigObjectSet(networkConfigObjectSetResp), err
}

// generated function to build the appropriate response types
func buildNetworkConfigObjectSet(response interface{}) []*nimbleos.NetworkConfig {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.NetworkConfig, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.NetworkConfig{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// ActivateNetconfig - Activate a network configuration.
func (objectSet *NetworkConfigObjectSet) ActivateNetconfigObjectSet(id *string, ignoreValidationMask *uint64) error {

	activateNetconfigUri := networkConfigPath
	activateNetconfigUri = activateNetconfigUri + "/" + *id
	activateNetconfigUri = activateNetconfigUri + "/actions/" + "activate_netconfig"

	payload := &struct {
		Id                   *string `json:"id,omitempty"`
		IgnoreValidationMask *uint64 `json:"ignore_validation_mask,omitempty"`
	}{
		id,
		ignoreValidationMask,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(activateNetconfigUri, payload, &emptyStruct)
	return err
}

// ValidateNetconfig - Validate a network configuration.
func (objectSet *NetworkConfigObjectSet) ValidateNetconfigObjectSet(id *string, ignoreValidationMask *uint64) error {

	validateNetconfigUri := networkConfigPath
	validateNetconfigUri = validateNetconfigUri + "/" + *id
	validateNetconfigUri = validateNetconfigUri + "/actions/" + "validate_netconfig"

	payload := &struct {
		Id                   *string `json:"id,omitempty"`
		IgnoreValidationMask *uint64 `json:"ignore_validation_mask,omitempty"`
	}{
		id,
		ignoreValidationMask,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(validateNetconfigUri, payload, &emptyStruct)
	return err
}
