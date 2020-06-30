// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Shelf Service - Disk shelf and head unit houses disks and controller.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetShelfsWithFields - method returns a array of pointers of type "Shelf" 
func (svc *ShelfService) GetShelfsWithFields(fields []string) ([]*model.Shelf, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateShelf - method creates a "Shelf"
func (svc *ShelfService) CreateShelf(obj *model.Shelf) (*model.Shelf, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditShelf - method modifies  the "Shelf" 
func (svc *ShelfService) EditShelf(id string, obj *model.Shelf) (*model.Shelf, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyShelf - private method for more than one element check. 
func onlyShelf(objs []*model.Shelf) (*model.Shelf, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Shelf found with the given filter")
	}

	return objs[0], nil
}

 
// GetShelfsByID - method returns associative a array of pointers of type "Shelf", filter by Id
func (svc *ShelfService) GetShelfsByID(pool *model.Pool, fields []string) (map[string]*model.Shelf, error) {
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
	objs, err := svc.GetShelfs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Shelf)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetShelfById - method returns a pointer to "Shelf"
func (svc *ShelfService) GetShelfById(id string) (*model.Shelf, error) {
	return svc.objectSet.GetObject(id)
}


