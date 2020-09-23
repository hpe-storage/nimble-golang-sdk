// Copyright 2020 Hewlett Packard Enterprise Development LP

// update true
package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
	"strings"
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
func (objectSet *AlarmObjectSet) CreateObject(payload *nimbleos.Alarm) (*nimbleos.Alarm, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Alarm")
}

// UpdateObject Modify existing Alarm object
func (objectSet *AlarmObjectSet) UpdateObject(id string, payload *nimbleos.Alarm) (*nimbleos.Alarm, error) {
	resp, err := objectSet.Client.Put(alarmPath, id, payload, &nimbleos.Alarm{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)

			}
		}
		return nil, err
	}

	return resp.(*nimbleos.Alarm), err
}

// DeleteObject deletes the Alarm object with the specified ID
func (objectSet *AlarmObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(alarmPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Alarm object with the given ID
func (objectSet *AlarmObjectSet) GetObject(id string) (*nimbleos.Alarm, error) {
	resp, err := objectSet.Client.Get(alarmPath, id, &nimbleos.Alarm{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Alarm), err
}

// GetObjectList returns the list of Alarm objects
func (objectSet *AlarmObjectSet) GetObjectList() ([]*nimbleos.Alarm, error) {
	resp, err := objectSet.Client.List(alarmPath)
	if err != nil {
		return nil, err
	}
	return buildAlarmObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Alarm objects using the given params query info
func (objectSet *AlarmObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Alarm, error) {
	alarmObjectSetResp, err := objectSet.Client.ListFromParams(alarmPath, params)
	if err != nil {
		return nil, err
	}
	return buildAlarmObjectSet(alarmObjectSetResp), err
}

// generated function to build the appropriate response types
func buildAlarmObjectSet(response interface{}) []*nimbleos.Alarm {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Alarm, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Alarm{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Acknowledge - Acknowledge an alarm.
func (objectSet *AlarmObjectSet) Acknowledge(id *string, remindEvery *uint64, remindEveryUnit *nimbleos.NsPeriodUnit) error {

	acknowledgeUri := alarmPath
	acknowledgeUri = acknowledgeUri + "/" + *id
	acknowledgeUri = acknowledgeUri + "/actions/" + "acknowledge"

	payload := &struct {
		Id              *string                `json:"id,omitempty"`
		RemindEvery     *uint64                `json:"remind_every,omitempty"`
		RemindEveryUnit *nimbleos.NsPeriodUnit `json:"remind_every_unit,omitempty"`
	}{
		id,
		remindEvery,
		remindEveryUnit,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(acknowledgeUri, payload, &emptyStruct)
	return err
}

// Unacknowledge - Unacknowledge an alarm.
func (objectSet *AlarmObjectSet) Unacknowledge(id *string) error {

	unacknowledgeUri := alarmPath
	unacknowledgeUri = unacknowledgeUri + "/" + *id
	unacknowledgeUri = unacknowledgeUri + "/actions/" + "unacknowledge"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(unacknowledgeUri, payload, &emptyStruct)
	return err
}
