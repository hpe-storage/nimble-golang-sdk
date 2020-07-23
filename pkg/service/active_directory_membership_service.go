// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ActiveDirectoryMembership Service - Manages the storage array's membership with the Active Directory.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ActiveDirectoryMembershipService type
type ActiveDirectoryMembershipService struct {
	objectSet *client.ActiveDirectoryMembershipObjectSet
}

// NewActiveDirectoryMembershipService - method to initialize "ActiveDirectoryMembershipService"
func NewActiveDirectoryMembershipService(gs *NsGroupService) *ActiveDirectoryMembershipService {
	objectSet := gs.client.GetActiveDirectoryMembershipObjectSet()
	return &ActiveDirectoryMembershipService{objectSet: objectSet}
}

// GetActiveDirectoryMemberships - method returns a array of pointers of type "ActiveDirectoryMemberships"
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMemberships(params *param.GetParams) ([]*nimbleos.ActiveDirectoryMembership, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	activeDirectoryMembershipResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return activeDirectoryMembershipResp, nil
}

// CreateActiveDirectoryMembership - method creates a "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) CreateActiveDirectoryMembership(obj *nimbleos.ActiveDirectoryMembership) (*nimbleos.ActiveDirectoryMembership, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	activeDirectoryMembershipResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return activeDirectoryMembershipResp, nil
}

// UpdateActiveDirectoryMembership - method modifies  the "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) UpdateActiveDirectoryMembership(id string, obj *nimbleos.ActiveDirectoryMembership) (*nimbleos.ActiveDirectoryMembership, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	activeDirectoryMembershipResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return activeDirectoryMembershipResp, nil
}

// GetActiveDirectoryMembershipById - method returns a pointer to "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipById(id string) (*nimbleos.ActiveDirectoryMembership, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	activeDirectoryMembershipResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return activeDirectoryMembershipResp, nil
}

// GetActiveDirectoryMembershipByName - method returns a pointer "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) GetActiveDirectoryMembershipByName(name string) (*nimbleos.ActiveDirectoryMembership, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	activeDirectoryMembershipResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(activeDirectoryMembershipResp) == 0 {
		return nil, nil
	}

	return activeDirectoryMembershipResp[0], nil
}

// DeleteActiveDirectoryMembership - deletes the "ActiveDirectoryMembership"
func (svc *ActiveDirectoryMembershipService) DeleteActiveDirectoryMembership(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
