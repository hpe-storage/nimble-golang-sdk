// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// VolumeFamily Service - A volume family contains all the volumes, snapshots, and clones derived from and including a root volume.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// VolumeFamilyService type 
type VolumeFamilyService struct {
	objectSet *client.VolumeFamilyObjectSet
}

// NewVolumeFamilyService - method to initialize "VolumeFamilyService" 
func NewVolumeFamilyService(gs *NsGroupService) (*VolumeFamilyService) {
	objectSet := gs.client.GetVolumeFamilyObjectSet()
	return &VolumeFamilyService{objectSet: objectSet}
}

// GetVolumeFamilys - method returns a array of pointers of type "VolumeFamilys"
func (svc *VolumeFamilyService) GetVolumeFamilys(params *util.GetParams) ([]*model.VolumeFamily, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetVolumeFamilysWithFields - method returns a array of pointers of type "VolumeFamily" 
func (svc *VolumeFamilyService) GetVolumeFamilysWithFields(fields []string) ([]*model.VolumeFamily, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateVolumeFamily - method creates a "VolumeFamily"
func (svc *VolumeFamilyService) CreateVolumeFamily(obj *model.VolumeFamily) (*model.VolumeFamily, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditVolumeFamily - method modifies  the "VolumeFamily" 
func (svc *VolumeFamilyService) EditVolumeFamily(id string, obj *model.VolumeFamily) (*model.VolumeFamily, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyVolumeFamily - private method for more than one element check. 
func onlyVolumeFamily(objs []*model.VolumeFamily) (*model.VolumeFamily, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one VolumeFamily found with the given filter")
	}

	return objs[0], nil
}

 
// GetVolumeFamilysByID - method returns associative a array of pointers of type "VolumeFamily", filter by Id
func (svc *VolumeFamilyService) GetVolumeFamilysByID(pool *model.Pool, fields []string) (map[string]*model.VolumeFamily, error) {
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
	objs, err := svc.GetVolumeFamilys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.VolumeFamily)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetVolumeFamilyById - method returns a pointer to "VolumeFamily"
func (svc *VolumeFamilyService) GetVolumeFamilyById(id string) (*model.VolumeFamily, error) {
	return svc.objectSet.GetObject(id)
}


