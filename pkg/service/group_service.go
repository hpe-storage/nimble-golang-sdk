// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Group Service - Group is a collection of arrays operating together organized into storage pools.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// GroupService type 
type GroupService struct {
	objectSet *client.GroupObjectSet
}

// NewGroupService - method to initialize "GroupService" 
func NewGroupService(gs *NsGroupService) (*GroupService) {
	objectSet := gs.client.GetGroupObjectSet()
	return &GroupService{objectSet: objectSet}
}

// GetGroups - method returns a array of pointers of type "Groups"
func (svc *GroupService) GetGroups(params *util.GetParams) ([]*model.Group, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateGroup - method creates a "Group"
func (svc *GroupService) CreateGroup(obj *model.Group) (*model.Group, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateGroup - method modifies  the "Group" 
func (svc *GroupService) UpdateGroup(id string, obj *model.Group) (*model.Group, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetGroupById - method returns a pointer to "Group"
func (svc *GroupService) GetGroupById(id string) (*model.Group, error) {
	return svc.objectSet.GetObject(id)
}

// GetGroupByName - method returns a pointer "Group" 
func (svc *GroupService) GetGroupByName(name string) (*model.Group, error) {
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

// DeleteGroup - deletes the "Group"
func (svc *GroupService) DeleteGroup(id string) error {
	return svc.objectSet.DeleteObject(id)
}
