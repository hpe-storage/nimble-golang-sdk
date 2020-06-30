// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Role Service - Retrieve roles and privileges for role-based access control.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// RoleService type 
type RoleService struct {
	objectSet *client.RoleObjectSet
}

// NewRoleService - method to initialize "RoleService" 
func NewRoleService(gs *NsGroupService) (*RoleService) {
	objectSet := gs.client.GetRoleObjectSet()
	return &RoleService{objectSet: objectSet}
}

// GetRoles - method returns a array of pointers of type "Roles"
func (svc *RoleService) GetRoles(params *util.GetParams) ([]*model.Role, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetRolesWithFields - method returns a array of pointers of type "Role" 
func (svc *RoleService) GetRolesWithFields(fields []string) ([]*model.Role, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateRole - method creates a "Role"
func (svc *RoleService) CreateRole(obj *model.Role) (*model.Role, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditRole - method modifies  the "Role" 
func (svc *RoleService) EditRole(id string, obj *model.Role) (*model.Role, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyRole - private method for more than one element check. 
func onlyRole(objs []*model.Role) (*model.Role, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Role found with the given filter")
	}

	return objs[0], nil
}

 
// GetRolesByID - method returns associative a array of pointers of type "Role", filter by Id
func (svc *RoleService) GetRolesByID(pool *model.Pool, fields []string) (map[string]*model.Role, error) {
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
	objs, err := svc.GetRoles(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Role)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetRoleById - method returns a pointer to "Role"
func (svc *RoleService) GetRoleById(id string) (*model.Role, error) {
	return svc.objectSet.GetObject(id)
}

// GetRolesByName - method returns a associative array of pointers of type "Role", filter by name 
func (svc *RoleService) GetRolesByName(pool *model.Pool, fields []string) (map[string]*model.Role, error) {
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
	objs, err := svc.GetRoles(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Role)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetRoleByName - method returns a pointer "Role" 
func (svc *RoleService) GetRoleByName(name string) (*model.Role, error) {
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
	return onlyRole(objs)
}	

