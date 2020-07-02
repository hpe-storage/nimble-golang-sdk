// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Show the API version.
const (
    versionPath = "versions"
)

// VersionObjectSet
type VersionObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Version object
func (objectSet *VersionObjectSet) CreateObject(payload *model.Version) (*model.Version, error) {
	response, err := objectSet.Client.Post(versionPath, payload)
	return response.(*model.Version), err
}

// UpdateObject Modify existing Version object
func (objectSet *VersionObjectSet) UpdateObject(id string, payload *model.Version) (*model.Version, error) {
	response, err := objectSet.Client.Put(versionPath, id, payload)
	return response.(*model.Version), err
}

// DeleteObject deletes the Version object with the specified ID
func (objectSet *VersionObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(versionPath, id)
}

// GetObject returns a Version object with the given ID
func (objectSet *VersionObjectSet) GetObject(id string) (*model.Version, error) {
	response, err := objectSet.Client.Get(versionPath, id, model.Version{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Version), err
}

// GetObjectList returns the list of Version objects
func (objectSet *VersionObjectSet) GetObjectList() ([]*model.Version, error) {
	response, err := objectSet.Client.List(versionPath)
	return buildVersionObjectSet(response), err
}

// GetObjectListFromParams returns the list of Version objects using the given params query info
func (objectSet *VersionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Version, error) {
	response, err := objectSet.Client.ListFromParams(versionPath, params)
	return buildVersionObjectSet(response), err
}

// generated function to build the appropriate response types
func buildVersionObjectSet(response interface{}) ([]*model.Version) {
	values := reflect.ValueOf(response)
	results := make([]*model.Version, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Version{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}