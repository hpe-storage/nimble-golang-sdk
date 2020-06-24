// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Configuration information for virtual appliance that provides highly available storage and compute.
 *
 */
const (
    hcClusterConfigPath = "hc_cluster_configs"
)

/**
 * HcClusterConfigObjectSet
*/
type HcClusterConfigObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new HcClusterConfig object
func (objectSet *HcClusterConfigObjectSet) CreateObject(payload *model.HcClusterConfig) (*model.HcClusterConfig, error) {
	response, err := objectSet.Client.Post(hcClusterConfigPath, payload)
	return response.(*model.HcClusterConfig), err
}

// UpdateObject Modify existing HcClusterConfig object
func (objectSet *HcClusterConfigObjectSet) UpdateObject(id string, payload *model.HcClusterConfig) (*model.HcClusterConfig, error) {
	response, err := objectSet.Client.Put(hcClusterConfigPath, id, payload)
	return response.(*model.HcClusterConfig), err
}

// DeleteObject deletes the HcClusterConfig object with the specified ID
func (objectSet *HcClusterConfigObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(hcClusterConfigPath, id)
}

// GetObject returns a HcClusterConfig object with the given ID
func (objectSet *HcClusterConfigObjectSet) GetObject(id string) (*model.HcClusterConfig, error) {
	response, err := objectSet.Client.Get(hcClusterConfigPath, id, model.HcClusterConfig{})
	if response == nil {
		return nil, err
	}
	return response.(*model.HcClusterConfig), err
}

// GetObjectList returns the list of HcClusterConfig objects
func (objectSet *HcClusterConfigObjectSet) GetObjectList() ([]*model.HcClusterConfig, error) {
	response, err := objectSet.Client.List(hcClusterConfigPath)
	return buildHcClusterConfigObjectSet(response), err
}

// GetObjectListFromParams returns the list of HcClusterConfig objects using the given params query info
func (objectSet *HcClusterConfigObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.HcClusterConfig, error) {
	response, err := objectSet.Client.ListFromParams(hcClusterConfigPath, params)
	return buildHcClusterConfigObjectSet(response), err
}

// generated function to build the appropriate response types
func buildHcClusterConfigObjectSet(response interface{}) ([]*model.HcClusterConfig) {
	values := reflect.ValueOf(response)
	results := make([]*model.HcClusterConfig, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.HcClusterConfig{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}