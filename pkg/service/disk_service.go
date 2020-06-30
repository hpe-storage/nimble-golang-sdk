// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Disk Service - Disks are used for storing user data.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// DiskService type 
type DiskService struct {
	objectSet *client.DiskObjectSet
}

// NewDiskService - method to initialize "DiskService" 
func NewDiskService(gs *NsGroupService) (*DiskService) {
	objectSet := gs.client.GetDiskObjectSet()
	return &DiskService{objectSet: objectSet}
}

// GetDisks - method returns a array of pointers of type "Disks"
func (svc *DiskService) GetDisks(params *util.GetParams) ([]*model.Disk, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetDisksWithFields - method returns a array of pointers of type "Disk" 
func (svc *DiskService) GetDisksWithFields(fields []string) ([]*model.Disk, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateDisk - method creates a "Disk"
func (svc *DiskService) CreateDisk(obj *model.Disk) (*model.Disk, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditDisk - method modifies  the "Disk" 
func (svc *DiskService) EditDisk(id string, obj *model.Disk) (*model.Disk, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyDisk - private method for more than one element check. 
func onlyDisk(objs []*model.Disk) (*model.Disk, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Disk found with the given filter")
	}

	return objs[0], nil
}

 
// GetDisksByID - method returns associative a array of pointers of type "Disk", filter by Id
func (svc *DiskService) GetDisksByID(pool *model.Pool, fields []string) (map[string]*model.Disk, error) {
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
	objs, err := svc.GetDisks(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Disk)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetDiskById - method returns a pointer to "Disk"
func (svc *DiskService) GetDiskById(id string) (*model.Disk, error) {
	return svc.objectSet.GetObject(id)
}


