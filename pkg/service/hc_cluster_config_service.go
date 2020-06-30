// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// HcClusterConfig Service - Configuration information for virtual appliance that provides highly available storage and compute.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// HcClusterConfigService type 
type HcClusterConfigService struct {
	objectSet *client.HcClusterConfigObjectSet
}

// NewHcClusterConfigService - method to initialize "HcClusterConfigService" 
func NewHcClusterConfigService(gs *NsGroupService) (*HcClusterConfigService) {
	objectSet := gs.client.GetHcClusterConfigObjectSet()
	return &HcClusterConfigService{objectSet: objectSet}
}

// GetHcClusterConfigs - method returns a array of pointers of type "HcClusterConfigs"
func (svc *HcClusterConfigService) GetHcClusterConfigs(params *util.GetParams) ([]*model.HcClusterConfig, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetHcClusterConfigsWithFields - method returns a array of pointers of type "HcClusterConfig" 
func (svc *HcClusterConfigService) GetHcClusterConfigsWithFields(fields []string) ([]*model.HcClusterConfig, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateHcClusterConfig - method creates a "HcClusterConfig"
func (svc *HcClusterConfigService) CreateHcClusterConfig(obj *model.HcClusterConfig) (*model.HcClusterConfig, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditHcClusterConfig - method modifies  the "HcClusterConfig" 
func (svc *HcClusterConfigService) EditHcClusterConfig(id string, obj *model.HcClusterConfig) (*model.HcClusterConfig, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyHcClusterConfig - private method for more than one element check. 
func onlyHcClusterConfig(objs []*model.HcClusterConfig) (*model.HcClusterConfig, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one HcClusterConfig found with the given filter")
	}

	return objs[0], nil
}

 
// GetHcClusterConfigsByID - method returns associative a array of pointers of type "HcClusterConfig", filter by Id
func (svc *HcClusterConfigService) GetHcClusterConfigsByID(pool *model.Pool, fields []string) (map[string]*model.HcClusterConfig, error) {
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
	objs, err := svc.GetHcClusterConfigs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.HcClusterConfig)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetHcClusterConfigById - method returns a pointer to "HcClusterConfig"
func (svc *HcClusterConfigService) GetHcClusterConfigById(id string) (*model.HcClusterConfig, error) {
	return svc.objectSet.GetObject(id)
}

// GetHcClusterConfigsByName - method returns a associative array of pointers of type "HcClusterConfig", filter by name 
func (svc *HcClusterConfigService) GetHcClusterConfigsByName(pool *model.Pool, fields []string) (map[string]*model.HcClusterConfig, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetHcClusterConfigs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.HcClusterConfig)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetHcClusterConfigByName - method returns a pointer "HcClusterConfig" 
func (svc *HcClusterConfigService) GetHcClusterConfigByName(name string) (*model.HcClusterConfig, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return onlyHcClusterConfig(objs)
}	

