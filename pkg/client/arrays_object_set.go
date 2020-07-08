// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *ArrayObjectSet) CreateObject(payload *model.Array) (*model.Array, error) {
	newPayload, err := model.EncodeArray(payload)
	arrayObjectSetResp, err := objectSet.Client.Post(arrayPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if arrayObjectSetResp == nil {
		return nil, nil
	}

	return model.DecodeArray(arrayObjectSetResp)
}

// UpdateObject Modify existing Array object
func (objectSet *ArrayObjectSet) UpdateObject(id string, payload *model.Array) (*model.Array, error) {
	newPayload, err := model.EncodeArray(payload)
	arrayObjectSetResp, err := objectSet.Client.Put(arrayPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if arrayObjectSetResp == nil {
		return nil, nil
	}
	return model.DecodeArray(arrayObjectSetResp)
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
func (objectSet *ArrayObjectSet) GetObject(id string) (*model.Array, error) {
	arrayObjectSetResp, err := objectSet.Client.Get(arrayPath, id, model.Array{})
	if err != nil {
		return nil, err
	}

	// null check
	if arrayObjectSetResp == nil {
		return nil, nil
	}
	return arrayObjectSetResp.(*model.Array), err
}

// GetObjectList returns the list of Array objects
func (objectSet *ArrayObjectSet) GetObjectList() ([]*model.Array, error) {
	arrayObjectSetResp, err := objectSet.Client.List(arrayPath)
	if err != nil {
		return nil, err
	}
	return buildArrayObjectSet(arrayObjectSetResp), err
}

// GetObjectListFromParams returns the list of Array objects using the given params query info
func (objectSet *ArrayObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Array, error) {
	arrayObjectSetResp, err := objectSet.Client.ListFromParams(arrayPath, params)
	if err != nil {
		return nil, err
	}
	return buildArrayObjectSet(arrayObjectSetResp), err
}

// generated function to build the appropriate response types
func buildArrayObjectSet(response interface{}) []*model.Array {
	values := reflect.ValueOf(response)
	results := make([]*model.Array, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Array{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
