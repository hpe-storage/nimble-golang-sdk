// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *AccessControlRecordObjectSet) CreateObject(payload *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	newPayload, err := model.EncodeAccessControlRecord(payload)
	accessControlRecordObjectSetResp, err := objectSet.Client.Post(accessControlRecordPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if accessControlRecordObjectSetResp == nil {
		return nil, nil
	}

	return model.DecodeAccessControlRecord(accessControlRecordObjectSetResp)
}

// UpdateObject Modify existing AccessControlRecord object
func (objectSet *AccessControlRecordObjectSet) UpdateObject(id string, payload *model.AccessControlRecord) (*model.AccessControlRecord, error) {
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
func (objectSet *AccessControlRecordObjectSet) GetObject(id string) (*model.AccessControlRecord, error) {
	accessControlRecordObjectSetResp, err := objectSet.Client.Get(accessControlRecordPath, id, model.AccessControlRecord{})
	if err != nil {
		return nil, err
	}

	// null check
	if accessControlRecordObjectSetResp == nil {
		return nil, nil
	}
	return accessControlRecordObjectSetResp.(*model.AccessControlRecord), err
}

// GetObjectList returns the list of AccessControlRecord objects
func (objectSet *AccessControlRecordObjectSet) GetObjectList() ([]*model.AccessControlRecord, error) {
	accessControlRecordObjectSetResp, err := objectSet.Client.List(accessControlRecordPath)
	if err != nil {
		return nil, err
	}
	return buildAccessControlRecordObjectSet(accessControlRecordObjectSetResp), err
}

// GetObjectListFromParams returns the list of AccessControlRecord objects using the given params query info
func (objectSet *AccessControlRecordObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.AccessControlRecord, error) {
	accessControlRecordObjectSetResp, err := objectSet.Client.ListFromParams(accessControlRecordPath, params)
	if err != nil {
		return nil, err
	}
	return buildAccessControlRecordObjectSet(accessControlRecordObjectSetResp), err
}

// generated function to build the appropriate response types
func buildAccessControlRecordObjectSet(response interface{}) []*model.AccessControlRecord {
	values := reflect.ValueOf(response)
	results := make([]*model.AccessControlRecord, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.AccessControlRecord{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
