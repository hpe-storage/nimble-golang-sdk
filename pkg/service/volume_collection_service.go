// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// VolumeCollection Service - Manage volume collections. Volume collections are logical groups of volumes that share protection characteristics such as snapshot and replication schedules. Volume collections
// can be created from scratch or based on predefined protection templates.

import (
	"fmt"
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
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	volumeCollectionResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return volumeCollectionResp,nil
}

// CreateVolumeCollection - method creates a "VolumeCollection"
func (svc *VolumeCollectionService) CreateVolumeCollection(obj *model.VolumeCollection) (*model.VolumeCollection, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	volumeCollectionResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return volumeCollectionResp,nil
}

// UpdateVolumeCollection - method modifies  the "VolumeCollection" 
func (svc *VolumeCollectionService) UpdateVolumeCollection(id string, obj *model.VolumeCollection) (*model.VolumeCollection, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	volumeCollectionResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return volumeCollectionResp,nil
}

// GetVolumeCollectionById - method returns a pointer to "VolumeCollection"
func (svc *VolumeCollectionService) GetVolumeCollectionById(id string) (*model.VolumeCollection, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	volumeCollectionResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return volumeCollectionResp,nil
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
	volumeCollectionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(volumeCollectionResp) == 0 {
    	return nil, nil
    }
    
	return volumeCollectionResp[0],nil
}	

// DeleteVolumeCollection - deletes the "VolumeCollection"
func (svc *VolumeCollectionService) DeleteVolumeCollection(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
