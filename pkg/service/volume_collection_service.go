// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// VolumeCollection Service - Manage volume collections. Volume collections are logical groups of volumes that share protection characteristics such as snapshot and replication schedules. Volume collections
// can be created from scratch or based on predefined protection templates.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// VolumeCollectionService type 
type VolumeCollectionService struct {
	objectSet *client.VolumeCollectionObjectSet
}

// NewVolumeCollectionService - method to initialize "VolumeCollectionService" 
func NewVolumeCollectionService(gs *NsGroupService) (*VolumeCollectionService) {
	objectSet := gs.client.GetVolumeCollectionObjectSet()
	return &VolumeCollectionService{objectSet: objectSet}
}

// GetVolumeCollections - method returns a array of pointers of type "VolumeCollections"
func (svc *VolumeCollectionService) GetVolumeCollections(params *util.GetParams) ([]*model.VolumeCollection, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateVolumeCollection - method creates a "VolumeCollection"
func (svc *VolumeCollectionService) CreateVolumeCollection(obj *model.VolumeCollection) (*model.VolumeCollection, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateVolumeCollection - method modifies  the "VolumeCollection" 
func (svc *VolumeCollectionService) UpdateVolumeCollection(id string, obj *model.VolumeCollection) (*model.VolumeCollection, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetVolumeCollectionById - method returns a pointer to "VolumeCollection"
func (svc *VolumeCollectionService) GetVolumeCollectionById(id string) (*model.VolumeCollection, error) {
	return svc.objectSet.GetObject(id)
}

// GetVolumeCollectionByName - method returns a pointer "VolumeCollection" 
func (svc *VolumeCollectionService) GetVolumeCollectionByName(name string) (*model.VolumeCollection, error) {
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

// DeleteVolumeCollection - deletes the "VolumeCollection"
func (svc *VolumeCollectionService) DeleteVolumeCollection(id string) error {
	return svc.objectSet.DeleteObject(id)
}
