// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on storage allocation.
 *
 */
const (
    volumePath = "volumes"
)

/**
 * VolumeObjectSet
*/
type VolumeObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Volume object
func (objectSet *VolumeObjectSet) CreateObject(payload *model.Volume) (*model.Volume, error) {
	response, err := objectSet.Client.Post(volumePath, payload)
	return response.(*model.Volume), err
}

// UpdateObject Modify existing Volume object
func (objectSet *VolumeObjectSet) UpdateObject(id string, payload *model.Volume) (*model.Volume, error) {
	response, err := objectSet.Client.Put(volumePath, id, payload)
	return response.(*model.Volume), err
}

// DeleteObject deletes the Volume object with the specified ID
func (objectSet *VolumeObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(volumePath, id)
}

// GetObject returns a Volume object with the given ID
func (objectSet *VolumeObjectSet) GetObject(id string) (*model.Volume, error) {
	response, err := objectSet.Client.Get(volumePath, id, model.Volume{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Volume), err
}

// GetObjectList returns the list of Volume objects
func (objectSet *VolumeObjectSet) GetObjectList() ([]*model.Volume, error) {
	response, err := objectSet.Client.List(volumePath)
	return buildVolumeObjectSet(response), err
}

// GetObjectListFromParams returns the list of Volume objects using the given params query info
func (objectSet *VolumeObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Volume, error) {
	response, err := objectSet.Client.ListFromParams(volumePath, params)
	return buildVolumeObjectSet(response), err
}

// generated function to build the appropriate response types
func buildVolumeObjectSet(response interface{}) ([]*model.Volume) {
	values := reflect.ValueOf(response)
	results := make([]*model.Volume, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Volume{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}