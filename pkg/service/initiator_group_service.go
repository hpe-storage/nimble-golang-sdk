// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// InitiatorGroup Service - Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through
// group membership.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// InitiatorGroupService type 
type InitiatorGroupService struct {
	objectSet *client.InitiatorGroupObjectSet
}

// NewInitiatorGroupService - method to initialize "InitiatorGroupService" 
func NewInitiatorGroupService(gs *NsGroupService) (*InitiatorGroupService) {
	objectSet := gs.client.GetInitiatorGroupObjectSet()
	return &InitiatorGroupService{objectSet: objectSet}
}

// GetInitiatorGroups - method returns a array of pointers of type "InitiatorGroups"
func (svc *InitiatorGroupService) GetInitiatorGroups(params *util.GetParams) ([]*model.InitiatorGroup, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateInitiatorGroup - method creates a "InitiatorGroup"
func (svc *InitiatorGroupService) CreateInitiatorGroup(obj *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateInitiatorGroup - method modifies  the "InitiatorGroup" 
func (svc *InitiatorGroupService) UpdateInitiatorGroup(id string, obj *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetInitiatorGroupById - method returns a pointer to "InitiatorGroup"
func (svc *InitiatorGroupService) GetInitiatorGroupById(id string) (*model.InitiatorGroup, error) {
	return svc.objectSet.GetObject(id)
}

// GetInitiatorGroupByName - method returns a pointer "InitiatorGroup" 
func (svc *InitiatorGroupService) GetInitiatorGroupByName(name string) (*model.InitiatorGroup, error) {
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

// DeleteInitiatorGroup - deletes the "InitiatorGroup"
func (svc *InitiatorGroupService) DeleteInitiatorGroup(id string) error {
	return svc.objectSet.DeleteObject(id)
}
