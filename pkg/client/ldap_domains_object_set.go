// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Manages the storage array's membership with LDAP servers.
const (
	ldapDomainPath = "ldap_domains"
)

// LdapDomainObjectSet
type LdapDomainObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new LdapDomain object
func (objectSet *LdapDomainObjectSet) CreateObject(payload *nimbleos.LdapDomain) (*nimbleos.LdapDomain, error) {
	resp, err := objectSet.Client.Post(ldapDomainPath, payload, &nimbleos.LdapDomain{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.LdapDomain), err
}

// UpdateObject Modify existing LdapDomain object
func (objectSet *LdapDomainObjectSet) UpdateObject(id string, payload *nimbleos.LdapDomain) (*nimbleos.LdapDomain, error) {
	resp, err := objectSet.Client.Put(ldapDomainPath, id, payload, &nimbleos.LdapDomain{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.LdapDomain), err
}

// DeleteObject deletes the LdapDomain object with the specified ID
func (objectSet *LdapDomainObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(ldapDomainPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a LdapDomain object with the given ID
func (objectSet *LdapDomainObjectSet) GetObject(id string) (*nimbleos.LdapDomain, error) {
	resp, err := objectSet.Client.Get(ldapDomainPath, id, &nimbleos.LdapDomain{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.LdapDomain), err
}

// GetObjectList returns the list of LdapDomain objects
func (objectSet *LdapDomainObjectSet) GetObjectList() ([]*nimbleos.LdapDomain, error) {
	resp, err := objectSet.Client.List(ldapDomainPath)
	if err != nil {
		return nil, err
	}
	return buildLdapDomainObjectSet(resp), err
}

// GetObjectListFromParams returns the list of LdapDomain objects using the given params query info
func (objectSet *LdapDomainObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.LdapDomain, error) {
	ldapDomainObjectSetResp, err := objectSet.Client.ListFromParams(ldapDomainPath, params)
	if err != nil {
		return nil, err
	}
	return buildLdapDomainObjectSet(ldapDomainObjectSetResp), err
}

// generated function to build the appropriate response types
func buildLdapDomainObjectSet(response interface{}) []*nimbleos.LdapDomain {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.LdapDomain, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.LdapDomain{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// TestUser - Tests whether the user exist in the given LDAP Domain.
func (objectSet *LdapDomainObjectSet) TestUser(id *string, user *string) (*nimbleos.NsLdapUser, error) {
	testUserUri := ldapDomainPath
	testUserUri = testUserUri + "/" + *id
	testUserUri = testUserUri + "/actions/" + "test_user"

	payload := &struct {
		Id   *string `json:"id,omitempty"`
		User *string `json:"user,omitempty"`
	}{
		id,
		user,
	}

	resp, err := objectSet.Client.Post(testUserUri, payload, &nimbleos.NsLdapUser{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsLdapUser), err
}

// TestGroup - Tests whether the user group exist in the given LDAP Domain.
func (objectSet *LdapDomainObjectSet) TestGroup(id *string, group *string) error {
	testGroupUri := ldapDomainPath
	testGroupUri = testGroupUri + "/" + *id
	testGroupUri = testGroupUri + "/actions/" + "test_group"

	payload := &struct {
		Id    *string `json:"id,omitempty"`
		Group *string `json:"group,omitempty"`
	}{
		id,
		group,
	}

	_, err := objectSet.Client.Post(testGroupUri, payload, &nimbleos.LdapDomain{})
	return err
}

// ReportStatus - Reports the LDAP connectivity status of the given LDAP ID.
func (objectSet *LdapDomainObjectSet) ReportStatus(id *string) (*nimbleos.NsLdapReportStatusReturn, error) {
	reportStatusUri := ldapDomainPath
	reportStatusUri = reportStatusUri + "/" + *id
	reportStatusUri = reportStatusUri + "/actions/" + "report_status"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	resp, err := objectSet.Client.Post(reportStatusUri, payload, &nimbleos.NsLdapReportStatusReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsLdapReportStatusReturn), err
}
