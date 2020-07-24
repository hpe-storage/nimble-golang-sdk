// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *AccessControlRecordObjectSet) CreateObject(payload *nimbleos.AccessControlRecord) (*nimbleos.AccessControlRecord, error) {
	resp, err := objectSet.Client.Post(accessControlRecordPath, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return resp.(*nimbleos.AccessControlRecord), err
}

// UpdateObject Modify existing AccessControlRecord object
func (objectSet *AccessControlRecordObjectSet) UpdateObject(id string, payload *nimbleos.AccessControlRecord) (*nimbleos.AccessControlRecord, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on AccessControlRecord")
}

// DeleteObject deletes the AccessControlRecord object with the specified ID
func (objectSet *AccessControlRecordObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(accessControlRecordPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a AccessControlRecord object with the given ID
func (objectSet *AccessControlRecordObjectSet) GetObject(id string) (*nimbleos.AccessControlRecord, error) {
	resp, err := objectSet.Client.Get(accessControlRecordPath, id, nimbleos.AccessControlRecord{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.AccessControlRecord), err
}

// GetObjectList returns the list of AccessControlRecord objects
func (objectSet *AccessControlRecordObjectSet) GetObjectList() ([]*nimbleos.AccessControlRecord, error) {
	resp, err := objectSet.Client.List(accessControlRecordPath)
	if err != nil {
		return nil, err
	}
	return buildAccessControlRecordObjectSet(resp), err
}

// GetObjectListFromParams returns the list of AccessControlRecord objects using the given params query info
func (objectSet *AccessControlRecordObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.AccessControlRecord, error) {
	accessControlRecordObjectSetResp, err := objectSet.Client.ListFromParams(accessControlRecordPath, params)
	if err != nil {
		return nil, err
	}
	return buildAccessControlRecordObjectSet(accessControlRecordObjectSetResp), err
}

// generated function to build the appropriate response types
func buildAccessControlRecordObjectSet(response interface{}) []*nimbleos.AccessControlRecord {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.AccessControlRecord, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.AccessControlRecord{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
