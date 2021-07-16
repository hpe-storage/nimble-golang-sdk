// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Subscribers are websocket based notification clients that can subscribe to interesting operations and events and recieve notifications whenever the subscribed to operations and
// events happen on the array.
const (
	subscriberPath = "subscribers"
)

// SubscriberObjectSet
type SubscriberObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Subscriber object
func (objectSet *SubscriberObjectSet) CreateObject(payload *nimbleos.Subscriber) (*nimbleos.Subscriber, error) {
	resp, err := objectSet.Client.Post(subscriberPath, payload, &nimbleos.Subscriber{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Subscriber), err
}

// UpdateObject Modify existing Subscriber object
func (objectSet *SubscriberObjectSet) UpdateObject(id string, payload *nimbleos.Subscriber) (*nimbleos.Subscriber, error) {
	resp, err := objectSet.Client.Put(subscriberPath, id, payload, &nimbleos.Subscriber{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Subscriber), err
}

// DeleteObject deletes the Subscriber object with the specified ID
func (objectSet *SubscriberObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(subscriberPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Subscriber object with the given ID
func (objectSet *SubscriberObjectSet) GetObject(id string) (*nimbleos.Subscriber, error) {
	resp, err := objectSet.Client.Get(subscriberPath, id, &nimbleos.Subscriber{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Subscriber), err
}

// GetObjectList returns the list of Subscriber objects
func (objectSet *SubscriberObjectSet) GetObjectList() ([]*nimbleos.Subscriber, error) {
	resp, err := objectSet.Client.List(subscriberPath)
	if err != nil {
		return nil, err
	}
	return buildSubscriberObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Subscriber objects using the given params query info
func (objectSet *SubscriberObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Subscriber, error) {
	subscriberObjectSetResp, err := objectSet.Client.ListFromParams(subscriberPath, params)
	if err != nil {
		return nil, err
	}
	return buildSubscriberObjectSet(subscriberObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSubscriberObjectSet(response interface{}) []*nimbleos.Subscriber {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Subscriber, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Subscriber{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
