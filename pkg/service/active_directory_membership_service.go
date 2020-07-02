// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ActiveDirectoryMembership Service - Manages the storage array's membership with the Active Directory.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateActiveDirectoryMembership - method creates a "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) CreateActiveDirectoryMembership(obj *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateActiveDirectoryMembership - method modifies  the "ActiveDirectoryMembership" 
func (svc *ActiveDirectoryMembershipService) UpdateActiveDirectoryMembership(id string, obj *model.ActiveDirectoryMembership) (*model.ActiveDirectoryMembership, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetActiveDirectoryMembershipById - method returns a pointer to "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipById(id string) (*model.ActiveDirectoryMembership, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteActiveDirectoryMembership - deletes the "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) DeleteActiveDirectoryMembership(id string) error {
	return svc.objectSet.DeleteObject(id)
}
