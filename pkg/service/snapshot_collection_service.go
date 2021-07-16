// Copyright 2021 Hewlett Packard Enterprise Development LP

package service

// SnapshotCollection Service - Snapshot collections are collections of scheduled snapshots that are taken from volumes sharing a volume collection. Snapshot collections are replicated in the order that the
// collections were taken.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// SnapshotCollectionService type
type SnapshotCollectionService struct {
	objectSet *client.SnapshotCollectionObjectSet
}

// NewSnapshotCollectionService - method to initialize "SnapshotCollectionService"
func NewSnapshotCollectionService(gs *NsGroupService) *SnapshotCollectionService {
	objectSet := gs.client.GetSnapshotCollectionObjectSet()
	return &SnapshotCollectionService{objectSet: objectSet}
}

// GetSnapshotCollections - method returns a array of pointers of type "SnapshotCollections"
func (svc *SnapshotCollectionService) GetSnapshotCollections(params *param.GetParams) ([]*nimbleos.SnapshotCollection, error) {
	snapshotCollectionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return snapshotCollectionResp, nil
}

// CreateSnapshotCollection - method creates a "SnapshotCollection"
func (svc *SnapshotCollectionService) CreateSnapshotCollection(obj *nimbleos.SnapshotCollection) (*nimbleos.SnapshotCollection, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateSnapshotCollection: invalid parameter specified, %v", obj)
	}

	snapshotCollectionResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return snapshotCollectionResp, nil
}

// UpdateSnapshotCollection - method modifies  the "SnapshotCollection"
func (svc *SnapshotCollectionService) UpdateSnapshotCollection(id string, obj *nimbleos.SnapshotCollection) (*nimbleos.SnapshotCollection, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateSnapshotCollection: invalid parameter specified, %v", obj)
	}

	snapshotCollectionResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return snapshotCollectionResp, nil
}

// GetSnapshotCollectionById - method returns a pointer to "SnapshotCollection"
func (svc *SnapshotCollectionService) GetSnapshotCollectionById(id string) (*nimbleos.SnapshotCollection, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetSnapshotCollectionById: invalid parameter specified, %v", id)
	}

	snapshotCollectionResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return snapshotCollectionResp, nil
}

// GetSnapshotCollectionByName - method returns a pointer "SnapshotCollection"
func (svc *SnapshotCollectionService) GetSnapshotCollectionByName(name string) (*nimbleos.SnapshotCollection, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	snapshotCollectionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(snapshotCollectionResp) == 0 {
		return nil, nil
	}

	return snapshotCollectionResp[0], nil
}

// DeleteSnapshotCollection - deletes the "SnapshotCollection"
func (svc *SnapshotCollectionService) DeleteSnapshotCollection(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteSnapshotCollection: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
