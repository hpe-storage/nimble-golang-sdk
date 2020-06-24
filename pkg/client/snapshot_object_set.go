// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Snapshots are point-in-time copies of a volume. Snapshots are managed the same way you manage volumes. In reality, snapshots are volumes: they can be accessed by initiators, are subject to the same controls, can be modified, and have the same restrictions as volumes. Snapshots can be cloned and replicated. The initial snapshot uses no space: it shares the original data with the source volume. Each successive snapshot captures the changes that have occurred on the volume. The changed blocks are compressed.
 *
 */
const (
    snapshotPath = "snapshots"
)

/**
 * SnapshotObjectSet
*/
type SnapshotObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Snapshot object
func (objectSet *SnapshotObjectSet) CreateObject(payload *model.Snapshot) (*model.Snapshot, error) {
	response, err := objectSet.Client.Post(snapshotPath, payload)
	return response.(*model.Snapshot), err
}

// UpdateObject Modify existing Snapshot object
func (objectSet *SnapshotObjectSet) UpdateObject(id string, payload *model.Snapshot) (*model.Snapshot, error) {
	response, err := objectSet.Client.Put(snapshotPath, id, payload)
	return response.(*model.Snapshot), err
}

// DeleteObject deletes the Snapshot object with the specified ID
func (objectSet *SnapshotObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(snapshotPath, id)
}

// GetObject returns a Snapshot object with the given ID
func (objectSet *SnapshotObjectSet) GetObject(id string) (*model.Snapshot, error) {
	response, err := objectSet.Client.Get(snapshotPath, id, model.Snapshot{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Snapshot), err
}

// GetObjectList returns the list of Snapshot objects
func (objectSet *SnapshotObjectSet) GetObjectList() ([]*model.Snapshot, error) {
	response, err := objectSet.Client.List(snapshotPath)
	return buildSnapshotObjectSet(response), err
}

// GetObjectListFromParams returns the list of Snapshot objects using the given params query info
func (objectSet *SnapshotObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Snapshot, error) {
	response, err := objectSet.Client.ListFromParams(snapshotPath, params)
	return buildSnapshotObjectSet(response), err
}

// generated function to build the appropriate response types
func buildSnapshotObjectSet(response interface{}) ([]*model.Snapshot) {
	values := reflect.ValueOf(response)
	results := make([]*model.Snapshot, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Snapshot{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}