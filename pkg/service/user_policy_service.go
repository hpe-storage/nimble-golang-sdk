// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// UserPolicy Service - Manages the password policies configured for the group.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// UserPolicyService type
type UserPolicyService struct {
	objectSet *client.UserPolicyObjectSet
}

// NewUserPolicyService - method to initialize "UserPolicyService"
func NewUserPolicyService(gs *NsGroupService) *UserPolicyService {
	objectSet := gs.client.GetUserPolicyObjectSet()
	return &UserPolicyService{objectSet: objectSet}
}

// GetUserPolicies - method returns a array of pointers of type "UserPolicies"
func (svc *UserPolicyService) GetUserPolicies(params *param.GetParams) ([]*nimbleos.UserPolicy, error) {
	userPolicyResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return userPolicyResp, nil
}

// CreateUserPolicy - method creates a "UserPolicy"
func (svc *UserPolicyService) CreateUserPolicy(obj *nimbleos.UserPolicy) (*nimbleos.UserPolicy, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateUserPolicy: invalid parameter specified, %v", obj)
	}

	userPolicyResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return userPolicyResp, nil
}

// UpdateUserPolicy - method modifies  the "UserPolicy"
func (svc *UserPolicyService) UpdateUserPolicy(id string, obj *nimbleos.UserPolicy) (*nimbleos.UserPolicy, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateUserPolicy: invalid parameter specified, %v", obj)
	}

	userPolicyResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return userPolicyResp, nil
}

// GetUserPolicyById - method returns a pointer to "UserPolicy"
func (svc *UserPolicyService) GetUserPolicyById(id string) (*nimbleos.UserPolicy, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetUserPolicyById: invalid parameter specified, %v", id)
	}

	userPolicyResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return userPolicyResp, nil
}

// DeleteUserPolicy - deletes the "UserPolicy"
func (svc *UserPolicyService) DeleteUserPolicy(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteUserPolicy: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
