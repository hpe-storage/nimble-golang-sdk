// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.
const (
	volumePath = "volumes"
)

// VolumeObjectSet
type VolumeObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Volume object
func (objectSet *VolumeObjectSet) CreateObject(payload *nimbleos.Volume) (*nimbleos.Volume, error) {
	resp, err := objectSet.Client.Post(volumePath, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return resp.(*nimbleos.Volume), err
}

// UpdateObject Modify existing Volume object
func (objectSet *VolumeObjectSet) UpdateObject(id string, payload *nimbleos.Volume) (*nimbleos.Volume, error) {
	resp, err := objectSet.Client.Put(volumePath, id, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Volume), err
}

// DeleteObject deletes the Volume object with the specified ID
func (objectSet *VolumeObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(volumePath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Volume object with the given ID
func (objectSet *VolumeObjectSet) GetObject(id string) (*nimbleos.Volume, error) {
	resp, err := objectSet.Client.Get(volumePath, id, nimbleos.Volume{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Volume), err
}

// GetObjectList returns the list of Volume objects
func (objectSet *VolumeObjectSet) GetObjectList() ([]*nimbleos.Volume, error) {
	resp, err := objectSet.Client.List(volumePath)
	if err != nil {
		return nil, err
	}
	return buildVolumeObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Volume objects using the given params query info
func (objectSet *VolumeObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Volume, error) {
	volumeObjectSetResp, err := objectSet.Client.ListFromParams(volumePath, params)
	if err != nil {
		return nil, err
	}
	return buildVolumeObjectSet(volumeObjectSetResp), err
}

// generated function to build the appropriate response types
func buildVolumeObjectSet(response interface{}) []*nimbleos.Volume {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Volume, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Volume{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
