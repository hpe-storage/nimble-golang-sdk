// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Event Service - View events.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// EventService type 
type EventService struct {
	objectSet *client.EventObjectSet
}

// NewEventService - method to initialize "EventService" 
func NewEventService(gs *NsGroupService) (*EventService) {
	objectSet := gs.client.GetEventObjectSet()
	return &EventService{objectSet: objectSet}
}

// GetEvents - method returns a array of pointers of type "Events"
func (svc *EventService) GetEvents(params *util.GetParams) ([]*model.Event, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateEvent - method creates a "Event"
func (svc *EventService) CreateEvent(obj *model.Event) (*model.Event, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateEvent - method modifies  the "Event" 
func (svc *EventService) UpdateEvent(id string, obj *model.Event) (*model.Event, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetEventById - method returns a pointer to "Event"
func (svc *EventService) GetEventById(id string) (*model.Event, error) {
	return svc.objectSet.GetObject(id)
}

// GetEventByName - method returns a pointer "Event" 
func (svc *EventService) GetEventByName(name string) (*model.Event, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteEvent - deletes the "Event"
func (svc *EventService) DeleteEvent(id string) error {
	return svc.objectSet.DeleteObject(id)
}
