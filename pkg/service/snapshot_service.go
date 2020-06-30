// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Snapshot Service - Snapshots are point-in-time copies of a volume. Snapshots are managed the same way you manage volumes. In reality, snapshots are volumes: they can be accessed by initiators, are
// subject to the same controls, can be modified, and have the same restrictions as volumes. Snapshots can be cloned and replicated. The initial snapshot uses no space: it shares
// the original data with the source volume. Each successive snapshot captures the changes that have occurred on the volume. The changed blocks are compressed.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetSnapshotsWithFields - method returns a array of pointers of type "Snapshot" 
func (svc *SnapshotService) GetSnapshotsWithFields(fields []string) ([]*model.Snapshot, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSnapshot - method creates a "Snapshot"
func (svc *SnapshotService) CreateSnapshot(obj *model.Snapshot) (*model.Snapshot, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSnapshot - method modifies  the "Snapshot" 
func (svc *SnapshotService) EditSnapshot(id string, obj *model.Snapshot) (*model.Snapshot, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySnapshot - private method for more than one element check. 
func onlySnapshot(objs []*model.Snapshot) (*model.Snapshot, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Snapshot found with the given filter")
	}

	return objs[0], nil
}

 
// GetSnapshotsByID - method returns associative a array of pointers of type "Snapshot", filter by Id
func (svc *SnapshotService) GetSnapshotsByID(pool *model.Pool, fields []string) (map[string]*model.Snapshot, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetSnapshots(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Snapshot)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetSnapshotById - method returns a pointer to "Snapshot"
func (svc *SnapshotService) GetSnapshotById(id string) (*model.Snapshot, error) {
	return svc.objectSet.GetObject(id)
}

// GetSnapshotsByName - method returns a associative array of pointers of type "Snapshot", filter by name 
func (svc *SnapshotService) GetSnapshotsByName(pool *model.Pool, fields []string) (map[string]*model.Snapshot, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetSnapshots(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Snapshot)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlySnapshot(objs)
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
	return onlySnapshot(objs)
}

