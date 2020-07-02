// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtectionTemplate Service - Manage protection templates. Protection templates are sets of snapshot schedules, replication schedules, and retention limits that can be used to prefill the protection
// information when creating new volume collections. A volume collection, once created, is not affected by edits to the protection template that was used to create it. All the
// volumes assigned to a volume collection use the same settings. You cannot edit or delete the predefined protection templates provided by storage array, but you can create custom
// protection templates as needed.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ProtectionTemplateService type 
type ProtectionTemplateService struct {
	objectSet *client.ProtectionTemplateObjectSet
}

// NewProtectionTemplateService - method to initialize "ProtectionTemplateService" 
func NewProtectionTemplateService(gs *NsGroupService) (*ProtectionTemplateService) {
	objectSet := gs.client.GetProtectionTemplateObjectSet()
	return &ProtectionTemplateService{objectSet: objectSet}
}

// GetProtectionTemplates - method returns a array of pointers of type "ProtectionTemplates"
func (svc *ProtectionTemplateService) GetProtectionTemplates(params *util.GetParams) ([]*model.ProtectionTemplate, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateProtectionTemplate - method creates a "ProtectionTemplate"
func (svc *ProtectionTemplateService) CreateProtectionTemplate(obj *model.ProtectionTemplate) (*model.ProtectionTemplate, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateProtectionTemplate - method modifies  the "ProtectionTemplate" 
func (svc *ProtectionTemplateService) UpdateProtectionTemplate(id string, obj *model.ProtectionTemplate) (*model.ProtectionTemplate, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetProtectionTemplateById - method returns a pointer to "ProtectionTemplate"
func (svc *ProtectionTemplateService) GetProtectionTemplateById(id string) (*model.ProtectionTemplate, error) {
	return svc.objectSet.GetObject(id)
}

// GetProtectionTemplateByName - method returns a pointer "ProtectionTemplate" 
func (svc *ProtectionTemplateService) GetProtectionTemplateByName(name string) (*model.ProtectionTemplate, error) {
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

// DeleteProtectionTemplate - deletes the "ProtectionTemplate"
func (svc *ProtectionTemplateService) DeleteProtectionTemplate(id string) error {
	return svc.objectSet.DeleteObject(id)
}
