// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelPort Service - Fibre Channel ports provide data access. This API provides the list of all Fibre Channel ports configured on the arrays.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// FibreChannelPortService type 
type FibreChannelPortService struct {
	objectSet *client.FibreChannelPortObjectSet
}

// NewFibreChannelPortService - method to initialize "FibreChannelPortService" 
func NewFibreChannelPortService(gs *NsGroupService) (*FibreChannelPortService) {
	objectSet := gs.client.GetFibreChannelPortObjectSet()
	return &FibreChannelPortService{objectSet: objectSet}
}

// GetFibreChannelPorts - method returns a array of pointers of type "FibreChannelPorts"
func (svc *FibreChannelPortService) GetFibreChannelPorts(params *util.GetParams) ([]*model.FibreChannelPort, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetFibreChannelPortsWithFields - method returns a array of pointers of type "FibreChannelPort" 
func (svc *FibreChannelPortService) GetFibreChannelPortsWithFields(fields []string) ([]*model.FibreChannelPort, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelPort - method creates a "FibreChannelPort"
func (svc *FibreChannelPortService) CreateFibreChannelPort(obj *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditFibreChannelPort - method modifies  the "FibreChannelPort" 
func (svc *FibreChannelPortService) EditFibreChannelPort(id string, obj *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyFibreChannelPort - private method for more than one element check. 
func onlyFibreChannelPort(objs []*model.FibreChannelPort) (*model.FibreChannelPort, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one FibreChannelPort found with the given filter")
	}

	return objs[0], nil
}

 
// GetFibreChannelPortsByID - method returns associative a array of pointers of type "FibreChannelPort", filter by Id
func (svc *FibreChannelPortService) GetFibreChannelPortsByID(pool *model.Pool, fields []string) (map[string]*model.FibreChannelPort, error) {
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
	objs, err := svc.GetFibreChannelPorts(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FibreChannelPort)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetFibreChannelPortById - method returns a pointer to "FibreChannelPort"
func (svc *FibreChannelPortService) GetFibreChannelPortById(id string) (*model.FibreChannelPort, error) {
	return svc.objectSet.GetObject(id)
}


