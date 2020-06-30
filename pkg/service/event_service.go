// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Event Service - View events.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetEventsWithFields - method returns a array of pointers of type "Event" 
func (svc *EventService) GetEventsWithFields(fields []string) ([]*model.Event, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateEvent - method creates a "Event"
func (svc *EventService) CreateEvent(obj *model.Event) (*model.Event, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditEvent - method modifies  the "Event" 
func (svc *EventService) EditEvent(id string, obj *model.Event) (*model.Event, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyEvent - private method for more than one element check. 
func onlyEvent(objs []*model.Event) (*model.Event, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Event found with the given filter")
	}

	return objs[0], nil
}

 
// GetEventsByID - method returns associative a array of pointers of type "Event", filter by Id
func (svc *EventService) GetEventsByID(pool *model.Pool, fields []string) (map[string]*model.Event, error) {
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
	objs, err := svc.GetEvents(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Event)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetEventById - method returns a pointer to "Event"
func (svc *EventService) GetEventById(id string) (*model.Event, error) {
	return svc.objectSet.GetObject(id)
}

// GetEventsByName - method returns a associative array of pointers of type "Event", filter by name 
func (svc *EventService) GetEventsByName(pool *model.Pool, fields []string) (map[string]*model.Event, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetEvents(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Event)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlyEvent(objs)
}	

