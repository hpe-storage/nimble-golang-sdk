// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Manage Challenge-Response Handshake Authentication Protocol (CHAP) user accounts. CHAP users are one method of access control for iSCSI initiators. Each CHAP user has a CHAP
// password, sometimes called a CHAP secret. The CHAP passwords must match on the array and on the iSCSI initiator in order for the array to authenicate the initiator and allow it
// access. The CHAP user information must exist on both the array and the iSCSI initiator. Target authentication gives security only for the specific iSCSI target.
const (
	chapUserPath = "chap_users"
)

// ChapUserObjectSet
type ChapUserObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new ChapUser object
func (objectSet *ChapUserObjectSet) CreateObject(payload *nimbleos.ChapUser) (*nimbleos.ChapUser, error) {
	resp, err := objectSet.Client.Post(chapUserPath, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return resp.(*nimbleos.ChapUser), err
}

// UpdateObject Modify existing ChapUser object
func (objectSet *ChapUserObjectSet) UpdateObject(id string, payload *nimbleos.ChapUser) (*nimbleos.ChapUser, error) {
	resp, err := objectSet.Client.Put(chapUserPath, id, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.ChapUser), err
}

// DeleteObject deletes the ChapUser object with the specified ID
func (objectSet *ChapUserObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(chapUserPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a ChapUser object with the given ID
func (objectSet *ChapUserObjectSet) GetObject(id string) (*nimbleos.ChapUser, error) {
	resp, err := objectSet.Client.Get(chapUserPath, id, nimbleos.ChapUser{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.ChapUser), err
}

// GetObjectList returns the list of ChapUser objects
func (objectSet *ChapUserObjectSet) GetObjectList() ([]*nimbleos.ChapUser, error) {
	resp, err := objectSet.Client.List(chapUserPath)
	if err != nil {
		return nil, err
	}
	return buildChapUserObjectSet(resp), err
}

// GetObjectListFromParams returns the list of ChapUser objects using the given params query info
func (objectSet *ChapUserObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ChapUser, error) {
	chapUserObjectSetResp, err := objectSet.Client.ListFromParams(chapUserPath, params)
	if err != nil {
		return nil, err
	}
	return buildChapUserObjectSet(chapUserObjectSetResp), err
}

// generated function to build the appropriate response types
func buildChapUserObjectSet(response interface{}) []*nimbleos.ChapUser {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.ChapUser, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.ChapUser{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
