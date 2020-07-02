// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Alarm Service - View alarms.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// AlarmService type 
type AlarmService struct {
	objectSet *client.AlarmObjectSet
}

// NewAlarmService - method to initialize "AlarmService" 
func NewAlarmService(gs *NsGroupService) (*AlarmService) {
	objectSet := gs.client.GetAlarmObjectSet()
	return &AlarmService{objectSet: objectSet}
}

// GetAlarms - method returns a array of pointers of type "Alarms"
func (svc *AlarmService) GetAlarms(params *util.GetParams) ([]*model.Alarm, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateAlarm - method creates a "Alarm"
func (svc *AlarmService) CreateAlarm(obj *model.Alarm) (*model.Alarm, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateAlarm - method modifies  the "Alarm" 
func (svc *AlarmService) UpdateAlarm(id string, obj *model.Alarm) (*model.Alarm, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetAlarmById - method returns a pointer to "Alarm"
func (svc *AlarmService) GetAlarmById(id string) (*model.Alarm, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteAlarm - deletes the "Alarm"
func (svc *AlarmService) DeleteAlarm(id string) error {
	return svc.objectSet.DeleteObject(id)
}
