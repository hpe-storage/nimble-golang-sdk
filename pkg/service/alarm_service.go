// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Alarm Service - View alarms.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetAlarmsWithFields - method returns a array of pointers of type "Alarm" 
func (svc *AlarmService) GetAlarmsWithFields(fields []string) ([]*model.Alarm, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateAlarm - method creates a "Alarm"
func (svc *AlarmService) CreateAlarm(obj *model.Alarm) (*model.Alarm, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditAlarm - method modifies  the "Alarm" 
func (svc *AlarmService) EditAlarm(id string, obj *model.Alarm) (*model.Alarm, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyAlarm - private method for more than one element check. 
func onlyAlarm(objs []*model.Alarm) (*model.Alarm, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Alarm found with the given filter")
	}

	return objs[0], nil
}

 
// GetAlarmsByID - method returns associative a array of pointers of type "Alarm", filter by Id
func (svc *AlarmService) GetAlarmsByID(pool *model.Pool, fields []string) (map[string]*model.Alarm, error) {
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
	objs, err := svc.GetAlarms(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Alarm)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetAlarmById - method returns a pointer to "Alarm"
func (svc *AlarmService) GetAlarmById(id string) (*model.Alarm, error) {
	return svc.objectSet.GetObject(id)
}


