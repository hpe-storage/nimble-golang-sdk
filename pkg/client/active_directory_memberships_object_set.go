// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manages the storage array's membership with the Active Directory.
const (
    activeDirectoryMembershipPath = "active_directory_memberships"
)

// ActiveDirectoryMembershipObjectSet
type ActiveDirectoryMembershipObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ActiveDirectoryMembership object
func (objectSet *ActiveDirectoryMembershipObjectSet) CreateObject(payload *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	activeDirectoryMembershipObjectSetResp, err := objectSet.Client.Post(activeDirectoryMembershipPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if activeDirectoryMembershipObjectSetResp == nil {
		return nil,nil
	}
	return activeDirectoryMembershipObjectSetResp.(*model.ActiveDirectoryMembership), err
}

// UpdateObject Modify existing ActiveDirectoryMembership object
func (objectSet *ActiveDirectoryMembershipObjectSet) UpdateObject(id string, payload *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	activeDirectoryMembershipObjectSetResp, err := objectSet.Client.Put(activeDirectoryMembershipPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if activeDirectoryMembershipObjectSetResp == nil {
		return nil,nil
	}
	return activeDirectoryMembershipObjectSetResp.(*model.ActiveDirectoryMembership), err
}

// DeleteObject deletes the ActiveDirectoryMembership object with the specified ID
func (objectSet *ActiveDirectoryMembershipObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on ActiveDirectoryMembership")
}

// GetObject returns a ActiveDirectoryMembership object with the given ID
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObject(id string) (*model.ActiveDirectoryMembership, error) {
	activeDirectoryMembershipObjectSetResp, err := objectSet.Client.Get(activeDirectoryMembershipPath, id, model.ActiveDirectoryMembership{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if activeDirectoryMembershipObjectSetResp == nil {
		return nil,nil
	}
	return activeDirectoryMembershipObjectSetResp.(*model.ActiveDirectoryMembership), err
}

// GetObjectList returns the list of ActiveDirectoryMembership objects
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObjectList() ([]*model.ActiveDirectoryMembership, error) {
	activeDirectoryMembershipObjectSetResp, err := objectSet.Client.List(activeDirectoryMembershipPath)
	if err != nil {
		return nil, err
	}
	return buildActiveDirectoryMembershipObjectSet(activeDirectoryMembershipObjectSetResp), err
}

// GetObjectListFromParams returns the list of ActiveDirectoryMembership objects using the given params query info
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ActiveDirectoryMembership, error) {
	activeDirectoryMembershipObjectSetResp, err := objectSet.Client.ListFromParams(activeDirectoryMembershipPath, params)
	if err != nil {
		return nil, err
	}
	return buildActiveDirectoryMembershipObjectSet(activeDirectoryMembershipObjectSetResp), err
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
