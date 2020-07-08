// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
)

// Provides the list of application categories that are available, to classify volumes depending on the applications that use them.
const (
	applicationCategoryPath = "application_categories"
)

// ApplicationCategoryObjectSet
type ApplicationCategoryObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new ApplicationCategory object
func (objectSet *ApplicationCategoryObjectSet) CreateObject(payload *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on ApplicationCategory")
}

// UpdateObject Modify existing ApplicationCategory object
func (objectSet *ApplicationCategoryObjectSet) UpdateObject(id string, payload *model.ApplicationCategory) (*model.ApplicationCategory, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on ApplicationCategory")
}

// DeleteObject deletes the ApplicationCategory object with the specified ID
func (objectSet *ApplicationCategoryObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on ApplicationCategory")
}

// GetObject returns a ApplicationCategory object with the given ID
func (objectSet *ApplicationCategoryObjectSet) GetObject(id string) (*model.ApplicationCategory, error) {
	applicationCategoryObjectSetResp, err := objectSet.Client.Get(applicationCategoryPath, id, model.ApplicationCategory{})
	if err != nil {
		return nil, err
	}

	// null check
	if applicationCategoryObjectSetResp == nil {
		return nil, nil
	}
	return applicationCategoryObjectSetResp.(*model.ApplicationCategory), err
}

// GetObjectList returns the list of ApplicationCategory objects
func (objectSet *ApplicationCategoryObjectSet) GetObjectList() ([]*model.ApplicationCategory, error) {
	applicationCategoryObjectSetResp, err := objectSet.Client.List(applicationCategoryPath)
	if err != nil {
		return nil, err
	}
	return buildApplicationCategoryObjectSet(applicationCategoryObjectSetResp), err
}

// GetObjectListFromParams returns the list of ApplicationCategory objects using the given params query info
func (objectSet *ApplicationCategoryObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ApplicationCategory, error) {
	applicationCategoryObjectSetResp, err := objectSet.Client.ListFromParams(applicationCategoryPath, params)
	if err != nil {
		return nil, err
	}
	return buildApplicationCategoryObjectSet(applicationCategoryObjectSetResp), err
}

// generated function to build the appropriate response types
func buildApplicationCategoryObjectSet(response interface{}) []*model.ApplicationCategory {
	values := reflect.ValueOf(response)
	results := make([]*model.ApplicationCategory, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ApplicationCategory{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
