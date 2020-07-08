// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Version Service - Show the API version.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// VersionService type
type VersionService struct {
	objectSet *client.VersionObjectSet
}

// NewVersionService - method to initialize "VersionService"
func NewVersionService(gs *NsGroupService) *VersionService {
	objectSet := gs.client.GetVersionObjectSet()
	return &VersionService{objectSet: objectSet}
}

// GetVersions - method returns a array of pointers of type "Versions"
func (svc *VersionService) GetVersions(params *util.GetParams) ([]*model.Version, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	versionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return versionResp, nil
}

// CreateVersion - method creates a "Version"
func (svc *VersionService) CreateVersion(obj *model.Version) (*model.Version, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	versionResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return versionResp, nil
}

// UpdateVersion - method modifies  the "Version"
func (svc *VersionService) UpdateVersion(id string, obj *model.Version) (*model.Version, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	versionResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return versionResp, nil
}

// GetVersionByName - method returns a pointer "Version"
func (svc *VersionService) GetVersionByName(name string) (*model.Version, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	versionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(versionResp) == 0 {
		return nil, nil
	}

	return versionResp[0], nil
}

// DeleteVersion - deletes the "Version"
func (svc *VersionService) DeleteVersion(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
