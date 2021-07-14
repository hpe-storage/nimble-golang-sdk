// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

    "fmt"
        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

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
func (objectSet *AuditLogObjectSet) CreateObject(payload *nimbleos.AuditLog) (*nimbleos.AuditLog, error) {
    return nil, fmt.Errorf("Unsupported operation 'create' on AuditLog")
}

// UpdateObject Modify existing AuditLog object
func (objectSet *AuditLogObjectSet) UpdateObject(id string, payload *nimbleos.AuditLog) (*nimbleos.AuditLog, error) {
    return nil, fmt.Errorf("Unsupported operation 'update' on AuditLog")
}

// DeleteObject deletes the AuditLog object with the specified ID
func (objectSet *AuditLogObjectSet) DeleteObject(id string) error {
    return fmt.Errorf("Unsupported operation 'delete' on AuditLog")
}

// GetObject returns a AuditLog object with the given ID
func (objectSet *AuditLogObjectSet) GetObject(id string) (*nimbleos.AuditLog, error) {
    resp, err:= objectSet.Client.Get(auditLogPath, id, &nimbleos.AuditLog{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.AuditLog), err
}

// GetObjectList returns the list of AuditLog objects
func (objectSet *AuditLogObjectSet) GetObjectList() ([]*nimbleos.AuditLog, error) {
    resp, err:= objectSet.Client.List(auditLogPath)
    if err != nil {
        return nil, err
    }
    return buildAuditLogObjectSet(resp), err
}

// GetObjectListFromParams returns the list of AuditLog objects using the given params query info
func (objectSet *AuditLogObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.AuditLog, error) {
    auditLogObjectSetResp,err:= objectSet.Client.ListFromParams(auditLogPath, params)
    if err != nil {
        return nil, err
    }
    return buildAuditLogObjectSet(auditLogObjectSetResp), err
}
// generated function to build the appropriate response types
func buildAuditLogObjectSet(response interface{}) ([]*nimbleos.AuditLog) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.AuditLog, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.AuditLog{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

