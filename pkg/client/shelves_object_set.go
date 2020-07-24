// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Disk shelf and head unit houses disks and controller.
const (
	shelfPath = "shelves"
)

// ShelfObjectSet
type ShelfObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Shelf object
func (objectSet *ShelfObjectSet) CreateObject(payload *nimbleos.Shelf) (*nimbleos.Shelf, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Shelf")
}

// UpdateObject Modify existing Shelf object
func (objectSet *ShelfObjectSet) UpdateObject(id string, payload *nimbleos.Shelf) (*nimbleos.Shelf, error) {
	resp, err := objectSet.Client.Put(shelfPath, id, payload, &nimbleos.Shelf{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Shelf), err
}

// DeleteObject deletes the Shelf object with the specified ID
func (objectSet *ShelfObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Shelf")
}

// GetObject returns a Shelf object with the given ID
func (objectSet *ShelfObjectSet) GetObject(id string) (*nimbleos.Shelf, error) {
	resp, err := objectSet.Client.Get(shelfPath, id, nimbleos.Shelf{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Shelf), err
}

// GetObjectList returns the list of Shelf objects
func (objectSet *ShelfObjectSet) GetObjectList() ([]*nimbleos.Shelf, error) {
	resp, err := objectSet.Client.List(shelfPath)
	if err != nil {
		return nil, err
	}
	return buildShelfObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Shelf objects using the given params query info
func (objectSet *ShelfObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Shelf, error) {
	shelfObjectSetResp, err := objectSet.Client.ListFromParams(shelfPath, params)
	if err != nil {
		return nil, err
	}
	return buildShelfObjectSet(shelfObjectSetResp), err
}

// generated function to build the appropriate response types
func buildShelfObjectSet(response interface{}) []*nimbleos.Shelf {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Shelf, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Shelf{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
