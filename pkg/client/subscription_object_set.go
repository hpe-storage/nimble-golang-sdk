package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification client.
 *
 */
const (
    subscriptionPath = "subscriptions"
)

/**
 * SubscriptionObjectSet
*/
type SubscriptionObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Subscription object
func (objectSet *SubscriptionObjectSet) CreateObject(payload *model.Subscription) (*model.Subscription, error) {
	response, err := objectSet.Client.Post(subscriptionPath, payload)
	return response.(*model.Subscription), err
}

// UpdateObject Modify existing Subscription object
func (objectSet *SubscriptionObjectSet) UpdateObject(id string, payload *model.Subscription) (*model.Subscription, error) {
	response, err := objectSet.Client.Put(subscriptionPath, id, payload)
	return response.(*model.Subscription), err
}

// DeleteObject deletes the Subscription object with the specified ID
func (objectSet *SubscriptionObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(subscriptionPath, id)
}

// GetObject returns a Subscription object with the given ID
func (objectSet *SubscriptionObjectSet) GetObject(id string) (*model.Subscription, error) {
	response, err := objectSet.Client.Get(subscriptionPath, id, model.Subscription{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Subscription), err
}

// GetObjectList returns the list of Subscription objects
func (objectSet *SubscriptionObjectSet) GetObjectList() ([]*model.Subscription, error) {
	response, err := objectSet.Client.List(subscriptionPath)
	return buildSubscriptionObjectSet(response), err
}

// GetObjectListFromParams returns the list of Subscription objects using the given params query info
func (objectSet *SubscriptionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Subscription, error) {
	response, err := objectSet.Client.ListFromParams(subscriptionPath, params)
	return buildSubscriptionObjectSet(response), err
}

// generated function to build the appropriate response types
func buildSubscriptionObjectSet(response interface{}) ([]*model.Subscription) {
	values := reflect.ValueOf(response)
	results := make([]*model.Subscription, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Subscription{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}