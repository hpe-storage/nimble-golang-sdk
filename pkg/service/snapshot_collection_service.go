// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SnapshotCollection Service - Snapshot collections are collections of scheduled snapshots that are taken from volumes sharing a volume collection. Snapshot collections are replicated in the order that the
// collections were taken.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// SnapshotCollectionService type 
type SnapshotCollectionService struct {
	objectSet *client.SnapshotCollectionObjectSet
}

// NewSnapshotCollectionService - method to initialize "SnapshotCollectionService" 
func NewSnapshotCollectionService(gs *NsGroupService) (*SnapshotCollectionService) {
	objectSet := gs.client.GetSnapshotCollectionObjectSet()
	return &SnapshotCollectionService{objectSet: objectSet}
}

// GetSnapshotCollections - method returns a array of pointers of type "SnapshotCollections"
func (svc *SnapshotCollectionService) GetSnapshotCollections(params *util.GetParams) ([]*model.SnapshotCollection, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSnapshotCollection - method creates a "SnapshotCollection"
func (svc *SnapshotCollectionService) CreateSnapshotCollection(obj *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateSnapshotCollection - method modifies  the "SnapshotCollection" 
func (svc *SnapshotCollectionService) UpdateSnapshotCollection(id string, obj *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetSnapshotCollectionById - method returns a pointer to "SnapshotCollection"
func (svc *SnapshotCollectionService) GetSnapshotCollectionById(id string) (*model.SnapshotCollection, error) {
	return svc.objectSet.GetObject(id)
}

// GetSnapshotCollectionByName - method returns a pointer "SnapshotCollection" 
func (svc *SnapshotCollectionService) GetSnapshotCollectionByName(name string) (*model.SnapshotCollection, error) {
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

// DeleteSnapshotCollection - deletes the "SnapshotCollection"
func (svc *SnapshotCollectionService) DeleteSnapshotCollection(id string) error {
	return svc.objectSet.DeleteObject(id)
}
