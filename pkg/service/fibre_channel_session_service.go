// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelSession Service - Fibre Channel session is created when Fibre Channel initiator connects to this group.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// FibreChannelSessionService type 
type FibreChannelSessionService struct {
	objectSet *client.FibreChannelSessionObjectSet
}

// NewFibreChannelSessionService - method to initialize "FibreChannelSessionService" 
func NewFibreChannelSessionService(gs *NsGroupService) (*FibreChannelSessionService) {
	objectSet := gs.client.GetFibreChannelSessionObjectSet()
	return &FibreChannelSessionService{objectSet: objectSet}
}

// GetFibreChannelSessions - method returns a array of pointers of type "FibreChannelSessions"
func (svc *FibreChannelSessionService) GetFibreChannelSessions(params *util.GetParams) ([]*model.FibreChannelSession, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetFibreChannelSessionsWithFields - method returns a array of pointers of type "FibreChannelSession" 
func (svc *FibreChannelSessionService) GetFibreChannelSessionsWithFields(fields []string) ([]*model.FibreChannelSession, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelSession - method creates a "FibreChannelSession"
func (svc *FibreChannelSessionService) CreateFibreChannelSession(obj *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditFibreChannelSession - method modifies  the "FibreChannelSession" 
func (svc *FibreChannelSessionService) EditFibreChannelSession(id string, obj *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyFibreChannelSession - private method for more than one element check. 
func onlyFibreChannelSession(objs []*model.FibreChannelSession) (*model.FibreChannelSession, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one FibreChannelSession found with the given filter")
	}

	return objs[0], nil
}

 
// GetFibreChannelSessionsByID - method returns associative a array of pointers of type "FibreChannelSession", filter by Id
func (svc *FibreChannelSessionService) GetFibreChannelSessionsByID(pool *model.Pool, fields []string) (map[string]*model.FibreChannelSession, error) {
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
	objs, err := svc.GetFibreChannelSessions(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.FibreChannelSession)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetFibreChannelSessionById - method returns a pointer to "FibreChannelSession"
func (svc *FibreChannelSessionService) GetFibreChannelSessionById(id string) (*model.FibreChannelSession, error) {
	return svc.objectSet.GetObject(id)
}


