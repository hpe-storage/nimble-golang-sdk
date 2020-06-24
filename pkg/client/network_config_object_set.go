package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.
 *
 */
const (
    networkConfigPath = "network_configs"
)

/**
 * NetworkConfigObjectSet
*/
type NetworkConfigObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new NetworkConfig object
func (objectSet *NetworkConfigObjectSet) CreateObject(payload *model.NetworkConfig) (*model.NetworkConfig, error) {
	response, err := objectSet.Client.Post(networkConfigPath, payload)
	return response.(*model.NetworkConfig), err
}

// UpdateObject Modify existing NetworkConfig object
func (objectSet *NetworkConfigObjectSet) UpdateObject(id string, payload *model.NetworkConfig) (*model.NetworkConfig, error) {
	response, err := objectSet.Client.Put(networkConfigPath, id, payload)
	return response.(*model.NetworkConfig), err
}

// DeleteObject deletes the NetworkConfig object with the specified ID
func (objectSet *NetworkConfigObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(networkConfigPath, id)
}

// GetObject returns a NetworkConfig object with the given ID
func (objectSet *NetworkConfigObjectSet) GetObject(id string) (*model.NetworkConfig, error) {
	response, err := objectSet.Client.Get(networkConfigPath, id, model.NetworkConfig{})
	if response == nil {
		return nil, err
	}
	return response.(*model.NetworkConfig), err
}

// GetObjectList returns the list of NetworkConfig objects
func (objectSet *NetworkConfigObjectSet) GetObjectList() ([]*model.NetworkConfig, error) {
	response, err := objectSet.Client.List(networkConfigPath)
	return buildNetworkConfigObjectSet(response), err
}

// GetObjectListFromParams returns the list of NetworkConfig objects using the given params query info
func (objectSet *NetworkConfigObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.NetworkConfig, error) {
	response, err := objectSet.Client.ListFromParams(networkConfigPath, params)
	return buildNetworkConfigObjectSet(response), err
}

// generated function to build the appropriate response types
func buildNetworkConfigObjectSet(response interface{}) ([]*model.NetworkConfig) {
	values := reflect.ValueOf(response)
	results := make([]*model.NetworkConfig, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.NetworkConfig{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}