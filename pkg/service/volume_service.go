// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Volume Service - Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// VolumeService type 
type VolumeService struct {
	objectSet *client.VolumeObjectSet
}

// NewVolumeService - method to initialize "VolumeService" 
func NewVolumeService(gs *NsGroupService) (*VolumeService) {
	objectSet := gs.client.GetVolumeObjectSet()
	return &VolumeService{objectSet: objectSet}
}

// GetVolumes - method returns a array of pointers of type "Volumes"
func (svc *VolumeService) GetVolumes(params *util.GetParams) ([]*model.Volume, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateVolume - method creates a "Volume"
func (svc *VolumeService) CreateVolume(obj *model.Volume) (*model.Volume, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateVolume - method modifies  the "Volume" 
func (svc *VolumeService) UpdateVolume(id string, obj *model.Volume) (*model.Volume, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetVolumeById - method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeById(id string) (*model.Volume, error) {
	return svc.objectSet.GetObject(id)
}

// GetVolumeByName - method returns a pointer "Volume" 
func (svc *VolumeService) GetVolumeByName(name string) (*model.Volume, error) {
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
// GetVolumeBySerialNumber method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeBySerialNumber(serialNumber string) (*model.Volume, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.SerialNumber,
			Operator:  util.EQUALS.String(),
			Value:     serialNumber,
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

//OnlineVolume - method makes the volume online 
func (svc *VolumeService) OnlineVolume(id string, force bool) (*model.Volume, error) {
	return svc.UpdateVolume(id, &model.Volume{
		Online: util.NewBool(true),
		Force:  util.NewBool(force),
	})
}

// OfflineVolume - makes the volume offline 
func (svc *VolumeService) OfflineVolume(id string, force bool) (*model.Volume, error) {
	return svc.UpdateVolume(id, &model.Volume{
		Online: util.NewBool(false),
		Force:  util.NewBool(force),
	})
}

// DeleteVolume - deletes the volume
func (svc *VolumeService) DeleteVolume(id string) error {
	_, err := svc.OfflineVolume(id, false)
	if err != nil {
		return err
	}

	return svc.objectSet.DeleteObject(id)
}

