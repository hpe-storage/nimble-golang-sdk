// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Manage performance policies. A performance policy is a set of optimizations including block size, compression, and caching, to ensure that the volume's performance is the best
// configuration for its intended use like databases or log files. By default, a volume uses the \\"default\\" performance policy, which is set to use 4096 byte blocks with full
// compression and caching enabled. For replicated volumes, the same performance policy must exist on each replication partner.
const (
	performancePolicyPath = "performance_policies"
)

// PerformancePolicyObjectSet
type PerformancePolicyObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new PerformancePolicy object
func (objectSet *PerformancePolicyObjectSet) CreateObject(payload *nimbleos.PerformancePolicy) (*nimbleos.PerformancePolicy, error) {
	resp, err := objectSet.Client.Post(performancePolicyPath, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return resp.(*nimbleos.PerformancePolicy), err
}

// UpdateObject Modify existing PerformancePolicy object
func (objectSet *PerformancePolicyObjectSet) UpdateObject(id string, payload *nimbleos.PerformancePolicy) (*nimbleos.PerformancePolicy, error) {
	resp, err := objectSet.Client.Put(performancePolicyPath, id, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.PerformancePolicy), err
}

// DeleteObject deletes the PerformancePolicy object with the specified ID
func (objectSet *PerformancePolicyObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(performancePolicyPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a PerformancePolicy object with the given ID
func (objectSet *PerformancePolicyObjectSet) GetObject(id string) (*nimbleos.PerformancePolicy, error) {
	resp, err := objectSet.Client.Get(performancePolicyPath, id, nimbleos.PerformancePolicy{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.PerformancePolicy), err
}

// GetObjectList returns the list of PerformancePolicy objects
func (objectSet *PerformancePolicyObjectSet) GetObjectList() ([]*nimbleos.PerformancePolicy, error) {
	resp, err := objectSet.Client.List(performancePolicyPath)
	if err != nil {
		return nil, err
	}
	return buildPerformancePolicyObjectSet(resp), err
}

// GetObjectListFromParams returns the list of PerformancePolicy objects using the given params query info
func (objectSet *PerformancePolicyObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.PerformancePolicy, error) {
	performancePolicyObjectSetResp, err := objectSet.Client.ListFromParams(performancePolicyPath, params)
	if err != nil {
		return nil, err
	}
	return buildPerformancePolicyObjectSet(performancePolicyObjectSetResp), err
}

// generated function to build the appropriate response types
func buildPerformancePolicyObjectSet(response interface{}) []*nimbleos.PerformancePolicy {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.PerformancePolicy, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.PerformancePolicy{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
