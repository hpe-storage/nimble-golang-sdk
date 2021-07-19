// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

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
	fibreChannelConfigResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return fibreChannelConfigResp, nil
}

// CreateFibreChannelConfig - method creates a "FibreChannelConfig"
func (svc *FibreChannelConfigService) CreateFibreChannelConfig(obj *nimbleos.FibreChannelConfig) (*nimbleos.FibreChannelConfig, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateFibreChannelConfig: invalid parameter specified, %v", obj)
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
		return nil, fmt.Errorf("UpdateFibreChannelConfig: invalid parameter specified, %v", obj)
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
		return nil, fmt.Errorf("GetFibreChannelConfigById: invalid parameter specified, %v", id)
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
		return fmt.Errorf("DeleteFibreChannelConfig: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// RegenerateFibreChannelConfig - regenerate Fibre Channel configuration.
//   Required parameters:
//       id - ID of the Fibre Channel configuration.
//       wwnnBaseStr - Base World Wide Node Name(WWNN).
//       precheck - Check if the interfaces are offline before regenerating the WWNN (World Wide Node Name).

func (svc *FibreChannelConfigService) RegenerateFibreChannelConfig(id string, wwnnBaseStr string, precheck bool) (*nimbleos.NsFcConfigRegenerateReturn, error) {

	if len(id) == 0 || len(wwnnBaseStr) == 0 {
		return nil, fmt.Errorf("RegenerateFibreChannelConfig: invalid parameter specified id: %v, wwnnBaseStr: %v ", id, wwnnBaseStr)
	}

	resp, err := svc.objectSet.Regenerate(&id, &wwnnBaseStr, &precheck)
	return resp, err
}

// HwUpgradeFibreChannelConfig - update Fibre Channel configuration after hardware changes.
//   Required parameters:
//       id - ID of the Fibre Channel configuration.

func (svc *FibreChannelConfigService) HwUpgradeFibreChannelConfig(id string) error {

	if len(id) == 0 {
		return fmt.Errorf("HwUpgradeFibreChannelConfig: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.HwUpgrade(&id)
	return err
}
