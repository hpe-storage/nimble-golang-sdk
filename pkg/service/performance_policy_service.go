// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// PerformancePolicy Service - Manage performance policies. A performance policy is a set of optimizations including block size, compression, and caching, to ensure that the volume's performance is the best
// configuration for its intended use like databases or log files. By default, a volume uses the \\"default\\" performance policy, which is set to use 4096 byte blocks with full
// compression and caching enabled. For replicated volumes, the same performance policy must exist on each replication partner.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// GetPerformancePolicies - method returns a array of pointers of type "PerformancePolicies"
func (svc *PerformancePolicyService) GetPerformancePolicies(params *util.GetParams) ([]*model.PerformancePolicy, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}


// CreatePerformancePolicy - method creates a "PerformancePolicy"
func (svc *PerformancePolicyService) CreatePerformancePolicy(obj *model.PerformancePolicy) (*model.PerformancePolicy, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdatePerformancePolicy - method modifies  the "PerformancePolicy" 
func (svc *PerformancePolicyService) UpdatePerformancePolicy(id string, obj *model.PerformancePolicy) (*model.PerformancePolicy, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetPerformancePolicyById - method returns a pointer to "PerformancePolicy"
func (svc *PerformancePolicyService) GetPerformancePolicyById(id string) (*model.PerformancePolicy, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeletePerformancePolicy - deletes the "PerformancePolicy"
func (svc *PerformancePolicyService) DeletePerformancePolicy(id string) error {
	return svc.objectSet.DeleteObject(id)
}
