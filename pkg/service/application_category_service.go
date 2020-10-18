// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ApplicationCategory Service - Provides the list of application categories that are available, to classify volumes depending on the applications that use them.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *ApplicationCategoryService) GetApplicationCategories(params *param.GetParams) ([]*nimbleos.ApplicationCategory, error) {
	applicationCategoryResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// CreateApplicationCategory - method creates a "ApplicationCategory"
func (svc *ApplicationCategoryService) CreateApplicationCategory(obj *nimbleos.ApplicationCategory) (*nimbleos.ApplicationCategory, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateApplicationCategory: invalid parameter specified, %v", obj)
	}

	applicationCategoryResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// UpdateApplicationCategory - method modifies  the "ApplicationCategory"
func (svc *ApplicationCategoryService) UpdateApplicationCategory(id string, obj *nimbleos.ApplicationCategory) (*nimbleos.ApplicationCategory, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateApplicationCategory: invalid parameter specified, %v", obj)
	}

	applicationCategoryResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// GetApplicationCategoryById - method returns a pointer to "ApplicationCategory"
func (svc *ApplicationCategoryService) GetApplicationCategoryById(id string) (*nimbleos.ApplicationCategory, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetApplicationCategoryById: invalid parameter specified, %v", id)
	}

	applicationCategoryResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return applicationCategoryResp, nil
}

// GetApplicationCategoryByName - method returns a pointer "ApplicationCategory"
func (svc *ApplicationCategoryService) GetApplicationCategoryByName(name string) (*nimbleos.ApplicationCategory, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
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
		return fmt.Errorf("DeleteApplicationCategory: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
