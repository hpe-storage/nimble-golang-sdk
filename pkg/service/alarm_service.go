// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Alarm Service - View alarms.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// AlarmService type
type AlarmService struct {
	objectSet *client.AlarmObjectSet
}

// NewAlarmService - method to initialize "AlarmService"
func NewAlarmService(gs *NsGroupService) *AlarmService {
	objectSet := gs.client.GetAlarmObjectSet()
	return &AlarmService{objectSet: objectSet}
}

// GetAlarms - method returns a array of pointers of type "Alarms"
func (svc *AlarmService) GetAlarms(params *util.GetParams) ([]*model.Alarm, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	alarmResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return alarmResp, nil
}

// CreateAlarm - method creates a "Alarm"
func (svc *AlarmService) CreateAlarm(obj *model.Alarm) (*model.Alarm, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	alarmResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return alarmResp, nil
}

// UpdateAlarm - method modifies  the "Alarm"
func (svc *AlarmService) UpdateAlarm(id string, obj *model.Alarm) (*model.Alarm, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	alarmResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return alarmResp, nil
}

// GetAlarmById - method returns a pointer to "Alarm"
func (svc *AlarmService) GetAlarmById(id string) (*model.Alarm, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	alarmResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return alarmResp, nil
}

// DeleteAlarm - deletes the "Alarm"
func (svc *AlarmService) DeleteAlarm(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
