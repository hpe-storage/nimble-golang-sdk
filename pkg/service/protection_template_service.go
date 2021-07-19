// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// ProtectionTemplate Service - Manage protection templates. Protection templates are sets of snapshot schedules, replication schedules, and retention limits that can be used to prefill the protection
// information when creating new volume collections. A volume collection, once created, is not affected by edits to the protection template that was used to create it. All the
// volumes assigned to a volume collection use the same settings. You cannot edit or delete the predefined protection templates provided by storage array, but you can create custom
// protection templates as needed.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ProtectionTemplateService type
type ProtectionTemplateService struct {
	objectSet *client.ProtectionTemplateObjectSet
}

// NewProtectionTemplateService - method to initialize "ProtectionTemplateService"
func NewProtectionTemplateService(gs *NsGroupService) *ProtectionTemplateService {
	objectSet := gs.client.GetProtectionTemplateObjectSet()
	return &ProtectionTemplateService{objectSet: objectSet}
}

// GetProtectionTemplates - method returns a array of pointers of type "ProtectionTemplates"
func (svc *ProtectionTemplateService) GetProtectionTemplates(params *param.GetParams) ([]*nimbleos.ProtectionTemplate, error) {
	protectionTemplateResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return protectionTemplateResp, nil
}

// CreateProtectionTemplate - method creates a "ProtectionTemplate"
func (svc *ProtectionTemplateService) CreateProtectionTemplate(obj *nimbleos.ProtectionTemplate) (*nimbleos.ProtectionTemplate, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateProtectionTemplate: invalid parameter specified, %v", obj)
	}

	protectionTemplateResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return protectionTemplateResp, nil
}

// UpdateProtectionTemplate - method modifies  the "ProtectionTemplate"
func (svc *ProtectionTemplateService) UpdateProtectionTemplate(id string, obj *nimbleos.ProtectionTemplate) (*nimbleos.ProtectionTemplate, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateProtectionTemplate: invalid parameter specified, %v", obj)
	}

	protectionTemplateResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return protectionTemplateResp, nil
}

// GetProtectionTemplateById - method returns a pointer to "ProtectionTemplate"
func (svc *ProtectionTemplateService) GetProtectionTemplateById(id string) (*nimbleos.ProtectionTemplate, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetProtectionTemplateById: invalid parameter specified, %v", id)
	}

	protectionTemplateResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return protectionTemplateResp, nil
}

// GetProtectionTemplateByName - method returns a pointer "ProtectionTemplate"
func (svc *ProtectionTemplateService) GetProtectionTemplateByName(name string) (*nimbleos.ProtectionTemplate, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	protectionTemplateResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(protectionTemplateResp) == 0 {
		return nil, nil
	}

	return protectionTemplateResp[0], nil
}

// DeleteProtectionTemplate - deletes the "ProtectionTemplate"
func (svc *ProtectionTemplateService) DeleteProtectionTemplate(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteProtectionTemplate: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
