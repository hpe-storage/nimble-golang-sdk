// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Version Service - Show the API version.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// VersionService type 
type VersionService struct {
	objectSet *client.VersionObjectSet
}

// NewVersionService - method to initialize "VersionService" 
func NewVersionService(gs *NsGroupService) (*VersionService) {
	objectSet := gs.client.GetVersionObjectSet()
	return &VersionService{objectSet: objectSet}
}

// GetVersions - method returns a array of pointers of type "Versions"
func (svc *VersionService) GetVersions(params *util.GetParams) ([]*model.Version, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateVersion - method creates a "Version"
func (svc *VersionService) CreateVersion(obj *model.Version) (*model.Version, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateVersion - method modifies  the "Version" 
func (svc *VersionService) UpdateVersion(id string, obj *model.Version) (*model.Version, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetVersionByName - method returns a pointer "Version" 
func (svc *VersionService) GetVersionByName(name string) (*model.Version, error) {
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

// DeleteVersion - deletes the "Version"
func (svc *VersionService) DeleteVersion(id string) error {
	return svc.objectSet.DeleteObject(id)
}
