// Copyright 2021 Hewlett Packard Enterprise Development LP

package service

// Subscriber Service - Subscribers are websocket based notification clients that can subscribe to interesting operations and events and recieve notifications whenever the subscribed to operations and
// events happen on the array.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// SubscriberService type
type SubscriberService struct {
	objectSet *client.SubscriberObjectSet
}

// NewSubscriberService - method to initialize "SubscriberService"
func NewSubscriberService(gs *NsGroupService) *SubscriberService {
	objectSet := gs.client.GetSubscriberObjectSet()
	return &SubscriberService{objectSet: objectSet}
}

// GetSubscribers - method returns a array of pointers of type "Subscribers"
func (svc *SubscriberService) GetSubscribers(params *param.GetParams) ([]*nimbleos.Subscriber, error) {
	subscriberResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return subscriberResp, nil
}

// CreateSubscriber - method creates a "Subscriber"
func (svc *SubscriberService) CreateSubscriber(obj *nimbleos.Subscriber) (*nimbleos.Subscriber, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateSubscriber: invalid parameter specified, %v", obj)
	}

	subscriberResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return subscriberResp, nil
}

// UpdateSubscriber - method modifies  the "Subscriber"
func (svc *SubscriberService) UpdateSubscriber(id string, obj *nimbleos.Subscriber) (*nimbleos.Subscriber, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateSubscriber: invalid parameter specified, %v", obj)
	}

	subscriberResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return subscriberResp, nil
}

// GetSubscriberById - method returns a pointer to "Subscriber"
func (svc *SubscriberService) GetSubscriberById(id string) (*nimbleos.Subscriber, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetSubscriberById: invalid parameter specified, %v", id)
	}

	subscriberResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return subscriberResp, nil
}

// DeleteSubscriber - deletes the "Subscriber"
func (svc *SubscriberService) DeleteSubscriber(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteSubscriber: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
