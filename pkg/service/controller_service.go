// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Controller Service - Controller is a redundant collection of hardware capable of running the array software.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// ControllerService type 
type ControllerService struct {
	objectSet *client.ControllerObjectSet
}

// NewControllerService - method to initialize "ControllerService" 
func NewControllerService(gs *NsGroupService) (*ControllerService) {
	objectSet := gs.client.GetControllerObjectSet()
	return &ControllerService{objectSet: objectSet}
}

// GetControllers - method returns a array of pointers of type "Controllers"
func (svc *ControllerService) GetControllers(params *util.GetParams) ([]*model.Controller, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetControllersWithFields - method returns a array of pointers of type "Controller" 
func (svc *ControllerService) GetControllersWithFields(fields []string) ([]*model.Controller, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateController - method creates a "Controller"
func (svc *ControllerService) CreateController(obj *model.Controller) (*model.Controller, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditController - method modifies  the "Controller" 
func (svc *ControllerService) EditController(id string, obj *model.Controller) (*model.Controller, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyController - private method for more than one element check. 
func onlyController(objs []*model.Controller) (*model.Controller, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Controller found with the given filter")
	}

	return objs[0], nil
}

 
// GetControllersByID - method returns associative a array of pointers of type "Controller", filter by Id
func (svc *ControllerService) GetControllersByID(pool *model.Pool, fields []string) (map[string]*model.Controller, error) {
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
	objs, err := svc.GetControllers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Controller)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetControllerById - method returns a pointer to "Controller"
func (svc *ControllerService) GetControllerById(id string) (*model.Controller, error) {
	return svc.objectSet.GetObject(id)
}

// GetControllersByName - method returns a associative array of pointers of type "Controller", filter by name 
func (svc *ControllerService) GetControllersByName(pool *model.Pool, fields []string) (map[string]*model.Controller, error) {
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
	objs, err := svc.GetControllers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Controller)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetControllerByName - method returns a pointer "Controller" 
func (svc *ControllerService) GetControllerByName(name string) (*model.Controller, error) {
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
	return onlyController(objs)
}	

