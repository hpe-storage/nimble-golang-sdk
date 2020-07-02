// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// AuditLog Service - View audit log.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (svc *AuditLogService) GetAuditLogs(params *util.GetParams) ([]*model.AuditLog, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateAuditLog - method creates a "AuditLog"
func (svc *AuditLogService) CreateAuditLog(obj *model.AuditLog) (*model.AuditLog, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateAuditLog - method modifies  the "AuditLog" 
func (svc *AuditLogService) UpdateAuditLog(id string, obj *model.AuditLog) (*model.AuditLog, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetAuditLogById - method returns a pointer to "AuditLog"
func (svc *AuditLogService) GetAuditLogById(id string) (*model.AuditLog, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteAuditLog - deletes the "AuditLog"
func (svc *AuditLogService) DeleteAuditLog(id string) error {
	return svc.objectSet.DeleteObject(id)
}
