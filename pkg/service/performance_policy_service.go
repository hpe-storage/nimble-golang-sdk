// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// PerformancePolicy Service - Manage performance policies. A performance policy is a set of optimizations including block size, compression, and caching, to ensure that the volume's performance is the best
// configuration for its intended use like databases or log files. By default, a volume uses the \\"default\\" performance policy, which is set to use 4096 byte blocks with full
// compression and caching enabled. For replicated volumes, the same performance policy must exist on each replication partner.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// PerformancePolicyService type
type PerformancePolicyService struct {
	objectSet *client.PerformancePolicyObjectSet
}

// NewPerformancePolicyService - method to initialize "PerformancePolicyService"
func NewPerformancePolicyService(gs *NsGroupService) *PerformancePolicyService {
	objectSet := gs.client.GetPerformancePolicyObjectSet()
	return &PerformancePolicyService{objectSet: objectSet}
}

// GetPerformancePolicies - method returns a array of pointers of type "PerformancePolicies"
func (svc *PerformancePolicyService) GetPerformancePolicies(params *param.GetParams) ([]*nimbleos.PerformancePolicy, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	performancePolicyResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return performancePolicyResp, nil
}

// CreatePerformancePolicy - method creates a "PerformancePolicy"
func (svc *PerformancePolicyService) CreatePerformancePolicy(obj *nimbleos.PerformancePolicy) (*nimbleos.PerformancePolicy, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	performancePolicyResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return performancePolicyResp, nil
}

// UpdatePerformancePolicy - method modifies  the "PerformancePolicy"
func (svc *PerformancePolicyService) UpdatePerformancePolicy(id string, obj *nimbleos.PerformancePolicy) (*nimbleos.PerformancePolicy, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	performancePolicyResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return performancePolicyResp, nil
}

// GetPerformancePolicyById - method returns a pointer to "PerformancePolicy"
func (svc *PerformancePolicyService) GetPerformancePolicyById(id string) (*nimbleos.PerformancePolicy, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	performancePolicyResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return performancePolicyResp, nil
}

// GetPerformancePolicyByName - method returns a pointer "PerformancePolicy"
func (svc *PerformancePolicyService) GetPerformancePolicyByName(name string) (*nimbleos.PerformancePolicy, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	performancePolicyResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(performancePolicyResp) == 0 {
		return nil, nil
	}

	return performancePolicyResp[0], nil
}

// DeletePerformancePolicy - deletes the "PerformancePolicy"
func (svc *PerformancePolicyService) DeletePerformancePolicy(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
