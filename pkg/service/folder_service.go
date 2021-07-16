// Copyright 2021 Hewlett Packard Enterprise Development LP

package service

// Folder Service - Folders are a way to group volumes, as well as a way to apply space constraints to them.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// FolderService type
type FolderService struct {
	objectSet *client.FolderObjectSet
}

// NewFolderService - method to initialize "FolderService"
func NewFolderService(gs *NsGroupService) *FolderService {
	objectSet := gs.client.GetFolderObjectSet()
	return &FolderService{objectSet: objectSet}
}

// GetFolders - method returns a array of pointers of type "Folders"
func (svc *FolderService) GetFolders(params *param.GetParams) ([]*nimbleos.Folder, error) {
	folderResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return folderResp, nil
}

// CreateFolder - method creates a "Folder"
func (svc *FolderService) CreateFolder(obj *nimbleos.Folder) (*nimbleos.Folder, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateFolder: invalid parameter specified, %v", obj)
	}

	folderResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return folderResp, nil
}

// UpdateFolder - method modifies  the "Folder"
func (svc *FolderService) UpdateFolder(id string, obj *nimbleos.Folder) (*nimbleos.Folder, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateFolder: invalid parameter specified, %v", obj)
	}

	folderResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return folderResp, nil
}

// GetFolderById - method returns a pointer to "Folder"
func (svc *FolderService) GetFolderById(id string) (*nimbleos.Folder, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetFolderById: invalid parameter specified, %v", id)
	}

	folderResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return folderResp, nil
}

// GetFolderByName - method returns a pointer "Folder"
func (svc *FolderService) GetFolderByName(name string) (*nimbleos.Folder, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
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

	return folderResp[0], nil
}

// DeleteFolder - deletes the "Folder"
func (svc *FolderService) DeleteFolder(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteFolder: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// SetDedupeFolder - set dedupe enabled/disabled for all applicable volumes inside a folder.
//   Required parameters:
//       dedupeEnabled - Enable/disable dedupe.
//       id - Folder containing the volumes to enable/disable dedupe on.

func (svc *FolderService) SetDedupeFolder(dedupeEnabled bool, id string) error {

	if len(id) == 0 {
		return fmt.Errorf("SetDedupeFolder: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.SetDedupe(&dedupeEnabled, &id)
	return err
}
