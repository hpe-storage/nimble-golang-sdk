// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Manage initiators in initiator groups. An initiator group has a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.
const (
    initiatorPath = "initiators"
)


// InitiatorObjectSet
type InitiatorObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Initiator object
func (objectSet *InitiatorObjectSet) CreateObject(payload *model.Initiator) (*model.Initiator, error) {
	response, err := objectSet.Client.Post(initiatorPath, payload)
	return response.(*model.Initiator), err
}

// UpdateObject Modify existing Initiator object
func (objectSet *InitiatorObjectSet) UpdateObject(id string, payload *model.Initiator) (*model.Initiator, error) {
	response, err := objectSet.Client.Put(initiatorPath, id, payload)
	return response.(*model.Initiator), err
}

// DeleteObject deletes the Initiator object with the specified ID
func (objectSet *InitiatorObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(initiatorPath, id)
}

// GetObject returns a Initiator object with the given ID
func (objectSet *InitiatorObjectSet) GetObject(id string) (*model.Initiator, error) {
	response, err := objectSet.Client.Get(initiatorPath, id, model.Initiator{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Initiator), err
}

// GetObjectList returns the list of Initiator objects
func (objectSet *InitiatorObjectSet) GetObjectList() ([]*model.Initiator, error) {
	response, err := objectSet.Client.List(initiatorPath)
	return buildInitiatorObjectSet(response), err
}

// GetObjectListFromParams returns the list of Initiator objects using the given params query info
func (objectSet *InitiatorObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Initiator, error) {
	response, err := objectSet.Client.ListFromParams(initiatorPath, params)
	return buildInitiatorObjectSet(response), err
}

// generated function to build the appropriate response types
func buildInitiatorObjectSet(response interface{}) ([]*model.Initiator) {
	values := reflect.ValueOf(response)
	results := make([]*model.Initiator, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Initiator{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}