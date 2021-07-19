// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// Subscription Service - Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification
// client.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// SubscriptionService type
type SubscriptionService struct {
	objectSet *client.SubscriptionObjectSet
}

// NewSubscriptionService - method to initialize "SubscriptionService"
func NewSubscriptionService(gs *NsGroupService) *SubscriptionService {
	objectSet := gs.client.GetSubscriptionObjectSet()
	return &SubscriptionService{objectSet: objectSet}
}

// GetSubscriptions - method returns a array of pointers of type "Subscriptions"
func (svc *SubscriptionService) GetSubscriptions(params *param.GetParams) ([]*nimbleos.Subscription, error) {
	subscriptionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return subscriptionResp, nil
}

// CreateSubscription - method creates a "Subscription"
func (svc *SubscriptionService) CreateSubscription(obj *nimbleos.Subscription) (*nimbleos.Subscription, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateSubscription: invalid parameter specified, %v", obj)
	}

	subscriptionResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return subscriptionResp, nil
}

// UpdateSubscription - method modifies  the "Subscription"
func (svc *SubscriptionService) UpdateSubscription(id string, obj *nimbleos.Subscription) (*nimbleos.Subscription, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateSubscription: invalid parameter specified, %v", obj)
	}

	subscriptionResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return subscriptionResp, nil
}

// GetSubscriptionById - method returns a pointer to "Subscription"
func (svc *SubscriptionService) GetSubscriptionById(id string) (*nimbleos.Subscription, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetSubscriptionById: invalid parameter specified, %v", id)
	}

	subscriptionResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return subscriptionResp, nil
}

// DeleteSubscription - deletes the "Subscription"
func (svc *SubscriptionService) DeleteSubscription(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteSubscription: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
