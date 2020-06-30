// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// User Service - Represents users configured to manage the system.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// UserService type 
type UserService struct {
	objectSet *client.UserObjectSet
}

// NewUserService - method to initialize "UserService" 
func NewUserService(gs *NsGroupService) (*UserService) {
	objectSet := gs.client.GetUserObjectSet()
	return &UserService{objectSet: objectSet}
}

// GetUsers - method returns a array of pointers of type "Users"
func (svc *UserService) GetUsers(params *util.GetParams) ([]*model.User, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetUsersWithFields - method returns a array of pointers of type "User" 
func (svc *UserService) GetUsersWithFields(fields []string) ([]*model.User, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateUser - method creates a "User"
func (svc *UserService) CreateUser(obj *model.User) (*model.User, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditUser - method modifies  the "User" 
func (svc *UserService) EditUser(id string, obj *model.User) (*model.User, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyUser - private method for more than one element check. 
func onlyUser(objs []*model.User) (*model.User, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one User found with the given filter")
	}

	return objs[0], nil
}

 
// GetUsersByID - method returns associative a array of pointers of type "User", filter by Id
func (svc *UserService) GetUsersByID(pool *model.Pool, fields []string) (map[string]*model.User, error) {
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
	objs, err := svc.GetUsers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.User)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetUserById - method returns a pointer to "User"
func (svc *UserService) GetUserById(id string) (*model.User, error) {
	return svc.objectSet.GetObject(id)
}

// GetUsersByName - method returns a associative array of pointers of type "User", filter by name 
func (svc *UserService) GetUsersByName(pool *model.Pool, fields []string) (map[string]*model.User, error) {
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
	objs, err := svc.GetUsers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.User)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetUserByName - method returns a pointer "User" 
func (svc *UserService) GetUserByName(name string) (*model.User, error) {
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
	return onlyUser(objs)
}	

