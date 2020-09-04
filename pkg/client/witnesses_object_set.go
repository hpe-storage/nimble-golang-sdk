// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *WitnessObjectSet) CreateObject(payload *nimbleos.Witness) (*nimbleos.Witness, error) {
	resp, err := objectSet.Client.Post(witnessPath, payload, &nimbleos.Witness{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Witness), err
}

// UpdateObject Modify existing Witness object
func (objectSet *WitnessObjectSet) UpdateObject(id string, payload *nimbleos.Witness) (*nimbleos.Witness, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Witness")
}

// DeleteObject deletes the Witness object with the specified ID
func (objectSet *WitnessObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(witnessPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Witness object with the given ID
func (objectSet *WitnessObjectSet) GetObject(id string) (*nimbleos.Witness, error) {
	resp, err := objectSet.Client.Get(witnessPath, id, nimbleos.Witness{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Witness), err
}

// GetObjectList returns the list of Witness objects
func (objectSet *WitnessObjectSet) GetObjectList() ([]*nimbleos.Witness, error) {
	resp, err := objectSet.Client.List(witnessPath)
	if err != nil {
		return nil, err
	}
	return buildWitnessObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Witness objects using the given params query info
func (objectSet *WitnessObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Witness, error) {
	witnessObjectSetResp, err := objectSet.Client.ListFromParams(witnessPath, params)
	if err != nil {
		return nil, err
	}
	return buildWitnessObjectSet(witnessObjectSetResp), err
}

// generated function to build the appropriate response types
func buildWitnessObjectSet(response interface{}) []*nimbleos.Witness {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Witness, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Witness{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Test - Tests and validates witness configuration between the array and the witness.
func (objectSet *WitnessObjectSet) Test(id *string) (*[]*nimbleos.NsWitnessTestResponse, error) {

	testUri := witnessPath
	testUri = testUri + "/" + *id
	testUri = testUri + "/actions/" + "test"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	resp, err := objectSet.Client.Post(testUri, payload, &emptyStruct)
	if err != nil {
		return nil, err
	}

	return resp.(*[]*nimbleos.NsWitnessTestResponse), err
}
