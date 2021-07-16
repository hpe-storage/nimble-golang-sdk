// Copyright 2021 Hewlett Packard Enterprise Development LP

package service

// Snapshot Service - Snapshots are point-in-time copies of a volume. Snapshots are managed the same way you manage volumes. In reality, snapshots are volumes: they can be accessed by initiators, are
// subject to the same controls, can be modified, and have the same restrictions as volumes. Snapshots can be cloned and replicated. The initial snapshot uses no space: it shares
// the original data with the source volume. Each successive snapshot captures the changes that have occurred on the volume. The changed blocks are compressed.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// SnapshotService type
type SnapshotService struct {
	objectSet *client.SnapshotObjectSet
}

// NewSnapshotService - method to initialize "SnapshotService"
func NewSnapshotService(gs *NsGroupService) *SnapshotService {
	objectSet := gs.client.GetSnapshotObjectSet()
	return &SnapshotService{objectSet: objectSet}
}

// GetSnapshots - method returns a array of pointers of type "Snapshots"
func (svc *SnapshotService) GetSnapshots(params *param.GetParams) ([]*nimbleos.Snapshot, error) {
	snapshotResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return snapshotResp, nil
}

// CreateSnapshot - method creates a "Snapshot"
func (svc *SnapshotService) CreateSnapshot(obj *nimbleos.Snapshot) (*nimbleos.Snapshot, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateSnapshot: invalid parameter specified, %v", obj)
	}

	snapshotResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return snapshotResp, nil
}

// UpdateSnapshot - method modifies  the "Snapshot"
func (svc *SnapshotService) UpdateSnapshot(id string, obj *nimbleos.Snapshot) (*nimbleos.Snapshot, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateSnapshot: invalid parameter specified, %v", obj)
	}

	snapshotResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return snapshotResp, nil
}

// GetSnapshotById - method returns a pointer to "Snapshot"
func (svc *SnapshotService) GetSnapshotById(id string) (*nimbleos.Snapshot, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetSnapshotById: invalid parameter specified, %v", id)
	}

	snapshotResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return snapshotResp, nil
}

// GetSnapshotByName - method returns a list of pointers to "Snapshot"
func (svc *SnapshotService) GetSnapshotByName(volName string) ([]*nimbleos.Snapshot, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.SnapshotFields.VolName,
			Operator:  param.EQUALS.String(),
			Value:     volName,
		},
	}
	snapshotResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(snapshotResp) == 0 {
		return nil, nil
	}
	return snapshotResp, nil
}

// GetSnapshotBySerialNumber method returns a pointer to "Snapshot"
func (svc *SnapshotService) GetSnapshotBySerialNumber(serialNumber string) (*nimbleos.Snapshot, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.SnapshotFields.SerialNumber,
			Operator:  param.EQUALS.String(),
			Value:     serialNumber,
		},
	}
	snapshotResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	if len(snapshotResp) == 0 {
		return nil, nil
	}

	return snapshotResp[0], nil
}

// DeleteSnapshot - deletes the "Snapshot"
func (svc *SnapshotService) DeleteSnapshot(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteSnapshot: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// BulkCreateSnapshots - create snapshots on the given set of volumes.
//   Required parameters:
//       snapVolList - List of volumes to snapshot and corresponding snapshot creation attributes. VSS application-synchronized snapshot must specify the 'writable' parameter and set it to true.
//       replicate - Allow snapshot to be replicated.
//       vssSnap - VSS app-synchronized snapshot; we don't support creation of non app-synchronized sanpshots through this interface; must be set to true.

func (svc *SnapshotService) BulkCreateSnapshots(snapVolList []*nimbleos.NsSnapVol, replicate bool, vssSnap bool) (*nimbleos.NsSnapVolListReturn, error) {

	if len(snapVolList) == 0 {
		return nil, fmt.Errorf("BulkCreateSnapshot: invalid parameter specified snapVolList: %v ", snapVolList)
	}

	resp, err := svc.objectSet.BulkCreate(snapVolList, &replicate, &vssSnap)
	return resp, err
}
