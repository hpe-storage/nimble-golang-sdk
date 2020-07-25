// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *ActiveDirectoryMembershipObjectSet) CreateObject(payload *nimbleos.ActiveDirectoryMembership) (*nimbleos.ActiveDirectoryMembership, error) {
	resp, err := objectSet.Client.Post(activeDirectoryMembershipPath, payload, &nimbleos.ActiveDirectoryMembership{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.ActiveDirectoryMembership), err
}

// UpdateObject Modify existing ActiveDirectoryMembership object
func (objectSet *ActiveDirectoryMembershipObjectSet) UpdateObject(id string, payload *nimbleos.ActiveDirectoryMembership) (*nimbleos.ActiveDirectoryMembership, error) {
	resp, err := objectSet.Client.Put(activeDirectoryMembershipPath, id, payload, &nimbleos.ActiveDirectoryMembership{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.ActiveDirectoryMembership), err
}

// DeleteObject deletes the ActiveDirectoryMembership object with the specified ID
func (objectSet *ActiveDirectoryMembershipObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on ActiveDirectoryMembership")
}

// GetObject returns a ActiveDirectoryMembership object with the given ID
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObject(id string) (*nimbleos.ActiveDirectoryMembership, error) {
	resp, err := objectSet.Client.Get(activeDirectoryMembershipPath, id, nimbleos.ActiveDirectoryMembership{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.ActiveDirectoryMembership), err
}

// GetObjectList returns the list of ActiveDirectoryMembership objects
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObjectList() ([]*nimbleos.ActiveDirectoryMembership, error) {
	resp, err := objectSet.Client.List(activeDirectoryMembershipPath)
	if err != nil {
		return nil, err
	}
	return buildActiveDirectoryMembershipObjectSet(resp), err
}

// GetObjectListFromParams returns the list of ActiveDirectoryMembership objects using the given params query info
func (objectSet *ActiveDirectoryMembershipObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ActiveDirectoryMembership, error) {
	activeDirectoryMembershipObjectSetResp, err := objectSet.Client.ListFromParams(activeDirectoryMembershipPath, params)
	if err != nil {
		return nil, err
	}
	return buildActiveDirectoryMembershipObjectSet(activeDirectoryMembershipObjectSetResp), err
}

// generated function to build the appropriate response types
func buildActiveDirectoryMembershipObjectSet(response interface{}) []*nimbleos.ActiveDirectoryMembership {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.ActiveDirectoryMembership, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.ActiveDirectoryMembership{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
