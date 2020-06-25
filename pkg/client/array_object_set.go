// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Retrieve information of specified arrays. The array is the management and configuration for the underlying physical hardware array box.
const (
    arrayPath = "arrays"
)


// ArrayObjectSet
type ArrayObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Array object
func (objectSet *ArrayObjectSet) CreateObject(payload *model.Array) (*model.Array, error) {
	response, err := objectSet.Client.Post(arrayPath, payload)
	return response.(*model.Array), err
}

// UpdateObject Modify existing Array object
func (objectSet *ArrayObjectSet) UpdateObject(id string, payload *model.Array) (*model.Array, error) {
	response, err := objectSet.Client.Put(arrayPath, id, payload)
	return response.(*model.Array), err
}

// DeleteObject deletes the Array object with the specified ID
func (objectSet *ArrayObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(arrayPath, id)
}

// GetObject returns a Array object with the given ID
func (objectSet *ArrayObjectSet) GetObject(id string) (*model.Array, error) {
	response, err := objectSet.Client.Get(arrayPath, id, model.Array{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Array), err
}

// GetObjectList returns the list of Array objects
func (objectSet *ArrayObjectSet) GetObjectList() ([]*model.Array, error) {
	response, err := objectSet.Client.List(arrayPath)
	return buildArrayObjectSet(response), err
}

// GetObjectListFromParams returns the list of Array objects using the given params query info
func (objectSet *ArrayObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Array, error) {
	response, err := objectSet.Client.ListFromParams(arrayPath, params)
	return buildArrayObjectSet(response), err
}

// generated function to build the appropriate response types
func buildArrayObjectSet(response interface{}) ([]*model.Array) {
	values := reflect.ValueOf(response)
	results := make([]*model.Array, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Array{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}