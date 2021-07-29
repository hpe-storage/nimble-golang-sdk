// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *DiskObjectSet) CreateObject(payload *nimbleos.Disk) (*nimbleos.Disk, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Disk")
}

// UpdateObject Modify existing Disk object
func (objectSet *DiskObjectSet) UpdateObject(id string, payload *nimbleos.Disk) (*nimbleos.Disk, error) {
	resp, err := objectSet.Client.Put(diskPath, id, payload, &nimbleos.Disk{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Disk), err
}

// DeleteObject deletes the Disk object with the specified ID
func (objectSet *DiskObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Disk")
}

// GetObject returns a Disk object with the given ID
func (objectSet *DiskObjectSet) GetObject(id string) (*nimbleos.Disk, error) {
	resp, err := objectSet.Client.Get(diskPath, id, &nimbleos.Disk{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Disk), err
}

// GetObjectList returns the list of Disk objects
func (objectSet *DiskObjectSet) GetObjectList() ([]*nimbleos.Disk, error) {
	resp, err := objectSet.Client.List(diskPath)
	if err != nil {
		return nil, err
	}
	return buildDiskObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Disk objects using the given params query info
func (objectSet *DiskObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Disk, error) {
	diskObjectSetResp, err := objectSet.Client.ListFromParams(diskPath, params)
	if err != nil {
		return nil, err
	}
	return buildDiskObjectSet(diskObjectSetResp), err
}

// generated function to build the appropriate response types
func buildDiskObjectSet(response interface{}) []*nimbleos.Disk {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Disk, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Disk{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
