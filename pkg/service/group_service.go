// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Group Service - Group is a collection of arrays operating together organized into storage pools.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// GroupService type
type GroupService struct {
	objectSet *client.GroupObjectSet
}

// NewGroupService - method to initialize "GroupService"
func NewGroupService(gs *NsGroupService) *GroupService {
	objectSet := gs.client.GetGroupObjectSet()
	return &GroupService{objectSet: objectSet}
}

// GetGroups - method returns a array of pointers of type "Groups"
func (svc *GroupService) GetGroups(params *util.GetParams) ([]*model.Group, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	groupResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return groupResp, nil
}

// CreateGroup - method creates a "Group"
func (svc *GroupService) CreateGroup(obj *model.Group) (*model.Group, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	groupResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return groupResp, nil
}

// UpdateGroup - method modifies  the "Group"
func (svc *GroupService) UpdateGroup(id string, obj *model.Group) (*model.Group, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	groupResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return groupResp, nil
}

// GetGroupById - method returns a pointer to "Group"
func (svc *GroupService) GetGroupById(id string) (*model.Group, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	groupResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return groupResp, nil
}

// GetGroupByName - method returns a pointer "Group"
func (svc *GroupService) GetGroupByName(name string) (*model.Group, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	groupResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(groupResp) == 0 {
		return nil, nil
	}

	return groupResp[0], nil
}

// DeleteGroup - deletes the "Group"
func (svc *GroupService) DeleteGroup(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
