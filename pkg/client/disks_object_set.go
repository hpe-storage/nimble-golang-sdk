// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Disks are used for storing user data.
const (
    diskPath = "disks"
)

// DiskObjectSet
type DiskObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Disk object
func (objectSet *DiskObjectSet) CreateObject(payload *model.Disk) (*model.Disk, error) {
	response, err := objectSet.Client.Post(diskPath, payload)
	return response.(*model.Disk), err
}

// UpdateObject Modify existing Disk object
func (objectSet *DiskObjectSet) UpdateObject(id string, payload *model.Disk) (*model.Disk, error) {
	response, err := objectSet.Client.Put(diskPath, id, payload)
	return response.(*model.Disk), err
}

// DeleteObject deletes the Disk object with the specified ID
func (objectSet *DiskObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(diskPath, id)
}

// GetObject returns a Disk object with the given ID
func (objectSet *DiskObjectSet) GetObject(id string) (*model.Disk, error) {
	response, err := objectSet.Client.Get(diskPath, id, model.Disk{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Disk), err
}

// GetObjectList returns the list of Disk objects
func (objectSet *DiskObjectSet) GetObjectList() ([]*model.Disk, error) {
	response, err := objectSet.Client.List(diskPath)
	return buildDiskObjectSet(response), err
}

// GetObjectListFromParams returns the list of Disk objects using the given params query info
func (objectSet *DiskObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Disk, error) {
	response, err := objectSet.Client.ListFromParams(diskPath, params)
	return buildDiskObjectSet(response), err
}

// generated function to build the appropriate response types
func buildDiskObjectSet(response interface{}) ([]*model.Disk) {
	values := reflect.ValueOf(response)
	results := make([]*model.Disk, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Disk{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}