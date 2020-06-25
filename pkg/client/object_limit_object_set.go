// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// List the maximum limits and warning thresholds for number of objects in the storage group.
const (
    objectLimitPath = "object_limits"
)


// ObjectLimitObjectSet
type ObjectLimitObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ObjectLimit object
func (objectSet *ObjectLimitObjectSet) CreateObject(payload *model.ObjectLimit) (*model.ObjectLimit, error) {
	response, err := objectSet.Client.Post(objectLimitPath, payload)
	return response.(*model.ObjectLimit), err
}

// UpdateObject Modify existing ObjectLimit object
func (objectSet *ObjectLimitObjectSet) UpdateObject(id string, payload *model.ObjectLimit) (*model.ObjectLimit, error) {
	response, err := objectSet.Client.Put(objectLimitPath, id, payload)
	return response.(*model.ObjectLimit), err
}

// DeleteObject deletes the ObjectLimit object with the specified ID
func (objectSet *ObjectLimitObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(objectLimitPath, id)
}

// GetObject returns a ObjectLimit object with the given ID
func (objectSet *ObjectLimitObjectSet) GetObject(id string) (*model.ObjectLimit, error) {
	response, err := objectSet.Client.Get(objectLimitPath, id, model.ObjectLimit{})
	if response == nil {
		return nil, err
	}
	return response.(*model.ObjectLimit), err
}

// GetObjectList returns the list of ObjectLimit objects
func (objectSet *ObjectLimitObjectSet) GetObjectList() ([]*model.ObjectLimit, error) {
	response, err := objectSet.Client.List(objectLimitPath)
	return buildObjectLimitObjectSet(response), err
}

// GetObjectListFromParams returns the list of ObjectLimit objects using the given params query info
func (objectSet *ObjectLimitObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ObjectLimit, error) {
	response, err := objectSet.Client.ListFromParams(objectLimitPath, params)
	return buildObjectLimitObjectSet(response), err
}

// generated function to build the appropriate response types
func buildObjectLimitObjectSet(response interface{}) ([]*model.ObjectLimit) {
	values := reflect.ValueOf(response)
	results := make([]*model.ObjectLimit, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ObjectLimit{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}