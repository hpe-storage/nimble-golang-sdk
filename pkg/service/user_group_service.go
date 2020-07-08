// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// UserGroup Service - Represents Active Directory groups configured to manage the system.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// UserGroupService type
type UserGroupService struct {
	objectSet *client.UserGroupObjectSet
}

// NewUserGroupService - method to initialize "UserGroupService"
func NewUserGroupService(gs *NsGroupService) *UserGroupService {
	objectSet := gs.client.GetUserGroupObjectSet()
	return &UserGroupService{objectSet: objectSet}
}

// GetUserGroups - method returns a array of pointers of type "UserGroups"
func (svc *UserGroupService) GetUserGroups(params *util.GetParams) ([]*model.UserGroup, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	userGroupResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// CreateUserGroup - method creates a "UserGroup"
func (svc *UserGroupService) CreateUserGroup(obj *model.UserGroup) (*model.UserGroup, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	userGroupResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// UpdateUserGroup - method modifies  the "UserGroup"
func (svc *UserGroupService) UpdateUserGroup(id string, obj *model.UserGroup) (*model.UserGroup, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	userGroupResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// GetUserGroupById - method returns a pointer to "UserGroup"
func (svc *UserGroupService) GetUserGroupById(id string) (*model.UserGroup, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	userGroupResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// GetUserGroupByName - method returns a pointer "UserGroup"
func (svc *UserGroupService) GetUserGroupByName(name string) (*model.UserGroup, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	userGroupResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(userGroupResp) == 0 {
		return nil, nil
	}

	return userGroupResp[0], nil
}

// DeleteUserGroup - deletes the "UserGroup"
func (svc *UserGroupService) DeleteUserGroup(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
