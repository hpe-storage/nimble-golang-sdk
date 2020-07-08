// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// NetworkConfig Service - Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// NetworkConfigService type
type NetworkConfigService struct {
	objectSet *client.NetworkConfigObjectSet
}

// NewNetworkConfigService - method to initialize "NetworkConfigService"
func NewNetworkConfigService(gs *NsGroupService) *NetworkConfigService {
	objectSet := gs.client.GetNetworkConfigObjectSet()
	return &NetworkConfigService{objectSet: objectSet}
}

// GetNetworkConfigs - method returns a array of pointers of type "NetworkConfigs"
func (svc *NetworkConfigService) GetNetworkConfigs(params *util.GetParams) ([]*model.NetworkConfig, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	networkConfigResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return networkConfigResp, nil
}

// CreateNetworkConfig - method creates a "NetworkConfig"
func (svc *NetworkConfigService) CreateNetworkConfig(obj *model.NetworkConfig) (*model.NetworkConfig, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	networkConfigResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return networkConfigResp, nil
}

// UpdateNetworkConfig - method modifies  the "NetworkConfig"
func (svc *NetworkConfigService) UpdateNetworkConfig(id string, obj *model.NetworkConfig) (*model.NetworkConfig, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	networkConfigResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return networkConfigResp, nil
}

// GetNetworkConfigById - method returns a pointer to "NetworkConfig"
func (svc *NetworkConfigService) GetNetworkConfigById(id string) (*model.NetworkConfig, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	networkConfigResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return networkConfigResp, nil
}

// DeleteNetworkConfig - deletes the "NetworkConfig"
func (svc *NetworkConfigService) DeleteNetworkConfig(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
