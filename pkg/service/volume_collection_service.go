// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// VolumeCollection Service - Manage volume collections. Volume collections are logical groups of volumes that share protection characteristics such as snapshot and replication schedules. Volume collections
// can be created from scratch or based on predefined protection templates.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// VolumeCollectionService type 
type VolumeCollectionService struct {
	objectSet *client.VolumeCollectionObjectSet
}

// NewVolumeCollectionService - method to initialize "VolumeCollectionService" 
func NewVolumeCollectionService(gs *NsGroupService) (*VolumeCollectionService) {
	objectSet := gs.client.GetVolumeCollectionObjectSet()
	return &VolumeCollectionService{objectSet: objectSet}
}

// GetVolumeCollections - method returns a array of pointers of type "VolumeCollections"
func (svc *VolumeCollectionService) GetVolumeCollections(params *util.GetParams) ([]*model.VolumeCollection, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetVolumeCollectionsWithFields - method returns a array of pointers of type "VolumeCollection" 
func (svc *VolumeCollectionService) GetVolumeCollectionsWithFields(fields []string) ([]*model.VolumeCollection, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateVolumeCollection - method creates a "VolumeCollection"
func (svc *VolumeCollectionService) CreateVolumeCollection(obj *model.VolumeCollection) (*model.VolumeCollection, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditVolumeCollection - method modifies  the "VolumeCollection" 
func (svc *VolumeCollectionService) EditVolumeCollection(id string, obj *model.VolumeCollection) (*model.VolumeCollection, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyVolumeCollection - private method for more than one element check. 
func onlyVolumeCollection(objs []*model.VolumeCollection) (*model.VolumeCollection, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one VolumeCollection found with the given filter")
	}

	return objs[0], nil
}

 
// GetVolumeCollectionsByID - method returns associative a array of pointers of type "VolumeCollection", filter by Id
func (svc *VolumeCollectionService) GetVolumeCollectionsByID(pool *model.Pool, fields []string) (map[string]*model.VolumeCollection, error) {
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
	objs, err := svc.GetVolumeCollections(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.VolumeCollection)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetVolumeCollectionById - method returns a pointer to "VolumeCollection"
func (svc *VolumeCollectionService) GetVolumeCollectionById(id string) (*model.VolumeCollection, error) {
	return svc.objectSet.GetObject(id)
}

// GetVolumeCollectionsByName - method returns a associative array of pointers of type "VolumeCollection", filter by name 
func (svc *VolumeCollectionService) GetVolumeCollectionsByName(pool *model.Pool, fields []string) (map[string]*model.VolumeCollection, error) {
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
	objs, err := svc.GetVolumeCollections(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.VolumeCollection)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetVolumeCollectionByName - method returns a pointer "VolumeCollection" 
func (svc *VolumeCollectionService) GetVolumeCollectionByName(name string) (*model.VolumeCollection, error) {
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
	return onlyVolumeCollection(objs)
}	

