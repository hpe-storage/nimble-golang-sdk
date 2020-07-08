// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Volume Service - Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// VolumeService type
type VolumeService struct {
	objectSet *client.VolumeObjectSet
}

// NewVolumeService - method to initialize "VolumeService"
func NewVolumeService(gs *NsGroupService) *VolumeService {
	objectSet := gs.client.GetVolumeObjectSet()
	return &VolumeService{objectSet: objectSet}
}

// GetVolumes - method returns a array of pointers of type "Volumes"
func (svc *VolumeService) GetVolumes(params *util.GetParams) ([]*model.Volume, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	volumeResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// CreateVolume - method creates a "Volume"
func (svc *VolumeService) CreateVolume(obj *model.Volume) (*model.Volume, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	volumeResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// UpdateVolume - method modifies  the "Volume"
func (svc *VolumeService) UpdateVolume(id string, obj *model.Volume) (*model.Volume, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	volumeResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// GetVolumeById - method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeById(id string) (*model.Volume, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	volumeResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// GetVolumeByName - method returns a pointer "Volume"
func (svc *VolumeService) GetVolumeByName(name string) (*model.Volume, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	volumeResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(volumeResp) == 0 {
		return nil, nil
	}

	return volumeResp[0], nil
}

// GetVolumeBySerialNumber method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeBySerialNumber(serialNumber string) (*model.Volume, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.SerialNumber,
			Operator:  util.EQUALS.String(),
			Value:     serialNumber,
		},
	}
	volumeResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	if len(volumeResp) == 0 {
		return nil, nil
	}

	return volumeResp[0], nil
}

//OnlineVolume - method makes the volume online
func (svc *VolumeService) OnlineVolume(id string, force bool) (*model.Volume, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %s", id)
	}

	volumeResp, err := svc.UpdateVolume(id, &model.Volume{
		Online: util.NewBool(true),
		Force:  util.NewBool(force),
	})
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// OfflineVolume - makes the volume offline
func (svc *VolumeService) OfflineVolume(id string, force bool) (*model.Volume, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %s", id)
	}

	volumeResp, err := svc.UpdateVolume(id, &model.Volume{
		Online: util.NewBool(false),
		Force:  util.NewBool(force),
	})
	if err != nil {
		return nil, err
	}
	return volumeResp, nil

}

// DeleteVolume - deletes the volume
func (svc *VolumeService) DeleteVolume(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	_, err := svc.OfflineVolume(id, false)
	if err != nil {
		return err
	}

	err = svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
