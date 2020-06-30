// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ActiveDirectoryMembership Service - Manages the storage array's membership with the Active Directory.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// ActiveDirectoryMembershipService type 
type ActiveDirectoryMembershipService struct {
	objectSet *client.ActiveDirectoryMembershipObjectSet
}

// NewActiveDirectoryMembershipService - method to initialize "ActiveDirectoryMembershipService" 
func NewActiveDirectoryMembershipService(gs *NsGroupService) (*ActiveDirectoryMembershipService) {
	objectSet := gs.client.GetActiveDirectoryMembershipObjectSet()
	return &ActiveDirectoryMembershipService{objectSet: objectSet}
}

// GetActiveDirectoryMemberships - method returns a array of pointers of type "ActiveDirectoryMemberships"
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMemberships(params *util.GetParams) ([]*model.ActiveDirectoryMembership, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetActiveDirectoryMembershipsWithFields - method returns a array of pointers of type "ActiveDirectoryMembership" 
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipsWithFields(fields []string) ([]*model.ActiveDirectoryMembership, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateActiveDirectoryMembership - method creates a "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) CreateActiveDirectoryMembership(obj *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditActiveDirectoryMembership - method modifies  the "ActiveDirectoryMembership" 
func (svc *ActiveDirectoryMembershipService) EditActiveDirectoryMembership(id string, obj *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyActiveDirectoryMembership - private method for more than one element check. 
func onlyActiveDirectoryMembership(objs []*model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ActiveDirectoryMembership found with the given filter")
	}

	return objs[0], nil
}

 
// GetActiveDirectoryMembershipsByID - method returns associative a array of pointers of type "ActiveDirectoryMembership", filter by Id
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipsByID(pool *model.Pool, fields []string) (map[string]*model.ActiveDirectoryMembership, error) {
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
	objs, err := svc.GetActiveDirectoryMemberships(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ActiveDirectoryMembership)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetActiveDirectoryMembershipById - method returns a pointer to "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipById(id string) (*model.ActiveDirectoryMembership, error) {
	return svc.objectSet.GetObject(id)
}

// GetActiveDirectoryMembershipsByName - method returns a associative array of pointers of type "ActiveDirectoryMembership", filter by name 
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipsByName(pool *model.Pool, fields []string) (map[string]*model.ActiveDirectoryMembership, error) {
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
	objs, err := svc.GetActiveDirectoryMemberships(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ActiveDirectoryMembership)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetActiveDirectoryMembershipByName - method returns a pointer "ActiveDirectoryMembership" 
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipByName(name string) (*model.ActiveDirectoryMembership, error) {
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
	return onlyActiveDirectoryMembership(objs)
}	

