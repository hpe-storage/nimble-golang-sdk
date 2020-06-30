// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Volume Service - Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetVolumesWithFields - method returns a array of pointers of type "Volume" 
func (svc *VolumeService) GetVolumesWithFields(fields []string) ([]*model.Volume, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateVolume - method creates a "Volume"
func (svc *VolumeService) CreateVolume(obj *model.Volume) (*model.Volume, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditVolume - method modifies  the "Volume" 
func (svc *VolumeService) EditVolume(id string, obj *model.Volume) (*model.Volume, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyVolume - private method for more than one element check. 
func onlyVolume(objs []*model.Volume) (*model.Volume, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Volume found with the given filter")
	}

	return objs[0], nil
}

 
// GetVolumesByID - method returns associative a array of pointers of type "Volume", filter by Id
func (svc *VolumeService) GetVolumesByID(pool *model.Pool, fields []string) (map[string]*model.Volume, error) {
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
	objs, err := svc.GetVolumes(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Volume)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetVolumeById - method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeById(id string) (*model.Volume, error) {
	return svc.objectSet.GetObject(id)
}

// GetVolumesByName - method returns a associative array of pointers of type "Volume", filter by name 
func (svc *VolumeService) GetVolumesByName(pool *model.Pool, fields []string) (map[string]*model.Volume, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetVolumes(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Volume)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlyVolume(objs)
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
	return onlyVolume(objs)
}

//OnlineVolume - method makes the volume online 
func (svc *VolumeService) OnlineVolume(id string, force bool) (*model.Volume, error) {
	return svc.EditVolume(id, &model.Volume{
		Online: util.NewBool(true),
		Force:  util.NewBool(force),
	})
}

// OfflineVolume - makes the volume offline 
func (svc *VolumeService) OfflineVolume(id string, force bool) (*model.Volume, error) {
	return svc.EditVolume(id, &model.Volume{
		Online: util.NewBool(false),
		Force:  util.NewBool(force),
	})
}

// DestroyVolume - deletes the volume
func (svc *VolumeService) DestroyVolume(id string) error {
	_, err := svc.OfflineVolume(id, false)
	if err != nil {
		return err
	}

	return svc.objectSet.DeleteObject(id)
}

