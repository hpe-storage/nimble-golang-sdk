// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Subscription Service - Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification
// client.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// SubscriptionService type 
type SubscriptionService struct {
	objectSet *client.SubscriptionObjectSet
}

// NewSubscriptionService - method to initialize "SubscriptionService" 
func NewSubscriptionService(gs *NsGroupService) (*SubscriptionService) {
	objectSet := gs.client.GetSubscriptionObjectSet()
	return &SubscriptionService{objectSet: objectSet}
}

// GetSubscriptions - method returns a array of pointers of type "Subscriptions"
func (svc *SubscriptionService) GetSubscriptions(params *util.GetParams) ([]*model.Subscription, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetSubscriptionsWithFields - method returns a array of pointers of type "Subscription" 
func (svc *SubscriptionService) GetSubscriptionsWithFields(fields []string) ([]*model.Subscription, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSubscription - method creates a "Subscription"
func (svc *SubscriptionService) CreateSubscription(obj *model.Subscription) (*model.Subscription, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSubscription - method modifies  the "Subscription" 
func (svc *SubscriptionService) EditSubscription(id string, obj *model.Subscription) (*model.Subscription, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySubscription - private method for more than one element check. 
func onlySubscription(objs []*model.Subscription) (*model.Subscription, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Subscription found with the given filter")
	}

	return objs[0], nil
}

 
// GetSubscriptionsByID - method returns associative a array of pointers of type "Subscription", filter by Id
func (svc *SubscriptionService) GetSubscriptionsByID(pool *model.Pool, fields []string) (map[string]*model.Subscription, error) {
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
	objs, err := svc.GetSubscriptions(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Subscription)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetSubscriptionById - method returns a pointer to "Subscription"
func (svc *SubscriptionService) GetSubscriptionById(id string) (*model.Subscription, error) {
	return svc.objectSet.GetObject(id)
}


