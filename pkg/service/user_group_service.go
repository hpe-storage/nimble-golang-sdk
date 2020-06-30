// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// UserGroup Service - Represents Active Directory groups configured to manage the system.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// UserGroupService type 
type UserGroupService struct {
	objectSet *client.UserGroupObjectSet
}

// NewUserGroupService - method to initialize "UserGroupService" 
func NewUserGroupService(gs *NsGroupService) (*UserGroupService) {
	objectSet := gs.client.GetUserGroupObjectSet()
	return &UserGroupService{objectSet: objectSet}
}

// GetUserGroups - method returns a array of pointers of type "UserGroups"
func (svc *UserGroupService) GetUserGroups(params *util.GetParams) ([]*model.UserGroup, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetUserGroupsWithFields - method returns a array of pointers of type "UserGroup" 
func (svc *UserGroupService) GetUserGroupsWithFields(fields []string) ([]*model.UserGroup, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateUserGroup - method creates a "UserGroup"
func (svc *UserGroupService) CreateUserGroup(obj *model.UserGroup) (*model.UserGroup, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditUserGroup - method modifies  the "UserGroup" 
func (svc *UserGroupService) EditUserGroup(id string, obj *model.UserGroup) (*model.UserGroup, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyUserGroup - private method for more than one element check. 
func onlyUserGroup(objs []*model.UserGroup) (*model.UserGroup, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one UserGroup found with the given filter")
	}

	return objs[0], nil
}

 
// GetUserGroupsByID - method returns associative a array of pointers of type "UserGroup", filter by Id
func (svc *UserGroupService) GetUserGroupsByID(pool *model.Pool, fields []string) (map[string]*model.UserGroup, error) {
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
	objs, err := svc.GetUserGroups(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.UserGroup)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetUserGroupById - method returns a pointer to "UserGroup"
func (svc *UserGroupService) GetUserGroupById(id string) (*model.UserGroup, error) {
	return svc.objectSet.GetObject(id)
}

// GetUserGroupsByName - method returns a associative array of pointers of type "UserGroup", filter by name 
func (svc *UserGroupService) GetUserGroupsByName(pool *model.Pool, fields []string) (map[string]*model.UserGroup, error) {
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
	objs, err := svc.GetUserGroups(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.UserGroup)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetUserGroupByName - method returns a pointer "UserGroup" 
func (svc *UserGroupService) GetUserGroupByName(name string) (*model.UserGroup, error) {
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
	return onlyUserGroup(objs)
}	

