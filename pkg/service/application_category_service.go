// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ApplicationCategory Service - Provides the list of application categories that are available, to classify volumes depending on the applications that use them.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// GetApplicationCategories - method returns a array of pointers of type "ApplicationCategories"
func (svc *ApplicationCategoryService) GetApplicationCategories(params *util.GetParams) ([]*model.ApplicationCategory, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}


// CreateApplicationCategory - method creates a "ApplicationCategory"
func (svc *ApplicationCategoryService) CreateApplicationCategory(obj *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateApplicationCategory - method modifies  the "ApplicationCategory" 
func (svc *ApplicationCategoryService) UpdateApplicationCategory(id string, obj *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetApplicationCategoryById - method returns a pointer to "ApplicationCategory"
func (svc *ApplicationCategoryService) GetApplicationCategoryById(id string) (*model.ApplicationCategory, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteApplicationCategory - deletes the "ApplicationCategory"
func (svc *ApplicationCategoryService) DeleteApplicationCategory(id string) error {
	return svc.objectSet.DeleteObject(id)
}
