// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// View alarms.
const (
    alarmPath = "alarms"
)

// AlarmObjectSet
type AlarmObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Alarm object
func (objectSet *AlarmObjectSet) CreateObject(payload *model.Alarm) (*model.Alarm, error) {
	alarmObjectSetResp, err := objectSet.Client.Post(alarmPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if alarmObjectSetResp == nil {
		return nil,nil
	}
	return alarmObjectSetResp.(*model.Alarm), err
}

// UpdateObject Modify existing Alarm object
func (objectSet *AlarmObjectSet) UpdateObject(id string, payload *model.Alarm) (*model.Alarm, error) {
	alarmObjectSetResp, err := objectSet.Client.Put(alarmPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if alarmObjectSetResp == nil {
		return nil,nil
	}
	return alarmObjectSetResp.(*model.Alarm), err
}

// DeleteObject deletes the Alarm object with the specified ID
func (objectSet *AlarmObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(alarmPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a Alarm object with the given ID
func (objectSet *AlarmObjectSet) GetObject(id string) (*model.Alarm, error) {
	alarmObjectSetResp, err := objectSet.Client.Get(alarmPath, id, model.Alarm{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if alarmObjectSetResp == nil {
		return nil,nil
	}
	return alarmObjectSetResp.(*model.Alarm), err
}

// GetObjectList returns the list of Alarm objects
func (objectSet *AlarmObjectSet) GetObjectList() ([]*model.Alarm, error) {
	alarmObjectSetResp, err := objectSet.Client.List(alarmPath)
	if err != nil {
		return nil, err
	}
	return buildAlarmObjectSet(alarmObjectSetResp), err
}

// GetObjectListFromParams returns the list of Alarm objects using the given params query info
func (objectSet *AlarmObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Alarm, error) {
	alarmObjectSetResp, err := objectSet.Client.ListFromParams(alarmPath, params)
	if err != nil {
		return nil, err
	}
	return buildAlarmObjectSet(alarmObjectSetResp), err
}

// generated function to build the appropriate response types
func buildAlarmObjectSet(response interface{}) ([]*model.Alarm) {
	values := reflect.ValueOf(response)
	results := make([]*model.Alarm, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Alarm{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}