// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Retrieve information of specified arrays. The array is the management and configuration for the underlying physical hardware array box.
const (
	arrayPath = "arrays"
)

// ArrayObjectSet
type ArrayObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Array object
func (objectSet *ArrayObjectSet) CreateObject(payload *nimbleos.Array) (*nimbleos.Array, error) {
	encodedPayload, err := nimbleos.EncodeArray(payload)
	if err != nil {
		return nil, err
	}
	resp, err := objectSet.Client.Post(arrayPath, encodedPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return nimbleos.DecodeArray(resp)
}

// UpdateObject Modify existing Array object
func (objectSet *ArrayObjectSet) UpdateObject(id string, payload *nimbleos.Array) (*nimbleos.Array, error) {
	newPayload, err := nimbleos.EncodeArray(payload)
	resp, err := objectSet.Client.Put(arrayPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return nimbleos.DecodeArray(resp)
}

// DeleteObject deletes the Array object with the specified ID
func (objectSet *ArrayObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(arrayPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Array object with the given ID
func (objectSet *ArrayObjectSet) GetObject(id string) (*nimbleos.Array, error) {
	resp, err := objectSet.Client.Get(arrayPath, id, nimbleos.Array{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Array), err
}

// GetObjectList returns the list of Array objects
func (objectSet *ArrayObjectSet) GetObjectList() ([]*nimbleos.Array, error) {
	resp, err := objectSet.Client.List(arrayPath)
	if err != nil {
		return nil, err
	}
	return buildArrayObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Array objects using the given params query info
func (objectSet *ArrayObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Array, error) {
	arrayObjectSetResp, err := objectSet.Client.ListFromParams(arrayPath, params)
	if err != nil {
		return nil, err
	}
	return buildArrayObjectSet(arrayObjectSetResp), err
}

// generated function to build the appropriate response types
func buildArrayObjectSet(response interface{}) []*nimbleos.Array {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Array, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Array{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
