// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelInitiatorAlias Service - This API provides the alias information for Fibre Channel initiators.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// FibreChannelInitiatorAliasService type 
type FibreChannelInitiatorAliasService struct {
	objectSet *client.FibreChannelInitiatorAliasObjectSet
}

// NewFibreChannelInitiatorAliasService - method to initialize "FibreChannelInitiatorAliasService" 
func NewFibreChannelInitiatorAliasService(gs *NsGroupService) (*FibreChannelInitiatorAliasService) {
	objectSet := gs.client.GetFibreChannelInitiatorAliasObjectSet()
	return &FibreChannelInitiatorAliasService{objectSet: objectSet}
}

// GetFibreChannelInitiatorAliass - method returns a array of pointers of type "FibreChannelInitiatorAliass"
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliass(params *util.GetParams) ([]*model.FibreChannelInitiatorAlias, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetFibreChannelInitiatorAliassWithFields - method returns a array of pointers of type "FibreChannelInitiatorAlias" 
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliassWithFields(fields []string) ([]*model.FibreChannelInitiatorAlias, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelInitiatorAlias - method creates a "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) CreateFibreChannelInitiatorAlias(obj *model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditFibreChannelInitiatorAlias - method modifies  the "FibreChannelInitiatorAlias" 
func (svc *FibreChannelInitiatorAliasService) EditFibreChannelInitiatorAlias(id string, obj *model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyFibreChannelInitiatorAlias - private method for more than one element check. 
func onlyFibreChannelInitiatorAlias(objs []*model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one FibreChannelInitiatorAlias found with the given filter")
	}

	return objs[0], nil
}

 
// GetFibreChannelInitiatorAliassByID - method returns associative a array of pointers of type "FibreChannelInitiatorAlias", filter by Id
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliassByID(pool *model.Pool, fields []string) (map[string]*model.FibreChannelInitiatorAlias, error) {
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
	objs, err := svc.GetFibreChannelInitiatorAliass(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FibreChannelInitiatorAlias)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetFibreChannelInitiatorAliasById - method returns a pointer to "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliasById(id string) (*model.FibreChannelInitiatorAlias, error) {
	return svc.objectSet.GetObject(id)
}


