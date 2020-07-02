// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// User Service - Represents users configured to manage the system.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateUser - method creates a "User"
func (svc *UserService) CreateUser(obj *model.User) (*model.User, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateUser - method modifies  the "User" 
func (svc *UserService) UpdateUser(id string, obj *model.User) (*model.User, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetUserById - method returns a pointer to "User"
func (svc *UserService) GetUserById(id string) (*model.User, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteUser - deletes the "User"
func (svc *UserService) DeleteUser(id string) error {
	return svc.objectSet.DeleteObject(id)
}
