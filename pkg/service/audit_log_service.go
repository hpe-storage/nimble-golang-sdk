// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// AuditLog Service - View audit log.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetAuditLogsWithFields - method returns a array of pointers of type "AuditLog" 
func (svc *AuditLogService) GetAuditLogsWithFields(fields []string) ([]*model.AuditLog, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateAuditLog - method creates a "AuditLog"
func (svc *AuditLogService) CreateAuditLog(obj *model.AuditLog) (*model.AuditLog, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditAuditLog - method modifies  the "AuditLog" 
func (svc *AuditLogService) EditAuditLog(id string, obj *model.AuditLog) (*model.AuditLog, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyAuditLog - private method for more than one element check. 
func onlyAuditLog(objs []*model.AuditLog) (*model.AuditLog, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one AuditLog found with the given filter")
	}

	return objs[0], nil
}

 
// GetAuditLogsByID - method returns associative a array of pointers of type "AuditLog", filter by Id
func (svc *AuditLogService) GetAuditLogsByID(pool *model.Pool, fields []string) (map[string]*model.AuditLog, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetAuditLogs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.AuditLog)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetAuditLogById - method returns a pointer to "AuditLog"
func (svc *AuditLogService) GetAuditLogById(id string) (*model.AuditLog, error) {
	return svc.objectSet.GetObject(id)
}


