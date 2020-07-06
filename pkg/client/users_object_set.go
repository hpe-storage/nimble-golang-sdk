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
	userObjectSetResp, err := objectSet.Client.Post(userPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if userObjectSetResp == nil {
		return nil,nil
	}
	return userObjectSetResp.(*model.User), err
}

// UpdateObject Modify existing User object
func (objectSet *UserObjectSet) UpdateObject(id string, payload *model.User) (*model.User, error) {
	userObjectSetResp, err := objectSet.Client.Put(userPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if userObjectSetResp == nil {
		return nil,nil
	}
	return userObjectSetResp.(*model.User), err
}

// DeleteObject deletes the User object with the specified ID
func (objectSet *UserObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(userPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a User object with the given ID
func (objectSet *UserObjectSet) GetObject(id string) (*model.User, error) {
	userObjectSetResp, err := objectSet.Client.Get(userPath, id, model.User{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if userObjectSetResp == nil {
		return nil,nil
	}
	return userObjectSetResp.(*model.User), err
}

// GetObjectList returns the list of User objects
func (objectSet *UserObjectSet) GetObjectList() ([]*model.User, error) {
	userObjectSetResp, err := objectSet.Client.List(userPath)
	if err != nil {
		return nil, err
	}
	return buildUserObjectSet(userObjectSetResp), err
}

// GetObjectListFromParams returns the list of User objects using the given params query info
func (objectSet *UserObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.User, error) {
	userObjectSetResp, err := objectSet.Client.ListFromParams(userPath, params)
	if err != nil {
		return nil, err
	}
	return buildUserObjectSet(userObjectSetResp), err
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
