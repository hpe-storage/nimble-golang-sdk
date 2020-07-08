// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *DiskObjectSet) CreateObject(payload *model.Disk) (*model.Disk, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Disk")
}

// UpdateObject Modify existing Disk object
func (objectSet *DiskObjectSet) UpdateObject(id string, payload *model.Disk) (*model.Disk, error) {
	newPayload, err := model.EncodeDisk(payload)
	diskObjectSetResp, err := objectSet.Client.Put(diskPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if diskObjectSetResp == nil {
		return nil, nil
	}
	return model.DecodeDisk(diskObjectSetResp)
}

// DeleteObject deletes the Disk object with the specified ID
func (objectSet *DiskObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Disk")
}

// GetObject returns a Disk object with the given ID
func (objectSet *DiskObjectSet) GetObject(id string) (*model.Disk, error) {
	diskObjectSetResp, err := objectSet.Client.Get(diskPath, id, model.Disk{})
	if err != nil {
		return nil, err
	}

	// null check
	if diskObjectSetResp == nil {
		return nil, nil
	}
	return diskObjectSetResp.(*model.Disk), err
}

// GetObjectList returns the list of Disk objects
func (objectSet *DiskObjectSet) GetObjectList() ([]*model.Disk, error) {
	diskObjectSetResp, err := objectSet.Client.List(diskPath)
	if err != nil {
		return nil, err
	}
	return buildDiskObjectSet(diskObjectSetResp), err
}

// GetObjectListFromParams returns the list of Disk objects using the given params query info
func (objectSet *DiskObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Disk, error) {
	diskObjectSetResp, err := objectSet.Client.ListFromParams(diskPath, params)
	if err != nil {
		return nil, err
	}
	return buildDiskObjectSet(diskObjectSetResp), err
}

// generated function to build the appropriate response types
func buildDiskObjectSet(response interface{}) []*model.Disk {
	values := reflect.ValueOf(response)
	results := make([]*model.Disk, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Disk{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
