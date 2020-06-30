// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Method to help log events from outside of storage array to provide context for troubleshooting host-side or array-side issues.
const (
    debugLogPath = "debug_log"
)

// DebugLogObjectSet
type DebugLogObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new DebugLog object
func (objectSet *DebugLogObjectSet) CreateObject(payload *model.DebugLog) (*model.DebugLog, error) {
	response, err := objectSet.Client.Post(debugLogPath, payload)
	return response.(*model.DebugLog), err
}

// UpdateObject Modify existing DebugLog object
func (objectSet *DebugLogObjectSet) UpdateObject(id string, payload *model.DebugLog) (*model.DebugLog, error) {
	response, err := objectSet.Client.Put(debugLogPath, id, payload)
	return response.(*model.DebugLog), err
}

// DeleteObject deletes the DebugLog object with the specified ID
func (objectSet *DebugLogObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(debugLogPath, id)
}

// GetObject returns a DebugLog object with the given ID
func (objectSet *DebugLogObjectSet) GetObject(id string) (*model.DebugLog, error) {
	response, err := objectSet.Client.Get(debugLogPath, id, model.DebugLog{})
	if response == nil {
		return nil, err
	}
	return response.(*model.DebugLog), err
}

// GetObjectList returns the list of DebugLog objects
func (objectSet *DebugLogObjectSet) GetObjectList() ([]*model.DebugLog, error) {
	response, err := objectSet.Client.List(debugLogPath)
	return buildDebugLogObjectSet(response), err
}

// GetObjectListFromParams returns the list of DebugLog objects using the given params query info
func (objectSet *DebugLogObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.DebugLog, error) {
	response, err := objectSet.Client.ListFromParams(debugLogPath, params)
	return buildDebugLogObjectSet(response), err
}

// generated function to build the appropriate response types
func buildDebugLogObjectSet(response interface{}) ([]*model.DebugLog) {
	values := reflect.ValueOf(response)
	results := make([]*model.DebugLog, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.DebugLog{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}