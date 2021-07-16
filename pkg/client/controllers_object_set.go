// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *ControllerObjectSet) CreateObject(payload *nimbleos.Controller) (*nimbleos.Controller, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Controller")
}

// UpdateObject Modify existing Controller object
func (objectSet *ControllerObjectSet) UpdateObject(id string, payload *nimbleos.Controller) (*nimbleos.Controller, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Controller")
}

// DeleteObject deletes the Controller object with the specified ID
func (objectSet *ControllerObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Controller")
}

// GetObject returns a Controller object with the given ID
func (objectSet *ControllerObjectSet) GetObject(id string) (*nimbleos.Controller, error) {
	resp, err := objectSet.Client.Get(controllerPath, id, &nimbleos.Controller{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Controller), err
}

// GetObjectList returns the list of Controller objects
func (objectSet *ControllerObjectSet) GetObjectList() ([]*nimbleos.Controller, error) {
	resp, err := objectSet.Client.List(controllerPath)
	if err != nil {
		return nil, err
	}
	return buildControllerObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Controller objects using the given params query info
func (objectSet *ControllerObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Controller, error) {
	controllerObjectSetResp, err := objectSet.Client.ListFromParams(controllerPath, params)
	if err != nil {
		return nil, err
	}
	return buildControllerObjectSet(controllerObjectSetResp), err
}

// generated function to build the appropriate response types
func buildControllerObjectSet(response interface{}) []*nimbleos.Controller {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Controller, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Controller{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
