// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * View events.
 *
 */
const (
    eventPath = "events"
)

/**
 * EventObjectSet
*/
type EventObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Event object
func (objectSet *EventObjectSet) CreateObject(payload *model.Event) (*model.Event, error) {
	response, err := objectSet.Client.Post(eventPath, payload)
	return response.(*model.Event), err
}

// UpdateObject Modify existing Event object
func (objectSet *EventObjectSet) UpdateObject(id string, payload *model.Event) (*model.Event, error) {
	response, err := objectSet.Client.Put(eventPath, id, payload)
	return response.(*model.Event), err
}

// DeleteObject deletes the Event object with the specified ID
func (objectSet *EventObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(eventPath, id)
}

// GetObject returns a Event object with the given ID
func (objectSet *EventObjectSet) GetObject(id string) (*model.Event, error) {
	response, err := objectSet.Client.Get(eventPath, id, model.Event{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Event), err
}

// GetObjectList returns the list of Event objects
func (objectSet *EventObjectSet) GetObjectList() ([]*model.Event, error) {
	response, err := objectSet.Client.List(eventPath)
	return buildEventObjectSet(response), err
}

// GetObjectListFromParams returns the list of Event objects using the given params query info
func (objectSet *EventObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Event, error) {
	response, err := objectSet.Client.ListFromParams(eventPath, params)
	return buildEventObjectSet(response), err
}

// generated function to build the appropriate response types
func buildEventObjectSet(response interface{}) ([]*model.Event) {
	values := reflect.ValueOf(response)
	results := make([]*model.Event, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Event{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}