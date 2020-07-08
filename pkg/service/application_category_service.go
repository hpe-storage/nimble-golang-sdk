// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ApplicationCategory Service - Provides the list of application categories that are available, to classify volumes depending on the applications that use them.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ApplicationCategoryService type
type ApplicationCategoryService struct {
	objectSet *client.ApplicationCategoryObjectSet
}

// NewApplicationCategoryService - method to initialize "ApplicationCategoryService"
func NewApplicationCategoryService(gs *NsGroupService) *ApplicationCategoryService {
	objectSet := gs.client.GetApplicationCategoryObjectSet()
	return &ApplicationCategoryService{objectSet: objectSet}
}

// GetApplicationCategories - method returns a array of pointers of type "ApplicationCategories"
func (svc *ApplicationCategoryService) GetApplicationCategories(params *util.GetParams) ([]*model.ApplicationCategory, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	applicationCategoryResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// CreateApplicationCategory - method creates a "ApplicationCategory"
func (svc *ApplicationCategoryService) CreateApplicationCategory(obj *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	applicationCategoryResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// UpdateApplicationCategory - method modifies  the "ApplicationCategory"
func (svc *ApplicationCategoryService) UpdateApplicationCategory(id string, obj *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	applicationCategoryResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// GetApplicationCategoryById - method returns a pointer to "ApplicationCategory"
func (svc *ApplicationCategoryService) GetApplicationCategoryById(id string) (*model.ApplicationCategory, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	applicationCategoryResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// GetApplicationCategoryByName - method returns a pointer "ApplicationCategory"
func (svc *ApplicationCategoryService) GetApplicationCategoryByName(name string) (*model.ApplicationCategory, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	applicationCategoryResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(applicationCategoryResp) == 0 {
		return nil, nil
	}

	return applicationCategoryResp[0], nil
}

// DeleteApplicationCategory - deletes the "ApplicationCategory"
func (svc *ApplicationCategoryService) DeleteApplicationCategory(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
