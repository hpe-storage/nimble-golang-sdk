// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Shelf Service - Disk shelf and head unit houses disks and controller.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ShelfService type
type ShelfService struct {
	objectSet *client.ShelfObjectSet
}

// NewShelfService - method to initialize "ShelfService"
func NewShelfService(gs *NsGroupService) *ShelfService {
	objectSet := gs.client.GetShelfObjectSet()
	return &ShelfService{objectSet: objectSet}
}

// GetShelfs - method returns a array of pointers of type "Shelfs"
func (svc *ShelfService) GetShelfs(params *util.GetParams) ([]*model.Shelf, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	shelfResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// CreateShelf - method creates a "Shelf"
func (svc *ShelfService) CreateShelf(obj *model.Shelf) (*model.Shelf, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	shelfResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// UpdateShelf - method modifies  the "Shelf"
func (svc *ShelfService) UpdateShelf(id string, obj *model.Shelf) (*model.Shelf, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	shelfResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// GetShelfById - method returns a pointer to "Shelf"
func (svc *ShelfService) GetShelfById(id string) (*model.Shelf, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	shelfResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// DeleteShelf - deletes the "Shelf"
func (svc *ShelfService) DeleteShelf(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
