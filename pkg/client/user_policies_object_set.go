// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manages the password policies configured for the group.
const (
    userPolicyPath = "user_policies"
)

// UserPolicyObjectSet
type UserPolicyObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new UserPolicy object
func (objectSet *UserPolicyObjectSet) CreateObject(payload *model.UserPolicy) (*model.UserPolicy, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on UserPolicy")
}

// UpdateObject Modify existing UserPolicy object
func (objectSet *UserPolicyObjectSet) UpdateObject(id string, payload *model.UserPolicy) (*model.UserPolicy, error) {
	userPolicyObjectSetResp, err := objectSet.Client.Put(userPolicyPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if userPolicyObjectSetResp == nil {
		return nil,nil
	}
	return userPolicyObjectSetResp.(*model.UserPolicy), err
}

// DeleteObject deletes the UserPolicy object with the specified ID
func (objectSet *UserPolicyObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on UserPolicy")
}

// GetObject returns a UserPolicy object with the given ID
func (objectSet *UserPolicyObjectSet) GetObject(id string) (*model.UserPolicy, error) {
	userPolicyObjectSetResp, err := objectSet.Client.Get(userPolicyPath, id, model.UserPolicy{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if userPolicyObjectSetResp == nil {
		return nil,nil
	}
	return userPolicyObjectSetResp.(*model.UserPolicy), err
}

// GetObjectList returns the list of UserPolicy objects
func (objectSet *UserPolicyObjectSet) GetObjectList() ([]*model.UserPolicy, error) {
	userPolicyObjectSetResp, err := objectSet.Client.List(userPolicyPath)
	if err != nil {
		return nil, err
	}
	return buildUserPolicyObjectSet(userPolicyObjectSetResp), err
}

// GetObjectListFromParams returns the list of UserPolicy objects using the given params query info
func (objectSet *UserPolicyObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.UserPolicy, error) {
	userPolicyObjectSetResp, err := objectSet.Client.ListFromParams(userPolicyPath, params)
	if err != nil {
		return nil, err
	}
	return buildUserPolicyObjectSet(userPolicyObjectSetResp), err
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
