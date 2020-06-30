// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SnapshotCollection Service - Snapshot collections are collections of scheduled snapshots that are taken from volumes sharing a volume collection. Snapshot collections are replicated in the order that the
// collections were taken.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetSnapshotCollectionsWithFields - method returns a array of pointers of type "SnapshotCollection" 
func (svc *SnapshotCollectionService) GetSnapshotCollectionsWithFields(fields []string) ([]*model.SnapshotCollection, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSnapshotCollection - method creates a "SnapshotCollection"
func (svc *SnapshotCollectionService) CreateSnapshotCollection(obj *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSnapshotCollection - method modifies  the "SnapshotCollection" 
func (svc *SnapshotCollectionService) EditSnapshotCollection(id string, obj *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySnapshotCollection - private method for more than one element check. 
func onlySnapshotCollection(objs []*model.SnapshotCollection) (*model.SnapshotCollection, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one SnapshotCollection found with the given filter")
	}

	return objs[0], nil
}

 
// GetSnapshotCollectionsByID - method returns associative a array of pointers of type "SnapshotCollection", filter by Id
func (svc *SnapshotCollectionService) GetSnapshotCollectionsByID(pool *model.Pool, fields []string) (map[string]*model.SnapshotCollection, error) {
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
	objs, err := svc.GetSnapshotCollections(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.SnapshotCollection)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetSnapshotCollectionById - method returns a pointer to "SnapshotCollection"
func (svc *SnapshotCollectionService) GetSnapshotCollectionById(id string) (*model.SnapshotCollection, error) {
	return svc.objectSet.GetObject(id)
}

// GetSnapshotCollectionsByName - method returns a associative array of pointers of type "SnapshotCollection", filter by name 
func (svc *SnapshotCollectionService) GetSnapshotCollectionsByName(pool *model.Pool, fields []string) (map[string]*model.SnapshotCollection, error) {
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
	objs, err := svc.GetSnapshotCollections(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.SnapshotCollection)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlySnapshotCollection(objs)
}	

