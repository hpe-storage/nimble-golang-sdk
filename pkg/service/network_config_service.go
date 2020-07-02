// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// NetworkConfig Service - Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// NetworkConfigService type 
type NetworkConfigService struct {
	objectSet *client.NetworkConfigObjectSet
}

// NewNetworkConfigService - method to initialize "NetworkConfigService" 
func NewNetworkConfigService(gs *NsGroupService) (*NetworkConfigService) {
	objectSet := gs.client.GetNetworkConfigObjectSet()
	return &NetworkConfigService{objectSet: objectSet}
}

// GetNetworkConfigs - method returns a array of pointers of type "NetworkConfigs"
func (svc *NetworkConfigService) GetNetworkConfigs(params *util.GetParams) ([]*model.NetworkConfig, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateNetworkConfig - method creates a "NetworkConfig"
func (svc *NetworkConfigService) CreateNetworkConfig(obj *model.NetworkConfig) (*model.NetworkConfig, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateNetworkConfig - method modifies  the "NetworkConfig" 
func (svc *NetworkConfigService) UpdateNetworkConfig(id string, obj *model.NetworkConfig) (*model.NetworkConfig, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetNetworkConfigById - method returns a pointer to "NetworkConfig"
func (svc *NetworkConfigService) GetNetworkConfigById(id string) (*model.NetworkConfig, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteNetworkConfig - deletes the "NetworkConfig"
func (svc *NetworkConfigService) DeleteNetworkConfig(id string) error {
	return svc.objectSet.DeleteObject(id)
}
