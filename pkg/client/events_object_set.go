// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (objectSet *EventObjectSet) CreateObject(payload *nimbleos.Event) (*nimbleos.Event, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Event")
}

// UpdateObject Modify existing Event object
func (objectSet *EventObjectSet) UpdateObject(id string, payload *nimbleos.Event) (*nimbleos.Event, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Event")
}

// DeleteObject deletes the Event object with the specified ID
func (objectSet *EventObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Event")
}

// GetObject returns a Event object with the given ID
func (objectSet *EventObjectSet) GetObject(id string) (*nimbleos.Event, error) {
	resp, err := objectSet.Client.Get(eventPath, id, &nimbleos.Event{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Event), err
}

// GetObjectList returns the list of Event objects
func (objectSet *EventObjectSet) GetObjectList() ([]*nimbleos.Event, error) {
	resp, err := objectSet.Client.List(eventPath)
	if err != nil {
		return nil, err
	}
	return buildEventObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Event objects using the given params query info
func (objectSet *EventObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Event, error) {
	eventObjectSetResp, err := objectSet.Client.ListFromParams(eventPath, params)
	if err != nil {
		return nil, err
	}
	return buildEventObjectSet(eventObjectSetResp), err
}

// generated function to build the appropriate response types
func buildEventObjectSet(response interface{}) []*nimbleos.Event {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Event, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Event{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
