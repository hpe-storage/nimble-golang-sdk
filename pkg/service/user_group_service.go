// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// UserGroup Service - Represents Active Directory groups configured to manage the system.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateUserGroup - method creates a "UserGroup"
func (svc *UserGroupService) CreateUserGroup(obj *model.UserGroup) (*model.UserGroup, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateUserGroup - method modifies  the "UserGroup" 
func (svc *UserGroupService) UpdateUserGroup(id string, obj *model.UserGroup) (*model.UserGroup, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetUserGroupById - method returns a pointer to "UserGroup"
func (svc *UserGroupService) GetUserGroupById(id string) (*model.UserGroup, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteUserGroup - deletes the "UserGroup"
func (svc *UserGroupService) DeleteUserGroup(id string) error {
	return svc.objectSet.DeleteObject(id)
}
