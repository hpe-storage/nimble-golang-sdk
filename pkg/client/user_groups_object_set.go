// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *UserGroupObjectSet) CreateObject(payload *model.UserGroup) (*model.UserGroup, error) {
	userGroupObjectSetResp, err := objectSet.Client.Post(userGroupPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if userGroupObjectSetResp == nil {
		return nil,nil
	}
	return userGroupObjectSetResp.(*model.UserGroup), err
}

// UpdateObject Modify existing UserGroup object
func (objectSet *UserGroupObjectSet) UpdateObject(id string, payload *model.UserGroup) (*model.UserGroup, error) {
	userGroupObjectSetResp, err := objectSet.Client.Put(userGroupPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if userGroupObjectSetResp == nil {
		return nil,nil
	}
	return userGroupObjectSetResp.(*model.UserGroup), err
}

// DeleteObject deletes the UserGroup object with the specified ID
func (objectSet *UserGroupObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(userGroupPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a UserGroup object with the given ID
func (objectSet *UserGroupObjectSet) GetObject(id string) (*model.UserGroup, error) {
	userGroupObjectSetResp, err := objectSet.Client.Get(userGroupPath, id, model.UserGroup{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if userGroupObjectSetResp == nil {
		return nil,nil
	}
	return userGroupObjectSetResp.(*model.UserGroup), err
}

// GetObjectList returns the list of UserGroup objects
func (objectSet *UserGroupObjectSet) GetObjectList() ([]*model.UserGroup, error) {
	userGroupObjectSetResp, err := objectSet.Client.List(userGroupPath)
	if err != nil {
		return nil, err
	}
	return buildUserGroupObjectSet(userGroupObjectSetResp), err
}

// GetObjectListFromParams returns the list of UserGroup objects using the given params query info
func (objectSet *UserGroupObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.UserGroup, error) {
	userGroupObjectSetResp, err := objectSet.Client.ListFromParams(userGroupPath, params)
	if err != nil {
		return nil, err
	}
	return buildUserGroupObjectSet(userGroupObjectSetResp), err
}
// generated function to build the appropriate response types
func buildUserGroupObjectSet(response interface{}) ([]*model.UserGroup) {
	values := reflect.ValueOf(response)
	results := make([]*model.UserGroup, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.UserGroup{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
