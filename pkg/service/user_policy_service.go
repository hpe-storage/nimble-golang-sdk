// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// UserPolicy Service - Manages the password policies configured for the group.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// UserPolicyService type 
type UserPolicyService struct {
	objectSet *client.UserPolicyObjectSet
}

// NewUserPolicyService - method to initialize "UserPolicyService" 
func NewUserPolicyService(gs *NsGroupService) (*UserPolicyService) {
	objectSet := gs.client.GetUserPolicyObjectSet()
	return &UserPolicyService{objectSet: objectSet}
}

// GetUserPolicies - method returns a array of pointers of type "UserPolicies"
func (svc *UserPolicyService) GetUserPolicies(params *util.GetParams) ([]*model.UserPolicy, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}


// CreateUserPolicy - method creates a "UserPolicy"
func (svc *UserPolicyService) CreateUserPolicy(obj *model.UserPolicy) (*model.UserPolicy, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateUserPolicy - method modifies  the "UserPolicy" 
func (svc *UserPolicyService) UpdateUserPolicy(id string, obj *model.UserPolicy) (*model.UserPolicy, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetUserPolicyById - method returns a pointer to "UserPolicy"
func (svc *UserPolicyService) GetUserPolicyById(id string) (*model.UserPolicy, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteUserPolicy - deletes the "UserPolicy"
func (svc *UserPolicyService) DeleteUserPolicy(id string) error {
	return svc.objectSet.DeleteObject(id)
}
