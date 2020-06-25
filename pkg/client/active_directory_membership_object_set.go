// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Manages the storage array&#39;s membership with the Active Directory.
const (
    activeDirectoryMembershipPath = "active_directory_memberships"
)


// ActiveDirectoryMembershipObjectSet
type ActiveDirectoryMembershipObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ActiveDirectoryMembership object
func (objectSet *ActiveDirectoryMembershipObjectSet) CreateObject(payload *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	response, err := objectSet.Client.Post(activeDirectoryMembershipPath, payload)
	return response.(*model.ActiveDirectoryMembership), err
}

// UpdateObject Modify existing ActiveDirectoryMembership object
func (objectSet *ActiveDirectoryMembershipObjectSet) UpdateObject(id string, payload *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	response, err := objectSet.Client.Put(activeDirectoryMembershipPath, id, payload)
	return response.(*model.ActiveDirectoryMembership), err
}

// DeleteObject deletes the ActiveDirectoryMembership object with the specified ID
func (objectSet *ActiveDirectoryMembershipObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(activeDirectoryMembershipPath, id)
}

// GetObject returns a ActiveDirectoryMembership object with the given ID
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObject(id string) (*model.ActiveDirectoryMembership, error) {
	response, err := objectSet.Client.Get(activeDirectoryMembershipPath, id, model.ActiveDirectoryMembership{})
	if response == nil {
		return nil, err
	}
	return response.(*model.ActiveDirectoryMembership), err
}

// GetObjectList returns the list of ActiveDirectoryMembership objects
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObjectList() ([]*model.ActiveDirectoryMembership, error) {
	response, err := objectSet.Client.List(activeDirectoryMembershipPath)
	return buildActiveDirectoryMembershipObjectSet(response), err
}

// GetObjectListFromParams returns the list of ActiveDirectoryMembership objects using the given params query info
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ActiveDirectoryMembership, error) {
	response, err := objectSet.Client.ListFromParams(activeDirectoryMembershipPath, params)
	return buildActiveDirectoryMembershipObjectSet(response), err
}

// generated function to build the appropriate response types
func buildActiveDirectoryMembershipObjectSet(response interface{}) ([]*model.ActiveDirectoryMembership) {
	values := reflect.ValueOf(response)
	results := make([]*model.ActiveDirectoryMembership, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ActiveDirectoryMembership{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}