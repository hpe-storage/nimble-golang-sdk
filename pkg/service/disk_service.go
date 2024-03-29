// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// Disk Service - Disks are used for storing user data.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// DiskService type
type DiskService struct {
	objectSet *client.DiskObjectSet
}

// NewDiskService - method to initialize "DiskService"
func NewDiskService(gs *NsGroupService) *DiskService {
	objectSet := gs.client.GetDiskObjectSet()
	return &DiskService{objectSet: objectSet}
}

// GetDisks - method returns a array of pointers of type "Disks"
func (svc *DiskService) GetDisks(params *param.GetParams) ([]*nimbleos.Disk, error) {
	diskResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return diskResp, nil
}

// CreateDisk - method creates a "Disk"
func (svc *DiskService) CreateDisk(obj *nimbleos.Disk) (*nimbleos.Disk, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateDisk: invalid parameter specified, %v", obj)
	}

	diskResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return diskResp, nil
}

// UpdateDisk - method modifies  the "Disk"
func (svc *DiskService) UpdateDisk(id string, obj *nimbleos.Disk) (*nimbleos.Disk, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateDisk: invalid parameter specified, %v", obj)
	}

	diskResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return diskResp, nil
}

// GetDiskById - method returns a pointer to "Disk"
func (svc *DiskService) GetDiskById(id string) (*nimbleos.Disk, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetDiskById: invalid parameter specified, %v", id)
	}

	diskResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return diskResp, nil
}

// DeleteDisk - deletes the "Disk"
func (svc *DiskService) DeleteDisk(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteDisk: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
