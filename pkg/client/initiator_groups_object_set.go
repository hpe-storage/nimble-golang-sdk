// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through
// group membership.
const (
	initiatorGroupPath = "initiator_groups"
)

// InitiatorGroupObjectSet
type InitiatorGroupObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new InitiatorGroup object
func (objectSet *InitiatorGroupObjectSet) CreateObject(payload *nimbleos.InitiatorGroup) (*nimbleos.InitiatorGroup, error) {
	newPayload, err := nimbleos.EncodeInitiatorGroup(payload)
	resp, err := objectSet.Client.Post(initiatorGroupPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return nimbleos.DecodeInitiatorGroup(resp)
}

// UpdateObject Modify existing InitiatorGroup object
func (objectSet *InitiatorGroupObjectSet) UpdateObject(id string, payload *nimbleos.InitiatorGroup) (*nimbleos.InitiatorGroup, error) {
	newPayload, err := nimbleos.EncodeInitiatorGroup(payload)
	resp, err := objectSet.Client.Put(initiatorGroupPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return nimbleos.DecodeInitiatorGroup(resp)
}

// DeleteObject deletes the InitiatorGroup object with the specified ID
func (objectSet *InitiatorGroupObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(initiatorGroupPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a InitiatorGroup object with the given ID
func (objectSet *InitiatorGroupObjectSet) GetObject(id string) (*nimbleos.InitiatorGroup, error) {
	initiatorGroupObjectSetResp, err := objectSet.Client.Get(initiatorGroupPath, id, nimbleos.InitiatorGroup{})
	if err != nil {
		return nil, err
	}

	// null check
	if initiatorGroupObjectSetResp == nil {
		return nil, nil
	}
	return initiatorGroupObjectSetResp.(*nimbleos.InitiatorGroup), err
}

// GetObjectList returns the list of InitiatorGroup objects
func (objectSet *InitiatorGroupObjectSet) GetObjectList() ([]*nimbleos.InitiatorGroup, error) {
	initiatorGroupObjectSetResp, err := objectSet.Client.List(initiatorGroupPath)
	if err != nil {
		return nil, err
	}
	return buildInitiatorGroupObjectSet(initiatorGroupObjectSetResp), err
}

// GetObjectListFromParams returns the list of InitiatorGroup objects using the given params query info
func (objectSet *InitiatorGroupObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.InitiatorGroup, error) {
	initiatorGroupObjectSetResp, err := objectSet.Client.ListFromParams(initiatorGroupPath, params)
	if err != nil {
		return nil, err
	}
	return buildInitiatorGroupObjectSet(initiatorGroupObjectSetResp), err
}

// generated function to build the appropriate response types
func buildInitiatorGroupObjectSet(response interface{}) []*nimbleos.InitiatorGroup {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.InitiatorGroup, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.InitiatorGroup{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
