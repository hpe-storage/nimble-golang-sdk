// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *VolumeCollectionObjectSet) CreateObject(payload *nimbleos.VolumeCollection) (*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.Post(volumeCollectionPath, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return resp.(*nimbleos.VolumeCollection), err
}

// UpdateObject Modify existing VolumeCollection object
func (objectSet *VolumeCollectionObjectSet) UpdateObject(id string, payload *nimbleos.VolumeCollection) (*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.Put(volumeCollectionPath, id, payload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.VolumeCollection), err
}

// DeleteObject deletes the VolumeCollection object with the specified ID
func (objectSet *VolumeCollectionObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(volumeCollectionPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a VolumeCollection object with the given ID
func (objectSet *VolumeCollectionObjectSet) GetObject(id string) (*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.Get(volumeCollectionPath, id, nimbleos.VolumeCollection{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.VolumeCollection), err
}

// GetObjectList returns the list of VolumeCollection objects
func (objectSet *VolumeCollectionObjectSet) GetObjectList() ([]*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.List(volumeCollectionPath)
	if err != nil {
		return nil, err
	}
	return buildVolumeCollectionObjectSet(resp), err
}

// GetObjectListFromParams returns the list of VolumeCollection objects using the given params query info
func (objectSet *VolumeCollectionObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.VolumeCollection, error) {
	volumeCollectionObjectSetResp, err := objectSet.Client.ListFromParams(volumeCollectionPath, params)
	if err != nil {
		return nil, err
	}
	return buildVolumeCollectionObjectSet(volumeCollectionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildVolumeCollectionObjectSet(response interface{}) []*nimbleos.VolumeCollection {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.VolumeCollection, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.VolumeCollection{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
