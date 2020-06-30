// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtectionTemplate Service - Manage protection templates. Protection templates are sets of snapshot schedules, replication schedules, and retention limits that can be used to prefill the protection
// information when creating new volume collections. A volume collection, once created, is not affected by edits to the protection template that was used to create it. All the
// volumes assigned to a volume collection use the same settings. You cannot edit or delete the predefined protection templates provided by storage array, but you can create custom
// protection templates as needed.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetProtectionTemplatesWithFields - method returns a array of pointers of type "ProtectionTemplate" 
func (svc *ProtectionTemplateService) GetProtectionTemplatesWithFields(fields []string) ([]*model.ProtectionTemplate, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateProtectionTemplate - method creates a "ProtectionTemplate"
func (svc *ProtectionTemplateService) CreateProtectionTemplate(obj *model.ProtectionTemplate) (*model.ProtectionTemplate, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditProtectionTemplate - method modifies  the "ProtectionTemplate" 
func (svc *ProtectionTemplateService) EditProtectionTemplate(id string, obj *model.ProtectionTemplate) (*model.ProtectionTemplate, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyProtectionTemplate - private method for more than one element check. 
func onlyProtectionTemplate(objs []*model.ProtectionTemplate) (*model.ProtectionTemplate, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ProtectionTemplate found with the given filter")
	}

	return objs[0], nil
}

 
// GetProtectionTemplatesByID - method returns associative a array of pointers of type "ProtectionTemplate", filter by Id
func (svc *ProtectionTemplateService) GetProtectionTemplatesByID(pool *model.Pool, fields []string) (map[string]*model.ProtectionTemplate, error) {
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
	objs, err := svc.GetProtectionTemplates(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ProtectionTemplate)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetProtectionTemplateById - method returns a pointer to "ProtectionTemplate"
func (svc *ProtectionTemplateService) GetProtectionTemplateById(id string) (*model.ProtectionTemplate, error) {
	return svc.objectSet.GetObject(id)
}

// GetProtectionTemplatesByName - method returns a associative array of pointers of type "ProtectionTemplate", filter by name 
func (svc *ProtectionTemplateService) GetProtectionTemplatesByName(pool *model.Pool, fields []string) (map[string]*model.ProtectionTemplate, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetProtectionTemplates(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ProtectionTemplate)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlyProtectionTemplate(objs)
}	

