// Copyright 2021 Hewlett Packard Enterprise Development LP

package service

// Version Service - Show the API version.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *VersionService) GetVersions(params *param.GetParams) ([]*nimbleos.Version, error) {
	versionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return versionResp, nil
}

// CreateVersion - method creates a "Version"
func (svc *VersionService) CreateVersion(obj *nimbleos.Version) (*nimbleos.Version, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateVersion: invalid parameter specified, %v", obj)
	}

	versionResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return versionResp, nil
}

// UpdateVersion - method modifies  the "Version"
func (svc *VersionService) UpdateVersion(id string, obj *nimbleos.Version) (*nimbleos.Version, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateVersion: invalid parameter specified, %v", obj)
	}

	versionResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return versionResp, nil
}

// GetVersionByName - method returns a pointer "Version"
func (svc *VersionService) GetVersionByName(name string) (*nimbleos.Version, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
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
		return fmt.Errorf("DeleteVersion: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
