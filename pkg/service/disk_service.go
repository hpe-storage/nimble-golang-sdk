// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Disk Service - Disks are used for storing user data.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (svc *DiskService) GetDisks(params *util.GetParams) ([]*model.Disk, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	diskResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return diskResp, nil
}

// CreateDisk - method creates a "Disk"
func (svc *DiskService) CreateDisk(obj *model.Disk) (*model.Disk, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	diskResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return diskResp, nil
}

// UpdateDisk - method modifies  the "Disk"
func (svc *DiskService) UpdateDisk(id string, obj *model.Disk) (*model.Disk, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	diskResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return diskResp, nil
}

// GetDiskById - method returns a pointer to "Disk"
func (svc *DiskService) GetDiskById(id string) (*model.Disk, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
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
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
