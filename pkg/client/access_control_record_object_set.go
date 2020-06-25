// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Manage access control records for volumes.
const (
    accessControlRecordPath = "access_control_records"
)


// AccessControlRecordObjectSet
type AccessControlRecordObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new AccessControlRecord object
func (objectSet *AccessControlRecordObjectSet) CreateObject(payload *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	response, err := objectSet.Client.Post(accessControlRecordPath, payload)
	return response.(*model.AccessControlRecord), err
}

// UpdateObject Modify existing AccessControlRecord object
func (objectSet *AccessControlRecordObjectSet) UpdateObject(id string, payload *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	response, err := objectSet.Client.Put(accessControlRecordPath, id, payload)
	return response.(*model.AccessControlRecord), err
}

// DeleteObject deletes the AccessControlRecord object with the specified ID
func (objectSet *AccessControlRecordObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(accessControlRecordPath, id)
}

// GetObject returns a AccessControlRecord object with the given ID
func (objectSet *AccessControlRecordObjectSet) GetObject(id string) (*model.AccessControlRecord, error) {
	response, err := objectSet.Client.Get(accessControlRecordPath, id, model.AccessControlRecord{})
	if response == nil {
		return nil, err
	}
	return response.(*model.AccessControlRecord), err
}

// GetObjectList returns the list of AccessControlRecord objects
func (objectSet *AccessControlRecordObjectSet) GetObjectList() ([]*model.AccessControlRecord, error) {
	response, err := objectSet.Client.List(accessControlRecordPath)
	return buildAccessControlRecordObjectSet(response), err
}

// GetObjectListFromParams returns the list of AccessControlRecord objects using the given params query info
func (objectSet *AccessControlRecordObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.AccessControlRecord, error) {
	response, err := objectSet.Client.ListFromParams(accessControlRecordPath, params)
	return buildAccessControlRecordObjectSet(response), err
}

// generated function to build the appropriate response types
func buildAccessControlRecordObjectSet(response interface{}) ([]*model.AccessControlRecord) {
	values := reflect.ValueOf(response)
	results := make([]*model.AccessControlRecord, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.AccessControlRecord{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}