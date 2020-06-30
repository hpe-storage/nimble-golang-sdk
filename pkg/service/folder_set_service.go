// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FolderSet Service - Folder set represents a set of folder each on separate pools that represent a group-scope datastore spanning the entire group.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// FolderSetService type 
type FolderSetService struct {
	objectSet *client.FolderSetObjectSet
}

// NewFolderSetService - method to initialize "FolderSetService" 
func NewFolderSetService(gs *NsGroupService) (*FolderSetService) {
	objectSet := gs.client.GetFolderSetObjectSet()
	return &FolderSetService{objectSet: objectSet}
}

// GetFolderSets - method returns a array of pointers of type "FolderSets"
func (svc *FolderSetService) GetFolderSets(params *util.GetParams) ([]*model.FolderSet, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetFolderSetsWithFields - method returns a array of pointers of type "FolderSet" 
func (svc *FolderSetService) GetFolderSetsWithFields(fields []string) ([]*model.FolderSet, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFolderSet - method creates a "FolderSet"
func (svc *FolderSetService) CreateFolderSet(obj *model.FolderSet) (*model.FolderSet, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditFolderSet - method modifies  the "FolderSet" 
func (svc *FolderSetService) EditFolderSet(id string, obj *model.FolderSet) (*model.FolderSet, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyFolderSet - private method for more than one element check. 
func onlyFolderSet(objs []*model.FolderSet) (*model.FolderSet, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one FolderSet found with the given filter")
	}

	return objs[0], nil
}

 
// GetFolderSetsByID - method returns associative a array of pointers of type "FolderSet", filter by Id
func (svc *FolderSetService) GetFolderSetsByID(pool *model.Pool, fields []string) (map[string]*model.FolderSet, error) {
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
	objs, err := svc.GetFolderSets(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FolderSet)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetFolderSetById - method returns a pointer to "FolderSet"
func (svc *FolderSetService) GetFolderSetById(id string) (*model.FolderSet, error) {
	return svc.objectSet.GetObject(id)
}

// GetFolderSetsByName - method returns a associative array of pointers of type "FolderSet", filter by name 
func (svc *FolderSetService) GetFolderSetsByName(pool *model.Pool, fields []string) (map[string]*model.FolderSet, error) {
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
	objs, err := svc.GetFolderSets(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FolderSet)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetFolderSetByName - method returns a pointer "FolderSet" 
func (svc *FolderSetService) GetFolderSetByName(name string) (*model.FolderSet, error) {
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
	return onlyFolderSet(objs)
}	

