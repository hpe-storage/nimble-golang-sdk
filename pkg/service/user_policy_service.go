// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// UserPolicy Service - Manages the password policies configured for the group.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetUserPolicys - method returns a array of pointers of type "UserPolicys"
func (svc *UserPolicyService) GetUserPolicys(params *util.GetParams) ([]*model.UserPolicy, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetUserPolicysWithFields - method returns a array of pointers of type "UserPolicy" 
func (svc *UserPolicyService) GetUserPolicysWithFields(fields []string) ([]*model.UserPolicy, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateUserPolicy - method creates a "UserPolicy"
func (svc *UserPolicyService) CreateUserPolicy(obj *model.UserPolicy) (*model.UserPolicy, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditUserPolicy - method modifies  the "UserPolicy" 
func (svc *UserPolicyService) EditUserPolicy(id string, obj *model.UserPolicy) (*model.UserPolicy, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyUserPolicy - private method for more than one element check. 
func onlyUserPolicy(objs []*model.UserPolicy) (*model.UserPolicy, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one UserPolicy found with the given filter")
	}

	return objs[0], nil
}

 
// GetUserPolicysByID - method returns associative a array of pointers of type "UserPolicy", filter by Id
func (svc *UserPolicyService) GetUserPolicysByID(pool *model.Pool, fields []string) (map[string]*model.UserPolicy, error) {
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
	objs, err := svc.GetUserPolicys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.UserPolicy)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetUserPolicyById - method returns a pointer to "UserPolicy"
func (svc *UserPolicyService) GetUserPolicyById(id string) (*model.UserPolicy, error) {
	return svc.objectSet.GetObject(id)
}


