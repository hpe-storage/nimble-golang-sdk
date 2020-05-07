// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

const (
	volumePath = "volumes"
)

// VolumeObjectSet provides a wrapper to manage volumes from the client
type VolumeObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new volume
func (objectSet *VolumeObjectSet) CreateObject(payload *model.Volume) (*model.Volume, error) {
	response, err := objectSet.Client.Post(volumePath, payload)
	return response.(*model.Volume), err
}

func (objectSet *VolumeObjectSet) UpdateObject(id string, payload *model.Volume) (*model.Volume, error) {
	response, err := objectSet.Client.Put(volumePath, id, payload)
	return response.(*model.Volume), err
}

// DeleteObject deletes the volume with the specified ID
func (objectSet *VolumeObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(volumePath, id)
}

// GetObject returns a volume with the given ID
func (objectSet *VolumeObjectSet) GetObject(id string) (*model.Volume, error) {
	response, err := objectSet.Client.Get(volumePath, id, model.Volume{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Volume), err
}

// GetObjectList returns the list of volume objects
func (objectSet *VolumeObjectSet) GetObjectList() ([]*model.Volume, error) {
	response, err := objectSet.Client.List(volumePath)
	return buildVolumes(response), err
}

// GetObjectListFromParams returns the list of volume objects using the given params query info
func (objectSet *VolumeObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Volume, error) {
	response, err := objectSet.Client.ListFromParams(volumePath, params)
	return buildVolumes(response), err
}

// generated function to build the appropriate response types
func buildVolumes(response interface{}) []*model.Volume {
	values := reflect.ValueOf(response)
	results := make([]*model.Volume, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Volume{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
