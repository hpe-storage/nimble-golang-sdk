// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manage access control records for access control records.
 *
 */
const (
	accessControlRecordPath = "access_control_records"
)

// AccessControlRecordObjectSet provides a wrapper to manage access control records from the client
type AccessControlRecordObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new access control record
func (objectSet *AccessControlRecordObjectSet) CreateObject(payload *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	response, err := objectSet.Client.Post(accessControlRecordPath, payload)
	return response.(*model.AccessControlRecord), err
}

// GetObject returns an access control record with the given ID
func (objectSet *AccessControlRecordObjectSet) GetObject(id string) (*model.AccessControlRecord, error) {
	response, err := objectSet.Client.Get(accessControlRecordPath, id, model.AccessControlRecord{})
	if response == nil {
		return nil, err
	}
	return response.(*model.AccessControlRecord), err
}

// DeleteObject deletes the access control record with the specified ID
func (objectSet *AccessControlRecordObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(accessControlRecordPath, id)
}

// GetObjectList returns the list of access control record objects
func (objectSet *AccessControlRecordObjectSet) GetObjectList() ([]*model.AccessControlRecord, error) {
	response, err := objectSet.Client.List(accessControlRecordPath)
	return buildAccessControlRecords(response), err
}

// GetObjectListFromParams returns the list of access control records objects using the given params query info
func (objectSet *AccessControlRecordObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.AccessControlRecord, error) {
	response, err := objectSet.Client.ListFromParams(volumePath, params)
	return buildAccessControlRecords(response), err
}

// generated function to build the appropriate response types
func buildAccessControlRecords(response interface{}) []*model.AccessControlRecord {
	values := reflect.ValueOf(response)
	results := make([]*model.AccessControlRecord, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.AccessControlRecord{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
