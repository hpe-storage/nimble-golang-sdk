// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Event Service - View events.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (svc *EventService) GetEvents(params *util.GetParams) ([]*model.Event, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	eventResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// CreateEvent - method creates a "Event"
func (svc *EventService) CreateEvent(obj *model.Event) (*model.Event, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	eventResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// UpdateEvent - method modifies  the "Event"
func (svc *EventService) UpdateEvent(id string, obj *model.Event) (*model.Event, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	eventResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// GetEventById - method returns a pointer to "Event"
func (svc *EventService) GetEventById(id string) (*model.Event, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	eventResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return eventResp, nil
}

// GetEventByName - method returns a pointer "Event"
func (svc *EventService) GetEventByName(name string) (*model.Event, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
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
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
