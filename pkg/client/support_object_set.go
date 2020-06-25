// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// View and alter support-based parameters.
const (
    supportPath = "support"
)


// SupportObjectSet
type SupportObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Support object
func (objectSet *SupportObjectSet) CreateObject(payload *model.Support) (*model.Support, error) {
	response, err := objectSet.Client.Post(supportPath, payload)
	return response.(*model.Support), err
}

// UpdateObject Modify existing Support object
func (objectSet *SupportObjectSet) UpdateObject(id string, payload *model.Support) (*model.Support, error) {
	response, err := objectSet.Client.Put(supportPath, id, payload)
	return response.(*model.Support), err
}

// DeleteObject deletes the Support object with the specified ID
func (objectSet *SupportObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(supportPath, id)
}

// GetObject returns a Support object with the given ID
func (objectSet *SupportObjectSet) GetObject(id string) (*model.Support, error) {
	response, err := objectSet.Client.Get(supportPath, id, model.Support{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Support), err
}

// GetObjectList returns the list of Support objects
func (objectSet *SupportObjectSet) GetObjectList() ([]*model.Support, error) {
	response, err := objectSet.Client.List(supportPath)
	return buildSupportObjectSet(response), err
}

// GetObjectListFromParams returns the list of Support objects using the given params query info
func (objectSet *SupportObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Support, error) {
	response, err := objectSet.Client.ListFromParams(supportPath, params)
	return buildSupportObjectSet(response), err
}

// generated function to build the appropriate response types
func buildSupportObjectSet(response interface{}) ([]*model.Support) {
	values := reflect.ValueOf(response)
	results := make([]*model.Support, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Support{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}