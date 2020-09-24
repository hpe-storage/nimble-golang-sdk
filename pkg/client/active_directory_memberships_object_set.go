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
	resp, err := objectSet.Client.Get(activeDirectoryMembershipPath, id, &nimbleos.ActiveDirectoryMembership{})
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

// List of supported actions on object sets

// Remove - Leaves the Active Directory domain.
func (objectSet *ActiveDirectoryMembershipObjectSet) Remove(id *string, user *string, password *string, force *bool) error {
	removeUri := activeDirectoryMembershipPath
	removeUri = removeUri + "/" + *id
	removeUri = removeUri + "/actions/" + "remove"

	payload := &struct {
		Id       *string `json:"id,omitempty"`
		User     *string `json:"user,omitempty"`
		Password *string `json:"password,omitempty"`
		Force    *bool   `json:"force,omitempty"`
	}{
		id,
		user,
		password,
		force,
	}

	_, err := objectSet.Client.Post(removeUri, payload, &nimbleos.ActiveDirectoryMembership{})
	return err
}

// ReportStatus - Reports the detail status of the Active Directory domain.
func (objectSet *ActiveDirectoryMembershipObjectSet) ReportStatus(id *string) (*nimbleos.NsADReportStatusReturn, error) {
	reportStatusUri := activeDirectoryMembershipPath
	reportStatusUri = reportStatusUri + "/" + *id
	reportStatusUri = reportStatusUri + "/actions/" + "report_status"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	resp, err := objectSet.Client.Post(reportStatusUri, payload, &nimbleos.NsADReportStatusReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsADReportStatusReturn), err
}

// TestUser - Tests whether the user exist in the Active Directory. If the user is present, then the user's group and role information is reported.
func (objectSet *ActiveDirectoryMembershipObjectSet) TestUser(id *string, name *string) (*nimbleos.NsADTestUserReturn, error) {
	testUserUri := activeDirectoryMembershipPath
	testUserUri = testUserUri + "/" + *id
	testUserUri = testUserUri + "/actions/" + "test_user"

	payload := &struct {
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}{
		id,
		name,
	}

	resp, err := objectSet.Client.Post(testUserUri, payload, &nimbleos.NsADTestUserReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsADTestUserReturn), err
}

// TestGroup - Tests whether the user group exist in the Active Directory.
func (objectSet *ActiveDirectoryMembershipObjectSet) TestGroup(id *string, name *string) error {
	testGroupUri := activeDirectoryMembershipPath
	testGroupUri = testGroupUri + "/" + *id
	testGroupUri = testGroupUri + "/actions/" + "test_group"

	payload := &struct {
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}{
		id,
		name,
	}

	_, err := objectSet.Client.Post(testGroupUri, payload, &nimbleos.ActiveDirectoryMembership{})
	return err
}
