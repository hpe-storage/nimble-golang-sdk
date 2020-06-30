// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Subscriber Service - Subscribers are websocket based notification clients that can subscribe to interesting operations and events and recieve notifications whenever the subscribed to operations and
// events happen on the array.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// SubscriberService type 
type SubscriberService struct {
	objectSet *client.SubscriberObjectSet
}

// NewSubscriberService - method to initialize "SubscriberService" 
func NewSubscriberService(gs *NsGroupService) (*SubscriberService) {
	objectSet := gs.client.GetSubscriberObjectSet()
	return &SubscriberService{objectSet: objectSet}
}

// GetSubscribers - method returns a array of pointers of type "Subscribers"
func (svc *SubscriberService) GetSubscribers(params *util.GetParams) ([]*model.Subscriber, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetSubscribersWithFields - method returns a array of pointers of type "Subscriber" 
func (svc *SubscriberService) GetSubscribersWithFields(fields []string) ([]*model.Subscriber, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSubscriber - method creates a "Subscriber"
func (svc *SubscriberService) CreateSubscriber(obj *model.Subscriber) (*model.Subscriber, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSubscriber - method modifies  the "Subscriber" 
func (svc *SubscriberService) EditSubscriber(id string, obj *model.Subscriber) (*model.Subscriber, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySubscriber - private method for more than one element check. 
func onlySubscriber(objs []*model.Subscriber) (*model.Subscriber, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Subscriber found with the given filter")
	}

	return objs[0], nil
}

 
// GetSubscribersByID - method returns associative a array of pointers of type "Subscriber", filter by Id
func (svc *SubscriberService) GetSubscribersByID(pool *model.Pool, fields []string) (map[string]*model.Subscriber, error) {
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
	objs, err := svc.GetSubscribers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Subscriber)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetSubscriberById - method returns a pointer to "Subscriber"
func (svc *SubscriberService) GetSubscriberById(id string) (*model.Subscriber, error) {
	return svc.objectSet.GetObject(id)
}


