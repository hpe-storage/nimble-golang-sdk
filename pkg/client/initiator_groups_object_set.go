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
	resp, err := objectSet.Client.Post(initiatorGroupPath, payload, &nimbleos.InitiatorGroup{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.InitiatorGroup), err
}

// UpdateObject Modify existing InitiatorGroup object
func (objectSet *InitiatorGroupObjectSet) UpdateObject(id string, payload *nimbleos.InitiatorGroup) (*nimbleos.InitiatorGroup, error) {
	resp, err := objectSet.Client.Put(initiatorGroupPath, id, payload, &nimbleos.InitiatorGroup{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.InitiatorGroup), err
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
	resp, err := objectSet.Client.Get(initiatorGroupPath, id, &nimbleos.InitiatorGroup{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.InitiatorGroup), err
}

// GetObjectList returns the list of InitiatorGroup objects
func (objectSet *InitiatorGroupObjectSet) GetObjectList() ([]*nimbleos.InitiatorGroup, error) {
	resp, err := objectSet.Client.List(initiatorGroupPath)
	if err != nil {
		return nil, err
	}
	return buildInitiatorGroupObjectSet(resp), err
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

// List of supported actions on object sets

// SuggestLun - Suggest an LU number for the volume and initiator group combination.
func (objectSet *InitiatorGroupObjectSet) SuggestLun(id *string, volId *string) (*nimbleos.NsLunReturn, error) {

	suggestLunUri := initiatorGroupPath
	suggestLunUri = suggestLunUri + "/" + *id
	suggestLunUri = suggestLunUri + "/actions/" + "suggest_lun"

	payload := &struct {
		Id    *string `json:"id,omitempty"`
		VolId *string `json:"vol_id,omitempty"`
	}{
		id,
		volId,
	}

	resp, err := objectSet.Client.Post(suggestLunUri, payload, &nimbleos.NsLunReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsLunReturn), err
}

// ValidateLun - Validate an LU number for the volume and initiator group combination.
func (objectSet *InitiatorGroupObjectSet) ValidateLun(id *string, volId *string, lun *uint64) error {

	validateLunUri := initiatorGroupPath
	validateLunUri = validateLunUri + "/" + *id
	validateLunUri = validateLunUri + "/actions/" + "validate_lun"

	payload := &struct {
		Id    *string `json:"id,omitempty"`
		VolId *string `json:"vol_id,omitempty"`
		Lun   *uint64 `json:"lun,omitempty"`
	}{
		id,
		volId,
		lun,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(validateLunUri, payload, &emptyStruct)
	return err
}
