// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manage protection schedules used in protection templates.
 *
 */
const (
    protectionSchedulePath = "protection_schedules"
)

/**
 * ProtectionScheduleObjectSet
*/
type ProtectionScheduleObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ProtectionSchedule object
func (objectSet *ProtectionScheduleObjectSet) CreateObject(payload *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	response, err := objectSet.Client.Post(protectionSchedulePath, payload)
	return response.(*model.ProtectionSchedule), err
}

// UpdateObject Modify existing ProtectionSchedule object
func (objectSet *ProtectionScheduleObjectSet) UpdateObject(id string, payload *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	response, err := objectSet.Client.Put(protectionSchedulePath, id, payload)
	return response.(*model.ProtectionSchedule), err
}

// DeleteObject deletes the ProtectionSchedule object with the specified ID
func (objectSet *ProtectionScheduleObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(protectionSchedulePath, id)
}

// GetObject returns a ProtectionSchedule object with the given ID
func (objectSet *ProtectionScheduleObjectSet) GetObject(id string) (*model.ProtectionSchedule, error) {
	response, err := objectSet.Client.Get(protectionSchedulePath, id, model.ProtectionSchedule{})
	if response == nil {
		return nil, err
	}
	return response.(*model.ProtectionSchedule), err
}

// GetObjectList returns the list of ProtectionSchedule objects
func (objectSet *ProtectionScheduleObjectSet) GetObjectList() ([]*model.ProtectionSchedule, error) {
	response, err := objectSet.Client.List(protectionSchedulePath)
	return buildProtectionScheduleObjectSet(response), err
}

// GetObjectListFromParams returns the list of ProtectionSchedule objects using the given params query info
func (objectSet *ProtectionScheduleObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ProtectionSchedule, error) {
	response, err := objectSet.Client.ListFromParams(protectionSchedulePath, params)
	return buildProtectionScheduleObjectSet(response), err
}

// generated function to build the appropriate response types
func buildProtectionScheduleObjectSet(response interface{}) ([]*model.ProtectionSchedule) {
	values := reflect.ValueOf(response)
	results := make([]*model.ProtectionSchedule, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ProtectionSchedule{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}