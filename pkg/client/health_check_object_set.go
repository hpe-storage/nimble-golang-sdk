// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// View the health of the group of arrays.
const (
    healthCheckPath = "health_checks"
)


// HealthCheckObjectSet
type HealthCheckObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new HealthCheck object
func (objectSet *HealthCheckObjectSet) CreateObject(payload *model.HealthCheck) (*model.HealthCheck, error) {
	response, err := objectSet.Client.Post(healthCheckPath, payload)
	return response.(*model.HealthCheck), err
}

// UpdateObject Modify existing HealthCheck object
func (objectSet *HealthCheckObjectSet) UpdateObject(id string, payload *model.HealthCheck) (*model.HealthCheck, error) {
	response, err := objectSet.Client.Put(healthCheckPath, id, payload)
	return response.(*model.HealthCheck), err
}

// DeleteObject deletes the HealthCheck object with the specified ID
func (objectSet *HealthCheckObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(healthCheckPath, id)
}

// GetObject returns a HealthCheck object with the given ID
func (objectSet *HealthCheckObjectSet) GetObject(id string) (*model.HealthCheck, error) {
	response, err := objectSet.Client.Get(healthCheckPath, id, model.HealthCheck{})
	if response == nil {
		return nil, err
	}
	return response.(*model.HealthCheck), err
}

// GetObjectList returns the list of HealthCheck objects
func (objectSet *HealthCheckObjectSet) GetObjectList() ([]*model.HealthCheck, error) {
	response, err := objectSet.Client.List(healthCheckPath)
	return buildHealthCheckObjectSet(response), err
}

// GetObjectListFromParams returns the list of HealthCheck objects using the given params query info
func (objectSet *HealthCheckObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.HealthCheck, error) {
	response, err := objectSet.Client.ListFromParams(healthCheckPath, params)
	return buildHealthCheckObjectSet(response), err
}

// generated function to build the appropriate response types
func buildHealthCheckObjectSet(response interface{}) ([]*model.HealthCheck) {
	values := reflect.ValueOf(response)
	results := make([]*model.HealthCheck, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.HealthCheck{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}