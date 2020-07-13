// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelConfig Service - Manage group wide Fibre Channel configuration.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// FibreChannelConfigService type
type FibreChannelConfigService struct {
	objectSet *client.FibreChannelConfigObjectSet
}

// NewFibreChannelConfigService - method to initialize "FibreChannelConfigService"
func NewFibreChannelConfigService(gs *NsGroupService) *FibreChannelConfigService {
	objectSet := gs.client.GetFibreChannelConfigObjectSet()
	return &FibreChannelConfigService{objectSet: objectSet}
}

// GetFibreChannelConfigs - method returns a array of pointers of type "FibreChannelConfigs"
func (svc *FibreChannelConfigService) GetFibreChannelConfigs(params *param.GetParams) ([]*nimbleos.FibreChannelConfig, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	fibreChannelConfigResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return fibreChannelConfigResp, nil
}

// CreateFibreChannelConfig - method creates a "FibreChannelConfig"
func (svc *FibreChannelConfigService) CreateFibreChannelConfig(obj *nimbleos.FibreChannelConfig) (*nimbleos.FibreChannelConfig, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	fibreChannelConfigResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return fibreChannelConfigResp, nil
}

// UpdateFibreChannelConfig - method modifies  the "FibreChannelConfig"
func (svc *FibreChannelConfigService) UpdateFibreChannelConfig(id string, obj *nimbleos.FibreChannelConfig) (*nimbleos.FibreChannelConfig, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	fibreChannelConfigResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return fibreChannelConfigResp, nil
}

// GetFibreChannelConfigById - method returns a pointer to "FibreChannelConfig"
func (svc *FibreChannelConfigService) GetFibreChannelConfigById(id string) (*nimbleos.FibreChannelConfig, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	fibreChannelConfigResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return fibreChannelConfigResp, nil
}

// DeleteFibreChannelConfig - deletes the "FibreChannelConfig"
func (svc *FibreChannelConfigService) DeleteFibreChannelConfig(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
