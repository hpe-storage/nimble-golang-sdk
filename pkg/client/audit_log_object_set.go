// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
	auditLogObjectSetResp, err := objectSet.Client.Post(auditLogPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if auditLogObjectSetResp == nil {
		return nil,nil
	}
	return auditLogObjectSetResp.(*model.AuditLog), err
}

// UpdateObject Modify existing AuditLog object
func (objectSet *AuditLogObjectSet) UpdateObject(id string, payload *model.AuditLog) (*model.AuditLog, error) {
	auditLogObjectSetResp, err := objectSet.Client.Put(auditLogPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if auditLogObjectSetResp == nil {
		return nil,nil
	}
	return auditLogObjectSetResp.(*model.AuditLog), err
}

// DeleteObject deletes the AuditLog object with the specified ID
func (objectSet *AuditLogObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(auditLogPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a AuditLog object with the given ID
func (objectSet *AuditLogObjectSet) GetObject(id string) (*model.AuditLog, error) {
	auditLogObjectSetResp, err := objectSet.Client.Get(auditLogPath, id, model.AuditLog{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if auditLogObjectSetResp == nil {
		return nil,nil
	}
	return auditLogObjectSetResp.(*model.AuditLog), err
}

// GetObjectList returns the list of AuditLog objects
func (objectSet *AuditLogObjectSet) GetObjectList() ([]*model.AuditLog, error) {
	auditLogObjectSetResp, err := objectSet.Client.List(auditLogPath)
	if err != nil {
		return nil, err
	}
	return buildAuditLogObjectSet(auditLogObjectSetResp), err
}

// GetObjectListFromParams returns the list of AuditLog objects using the given params query info
func (objectSet *AuditLogObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.AuditLog, error) {
	auditLogObjectSetResp, err := objectSet.Client.ListFromParams(auditLogPath, params)
	if err != nil {
		return nil, err
	}
	return buildAuditLogObjectSet(auditLogObjectSetResp), err
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