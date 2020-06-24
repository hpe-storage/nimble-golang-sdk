// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * A volume family contains all the volumes, snapshots, and clones derived from and including a root volume.
 *
 */
const (
    volumeFamilyPath = "volume_families"
)

/**
 * VolumeFamilyObjectSet
*/
type VolumeFamilyObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new VolumeFamily object
func (objectSet *VolumeFamilyObjectSet) CreateObject(payload *model.VolumeFamily) (*model.VolumeFamily, error) {
	response, err := objectSet.Client.Post(volumeFamilyPath, payload)
	return response.(*model.VolumeFamily), err
}

// UpdateObject Modify existing VolumeFamily object
func (objectSet *VolumeFamilyObjectSet) UpdateObject(id string, payload *model.VolumeFamily) (*model.VolumeFamily, error) {
	response, err := objectSet.Client.Put(volumeFamilyPath, id, payload)
	return response.(*model.VolumeFamily), err
}

// DeleteObject deletes the VolumeFamily object with the specified ID
func (objectSet *VolumeFamilyObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(volumeFamilyPath, id)
}

// GetObject returns a VolumeFamily object with the given ID
func (objectSet *VolumeFamilyObjectSet) GetObject(id string) (*model.VolumeFamily, error) {
	response, err := objectSet.Client.Get(volumeFamilyPath, id, model.VolumeFamily{})
	if response == nil {
		return nil, err
	}
	return response.(*model.VolumeFamily), err
}

// GetObjectList returns the list of VolumeFamily objects
func (objectSet *VolumeFamilyObjectSet) GetObjectList() ([]*model.VolumeFamily, error) {
	response, err := objectSet.Client.List(volumeFamilyPath)
	return buildVolumeFamilyObjectSet(response), err
}

// GetObjectListFromParams returns the list of VolumeFamily objects using the given params query info
func (objectSet *VolumeFamilyObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.VolumeFamily, error) {
	response, err := objectSet.Client.ListFromParams(volumeFamilyPath, params)
	return buildVolumeFamilyObjectSet(response), err
}

// generated function to build the appropriate response types
func buildVolumeFamilyObjectSet(response interface{}) ([]*model.VolumeFamily) {
	values := reflect.ValueOf(response)
	results := make([]*model.VolumeFamily, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.VolumeFamily{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}