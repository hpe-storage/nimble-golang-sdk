// Copyright 2020 Hewlett Packard Enterprise Development LP

// create true

// update true
package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
	"strings"
)

// Represents Active Directory groups configured to manage the system.
const (
	userGroupPath = "user_groups"
)

// UserGroupObjectSet
type UserGroupObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new UserGroup object
func (objectSet *UserGroupObjectSet) CreateObject(payload *nimbleos.UserGroup) (*nimbleos.UserGroup, error) {
	resp, err := objectSet.Client.Post(userGroupPath, payload, &nimbleos.UserGroup{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)
			}
		}
		return nil, err
	}

	return resp.(*nimbleos.UserGroup), err
}

// UpdateObject Modify existing UserGroup object
func (objectSet *UserGroupObjectSet) UpdateObject(id string, payload *nimbleos.UserGroup) (*nimbleos.UserGroup, error) {
	resp, err := objectSet.Client.Put(userGroupPath, id, payload, &nimbleos.UserGroup{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)

			}
		}
		return nil, err
	}

	return resp.(*nimbleos.UserGroup), err
}

// DeleteObject deletes the UserGroup object with the specified ID
func (objectSet *UserGroupObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(userGroupPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a UserGroup object with the given ID
func (objectSet *UserGroupObjectSet) GetObject(id string) (*nimbleos.UserGroup, error) {
	resp, err := objectSet.Client.Get(userGroupPath, id, &nimbleos.UserGroup{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.UserGroup), err
}

// GetObjectList returns the list of UserGroup objects
func (objectSet *UserGroupObjectSet) GetObjectList() ([]*nimbleos.UserGroup, error) {
	resp, err := objectSet.Client.List(userGroupPath)
	if err != nil {
		return nil, err
	}
	return buildUserGroupObjectSet(resp), err
}

// GetObjectListFromParams returns the list of UserGroup objects using the given params query info
func (objectSet *UserGroupObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.UserGroup, error) {
	userGroupObjectSetResp, err := objectSet.Client.ListFromParams(userGroupPath, params)
	if err != nil {
		return nil, err
	}
	return buildUserGroupObjectSet(userGroupObjectSetResp), err
}

// generated function to build the appropriate response types
func buildUserGroupObjectSet(response interface{}) []*nimbleos.UserGroup {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.UserGroup, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.UserGroup{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
