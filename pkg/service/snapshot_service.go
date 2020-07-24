// Copyright 2020 Hewlett Packard Enterprise Development LP

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
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	snapshotResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return snapshotResp, nil
}

// CreateSnapshot - method creates a "Snapshot"
func (svc *SnapshotService) CreateSnapshot(obj *nimbleos.Snapshot) (*nimbleos.Snapshot, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
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
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
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
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	snapshotResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return snapshotResp, nil
}

// GetSnapshotByName - method returns a pointer "Snapshot"
func (svc *SnapshotService) GetSnapshotByName(name string) (*nimbleos.Snapshot, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
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
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
