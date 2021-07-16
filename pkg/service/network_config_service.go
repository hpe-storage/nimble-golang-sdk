// Copyright 2021 Hewlett Packard Enterprise Development LP

package service

// NetworkConfig Service - Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *NetworkConfigService) GetNetworkConfigs(params *param.GetParams) ([]*nimbleos.NetworkConfig, error) {
	networkConfigResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return networkConfigResp, nil
}

// CreateNetworkConfig - method creates a "NetworkConfig"
func (svc *NetworkConfigService) CreateNetworkConfig(obj *nimbleos.NetworkConfig) (*nimbleos.NetworkConfig, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateNetworkConfig: invalid parameter specified, %v", obj)
	}

	networkConfigResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return networkConfigResp, nil
}

// UpdateNetworkConfig - method modifies  the "NetworkConfig"
func (svc *NetworkConfigService) UpdateNetworkConfig(id string, obj *nimbleos.NetworkConfig) (*nimbleos.NetworkConfig, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateNetworkConfig: invalid parameter specified, %v", obj)
	}

	networkConfigResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return networkConfigResp, nil
}

// GetNetworkConfigById - method returns a pointer to "NetworkConfig"
func (svc *NetworkConfigService) GetNetworkConfigById(id string) (*nimbleos.NetworkConfig, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetNetworkConfigById: invalid parameter specified, %v", id)
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
		return fmt.Errorf("DeleteNetworkConfig: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// ActivateNetconfigNetworkConfig - activate a network configuration.
//   Required parameters:
//       id - ID of the netconfig to activate.
//       ignoreValidationMask - Whether to ignore validation or not.

func (svc *NetworkConfigService) ActivateNetconfigNetworkConfig(id string, ignoreValidationMask uint64) error {

	if len(id) == 0 {
		return fmt.Errorf("ActivateNetconfigNetworkConfig: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.ActivateNetconfig(&id, &ignoreValidationMask)
	return err
}

// ValidateNetconfigNetworkConfig - validate a network configuration.
//   Required parameters:
//       id - ID of the netconfig to validate.
//       ignoreValidationMask - Whether to ignore validation or not.

func (svc *NetworkConfigService) ValidateNetconfigNetworkConfig(id string, ignoreValidationMask uint64) error {

	if len(id) == 0 {
		return fmt.Errorf("ValidateNetconfigNetworkConfig: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.ValidateNetconfig(&id, &ignoreValidationMask)
	return err
}
