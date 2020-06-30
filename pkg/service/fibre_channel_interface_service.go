// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelInterface Service - Represent information of specified Fibre Channel interfaces. Fibre Channel interfaces are hosted on Fibre Channel ports to provide data access.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// FibreChannelInterfaceService type 
type FibreChannelInterfaceService struct {
	objectSet *client.FibreChannelInterfaceObjectSet
}

// NewFibreChannelInterfaceService - method to initialize "FibreChannelInterfaceService" 
func NewFibreChannelInterfaceService(gs *NsGroupService) (*FibreChannelInterfaceService) {
	objectSet := gs.client.GetFibreChannelInterfaceObjectSet()
	return &FibreChannelInterfaceService{objectSet: objectSet}
}

// GetFibreChannelInterfaces - method returns a array of pointers of type "FibreChannelInterfaces"
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaces(params *util.GetParams) ([]*model.FibreChannelInterface, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetFibreChannelInterfacesWithFields - method returns a array of pointers of type "FibreChannelInterface" 
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfacesWithFields(fields []string) ([]*model.FibreChannelInterface, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelInterface - method creates a "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) CreateFibreChannelInterface(obj *model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditFibreChannelInterface - method modifies  the "FibreChannelInterface" 
func (svc *FibreChannelInterfaceService) EditFibreChannelInterface(id string, obj *model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyFibreChannelInterface - private method for more than one element check. 
func onlyFibreChannelInterface(objs []*model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one FibreChannelInterface found with the given filter")
	}

	return objs[0], nil
}

 
// GetFibreChannelInterfacesByID - method returns associative a array of pointers of type "FibreChannelInterface", filter by Id
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfacesByID(pool *model.Pool, fields []string) (map[string]*model.FibreChannelInterface, error) {
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
	objs, err := svc.GetFibreChannelInterfaces(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FibreChannelInterface)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetFibreChannelInterfaceById - method returns a pointer to "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaceById(id string) (*model.FibreChannelInterface, error) {
	return svc.objectSet.GetObject(id)
}

// GetFibreChannelInterfacesByName - method returns a associative array of pointers of type "FibreChannelInterface", filter by name 
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfacesByName(pool *model.Pool, fields []string) (map[string]*model.FibreChannelInterface, error) {
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
	objs, err := svc.GetFibreChannelInterfaces(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FibreChannelInterface)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetFibreChannelInterfaceByName - method returns a pointer "FibreChannelInterface" 
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaceByName(name string) (*model.FibreChannelInterface, error) {
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
	return onlyFibreChannelInterface(objs)
}	

