// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *InitiatorGroupObjectSet) CreateObject(payload *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	newPayload, err := model.EncodeInitiatorGroup(payload)
	resp, err := objectSet.Client.Post(initiatorGroupPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return model.DecodeInitiatorGroup(resp)
}

// UpdateObject Modify existing InitiatorGroup object
func (objectSet *InitiatorGroupObjectSet) UpdateObject(id string, payload *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	newPayload, err := model.EncodeInitiatorGroup(payload)
	resp, err := objectSet.Client.Put(initiatorGroupPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return model.DecodeInitiatorGroup(resp)
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
func (objectSet *InitiatorGroupObjectSet) GetObject(id string) (*model.InitiatorGroup, error) {
	initiatorGroupObjectSetResp, err := objectSet.Client.Get(initiatorGroupPath, id, model.InitiatorGroup{})
	if err != nil {
		return nil, err
	}

	// null check
	if initiatorGroupObjectSetResp == nil {
		return nil, nil
	}
	return initiatorGroupObjectSetResp.(*model.InitiatorGroup), err
}

// GetObjectList returns the list of InitiatorGroup objects
func (objectSet *InitiatorGroupObjectSet) GetObjectList() ([]*model.InitiatorGroup, error) {
	initiatorGroupObjectSetResp, err := objectSet.Client.List(initiatorGroupPath)
	if err != nil {
		return nil, err
	}
	return buildInitiatorGroupObjectSet(initiatorGroupObjectSetResp), err
}

// GetObjectListFromParams returns the list of InitiatorGroup objects using the given params query info
func (objectSet *InitiatorGroupObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.InitiatorGroup, error) {
	initiatorGroupObjectSetResp, err := objectSet.Client.ListFromParams(initiatorGroupPath, params)
	if err != nil {
		return nil, err
	}
	return buildInitiatorGroupObjectSet(initiatorGroupObjectSetResp), err
}

// generated function to build the appropriate response types
func buildInitiatorGroupObjectSet(response interface{}) []*model.InitiatorGroup {
	values := reflect.ValueOf(response)
	results := make([]*model.InitiatorGroup, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.InitiatorGroup{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
