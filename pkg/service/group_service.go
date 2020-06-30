// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Group Service - Group is a collection of arrays operating together organized into storage pools.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// GroupService type 
type GroupService struct {
	objectSet *client.GroupObjectSet
}

// NewGroupService - method to initialize "GroupService" 
func NewGroupService(gs *NsGroupService) (*GroupService) {
	objectSet := gs.client.GetGroupObjectSet()
	return &GroupService{objectSet: objectSet}
}

// GetGroups - method returns a array of pointers of type "Groups"
func (svc *GroupService) GetGroups(params *util.GetParams) ([]*model.Group, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetGroupsWithFields - method returns a array of pointers of type "Group" 
func (svc *GroupService) GetGroupsWithFields(fields []string) ([]*model.Group, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateGroup - method creates a "Group"
func (svc *GroupService) CreateGroup(obj *model.Group) (*model.Group, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditGroup - method modifies  the "Group" 
func (svc *GroupService) EditGroup(id string, obj *model.Group) (*model.Group, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyGroup - private method for more than one element check. 
func onlyGroup(objs []*model.Group) (*model.Group, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Group found with the given filter")
	}

	return objs[0], nil
}

 
// GetGroupsByID - method returns associative a array of pointers of type "Group", filter by Id
func (svc *GroupService) GetGroupsByID(pool *model.Pool, fields []string) (map[string]*model.Group, error) {
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
	objs, err := svc.GetGroups(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Group)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetGroupById - method returns a pointer to "Group"
func (svc *GroupService) GetGroupById(id string) (*model.Group, error) {
	return svc.objectSet.GetObject(id)
}

// GetGroupsByName - method returns a associative array of pointers of type "Group", filter by name 
func (svc *GroupService) GetGroupsByName(pool *model.Pool, fields []string) (map[string]*model.Group, error) {
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
	objs, err := svc.GetGroups(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Group)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetGroupByName - method returns a pointer "Group" 
func (svc *GroupService) GetGroupByName(name string) (*model.Group, error) {
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
	return onlyGroup(objs)
}	

