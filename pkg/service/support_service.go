// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Support Service - View and alter support-based parameters.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// SupportService type 
type SupportService struct {
	objectSet *client.SupportObjectSet
}

// NewSupportService - method to initialize "SupportService" 
func NewSupportService(gs *NsGroupService) (*SupportService) {
	objectSet := gs.client.GetSupportObjectSet()
	return &SupportService{objectSet: objectSet}
}

// GetSupports - method returns a array of pointers of type "Supports"
func (svc *SupportService) GetSupports(params *util.GetParams) ([]*model.Support, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetSupportsWithFields - method returns a array of pointers of type "Support" 
func (svc *SupportService) GetSupportsWithFields(fields []string) ([]*model.Support, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSupport - method creates a "Support"
func (svc *SupportService) CreateSupport(obj *model.Support) (*model.Support, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSupport - method modifies  the "Support" 
func (svc *SupportService) EditSupport(id string, obj *model.Support) (*model.Support, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySupport - private method for more than one element check. 
func onlySupport(objs []*model.Support) (*model.Support, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Support found with the given filter")
	}

	return objs[0], nil
}

 
// GetSupportsByID - method returns associative a array of pointers of type "Support", filter by Id
func (svc *SupportService) GetSupportsByID(pool *model.Pool, fields []string) (map[string]*model.Support, error) {
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
	objs, err := svc.GetSupports(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Support)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetSupportById - method returns a pointer to "Support"
func (svc *SupportService) GetSupportById(id string) (*model.Support, error) {
	return svc.objectSet.GetObject(id)
}


