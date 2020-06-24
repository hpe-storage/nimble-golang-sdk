package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manage performance policies. A performance policy is a set of optimizations including block size, compression, and caching, to ensure that the volume&#39;s performance is the best configuration for its intended use like databases or log files. By default, a volume uses the \\&quot;default\\&quot; performance policy, which is set to use 4096 byte blocks with full compression and caching enabled. For replicated volumes, the same performance policy must exist on each replication partner.
 *
 */
const (
    performancePolicyPath = "performance_policies"
)

/**
 * PerformancePolicyObjectSet
*/
type PerformancePolicyObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new PerformancePolicy object
func (objectSet *PerformancePolicyObjectSet) CreateObject(payload *model.PerformancePolicy) (*model.PerformancePolicy, error) {
	response, err := objectSet.Client.Post(performancePolicyPath, payload)
	return response.(*model.PerformancePolicy), err
}

// UpdateObject Modify existing PerformancePolicy object
func (objectSet *PerformancePolicyObjectSet) UpdateObject(id string, payload *model.PerformancePolicy) (*model.PerformancePolicy, error) {
	response, err := objectSet.Client.Put(performancePolicyPath, id, payload)
	return response.(*model.PerformancePolicy), err
}

// DeleteObject deletes the PerformancePolicy object with the specified ID
func (objectSet *PerformancePolicyObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(performancePolicyPath, id)
}

// GetObject returns a PerformancePolicy object with the given ID
func (objectSet *PerformancePolicyObjectSet) GetObject(id string) (*model.PerformancePolicy, error) {
	response, err := objectSet.Client.Get(performancePolicyPath, id, model.PerformancePolicy{})
	if response == nil {
		return nil, err
	}
	return response.(*model.PerformancePolicy), err
}

// GetObjectList returns the list of PerformancePolicy objects
func (objectSet *PerformancePolicyObjectSet) GetObjectList() ([]*model.PerformancePolicy, error) {
	response, err := objectSet.Client.List(performancePolicyPath)
	return buildPerformancePolicyObjectSet(response), err
}

// GetObjectListFromParams returns the list of PerformancePolicy objects using the given params query info
func (objectSet *PerformancePolicyObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.PerformancePolicy, error) {
	response, err := objectSet.Client.ListFromParams(performancePolicyPath, params)
	return buildPerformancePolicyObjectSet(response), err
}

// generated function to build the appropriate response types
func buildPerformancePolicyObjectSet(response interface{}) ([]*model.PerformancePolicy) {
	values := reflect.ValueOf(response)
	results := make([]*model.PerformancePolicy, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.PerformancePolicy{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}