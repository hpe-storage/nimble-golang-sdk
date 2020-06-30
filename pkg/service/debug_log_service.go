// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// DebugLog Service - Method to help log events from outside of storage array to provide context for troubleshooting host-side or array-side issues.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// DebugLogService type 
type DebugLogService struct {
	objectSet *client.DebugLogObjectSet
}

// NewDebugLogService - method to initialize "DebugLogService" 
func NewDebugLogService(gs *NsGroupService) (*DebugLogService) {
	objectSet := gs.client.GetDebugLogObjectSet()
	return &DebugLogService{objectSet: objectSet}
}

// GetDebugLogs - method returns a array of pointers of type "DebugLogs"
func (svc *DebugLogService) GetDebugLogs(params *util.GetParams) ([]*model.DebugLog, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetDebugLogsWithFields - method returns a array of pointers of type "DebugLog" 
func (svc *DebugLogService) GetDebugLogsWithFields(fields []string) ([]*model.DebugLog, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateDebugLog - method creates a "DebugLog"
func (svc *DebugLogService) CreateDebugLog(obj *model.DebugLog) (*model.DebugLog, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditDebugLog - method modifies  the "DebugLog" 
func (svc *DebugLogService) EditDebugLog(id string, obj *model.DebugLog) (*model.DebugLog, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyDebugLog - private method for more than one element check. 
func onlyDebugLog(objs []*model.DebugLog) (*model.DebugLog, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one DebugLog found with the given filter")
	}

	return objs[0], nil
}


