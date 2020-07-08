// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// InitiatorGroup Service - Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through
// group membership.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// InitiatorGroupService type
type InitiatorGroupService struct {
	objectSet *client.InitiatorGroupObjectSet
}

// NewInitiatorGroupService - method to initialize "InitiatorGroupService"
func NewInitiatorGroupService(gs *NsGroupService) *InitiatorGroupService {
	objectSet := gs.client.GetInitiatorGroupObjectSet()
	return &InitiatorGroupService{objectSet: objectSet}
}

// GetInitiatorGroups - method returns a array of pointers of type "InitiatorGroups"
func (svc *InitiatorGroupService) GetInitiatorGroups(params *util.GetParams) ([]*model.InitiatorGroup, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	initiatorGroupResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return initiatorGroupResp, nil
}

// CreateInitiatorGroup - method creates a "InitiatorGroup"
func (svc *InitiatorGroupService) CreateInitiatorGroup(obj *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	initiatorGroupResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return initiatorGroupResp, nil
}

// UpdateInitiatorGroup - method modifies  the "InitiatorGroup"
func (svc *InitiatorGroupService) UpdateInitiatorGroup(id string, obj *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	initiatorGroupResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return initiatorGroupResp, nil
}

// GetInitiatorGroupById - method returns a pointer to "InitiatorGroup"
func (svc *InitiatorGroupService) GetInitiatorGroupById(id string) (*model.InitiatorGroup, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	initiatorGroupResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return initiatorGroupResp, nil
}

// GetInitiatorGroupByName - method returns a pointer "InitiatorGroup"
func (svc *InitiatorGroupService) GetInitiatorGroupByName(name string) (*model.InitiatorGroup, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	initiatorGroupResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(initiatorGroupResp) == 0 {
		return nil, nil
	}

	return initiatorGroupResp[0], nil
}

// DeleteInitiatorGroup - deletes the "InitiatorGroup"
func (svc *InitiatorGroupService) DeleteInitiatorGroup(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
