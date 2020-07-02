// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manage pools. Pools are an aggregation of arrays.
const (
    poolPath = "pools"
)

// PoolObjectSet
type PoolObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Pool object
func (objectSet *PoolObjectSet) CreateObject(payload *model.Pool) (*model.Pool, error) {
	response, err := objectSet.Client.Post(poolPath, payload)
	return response.(*model.Pool), err
}

// UpdateObject Modify existing Pool object
func (objectSet *PoolObjectSet) UpdateObject(id string, payload *model.Pool) (*model.Pool, error) {
	response, err := objectSet.Client.Put(poolPath, id, payload)
	return response.(*model.Pool), err
}

// DeleteObject deletes the Pool object with the specified ID
func (objectSet *PoolObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(poolPath, id)
}

// GetObject returns a Pool object with the given ID
func (objectSet *PoolObjectSet) GetObject(id string) (*model.Pool, error) {
	response, err := objectSet.Client.Get(poolPath, id, model.Pool{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Pool), err
}

// GetObjectList returns the list of Pool objects
func (objectSet *PoolObjectSet) GetObjectList() ([]*model.Pool, error) {
	response, err := objectSet.Client.List(poolPath)
	return buildPoolObjectSet(response), err
}

// GetObjectListFromParams returns the list of Pool objects using the given params query info
func (objectSet *PoolObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Pool, error) {
	response, err := objectSet.Client.ListFromParams(poolPath, params)
	return buildPoolObjectSet(response), err
}

// generated function to build the appropriate response types
func buildPoolObjectSet(response interface{}) ([]*model.Pool) {
	values := reflect.ValueOf(response)
	results := make([]*model.Pool, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Pool{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}