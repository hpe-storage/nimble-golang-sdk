// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (objectSet *SnapshotCollectionObjectSet) CreateObject(payload *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	snapshotCollectionObjectSetResp, err := objectSet.Client.Post(snapshotCollectionPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if snapshotCollectionObjectSetResp == nil {
		return nil,nil
	}
	return snapshotCollectionObjectSetResp.(*model.SnapshotCollection), err
}

// UpdateObject Modify existing SnapshotCollection object
func (objectSet *SnapshotCollectionObjectSet) UpdateObject(id string, payload *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	snapshotCollectionObjectSetResp, err := objectSet.Client.Put(snapshotCollectionPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if snapshotCollectionObjectSetResp == nil {
		return nil,nil
	}
	return snapshotCollectionObjectSetResp.(*model.SnapshotCollection), err
}

// DeleteObject deletes the SnapshotCollection object with the specified ID
func (objectSet *SnapshotCollectionObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(snapshotCollectionPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a SnapshotCollection object with the given ID
func (objectSet *SnapshotCollectionObjectSet) GetObject(id string) (*model.SnapshotCollection, error) {
	snapshotCollectionObjectSetResp, err := objectSet.Client.Get(snapshotCollectionPath, id, model.SnapshotCollection{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if snapshotCollectionObjectSetResp == nil {
		return nil,nil
	}
	return snapshotCollectionObjectSetResp.(*model.SnapshotCollection), err
}

// GetObjectList returns the list of SnapshotCollection objects
func (objectSet *SnapshotCollectionObjectSet) GetObjectList() ([]*model.SnapshotCollection, error) {
	snapshotCollectionObjectSetResp, err := objectSet.Client.List(snapshotCollectionPath)
	if err != nil {
		return nil, err
	}
	return buildSnapshotCollectionObjectSet(snapshotCollectionObjectSetResp), err
}

// GetObjectListFromParams returns the list of SnapshotCollection objects using the given params query info
func (objectSet *SnapshotCollectionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.SnapshotCollection, error) {
	snapshotCollectionObjectSetResp, err := objectSet.Client.ListFromParams(snapshotCollectionPath, params)
	if err != nil {
		return nil, err
	}
	return buildSnapshotCollectionObjectSet(snapshotCollectionObjectSetResp), err
}
// generated function to build the appropriate response types
func buildSnapshotCollectionObjectSet(response interface{}) ([]*model.SnapshotCollection) {
	values := reflect.ValueOf(response)
	results := make([]*model.SnapshotCollection, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.SnapshotCollection{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
