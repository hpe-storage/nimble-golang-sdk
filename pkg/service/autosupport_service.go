// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Autosupport Service - Get status of autosupport.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// AutosupportService type 
type AutosupportService struct {
	objectSet *client.AutosupportObjectSet
}

// NewAutosupportService - method to initialize "AutosupportService" 
func NewAutosupportService(gs *NsGroupService) (*AutosupportService) {
	objectSet := gs.client.GetAutosupportObjectSet()
	return &AutosupportService{objectSet: objectSet}
}

// GetAutosupports - method returns a array of pointers of type "Autosupports"
func (svc *AutosupportService) GetAutosupports(params *util.GetParams) ([]*model.Autosupport, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetAutosupportsWithFields - method returns a array of pointers of type "Autosupport" 
func (svc *AutosupportService) GetAutosupportsWithFields(fields []string) ([]*model.Autosupport, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateAutosupport - method creates a "Autosupport"
func (svc *AutosupportService) CreateAutosupport(obj *model.Autosupport) (*model.Autosupport, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditAutosupport - method modifies  the "Autosupport" 
func (svc *AutosupportService) EditAutosupport(id string, obj *model.Autosupport) (*model.Autosupport, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyAutosupport - private method for more than one element check. 
func onlyAutosupport(objs []*model.Autosupport) (*model.Autosupport, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Autosupport found with the given filter")
	}

	return objs[0], nil
}

 
// GetAutosupportsByID - method returns associative a array of pointers of type "Autosupport", filter by Id
func (svc *AutosupportService) GetAutosupportsByID(pool *model.Pool, fields []string) (map[string]*model.Autosupport, error) {
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
	objs, err := svc.GetAutosupports(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Autosupport)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetAutosupportById - method returns a pointer to "Autosupport"
func (svc *AutosupportService) GetAutosupportById(id string) (*model.Autosupport, error) {
	return svc.objectSet.GetObject(id)
}


