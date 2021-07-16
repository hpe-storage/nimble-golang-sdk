// Copyright 2021 Hewlett Packard Enterprise Development LP

package service

// Event Service - View events.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// EventService type
type EventService struct {
	objectSet *client.EventObjectSet
}

// NewEventService - method to initialize "EventService"
func NewEventService(gs *NsGroupService) *EventService {
	objectSet := gs.client.GetEventObjectSet()
	return &EventService{objectSet: objectSet}
}

// GetEvents - method returns a array of pointers of type "Events"
func (svc *EventService) GetEvents(params *param.GetParams) ([]*nimbleos.Event, error) {
	eventResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// CreateEvent - method creates a "Event"
func (svc *EventService) CreateEvent(obj *nimbleos.Event) (*nimbleos.Event, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateEvent: invalid parameter specified, %v", obj)
	}

	eventResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// UpdateEvent - method modifies  the "Event"
func (svc *EventService) UpdateEvent(id string, obj *nimbleos.Event) (*nimbleos.Event, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateEvent: invalid parameter specified, %v", obj)
	}

	eventResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// GetEventById - method returns a pointer to "Event"
func (svc *EventService) GetEventById(id string) (*nimbleos.Event, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetEventById: invalid parameter specified, %v", id)
	}

	eventResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// GetEventByName - method returns a pointer "Event"
func (svc *EventService) GetEventByName(name string) (*nimbleos.Event, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	eventResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(eventResp) == 0 {
		return nil, nil
	}

	return eventResp[0], nil
}

// DeleteEvent - deletes the "Event"
func (svc *EventService) DeleteEvent(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteEvent: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
