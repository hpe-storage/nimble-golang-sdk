// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelConfig Service - Manage group wide Fibre Channel configuration.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// FibreChannelConfigService type 
type FibreChannelConfigService struct {
	objectSet *client.FibreChannelConfigObjectSet
}

// NewFibreChannelConfigService - method to initialize "FibreChannelConfigService" 
func NewFibreChannelConfigService(gs *NsGroupService) (*FibreChannelConfigService) {
	objectSet := gs.client.GetFibreChannelConfigObjectSet()
	return &FibreChannelConfigService{objectSet: objectSet}
}

// GetFibreChannelConfigs - method returns a array of pointers of type "FibreChannelConfigs"
func (svc *FibreChannelConfigService) GetFibreChannelConfigs(params *util.GetParams) ([]*model.FibreChannelConfig, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelConfig - method creates a "FibreChannelConfig"
func (svc *FibreChannelConfigService) CreateFibreChannelConfig(obj *model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateFibreChannelConfig - method modifies  the "FibreChannelConfig" 
func (svc *FibreChannelConfigService) UpdateFibreChannelConfig(id string, obj *model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetFibreChannelConfigById - method returns a pointer to "FibreChannelConfig"
func (svc *FibreChannelConfigService) GetFibreChannelConfigById(id string) (*model.FibreChannelConfig, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteFibreChannelConfig - deletes the "FibreChannelConfig"
func (svc *FibreChannelConfigService) DeleteFibreChannelConfig(id string) error {
	return svc.objectSet.DeleteObject(id)
}
