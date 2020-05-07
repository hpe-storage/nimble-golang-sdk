// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

const (
	performancePolicyPath = "performance_policies"
)

// PerformancePolicyObjectSet provides a wrapper to manage performance policies from the client
type PerformancePolicyObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new performance policy
func (objectSet *PerformancePolicyObjectSet) CreateObject(payload *model.PerformancePolicy) (*model.PerformancePolicy, error) {
	response, err := objectSet.Client.Post(performancePolicyPath, payload)
	return response.(*model.PerformancePolicy), err
}

// GetObject returns a performance policy with the given ID
func (objectSet *PerformancePolicyObjectSet) GetObject(id string) (*model.PerformancePolicy, error) {
	response, err := objectSet.Client.Get(performancePolicyPath, id, model.PerformancePolicy{})
	if response == nil {
		return nil, err
	}
	return response.(*model.PerformancePolicy), err
}

// DeleteObject deletes the performance policy with the specified ID
func (objectSet *PerformancePolicyObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(performancePolicyPath, id)
}

// GetObjectList returns the list of performance policy objects
func (objectSet *PerformancePolicyObjectSet) GetObjectList() ([]*model.PerformancePolicy, error) {
	response, err := objectSet.Client.List(performancePolicyPath)
	return buildPerformancePolicies(response), err
}

// GetObjectListFromParams returns the list of performance policy objects using the given params query info
func (objectSet *PerformancePolicyObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.PerformancePolicy, error) {
	response, err := objectSet.Client.ListFromParams(performancePolicyPath, params)
	return buildPerformancePolicies(response), err
}

// generated function to build the appropriate response types
func buildPerformancePolicies(response interface{}) []*model.PerformancePolicy {
	values := reflect.ValueOf(response)
	results := make([]*model.PerformancePolicy, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.PerformancePolicy{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
