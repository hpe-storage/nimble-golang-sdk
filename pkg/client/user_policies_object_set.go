// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *UserPolicyObjectSet) CreateObject(payload *nimbleos.UserPolicy) (*nimbleos.UserPolicy, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on UserPolicy")
}

// UpdateObject Modify existing UserPolicy object
func (objectSet *UserPolicyObjectSet) UpdateObject(id string, payload *nimbleos.UserPolicy) (*nimbleos.UserPolicy, error) {
	resp, err := objectSet.Client.Put(userPolicyPath, id, payload, &nimbleos.UserPolicy{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.UserPolicy), err
}

// DeleteObject deletes the UserPolicy object with the specified ID
func (objectSet *UserPolicyObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on UserPolicy")
}

// GetObject returns a UserPolicy object with the given ID
func (objectSet *UserPolicyObjectSet) GetObject(id string) (*nimbleos.UserPolicy, error) {
	resp, err := objectSet.Client.Get(userPolicyPath, id, &nimbleos.UserPolicy{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.UserPolicy), err
}

// GetObjectList returns the list of UserPolicy objects
func (objectSet *UserPolicyObjectSet) GetObjectList() ([]*nimbleos.UserPolicy, error) {
	resp, err := objectSet.Client.List(userPolicyPath)
	if err != nil {
		return nil, err
	}
	return buildUserPolicyObjectSet(resp), err
}

// GetObjectListFromParams returns the list of UserPolicy objects using the given params query info
func (objectSet *UserPolicyObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.UserPolicy, error) {
	userPolicyObjectSetResp, err := objectSet.Client.ListFromParams(userPolicyPath, params)
	if err != nil {
		return nil, err
	}
	return buildUserPolicyObjectSet(userPolicyObjectSetResp), err
}

// generated function to build the appropriate response types
func buildUserPolicyObjectSet(response interface{}) []*nimbleos.UserPolicy {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.UserPolicy, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.UserPolicy{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
