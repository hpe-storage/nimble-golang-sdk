// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
)

// View events.
const (
	eventPath = "events"
)

// EventObjectSet
type EventObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Event object
func (objectSet *EventObjectSet) CreateObject(payload *model.Event) (*model.Event, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Event")
}

// UpdateObject Modify existing Event object
func (objectSet *EventObjectSet) UpdateObject(id string, payload *model.Event) (*model.Event, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Event")
}

// DeleteObject deletes the Event object with the specified ID
func (objectSet *EventObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Event")
}

// GetObject returns a Event object with the given ID
func (objectSet *EventObjectSet) GetObject(id string) (*model.Event, error) {
	eventObjectSetResp, err := objectSet.Client.Get(eventPath, id, model.Event{})
	if err != nil {
		return nil, err
	}

	// null check
	if eventObjectSetResp == nil {
		return nil, nil
	}
	return eventObjectSetResp.(*model.Event), err
}

// GetObjectList returns the list of Event objects
func (objectSet *EventObjectSet) GetObjectList() ([]*model.Event, error) {
	eventObjectSetResp, err := objectSet.Client.List(eventPath)
	if err != nil {
		return nil, err
	}
	return buildEventObjectSet(eventObjectSetResp), err
}

// GetObjectListFromParams returns the list of Event objects using the given params query info
func (objectSet *EventObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Event, error) {
	eventObjectSetResp, err := objectSet.Client.ListFromParams(eventPath, params)
	if err != nil {
		return nil, err
	}
	return buildEventObjectSet(eventObjectSetResp), err
}

// generated function to build the appropriate response types
func buildEventObjectSet(response interface{}) []*model.Event {
	values := reflect.ValueOf(response)
	results := make([]*model.Event, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Event{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
