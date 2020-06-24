package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manages the password policies configured for the group.
 *
 */
const (
    userPolicyPath = "user_policies"
)

/**
 * UserPolicyObjectSet
*/
type UserPolicyObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new UserPolicy object
func (objectSet *UserPolicyObjectSet) CreateObject(payload *model.UserPolicy) (*model.UserPolicy, error) {
	response, err := objectSet.Client.Post(userPolicyPath, payload)
	return response.(*model.UserPolicy), err
}

// UpdateObject Modify existing UserPolicy object
func (objectSet *UserPolicyObjectSet) UpdateObject(id string, payload *model.UserPolicy) (*model.UserPolicy, error) {
	response, err := objectSet.Client.Put(userPolicyPath, id, payload)
	return response.(*model.UserPolicy), err
}

// DeleteObject deletes the UserPolicy object with the specified ID
func (objectSet *UserPolicyObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(userPolicyPath, id)
}

// GetObject returns a UserPolicy object with the given ID
func (objectSet *UserPolicyObjectSet) GetObject(id string) (*model.UserPolicy, error) {
	response, err := objectSet.Client.Get(userPolicyPath, id, model.UserPolicy{})
	if response == nil {
		return nil, err
	}
	return response.(*model.UserPolicy), err
}

// GetObjectList returns the list of UserPolicy objects
func (objectSet *UserPolicyObjectSet) GetObjectList() ([]*model.UserPolicy, error) {
	response, err := objectSet.Client.List(userPolicyPath)
	return buildUserPolicyObjectSet(response), err
}

// GetObjectListFromParams returns the list of UserPolicy objects using the given params query info
func (objectSet *UserPolicyObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.UserPolicy, error) {
	response, err := objectSet.Client.ListFromParams(userPolicyPath, params)
	return buildUserPolicyObjectSet(response), err
}

// generated function to build the appropriate response types
func buildUserPolicyObjectSet(response interface{}) ([]*model.UserPolicy) {
	values := reflect.ValueOf(response)
	results := make([]*model.UserPolicy, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.UserPolicy{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}