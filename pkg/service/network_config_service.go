// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// NetworkConfig Service - Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetNetworkConfigsWithFields - method returns a array of pointers of type "NetworkConfig" 
func (svc *NetworkConfigService) GetNetworkConfigsWithFields(fields []string) ([]*model.NetworkConfig, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateNetworkConfig - method creates a "NetworkConfig"
func (svc *NetworkConfigService) CreateNetworkConfig(obj *model.NetworkConfig) (*model.NetworkConfig, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditNetworkConfig - method modifies  the "NetworkConfig" 
func (svc *NetworkConfigService) EditNetworkConfig(id string, obj *model.NetworkConfig) (*model.NetworkConfig, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyNetworkConfig - private method for more than one element check. 
func onlyNetworkConfig(objs []*model.NetworkConfig) (*model.NetworkConfig, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one NetworkConfig found with the given filter")
	}

	return objs[0], nil
}

 
// GetNetworkConfigsByID - method returns associative a array of pointers of type "NetworkConfig", filter by Id
func (svc *NetworkConfigService) GetNetworkConfigsByID(pool *model.Pool, fields []string) (map[string]*model.NetworkConfig, error) {
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
	objs, err := svc.GetNetworkConfigs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.NetworkConfig)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetNetworkConfigById - method returns a pointer to "NetworkConfig"
func (svc *NetworkConfigService) GetNetworkConfigById(id string) (*model.NetworkConfig, error) {
	return svc.objectSet.GetObject(id)
}


