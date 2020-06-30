// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// PerformancePolicy Service - Manage performance policies. A performance policy is a set of optimizations including block size, compression, and caching, to ensure that the volume's performance is the best
// configuration for its intended use like databases or log files. By default, a volume uses the \\"default\\" performance policy, which is set to use 4096 byte blocks with full
// compression and caching enabled. For replicated volumes, the same performance policy must exist on each replication partner.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// PerformancePolicyService type 
type PerformancePolicyService struct {
	objectSet *client.PerformancePolicyObjectSet
}

// NewPerformancePolicyService - method to initialize "PerformancePolicyService" 
func NewPerformancePolicyService(gs *NsGroupService) (*PerformancePolicyService) {
	objectSet := gs.client.GetPerformancePolicyObjectSet()
	return &PerformancePolicyService{objectSet: objectSet}
}

// GetPerformancePolicys - method returns a array of pointers of type "PerformancePolicys"
func (svc *PerformancePolicyService) GetPerformancePolicys(params *util.GetParams) ([]*model.PerformancePolicy, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetPerformancePolicysWithFields - method returns a array of pointers of type "PerformancePolicy" 
func (svc *PerformancePolicyService) GetPerformancePolicysWithFields(fields []string) ([]*model.PerformancePolicy, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreatePerformancePolicy - method creates a "PerformancePolicy"
func (svc *PerformancePolicyService) CreatePerformancePolicy(obj *model.PerformancePolicy) (*model.PerformancePolicy, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditPerformancePolicy - method modifies  the "PerformancePolicy" 
func (svc *PerformancePolicyService) EditPerformancePolicy(id string, obj *model.PerformancePolicy) (*model.PerformancePolicy, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyPerformancePolicy - private method for more than one element check. 
func onlyPerformancePolicy(objs []*model.PerformancePolicy) (*model.PerformancePolicy, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one PerformancePolicy found with the given filter")
	}

	return objs[0], nil
}

 
// GetPerformancePolicysByID - method returns associative a array of pointers of type "PerformancePolicy", filter by Id
func (svc *PerformancePolicyService) GetPerformancePolicysByID(pool *model.Pool, fields []string) (map[string]*model.PerformancePolicy, error) {
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
	objs, err := svc.GetPerformancePolicys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.PerformancePolicy)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetPerformancePolicyById - method returns a pointer to "PerformancePolicy"
func (svc *PerformancePolicyService) GetPerformancePolicyById(id string) (*model.PerformancePolicy, error) {
	return svc.objectSet.GetObject(id)
}

// GetPerformancePolicysByName - method returns a associative array of pointers of type "PerformancePolicy", filter by name 
func (svc *PerformancePolicyService) GetPerformancePolicysByName(pool *model.Pool, fields []string) (map[string]*model.PerformancePolicy, error) {
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
	objs, err := svc.GetPerformancePolicys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.PerformancePolicy)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetPerformancePolicyByName - method returns a pointer "PerformancePolicy" 
func (svc *PerformancePolicyService) GetPerformancePolicyByName(name string) (*model.PerformancePolicy, error) {
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
	return onlyPerformancePolicy(objs)
}	

