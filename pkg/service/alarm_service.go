// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Alarm Service - View alarms.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *AlarmService) GetAlarms(params *param.GetParams) ([]*nimbleos.Alarm, error) {
    alarmResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return alarmResp,nil
}

// CreateAlarm - method creates a "Alarm"
func (svc *AlarmService) CreateAlarm(obj *nimbleos.Alarm) (*nimbleos.Alarm, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateAlarm: invalid parameter specified, %v",obj)
    }

    alarmResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return alarmResp,nil
}

// UpdateAlarm - method modifies  the "Alarm"
func (svc *AlarmService) UpdateAlarm(id string, obj *nimbleos.Alarm) (*nimbleos.Alarm, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateAlarm: invalid parameter specified, %v",obj)
    }

    alarmResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return alarmResp,nil
}

// GetAlarmById - method returns a pointer to "Alarm"
func (svc *AlarmService) GetAlarmById(id string) (*nimbleos.Alarm, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetAlarmById: invalid parameter specified, %v",id)
    }

    alarmResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return alarmResp,nil
}

// DeleteAlarm - deletes the "Alarm"
func (svc *AlarmService) DeleteAlarm(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteAlarm: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

// AcknowledgeAlarm - acknowledge an alarm.
//   Required parameters:
//       id - ID of the acknowledged alarm.

//   Optional parameters:
//       remindEvery - Notification frequency unit.
//       remindEveryUnit - Period unit.

func (svc *AlarmService) AcknowledgeAlarm( id string , remindEvery *uint64 , remindEveryUnit *nimbleos.NsPeriodUnit )(error) {

     if  len(id) == 0  {
return fmt.Errorf("AcknowledgeAlarm: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.Acknowledge( &id , remindEvery , remindEveryUnit )
     return err
}
// UnacknowledgeAlarm - unacknowledge an alarm.
//   Required parameters:
//       id - ID of the acknowledged alarm.


func (svc *AlarmService) UnacknowledgeAlarm( id string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("UnacknowledgeAlarm: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.Unacknowledge( &id )
     return err
}
