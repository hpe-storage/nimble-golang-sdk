// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Subscribers are websocket based notification clients that can subscribe to interesting operations and events and recieve notifications whenever the subscribed to operations and events happen on the array.
 *
 */
const (
    subscriberPath = "subscribers"
)

/**
 * SubscriberObjectSet
*/
type SubscriberObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Subscriber object
func (objectSet *SubscriberObjectSet) CreateObject(payload *model.Subscriber) (*model.Subscriber, error) {
	response, err := objectSet.Client.Post(subscriberPath, payload)
	return response.(*model.Subscriber), err
}

// UpdateObject Modify existing Subscriber object
func (objectSet *SubscriberObjectSet) UpdateObject(id string, payload *model.Subscriber) (*model.Subscriber, error) {
	response, err := objectSet.Client.Put(subscriberPath, id, payload)
	return response.(*model.Subscriber), err
}

// DeleteObject deletes the Subscriber object with the specified ID
func (objectSet *SubscriberObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(subscriberPath, id)
}

// GetObject returns a Subscriber object with the given ID
func (objectSet *SubscriberObjectSet) GetObject(id string) (*model.Subscriber, error) {
	response, err := objectSet.Client.Get(subscriberPath, id, model.Subscriber{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Subscriber), err
}

// GetObjectList returns the list of Subscriber objects
func (objectSet *SubscriberObjectSet) GetObjectList() ([]*model.Subscriber, error) {
	response, err := objectSet.Client.List(subscriberPath)
	return buildSubscriberObjectSet(response), err
}

// GetObjectListFromParams returns the list of Subscriber objects using the given params query info
func (objectSet *SubscriberObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Subscriber, error) {
	response, err := objectSet.Client.ListFromParams(subscriberPath, params)
	return buildSubscriberObjectSet(response), err
}

// generated function to build the appropriate response types
func buildSubscriberObjectSet(response interface{}) ([]*model.Subscriber) {
	values := reflect.ValueOf(response)
	results := make([]*model.Subscriber, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Subscriber{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}