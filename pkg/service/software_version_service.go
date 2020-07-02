// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SoftwareVersion Service - Show the software version.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// SoftwareVersionService type 
type SoftwareVersionService struct {
	objectSet *client.SoftwareVersionObjectSet
}

// NewSoftwareVersionService - method to initialize "SoftwareVersionService" 
func NewSoftwareVersionService(gs *NsGroupService) (*SoftwareVersionService) {
	objectSet := gs.client.GetSoftwareVersionObjectSet()
	return &SoftwareVersionService{objectSet: objectSet}
}

// GetSoftwareVersions - method returns a array of pointers of type "SoftwareVersions"
func (svc *SoftwareVersionService) GetSoftwareVersions(params *util.GetParams) ([]*model.SoftwareVersion, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSoftwareVersion - method creates a "SoftwareVersion"
func (svc *SoftwareVersionService) CreateSoftwareVersion(obj *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateSoftwareVersion - method modifies  the "SoftwareVersion" 
func (svc *SoftwareVersionService) UpdateSoftwareVersion(id string, obj *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetSoftwareVersionByName - method returns a pointer "SoftwareVersion" 
func (svc *SoftwareVersionService) GetSoftwareVersionByName(name string) (*model.SoftwareVersion, error) {
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

// DeleteSoftwareVersion - deletes the "SoftwareVersion"
func (svc *SoftwareVersionService) DeleteSoftwareVersion(id string) error {
	return svc.objectSet.DeleteObject(id)
}
