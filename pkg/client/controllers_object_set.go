// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Controller is a redundant collection of hardware capable of running the array software.
const (
    controllerPath = "controllers"
)

// ControllerObjectSet
type ControllerObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Controller object
func (objectSet *ControllerObjectSet) CreateObject(payload *model.Controller) (*model.Controller, error) {
	response, err := objectSet.Client.Post(controllerPath, payload)
	return response.(*model.Controller), err
}

// UpdateObject Modify existing Controller object
func (objectSet *ControllerObjectSet) UpdateObject(id string, payload *model.Controller) (*model.Controller, error) {
	response, err := objectSet.Client.Put(controllerPath, id, payload)
	return response.(*model.Controller), err
}

// DeleteObject deletes the Controller object with the specified ID
func (objectSet *ControllerObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(controllerPath, id)
}

// GetObject returns a Controller object with the given ID
func (objectSet *ControllerObjectSet) GetObject(id string) (*model.Controller, error) {
	response, err := objectSet.Client.Get(controllerPath, id, model.Controller{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Controller), err
}

// GetObjectList returns the list of Controller objects
func (objectSet *ControllerObjectSet) GetObjectList() ([]*model.Controller, error) {
	response, err := objectSet.Client.List(controllerPath)
	return buildControllerObjectSet(response), err
}

// GetObjectListFromParams returns the list of Controller objects using the given params query info
func (objectSet *ControllerObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Controller, error) {
	response, err := objectSet.Client.ListFromParams(controllerPath, params)
	return buildControllerObjectSet(response), err
}

// generated function to build the appropriate response types
func buildControllerObjectSet(response interface{}) ([]*model.Controller) {
	values := reflect.ValueOf(response)
	results := make([]*model.Controller, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Controller{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}