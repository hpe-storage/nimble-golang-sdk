// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SoftwareVersion Service - Show the software version.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// SoftwareVersionService type
type SoftwareVersionService struct {
	objectSet *client.SoftwareVersionObjectSet
}

// NewSoftwareVersionService - method to initialize "SoftwareVersionService"
func NewSoftwareVersionService(gs *NsGroupService) *SoftwareVersionService {
	objectSet := gs.client.GetSoftwareVersionObjectSet()
	return &SoftwareVersionService{objectSet: objectSet}
}

// GetSoftwareVersions - method returns a array of pointers of type "SoftwareVersions"
func (svc *SoftwareVersionService) GetSoftwareVersions(params *util.GetParams) ([]*model.SoftwareVersion, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	softwareVersionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return softwareVersionResp, nil
}

// CreateSoftwareVersion - method creates a "SoftwareVersion"
func (svc *SoftwareVersionService) CreateSoftwareVersion(obj *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	softwareVersionResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return softwareVersionResp, nil
}

// UpdateSoftwareVersion - method modifies  the "SoftwareVersion"
func (svc *SoftwareVersionService) UpdateSoftwareVersion(id string, obj *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	softwareVersionResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return softwareVersionResp, nil
}

// GetSoftwareVersionByName - method returns a pointer "SoftwareVersion"
func (svc *SoftwareVersionService) GetSoftwareVersionByName(name string) (*model.SoftwareVersion, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	softwareVersionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(softwareVersionResp) == 0 {
		return nil, nil
	}

	return softwareVersionResp[0], nil
}

// DeleteSoftwareVersion - deletes the "SoftwareVersion"
func (svc *SoftwareVersionService) DeleteSoftwareVersion(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
