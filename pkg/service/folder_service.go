// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Folder Service - Folders are a way to group volumes, as well as a way to apply space constraints to them.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// FolderService type 
type FolderService struct {
	objectSet *client.FolderObjectSet
}

// NewFolderService - method to initialize "FolderService" 
func NewFolderService(gs *NsGroupService) (*FolderService) {
	objectSet := gs.client.GetFolderObjectSet()
	return &FolderService{objectSet: objectSet}
}

// GetFolders - method returns a array of pointers of type "Folders"
func (svc *FolderService) GetFolders(params *util.GetParams) ([]*model.Folder, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetFoldersWithFields - method returns a array of pointers of type "Folder" 
func (svc *FolderService) GetFoldersWithFields(fields []string) ([]*model.Folder, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFolder - method creates a "Folder"
func (svc *FolderService) CreateFolder(obj *model.Folder) (*model.Folder, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditFolder - method modifies  the "Folder" 
func (svc *FolderService) EditFolder(id string, obj *model.Folder) (*model.Folder, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyFolder - private method for more than one element check. 
func onlyFolder(objs []*model.Folder) (*model.Folder, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Folder found with the given filter")
	}

	return objs[0], nil
}

 
// GetFoldersByID - method returns associative a array of pointers of type "Folder", filter by Id
func (svc *FolderService) GetFoldersByID(pool *model.Pool, fields []string) (map[string]*model.Folder, error) {
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
	objs, err := svc.GetFolders(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Folder)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetFolderById - method returns a pointer to "Folder"
func (svc *FolderService) GetFolderById(id string) (*model.Folder, error) {
	return svc.objectSet.GetObject(id)
}

// GetFoldersByName - method returns a associative array of pointers of type "Folder", filter by name 
func (svc *FolderService) GetFoldersByName(pool *model.Pool, fields []string) (map[string]*model.Folder, error) {
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
	objs, err := svc.GetFolders(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Folder)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetFolderByName - method returns a pointer "Folder" 
func (svc *FolderService) GetFolderByName(name string) (*model.Folder, error) {
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
	return onlyFolder(objs)
}	

