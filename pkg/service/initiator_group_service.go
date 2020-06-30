// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// InitiatorGroup Service - Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through
// group membership.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// InitiatorGroupService type 
type InitiatorGroupService struct {
	objectSet *client.InitiatorGroupObjectSet
}

// NewInitiatorGroupService - method to initialize "InitiatorGroupService" 
func NewInitiatorGroupService(gs *NsGroupService) (*InitiatorGroupService) {
	objectSet := gs.client.GetInitiatorGroupObjectSet()
	return &InitiatorGroupService{objectSet: objectSet}
}

// GetInitiatorGroups - method returns a array of pointers of type "InitiatorGroups"
func (svc *InitiatorGroupService) GetInitiatorGroups(params *util.GetParams) ([]*model.InitiatorGroup, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetInitiatorGroupsWithFields - method returns a array of pointers of type "InitiatorGroup" 
func (svc *InitiatorGroupService) GetInitiatorGroupsWithFields(fields []string) ([]*model.InitiatorGroup, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateInitiatorGroup - method creates a "InitiatorGroup"
func (svc *InitiatorGroupService) CreateInitiatorGroup(obj *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditInitiatorGroup - method modifies  the "InitiatorGroup" 
func (svc *InitiatorGroupService) EditInitiatorGroup(id string, obj *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyInitiatorGroup - private method for more than one element check. 
func onlyInitiatorGroup(objs []*model.InitiatorGroup) (*model.InitiatorGroup, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one InitiatorGroup found with the given filter")
	}

	return objs[0], nil
}

 
// GetInitiatorGroupsByID - method returns associative a array of pointers of type "InitiatorGroup", filter by Id
func (svc *InitiatorGroupService) GetInitiatorGroupsByID(pool *model.Pool, fields []string) (map[string]*model.InitiatorGroup, error) {
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
	objs, err := svc.GetInitiatorGroups(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.InitiatorGroup)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetInitiatorGroupById - method returns a pointer to "InitiatorGroup"
func (svc *InitiatorGroupService) GetInitiatorGroupById(id string) (*model.InitiatorGroup, error) {
	return svc.objectSet.GetObject(id)
}

// GetInitiatorGroupsByName - method returns a associative array of pointers of type "InitiatorGroup", filter by name 
func (svc *InitiatorGroupService) GetInitiatorGroupsByName(pool *model.Pool, fields []string) (map[string]*model.InitiatorGroup, error) {
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
	objs, err := svc.GetInitiatorGroups(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.InitiatorGroup)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetInitiatorGroupByName - method returns a pointer "InitiatorGroup" 
func (svc *InitiatorGroupService) GetInitiatorGroupByName(name string) (*model.InitiatorGroup, error) {
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
	return onlyInitiatorGroup(objs)
}	

