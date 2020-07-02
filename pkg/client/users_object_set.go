// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Represents users configured to manage the system.
const (
    userPath = "users"
)

// UserObjectSet
type UserObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new User object
func (objectSet *UserObjectSet) CreateObject(payload *model.User) (*model.User, error) {
	response, err := objectSet.Client.Post(userPath, payload)
	return response.(*model.User), err
}

// UpdateObject Modify existing User object
func (objectSet *UserObjectSet) UpdateObject(id string, payload *model.User) (*model.User, error) {
	response, err := objectSet.Client.Put(userPath, id, payload)
	return response.(*model.User), err
}

// DeleteObject deletes the User object with the specified ID
func (objectSet *UserObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(userPath, id)
}

// GetObject returns a User object with the given ID
func (objectSet *UserObjectSet) GetObject(id string) (*model.User, error) {
	response, err := objectSet.Client.Get(userPath, id, model.User{})
	if response == nil {
		return nil, err
	}
	return response.(*model.User), err
}

// GetObjectList returns the list of User objects
func (objectSet *UserObjectSet) GetObjectList() ([]*model.User, error) {
	response, err := objectSet.Client.List(userPath)
	return buildUserObjectSet(response), err
}

// GetObjectListFromParams returns the list of User objects using the given params query info
func (objectSet *UserObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.User, error) {
	response, err := objectSet.Client.ListFromParams(userPath, params)
	return buildUserObjectSet(response), err
}

// generated function to build the appropriate response types
func buildUserObjectSet(response interface{}) ([]*model.User) {
	values := reflect.ValueOf(response)
	results := make([]*model.User, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.User{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}