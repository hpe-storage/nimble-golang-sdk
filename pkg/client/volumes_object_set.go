// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.
const (
	volumePath = "volumes"
)

// VolumeObjectSet
type VolumeObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Volume object
func (objectSet *VolumeObjectSet) CreateObject(payload *nimbleos.Volume) (*nimbleos.Volume, error) {
	resp, err := objectSet.Client.Post(volumePath, payload, &nimbleos.Volume{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Volume), err
}

// UpdateObject Modify existing Volume object
func (objectSet *VolumeObjectSet) UpdateObject(id string, payload *nimbleos.Volume) (*nimbleos.Volume, error) {
	resp, err := objectSet.Client.Put(volumePath, id, payload, &nimbleos.Volume{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Volume), err
}

// DeleteObject deletes the Volume object with the specified ID
func (objectSet *VolumeObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(volumePath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Volume object with the given ID
func (objectSet *VolumeObjectSet) GetObject(id string) (*nimbleos.Volume, error) {
	resp, err := objectSet.Client.Get(volumePath, id, &nimbleos.Volume{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Volume), err
}

// GetObjectList returns the list of Volume objects
func (objectSet *VolumeObjectSet) GetObjectList() ([]*nimbleos.Volume, error) {
	resp, err := objectSet.Client.List(volumePath)
	if err != nil {
		return nil, err
	}
	return buildVolumeObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Volume objects using the given params query info
func (objectSet *VolumeObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Volume, error) {
	volumeObjectSetResp, err := objectSet.Client.ListFromParams(volumePath, params)
	if err != nil {
		return nil, err
	}
	return buildVolumeObjectSet(volumeObjectSetResp), err
}

// generated function to build the appropriate response types
func buildVolumeObjectSet(response interface{}) []*nimbleos.Volume {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Volume, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Volume{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Restore - Restore volume data from a previous snapshot.
func (objectSet *VolumeObjectSet) Restore(id *string, baseSnapId *string) error {
	restoreUri := volumePath
	restoreUri = restoreUri + "/" + *id
	restoreUri = restoreUri + "/actions/" + "restore"

	payload := &struct {
		Id         *string `json:"id,omitempty"`
		BaseSnapId *string `json:"base_snap_id,omitempty"`
	}{
		id,
		baseSnapId,
	}

	_, err := objectSet.Client.Post(restoreUri, payload, &nimbleos.Volume{})
	return err
}

// Move - Move a volume and its related volumes to another pool. To change a single volume's folder assignment (while remaining in the same pool), use a volume update operation to change the folder_id attribute.
func (objectSet *VolumeObjectSet) Move(id *string, destPoolId *string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error) {
	moveUri := volumePath
	moveUri = moveUri + "/" + *id
	moveUri = moveUri + "/actions/" + "move"

	payload := &struct {
		Id         *string `json:"id,omitempty"`
		DestPoolId *string `json:"dest_pool_id,omitempty"`
		ForceVvol  *bool   `json:"force_vvol,omitempty"`
	}{
		id,
		destPoolId,
		forceVvol,
	}

	resp, err := objectSet.Client.Post(moveUri, payload, &nimbleos.NsVolumeListReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsVolumeListReturn), err
}

// BulkMove - Move volumes and their related volumes to another pool. To change a single volume's folder assignment (while remaining in the same pool), use a volume update operation to change the folder_id attribute.
func (objectSet *VolumeObjectSet) BulkMove(volIds []*string, destPoolId *string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error) {
	bulkMoveUri := volumePath
	bulkMoveUri = bulkMoveUri + "/actions/" + "bulk_move"

	payload := &struct {
		VolIds     []*string `json:"vol_ids,omitempty"`
		DestPoolId *string   `json:"dest_pool_id,omitempty"`
		ForceVvol  *bool     `json:"force_vvol,omitempty"`
	}{
		volIds,
		destPoolId,
		forceVvol,
	}

	resp, err := objectSet.Client.Post(bulkMoveUri, payload, &nimbleos.NsVolumeListReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsVolumeListReturn), err
}

// AbortMove - Abort the in-progress move of the specified volume to another pool.
func (objectSet *VolumeObjectSet) AbortMove(id *string) error {
	abortMoveUri := volumePath
	abortMoveUri = abortMoveUri + "/" + *id
	abortMoveUri = abortMoveUri + "/actions/" + "abort_move"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	_, err := objectSet.Client.Post(abortMoveUri, payload, &nimbleos.Volume{})
	return err
}

// BulkSetDedupe - Enable or disable dedupe on a list of volumes. If the volumes are not dedupe capable, the operation will fail for the specified volume.
func (objectSet *VolumeObjectSet) BulkSetDedupe(volIds []*string, dedupeEnabled *bool) error {
	bulkSetDedupeUri := volumePath
	bulkSetDedupeUri = bulkSetDedupeUri + "/actions/" + "bulk_set_dedupe"

	payload := &struct {
		VolIds        []*string `json:"vol_ids,omitempty"`
		DedupeEnabled *bool     `json:"dedupe_enabled,omitempty"`
	}{
		volIds,
		dedupeEnabled,
	}

	_, err := objectSet.Client.Post(bulkSetDedupeUri, payload, &nimbleos.Volume{})
	return err
}

// BulkSetOnlineAndOffline - Bring a list of volumes online or offline.
func (objectSet *VolumeObjectSet) BulkSetOnlineAndOffline(volIds []*string, online *bool) error {
	bulkSetOnlineAndOfflineUri := volumePath
	bulkSetOnlineAndOfflineUri = bulkSetOnlineAndOfflineUri + "/actions/" + "bulk_set_online_and_offline"

	payload := &struct {
		VolIds []*string `json:"vol_ids,omitempty"`
		Online *bool     `json:"online,omitempty"`
	}{
		volIds,
		online,
	}

	_, err := objectSet.Client.Post(bulkSetOnlineAndOfflineUri, payload, &nimbleos.Volume{})
	return err
}
