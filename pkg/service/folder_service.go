// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Folder Service - Folders are a way to group volumes, as well as a way to apply space constraints to them.

import (
	"fmt"
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
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	folderResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return folderResp,nil
}

// CreateFolder - method creates a "Folder"
func (svc *FolderService) CreateFolder(obj *model.Folder) (*model.Folder, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	folderResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return folderResp,nil
}

// UpdateFolder - method modifies  the "Folder" 
func (svc *FolderService) UpdateFolder(id string, obj *model.Folder) (*model.Folder, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	folderResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return folderResp,nil
}

// GetFolderById - method returns a pointer to "Folder"
func (svc *FolderService) GetFolderById(id string) (*model.Folder, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	folderResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return folderResp,nil
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
	folderResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(folderResp) == 0 {
    	return nil, nil
    }
    
	return folderResp[0],nil
}	

// DeleteFolder - deletes the "Folder"
func (svc *FolderService) DeleteFolder(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
