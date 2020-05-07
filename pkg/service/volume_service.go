// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"fmt"

	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

type VolumeService struct {
	objectSet *client.VolumeObjectSet
}

func NewVolumeService(gs *GroupService) (vs *VolumeService) {
	objectSet := gs.client.GetVolumeObjectSet()
	return &VolumeService{objectSet: objectSet}
}

func (vs *VolumeService) GetVolumes(params *util.GetParams) ([]*model.Volume, error) {
	return vs.objectSet.GetObjectListFromParams(params)
}

func (vs *VolumeService) GetVolumesWithFields(fields []string) ([]*model.Volume, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return vs.objectSet.GetObjectListFromParams(params)
}

func (vs *VolumeService) GetVolumesByName(pool *model.Pool, fields []string) (map[string]*model.Volume, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, pool.ID, false)
		params.WithSearchFilter(filter)
	}
	volumes, err := vs.GetVolumes(params)
	if err != nil {
		return nil, err
	}
	volumeMap := make(map[string]*model.Volume)
	for _, volume := range volumes {
		volumeMap[volume.Name] = volume
	}
	return volumeMap, nil
}

func (vs *VolumeService) GetVolumesByID(pool *model.Pool, fields []string) (map[string]*model.Volume, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, pool.ID, false)
		params.WithSearchFilter(filter)
	}
	volumes, err := vs.GetVolumes(params)
	if err != nil {
		return nil, err
	}
	volumeMap := make(map[string]*model.Volume)
	for _, volume := range volumes {
		volumeMap[volume.ID] = volume
	}
	return volumeMap, nil
}

func (vs *VolumeService) GetVolumeById(id string) (*model.Volume, error) {
	return vs.objectSet.GetObject(id)
}

func (vs *VolumeService) GetVolumeByName(name string) (*model.Volume, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	volumes, err := vs.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return onlyVolume(volumes)
}

func (vs *VolumeService) GetVolumeBySerialNumber(serialNumber string) (*model.Volume, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.SerialNumber,
			Operator:  util.EQUALS.String(),
			Value:     serialNumber,
		},
	}
	volumes, err := vs.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return onlyVolume(volumes)
}

func (vs *VolumeService) CreateVolume(volume *model.Volume) (*model.Volume, error) {
	// TODO: validate parameters
	return vs.objectSet.CreateObject(volume)
}

func (vs *VolumeService) EditVolume(id string, volume *model.Volume) (*model.Volume, error) {
	return vs.objectSet.UpdateObject(id, volume)
}

func (vs *VolumeService) OnlineVolume(id string, force bool) (*model.Volume, error) {
	return vs.EditVolume(id, &model.Volume{
		Online: util.NewBool(true),
		Force:  util.NewBool(force),
	})
}

func (vs *VolumeService) OfflineVolume(id string, force bool) (*model.Volume, error) {
	return vs.EditVolume(id, &model.Volume{
		Online: util.NewBool(false),
		Force:  util.NewBool(force),
	})
}

func (vs *VolumeService) DestroyVolume(id string) error {
	_, err := vs.OfflineVolume(id, false)
	if err != nil {
		return err
	}

	return vs.objectSet.DeleteObject(id)
}

// "private" functions
func onlyVolume(volumes []*model.Volume) (*model.Volume, error) {
	if len(volumes) == 0 {
		return nil, nil
	}

	if len(volumes) > 1 {
		return nil, fmt.Errorf("More than one volume found with the given filter")
	}

	return volumes[0], nil
}
