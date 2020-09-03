// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (objectSet *FibreChannelConfigObjectSet) CreateObject(payload *nimbleos.FibreChannelConfig) (*nimbleos.FibreChannelConfig, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelConfig")
}

// UpdateObject Modify existing FibreChannelConfig object
func (objectSet *FibreChannelConfigObjectSet) UpdateObject(id string, payload *nimbleos.FibreChannelConfig) (*nimbleos.FibreChannelConfig, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on FibreChannelConfig")
}

// DeleteObject deletes the FibreChannelConfig object with the specified ID
func (objectSet *FibreChannelConfigObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on FibreChannelConfig")
}

// GetObject returns a FibreChannelConfig object with the given ID
func (objectSet *FibreChannelConfigObjectSet) GetObject(id string) (*nimbleos.FibreChannelConfig, error) {
	resp, err := objectSet.Client.Get(fibreChannelConfigPath, id, nimbleos.FibreChannelConfig{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.FibreChannelConfig), err
}

// GetObjectList returns the list of FibreChannelConfig objects
func (objectSet *FibreChannelConfigObjectSet) GetObjectList() ([]*nimbleos.FibreChannelConfig, error) {
	resp, err := objectSet.Client.List(fibreChannelConfigPath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelConfigObjectSet(resp), err
}

// GetObjectListFromParams returns the list of FibreChannelConfig objects using the given params query info
func (objectSet *FibreChannelConfigObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.FibreChannelConfig, error) {
	fibreChannelConfigObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelConfigPath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelConfigObjectSet(fibreChannelConfigObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelConfigObjectSet(response interface{}) []*nimbleos.FibreChannelConfig {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.FibreChannelConfig, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.FibreChannelConfig{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Regenerate - Regenerate Fibre Channel configuration.
func (objectSet *FibreChannelConfigObjectSet) RegenerateObjectSet(id *string, wwnnBaseStr *string, precheck *bool) (*nimbleos.NsFcConfigRegenerateReturn, error) {

	regenerateUri := fibreChannelConfigPath
	regenerateUri = regenerateUri + "/" + *id
	regenerateUri = regenerateUri + "/actions/" + "regenerate"

	payload := &struct {
		Id          *string `json:"id,omitempty"`
		WwnnBaseStr *string `json:"wwnn_base_str,omitempty"`
		Precheck    *bool   `json:"precheck,omitempty"`
	}{
		id,
		wwnnBaseStr,
		precheck,
	}

	resp, err := objectSet.Client.Post(regenerateUri, payload, &nimbleos.NsFcConfigRegenerateReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsFcConfigRegenerateReturn), err
}

// HwUpgrade - Update Fibre Channel configuration after hardware changes.
func (objectSet *FibreChannelConfigObjectSet) HwUpgradeObjectSet(id *string) error {

	hwUpgradeUri := fibreChannelConfigPath
	hwUpgradeUri = hwUpgradeUri + "/" + *id
	hwUpgradeUri = hwUpgradeUri + "/actions/" + "hw_upgrade"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(hwUpgradeUri, payload, &emptyStruct)
	return err
}
