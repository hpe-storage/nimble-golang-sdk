// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"fmt"
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
	return nil, fmt.Errorf("Unsupported operation 'create' on Controller")
}

// UpdateObject Modify existing Controller object
func (objectSet *ControllerObjectSet) UpdateObject(id string, payload *model.Controller) (*model.Controller, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Controller")
}

// DeleteObject deletes the Controller object with the specified ID
func (objectSet *ControllerObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Controller")
}

// GetObject returns a Controller object with the given ID
func (objectSet *ControllerObjectSet) GetObject(id string) (*model.Controller, error) {
	controllerObjectSetResp, err := objectSet.Client.Get(controllerPath, id, model.Controller{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if controllerObjectSetResp == nil {
		return nil,nil
	}
	return controllerObjectSetResp.(*model.Controller), err
}

// GetObjectList returns the list of Controller objects
func (objectSet *ControllerObjectSet) GetObjectList() ([]*model.Controller, error) {
	controllerObjectSetResp, err := objectSet.Client.List(controllerPath)
	if err != nil {
		return nil, err
	}
	return buildControllerObjectSet(controllerObjectSetResp), err
}

// GetObjectListFromParams returns the list of Controller objects using the given params query info
func (objectSet *ControllerObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Controller, error) {
	controllerObjectSetResp, err := objectSet.Client.ListFromParams(controllerPath, params)
	if err != nil {
		return nil, err
	}
	return buildControllerObjectSet(controllerObjectSetResp), err
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
