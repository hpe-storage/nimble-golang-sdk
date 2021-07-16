// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *UserObjectSet) CreateObject(payload *nimbleos.User) (*nimbleos.User, error) {
	resp, err := objectSet.Client.Post(userPath, payload, &nimbleos.User{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.User), err
}

// UpdateObject Modify existing User object
func (objectSet *UserObjectSet) UpdateObject(id string, payload *nimbleos.User) (*nimbleos.User, error) {
	resp, err := objectSet.Client.Put(userPath, id, payload, &nimbleos.User{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.User), err
}

// DeleteObject deletes the User object with the specified ID
func (objectSet *UserObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(userPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a User object with the given ID
func (objectSet *UserObjectSet) GetObject(id string) (*nimbleos.User, error) {
	resp, err := objectSet.Client.Get(userPath, id, &nimbleos.User{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.User), err
}

// GetObjectList returns the list of User objects
func (objectSet *UserObjectSet) GetObjectList() ([]*nimbleos.User, error) {
	resp, err := objectSet.Client.List(userPath)
	if err != nil {
		return nil, err
	}
	return buildUserObjectSet(resp), err
}

// GetObjectListFromParams returns the list of User objects using the given params query info
func (objectSet *UserObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.User, error) {
	userObjectSetResp, err := objectSet.Client.ListFromParams(userPath, params)
	if err != nil {
		return nil, err
	}
	return buildUserObjectSet(userObjectSetResp), err
}

// generated function to build the appropriate response types
func buildUserObjectSet(response interface{}) []*nimbleos.User {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.User, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.User{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Unlock - Unlocks user account locked due to failed logins.
func (objectSet *UserObjectSet) Unlock(id *string) (*nimbleos.NsUserLockStatus, error) {
	unlockUri := userPath
	unlockUri = unlockUri + "/" + *id
	unlockUri = unlockUri + "/actions/" + "unlock"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	resp, err := objectSet.Client.Post(unlockUri, payload, &nimbleos.NsUserLockStatus{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsUserLockStatus), err
}
