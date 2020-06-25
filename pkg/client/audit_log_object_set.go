// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// View audit log.
const (
    auditLogPath = "audit_log"
)


// AuditLogObjectSet
type AuditLogObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new AuditLog object
func (objectSet *AuditLogObjectSet) CreateObject(payload *model.AuditLog) (*model.AuditLog, error) {
	response, err := objectSet.Client.Post(auditLogPath, payload)
	return response.(*model.AuditLog), err
}

// UpdateObject Modify existing AuditLog object
func (objectSet *AuditLogObjectSet) UpdateObject(id string, payload *model.AuditLog) (*model.AuditLog, error) {
	response, err := objectSet.Client.Put(auditLogPath, id, payload)
	return response.(*model.AuditLog), err
}

// DeleteObject deletes the AuditLog object with the specified ID
func (objectSet *AuditLogObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(auditLogPath, id)
}

// GetObject returns a AuditLog object with the given ID
func (objectSet *AuditLogObjectSet) GetObject(id string) (*model.AuditLog, error) {
	response, err := objectSet.Client.Get(auditLogPath, id, model.AuditLog{})
	if response == nil {
		return nil, err
	}
	return response.(*model.AuditLog), err
}

// GetObjectList returns the list of AuditLog objects
func (objectSet *AuditLogObjectSet) GetObjectList() ([]*model.AuditLog, error) {
	response, err := objectSet.Client.List(auditLogPath)
	return buildAuditLogObjectSet(response), err
}

// GetObjectListFromParams returns the list of AuditLog objects using the given params query info
func (objectSet *AuditLogObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.AuditLog, error) {
	response, err := objectSet.Client.ListFromParams(auditLogPath, params)
	return buildAuditLogObjectSet(response), err
}

// generated function to build the appropriate response types
func buildAuditLogObjectSet(response interface{}) ([]*model.AuditLog) {
	values := reflect.ValueOf(response)
	results := make([]*model.AuditLog, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.AuditLog{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}