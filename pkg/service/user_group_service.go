// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// UserGroup Service - Represents Active Directory groups configured to manage the system.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *UserGroupService) GetUserGroups(params *param.GetParams) ([]*nimbleos.UserGroup, error) {
	userGroupResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// CreateUserGroup - method creates a "UserGroup"
func (svc *UserGroupService) CreateUserGroup(obj *nimbleos.UserGroup) (*nimbleos.UserGroup, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateUserGroup: invalid parameter specified, %v", obj)
	}

	userGroupResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// UpdateUserGroup - method modifies  the "UserGroup"
func (svc *UserGroupService) UpdateUserGroup(id string, obj *nimbleos.UserGroup) (*nimbleos.UserGroup, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateUserGroup: invalid parameter specified, %v", obj)
	}

	userGroupResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// GetUserGroupById - method returns a pointer to "UserGroup"
func (svc *UserGroupService) GetUserGroupById(id string) (*nimbleos.UserGroup, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetUserGroupById: invalid parameter specified, %v", id)
	}

	userGroupResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return userGroupResp, nil
}

// GetUserGroupByName - method returns a pointer "UserGroup"
func (svc *UserGroupService) GetUserGroupByName(name string) (*nimbleos.UserGroup, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
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
		return fmt.Errorf("DeleteUserGroup: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
