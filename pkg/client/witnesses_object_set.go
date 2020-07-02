// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manage witness host configuration.
const (
    witnessPath = "witnesses"
)

// WitnessObjectSet
type WitnessObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Witness object
func (objectSet *WitnessObjectSet) CreateObject(payload *model.Witness) (*model.Witness, error) {
	response, err := objectSet.Client.Post(witnessPath, payload)
	return response.(*model.Witness), err
}

// UpdateObject Modify existing Witness object
func (objectSet *WitnessObjectSet) UpdateObject(id string, payload *model.Witness) (*model.Witness, error) {
	response, err := objectSet.Client.Put(witnessPath, id, payload)
	return response.(*model.Witness), err
}

// DeleteObject deletes the Witness object with the specified ID
func (objectSet *WitnessObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(witnessPath, id)
}

// GetObject returns a Witness object with the given ID
func (objectSet *WitnessObjectSet) GetObject(id string) (*model.Witness, error) {
	response, err := objectSet.Client.Get(witnessPath, id, model.Witness{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Witness), err
}

// GetObjectList returns the list of Witness objects
func (objectSet *WitnessObjectSet) GetObjectList() ([]*model.Witness, error) {
	response, err := objectSet.Client.List(witnessPath)
	return buildWitnessObjectSet(response), err
}

// GetObjectListFromParams returns the list of Witness objects using the given params query info
func (objectSet *WitnessObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Witness, error) {
	response, err := objectSet.Client.ListFromParams(witnessPath, params)
	return buildWitnessObjectSet(response), err
}

// generated function to build the appropriate response types
func buildWitnessObjectSet(response interface{}) ([]*model.Witness) {
	values := reflect.ValueOf(response)
	results := make([]*model.Witness, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Witness{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}