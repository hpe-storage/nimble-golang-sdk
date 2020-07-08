// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
)

// Snapshots are point-in-time copies of a volume. Snapshots are managed the same way you manage volumes. In reality, snapshots are volumes: they can be accessed by initiators, are
// subject to the same controls, can be modified, and have the same restrictions as volumes. Snapshots can be cloned and replicated. The initial snapshot uses no space: it shares
// the original data with the source volume. Each successive snapshot captures the changes that have occurred on the volume. The changed blocks are compressed.
const (
	snapshotPath = "snapshots"
)

// SnapshotObjectSet
type SnapshotObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Snapshot object
func (objectSet *SnapshotObjectSet) CreateObject(payload *model.Snapshot) (*model.Snapshot, error) {
	newPayload, err := model.EncodeSnapshot(payload)
	snapshotObjectSetResp, err := objectSet.Client.Post(snapshotPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if snapshotObjectSetResp == nil {
		return nil, nil
	}

	return model.DecodeSnapshot(snapshotObjectSetResp)
}

// UpdateObject Modify existing Snapshot object
func (objectSet *SnapshotObjectSet) UpdateObject(id string, payload *model.Snapshot) (*model.Snapshot, error) {
	newPayload, err := model.EncodeSnapshot(payload)
	snapshotObjectSetResp, err := objectSet.Client.Put(snapshotPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if snapshotObjectSetResp == nil {
		return nil, nil
	}
	return model.DecodeSnapshot(snapshotObjectSetResp)
}

// DeleteObject deletes the Snapshot object with the specified ID
func (objectSet *SnapshotObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(snapshotPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Snapshot object with the given ID
func (objectSet *SnapshotObjectSet) GetObject(id string) (*model.Snapshot, error) {
	snapshotObjectSetResp, err := objectSet.Client.Get(snapshotPath, id, model.Snapshot{})
	if err != nil {
		return nil, err
	}

	// null check
	if snapshotObjectSetResp == nil {
		return nil, nil
	}
	return snapshotObjectSetResp.(*model.Snapshot), err
}

// GetObjectList returns the list of Snapshot objects
func (objectSet *SnapshotObjectSet) GetObjectList() ([]*model.Snapshot, error) {
	snapshotObjectSetResp, err := objectSet.Client.List(snapshotPath)
	if err != nil {
		return nil, err
	}
	return buildSnapshotObjectSet(snapshotObjectSetResp), err
}

// GetObjectListFromParams returns the list of Snapshot objects using the given params query info
func (objectSet *SnapshotObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Snapshot, error) {
	snapshotObjectSetResp, err := objectSet.Client.ListFromParams(snapshotPath, params)
	if err != nil {
		return nil, err
	}
	return buildSnapshotObjectSet(snapshotObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSnapshotObjectSet(response interface{}) []*model.Snapshot {
	values := reflect.ValueOf(response)
	results := make([]*model.Snapshot, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Snapshot{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
