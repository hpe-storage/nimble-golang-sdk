// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ApplicationCategory Service - Provides the list of application categories that are available, to classify volumes depending on the applications that use them.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// ApplicationCategoryService type 
type ApplicationCategoryService struct {
	objectSet *client.ApplicationCategoryObjectSet
}

// NewApplicationCategoryService - method to initialize "ApplicationCategoryService" 
func NewApplicationCategoryService(gs *NsGroupService) (*ApplicationCategoryService) {
	objectSet := gs.client.GetApplicationCategoryObjectSet()
	return &ApplicationCategoryService{objectSet: objectSet}
}

// GetApplicationCategorys - method returns a array of pointers of type "ApplicationCategorys"
func (svc *ApplicationCategoryService) GetApplicationCategorys(params *util.GetParams) ([]*model.ApplicationCategory, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetApplicationCategorysWithFields - method returns a array of pointers of type "ApplicationCategory" 
func (svc *ApplicationCategoryService) GetApplicationCategorysWithFields(fields []string) ([]*model.ApplicationCategory, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateApplicationCategory - method creates a "ApplicationCategory"
func (svc *ApplicationCategoryService) CreateApplicationCategory(obj *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditApplicationCategory - method modifies  the "ApplicationCategory" 
func (svc *ApplicationCategoryService) EditApplicationCategory(id string, obj *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyApplicationCategory - private method for more than one element check. 
func onlyApplicationCategory(objs []*model.ApplicationCategory) (*model.ApplicationCategory, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ApplicationCategory found with the given filter")
	}

	return objs[0], nil
}

 
// GetApplicationCategorysByID - method returns associative a array of pointers of type "ApplicationCategory", filter by Id
func (svc *ApplicationCategoryService) GetApplicationCategorysByID(pool *model.Pool, fields []string) (map[string]*model.ApplicationCategory, error) {
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
	objs, err := svc.GetApplicationCategorys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ApplicationCategory)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetApplicationCategoryById - method returns a pointer to "ApplicationCategory"
func (svc *ApplicationCategoryService) GetApplicationCategoryById(id string) (*model.ApplicationCategory, error) {
	return svc.objectSet.GetObject(id)
}

// GetApplicationCategorysByName - method returns a associative array of pointers of type "ApplicationCategory", filter by name 
func (svc *ApplicationCategoryService) GetApplicationCategorysByName(pool *model.Pool, fields []string) (map[string]*model.ApplicationCategory, error) {
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
	objs, err := svc.GetApplicationCategorys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ApplicationCategory)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetApplicationCategoryByName - method returns a pointer "ApplicationCategory" 
func (svc *ApplicationCategoryService) GetApplicationCategoryByName(name string) (*model.ApplicationCategory, error) {
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
	return onlyApplicationCategory(objs)
}	

