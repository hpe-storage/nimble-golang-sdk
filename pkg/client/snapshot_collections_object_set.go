// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Snapshot collections are collections of scheduled snapshots that are taken from volumes sharing a volume collection. Snapshot collections are replicated in the order that the
// collections were taken.
const (
	snapshotCollectionPath = "snapshot_collections"
)

// SnapshotCollectionObjectSet
type SnapshotCollectionObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new SnapshotCollection object
func (objectSet *SnapshotCollectionObjectSet) CreateObject(payload *nimbleos.SnapshotCollection) (*nimbleos.SnapshotCollection, error) {
	resp, err := objectSet.Client.Post(snapshotCollectionPath, payload, &nimbleos.SnapshotCollection{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.SnapshotCollection), err
}

// UpdateObject Modify existing SnapshotCollection object
func (objectSet *SnapshotCollectionObjectSet) UpdateObject(id string, payload *nimbleos.SnapshotCollection) (*nimbleos.SnapshotCollection, error) {
	resp, err := objectSet.Client.Put(snapshotCollectionPath, id, payload, &nimbleos.SnapshotCollection{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.SnapshotCollection), err
}

// DeleteObject deletes the SnapshotCollection object with the specified ID
func (objectSet *SnapshotCollectionObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(snapshotCollectionPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a SnapshotCollection object with the given ID
func (objectSet *SnapshotCollectionObjectSet) GetObject(id string) (*nimbleos.SnapshotCollection, error) {
	resp, err := objectSet.Client.Get(snapshotCollectionPath, id, &nimbleos.SnapshotCollection{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.SnapshotCollection), err
}

// GetObjectList returns the list of SnapshotCollection objects
func (objectSet *SnapshotCollectionObjectSet) GetObjectList() ([]*nimbleos.SnapshotCollection, error) {
	resp, err := objectSet.Client.List(snapshotCollectionPath)
	if err != nil {
		return nil, err
	}
	return buildSnapshotCollectionObjectSet(resp), err
}

// GetObjectListFromParams returns the list of SnapshotCollection objects using the given params query info
func (objectSet *SnapshotCollectionObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.SnapshotCollection, error) {
	snapshotCollectionObjectSetResp, err := objectSet.Client.ListFromParams(snapshotCollectionPath, params)
	if err != nil {
		return nil, err
	}
	return buildSnapshotCollectionObjectSet(snapshotCollectionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSnapshotCollectionObjectSet(response interface{}) []*nimbleos.SnapshotCollection {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.SnapshotCollection, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.SnapshotCollection{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
