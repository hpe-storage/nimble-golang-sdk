// Copyright 2020 Hewlett Packard Enterprise Development LP

// create true

// update true
package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
	"strings"
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
func (objectSet *SnapshotObjectSet) CreateObject(payload *nimbleos.Snapshot) (*nimbleos.Snapshot, error) {
	resp, err := objectSet.Client.Post(snapshotPath, payload, &nimbleos.Snapshot{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)
			}
		}
		return nil, err
	}

	return resp.(*nimbleos.Snapshot), err
}

// UpdateObject Modify existing Snapshot object
func (objectSet *SnapshotObjectSet) UpdateObject(id string, payload *nimbleos.Snapshot) (*nimbleos.Snapshot, error) {
	resp, err := objectSet.Client.Put(snapshotPath, id, payload, &nimbleos.Snapshot{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)

			}
		}
		return nil, err
	}

	return resp.(*nimbleos.Snapshot), err
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
func (objectSet *SnapshotObjectSet) GetObject(id string) (*nimbleos.Snapshot, error) {
	resp, err := objectSet.Client.Get(snapshotPath, id, &nimbleos.Snapshot{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Snapshot), err
}

// GetObjectList returns the list of Snapshot objects
func (objectSet *SnapshotObjectSet) GetObjectList() ([]*nimbleos.Snapshot, error) {
	resp, err := objectSet.Client.List(snapshotPath)
	if err != nil {
		return nil, err
	}
	return buildSnapshotObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Snapshot objects using the given params query info
func (objectSet *SnapshotObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Snapshot, error) {
	snapshotObjectSetResp, err := objectSet.Client.ListFromParams(snapshotPath, params)
	if err != nil {
		return nil, err
	}
	return buildSnapshotObjectSet(snapshotObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSnapshotObjectSet(response interface{}) []*nimbleos.Snapshot {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Snapshot, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Snapshot{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// BulkCreate - Create snapshots on the given set of volumes.
func (objectSet *SnapshotObjectSet) BulkCreate(snapVolList []*nimbleos.NsSnapVol, replicate *bool, vssSnap *bool) (*nimbleos.NsSnapVolListReturn, error) {

	bulkCreateUri := snapshotPath
	bulkCreateUri = bulkCreateUri + "/actions/" + "bulk_create"

	payload := &struct {
		SnapVolList []*nimbleos.NsSnapVol `json:"snap_vol_list,omitempty"`
		Replicate   *bool                 `json:"replicate,omitempty"`
		VssSnap     *bool                 `json:"vss_snap,omitempty"`
	}{
		snapVolList,
		replicate,
		vssSnap,
	}

	resp, err := objectSet.Client.Post(bulkCreateUri, payload, &nimbleos.NsSnapVolListReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsSnapVolListReturn), err
}
