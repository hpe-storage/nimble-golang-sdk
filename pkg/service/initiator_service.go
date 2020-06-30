// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Initiator Service - Manage initiators in initiator groups. An initiator group has a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// InitiatorService type 
type InitiatorService struct {
	objectSet *client.InitiatorObjectSet
}

// NewInitiatorService - method to initialize "InitiatorService" 
func NewInitiatorService(gs *NsGroupService) (*InitiatorService) {
	objectSet := gs.client.GetInitiatorObjectSet()
	return &InitiatorService{objectSet: objectSet}
}

// GetInitiators - method returns a array of pointers of type "Initiators"
func (svc *InitiatorService) GetInitiators(params *util.GetParams) ([]*model.Initiator, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetInitiatorsWithFields - method returns a array of pointers of type "Initiator" 
func (svc *InitiatorService) GetInitiatorsWithFields(fields []string) ([]*model.Initiator, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateInitiator - method creates a "Initiator"
func (svc *InitiatorService) CreateInitiator(obj *model.Initiator) (*model.Initiator, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditInitiator - method modifies  the "Initiator" 
func (svc *InitiatorService) EditInitiator(id string, obj *model.Initiator) (*model.Initiator, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyInitiator - private method for more than one element check. 
func onlyInitiator(objs []*model.Initiator) (*model.Initiator, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Initiator found with the given filter")
	}

	return objs[0], nil
}

 
// GetInitiatorsByID - method returns associative a array of pointers of type "Initiator", filter by Id
func (svc *InitiatorService) GetInitiatorsByID(pool *model.Pool, fields []string) (map[string]*model.Initiator, error) {
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
	objs, err := svc.GetInitiators(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Initiator)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetInitiatorById - method returns a pointer to "Initiator"
func (svc *InitiatorService) GetInitiatorById(id string) (*model.Initiator, error) {
	return svc.objectSet.GetObject(id)
}


