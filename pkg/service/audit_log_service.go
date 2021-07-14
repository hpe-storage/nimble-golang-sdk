// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// AuditLog Service - View audit log.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// AuditLogService type
type AuditLogService struct {
    objectSet *client.AuditLogObjectSet
}

// NewAuditLogService - method to initialize "AuditLogService"
func NewAuditLogService(gs *NsGroupService) (*AuditLogService) {
    objectSet := gs.client.GetAuditLogObjectSet()
    return &AuditLogService{objectSet: objectSet}
}

// GetAuditLogs - method returns a array of pointers of type "AuditLogs"
func (svc *AuditLogService) GetAuditLogs(params *param.GetParams) ([]*nimbleos.AuditLog, error) {
    auditLogResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return auditLogResp,nil
}

// CreateAuditLog - method creates a "AuditLog"
func (svc *AuditLogService) CreateAuditLog(obj *nimbleos.AuditLog) (*nimbleos.AuditLog, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateAuditLog: invalid parameter specified, %v",obj)
    }

    auditLogResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return auditLogResp,nil
}

// UpdateAuditLog - method modifies  the "AuditLog"
func (svc *AuditLogService) UpdateAuditLog(id string, obj *nimbleos.AuditLog) (*nimbleos.AuditLog, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateAuditLog: invalid parameter specified, %v",obj)
    }

    auditLogResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return auditLogResp,nil
}

// GetAuditLogById - method returns a pointer to "AuditLog"
func (svc *AuditLogService) GetAuditLogById(id string) (*nimbleos.AuditLog, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetAuditLogById: invalid parameter specified, %v",id)
    }

    auditLogResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return auditLogResp,nil
}

// DeleteAuditLog - deletes the "AuditLog"
func (svc *AuditLogService) DeleteAuditLog(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteAuditLog: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

