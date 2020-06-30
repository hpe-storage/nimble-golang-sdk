// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelConfig Service - Manage group wide Fibre Channel configuration.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetFibreChannelConfigsWithFields - method returns a array of pointers of type "FibreChannelConfig" 
func (svc *FibreChannelConfigService) GetFibreChannelConfigsWithFields(fields []string) ([]*model.FibreChannelConfig, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelConfig - method creates a "FibreChannelConfig"
func (svc *FibreChannelConfigService) CreateFibreChannelConfig(obj *model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditFibreChannelConfig - method modifies  the "FibreChannelConfig" 
func (svc *FibreChannelConfigService) EditFibreChannelConfig(id string, obj *model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyFibreChannelConfig - private method for more than one element check. 
func onlyFibreChannelConfig(objs []*model.FibreChannelConfig) (*model.FibreChannelConfig, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one FibreChannelConfig found with the given filter")
	}

	return objs[0], nil
}

 
// GetFibreChannelConfigsByID - method returns associative a array of pointers of type "FibreChannelConfig", filter by Id
func (svc *FibreChannelConfigService) GetFibreChannelConfigsByID(pool *model.Pool, fields []string) (map[string]*model.FibreChannelConfig, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetFibreChannelConfigs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FibreChannelConfig)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetFibreChannelConfigById - method returns a pointer to "FibreChannelConfig"
func (svc *FibreChannelConfigService) GetFibreChannelConfigById(id string) (*model.FibreChannelConfig, error) {
	return svc.objectSet.GetObject(id)
}


