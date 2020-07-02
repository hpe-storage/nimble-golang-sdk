// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Disk Service - Disks are used for storing user data.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateDisk - method creates a "Disk"
func (svc *DiskService) CreateDisk(obj *model.Disk) (*model.Disk, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateDisk - method modifies  the "Disk" 
func (svc *DiskService) UpdateDisk(id string, obj *model.Disk) (*model.Disk, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetDiskById - method returns a pointer to "Disk"
func (svc *DiskService) GetDiskById(id string) (*model.Disk, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteDisk - deletes the "Disk"
func (svc *DiskService) DeleteDisk(id string) error {
	return svc.objectSet.DeleteObject(id)
}
