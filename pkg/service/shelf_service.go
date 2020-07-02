// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Shelf Service - Disk shelf and head unit houses disks and controller.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ShelfService type 
type ShelfService struct {
	objectSet *client.ShelfObjectSet
}

// NewShelfService - method to initialize "ShelfService" 
func NewShelfService(gs *NsGroupService) (*ShelfService) {
	objectSet := gs.client.GetShelfObjectSet()
	return &ShelfService{objectSet: objectSet}
}

// GetShelfs - method returns a array of pointers of type "Shelfs"
func (svc *ShelfService) GetShelfs(params *util.GetParams) ([]*model.Shelf, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateShelf - method creates a "Shelf"
func (svc *ShelfService) CreateShelf(obj *model.Shelf) (*model.Shelf, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateShelf - method modifies  the "Shelf" 
func (svc *ShelfService) UpdateShelf(id string, obj *model.Shelf) (*model.Shelf, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetShelfById - method returns a pointer to "Shelf"
func (svc *ShelfService) GetShelfById(id string) (*model.Shelf, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteShelf - deletes the "Shelf"
func (svc *ShelfService) DeleteShelf(id string) error {
	return svc.objectSet.DeleteObject(id)
}
