// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Snapshot Service - Snapshots are point-in-time copies of a volume. Snapshots are managed the same way you manage volumes. In reality, snapshots are volumes: they can be accessed by initiators, are
// subject to the same controls, can be modified, and have the same restrictions as volumes. Snapshots can be cloned and replicated. The initial snapshot uses no space: it shares
// the original data with the source volume. Each successive snapshot captures the changes that have occurred on the volume. The changed blocks are compressed.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// SnapshotService type 
type SnapshotService struct {
	objectSet *client.SnapshotObjectSet
}

// NewSnapshotService - method to initialize "SnapshotService" 
func NewSnapshotService(gs *NsGroupService) (*SnapshotService) {
	objectSet := gs.client.GetSnapshotObjectSet()
	return &SnapshotService{objectSet: objectSet}
}

// GetSnapshots - method returns a array of pointers of type "Snapshots"
func (svc *SnapshotService) GetSnapshots(params *util.GetParams) ([]*model.Snapshot, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSnapshot - method creates a "Snapshot"
func (svc *SnapshotService) CreateSnapshot(obj *model.Snapshot) (*model.Snapshot, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateSnapshot - method modifies  the "Snapshot" 
func (svc *SnapshotService) UpdateSnapshot(id string, obj *model.Snapshot) (*model.Snapshot, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetSnapshotById - method returns a pointer to "Snapshot"
func (svc *SnapshotService) GetSnapshotById(id string) (*model.Snapshot, error) {
	return svc.objectSet.GetObject(id)
}

// GetSnapshotByName - method returns a pointer "Snapshot" 
func (svc *SnapshotService) GetSnapshotByName(name string) (*model.Snapshot, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	
// GetSnapshotBySerialNumber method returns a pointer to "Snapshot"
func (svc *SnapshotService) GetSnapshotBySerialNumber(serialNumber string) (*model.Snapshot, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.SnapshotFields.SerialNumber,
			Operator:  util.EQUALS.String(),
			Value:     serialNumber,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}

// DeleteSnapshot - deletes the "Snapshot"
func (svc *SnapshotService) DeleteSnapshot(id string) error {
	return svc.objectSet.DeleteObject(id)
}
