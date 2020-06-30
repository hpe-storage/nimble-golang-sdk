// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// HealthCheck Service - View the health of the group of arrays.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// HealthCheckService type 
type HealthCheckService struct {
	objectSet *client.HealthCheckObjectSet
}

// NewHealthCheckService - method to initialize "HealthCheckService" 
func NewHealthCheckService(gs *NsGroupService) (*HealthCheckService) {
	objectSet := gs.client.GetHealthCheckObjectSet()
	return &HealthCheckService{objectSet: objectSet}
}

// GetHealthChecks - method returns a array of pointers of type "HealthChecks"
func (svc *HealthCheckService) GetHealthChecks(params *util.GetParams) ([]*model.HealthCheck, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetHealthChecksWithFields - method returns a array of pointers of type "HealthCheck" 
func (svc *HealthCheckService) GetHealthChecksWithFields(fields []string) ([]*model.HealthCheck, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateHealthCheck - method creates a "HealthCheck"
func (svc *HealthCheckService) CreateHealthCheck(obj *model.HealthCheck) (*model.HealthCheck, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditHealthCheck - method modifies  the "HealthCheck" 
func (svc *HealthCheckService) EditHealthCheck(id string, obj *model.HealthCheck) (*model.HealthCheck, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyHealthCheck - private method for more than one element check. 
func onlyHealthCheck(objs []*model.HealthCheck) (*model.HealthCheck, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one HealthCheck found with the given filter")
	}

	return objs[0], nil
}

 
// GetHealthChecksByID - method returns associative a array of pointers of type "HealthCheck", filter by Id
func (svc *HealthCheckService) GetHealthChecksByID(pool *model.Pool, fields []string) (map[string]*model.HealthCheck, error) {
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
	objs, err := svc.GetHealthChecks(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.HealthCheck)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetHealthCheckById - method returns a pointer to "HealthCheck"
func (svc *HealthCheckService) GetHealthCheckById(id string) (*model.HealthCheck, error) {
	return svc.objectSet.GetObject(id)
}


