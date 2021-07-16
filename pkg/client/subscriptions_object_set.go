// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification
// client.
const (
	subscriptionPath = "subscriptions"
)

// SubscriptionObjectSet
type SubscriptionObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Subscription object
func (objectSet *SubscriptionObjectSet) CreateObject(payload *nimbleos.Subscription) (*nimbleos.Subscription, error) {
	resp, err := objectSet.Client.Post(subscriptionPath, payload, &nimbleos.Subscription{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Subscription), err
}

// UpdateObject Modify existing Subscription object
func (objectSet *SubscriptionObjectSet) UpdateObject(id string, payload *nimbleos.Subscription) (*nimbleos.Subscription, error) {
	resp, err := objectSet.Client.Put(subscriptionPath, id, payload, &nimbleos.Subscription{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Subscription), err
}

// DeleteObject deletes the Subscription object with the specified ID
func (objectSet *SubscriptionObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(subscriptionPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Subscription object with the given ID
func (objectSet *SubscriptionObjectSet) GetObject(id string) (*nimbleos.Subscription, error) {
	resp, err := objectSet.Client.Get(subscriptionPath, id, &nimbleos.Subscription{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Subscription), err
}

// GetObjectList returns the list of Subscription objects
func (objectSet *SubscriptionObjectSet) GetObjectList() ([]*nimbleos.Subscription, error) {
	resp, err := objectSet.Client.List(subscriptionPath)
	if err != nil {
		return nil, err
	}
	return buildSubscriptionObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Subscription objects using the given params query info
func (objectSet *SubscriptionObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Subscription, error) {
	subscriptionObjectSetResp, err := objectSet.Client.ListFromParams(subscriptionPath, params)
	if err != nil {
		return nil, err
	}
	return buildSubscriptionObjectSet(subscriptionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSubscriptionObjectSet(response interface{}) []*nimbleos.Subscription {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Subscription, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Subscription{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
