// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Manage volume collections. Volume collections are logical groups of volumes that share protection characteristics such as snapshot and replication schedules. Volume collections
// can be created from scratch or based on predefined protection templates.
const (
    volumeCollectionPath = "volume_collections"
)

// VolumeCollectionObjectSet
type VolumeCollectionObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new VolumeCollection object
func (objectSet *VolumeCollectionObjectSet) CreateObject(payload *model.VolumeCollection) (*model.VolumeCollection, error) {
	response, err := objectSet.Client.Post(volumeCollectionPath, payload)
	return response.(*model.VolumeCollection), err
}

// UpdateObject Modify existing VolumeCollection object
func (objectSet *VolumeCollectionObjectSet) UpdateObject(id string, payload *model.VolumeCollection) (*model.VolumeCollection, error) {
	response, err := objectSet.Client.Put(volumeCollectionPath, id, payload)
	return response.(*model.VolumeCollection), err
}

// DeleteObject deletes the VolumeCollection object with the specified ID
func (objectSet *VolumeCollectionObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(volumeCollectionPath, id)
}

// GetObject returns a VolumeCollection object with the given ID
func (objectSet *VolumeCollectionObjectSet) GetObject(id string) (*model.VolumeCollection, error) {
	response, err := objectSet.Client.Get(volumeCollectionPath, id, model.VolumeCollection{})
	if response == nil {
		return nil, err
	}
	return response.(*model.VolumeCollection), err
}

// GetObjectList returns the list of VolumeCollection objects
func (objectSet *VolumeCollectionObjectSet) GetObjectList() ([]*model.VolumeCollection, error) {
	response, err := objectSet.Client.List(volumeCollectionPath)
	return buildVolumeCollectionObjectSet(response), err
}

// GetObjectListFromParams returns the list of VolumeCollection objects using the given params query info
func (objectSet *VolumeCollectionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.VolumeCollection, error) {
	response, err := objectSet.Client.ListFromParams(volumeCollectionPath, params)
	return buildVolumeCollectionObjectSet(response), err
}

// generated function to build the appropriate response types
func buildVolumeCollectionObjectSet(response interface{}) ([]*model.VolumeCollection) {
	values := reflect.ValueOf(response)
	results := make([]*model.VolumeCollection, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.VolumeCollection{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}