// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Folder Service - Folders are a way to group volumes, as well as a way to apply space constraints to them.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateFolder - method creates a "Folder"
func (svc *FolderService) CreateFolder(obj *model.Folder) (*model.Folder, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateFolder - method modifies  the "Folder" 
func (svc *FolderService) UpdateFolder(id string, obj *model.Folder) (*model.Folder, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetFolderById - method returns a pointer to "Folder"
func (svc *FolderService) GetFolderById(id string) (*model.Folder, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteFolder - deletes the "Folder"
func (svc *FolderService) DeleteFolder(id string) error {
	return svc.objectSet.DeleteObject(id)
}
