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
		return nil, fmt.Errorf("GetActiveDirectoryMemberships: invalid parameter specified, %v", params)
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
		return nil, fmt.Errorf("CreateActiveDirectoryMembership: invalid parameter specified, %v", obj)
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
		return nil, fmt.Errorf("UpdateActiveDirectoryMembership: invalid parameter specified, %v", obj)
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
		return nil, fmt.Errorf("GetActiveDirectoryMembershipById: invalid parameter specified, %v", id)
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
		return fmt.Errorf("DeleteActiveDirectoryMembership: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// RemoveActiveDirectoryMembership - leaves the Active Directory domain.
//   Required parameters:
//       id - ID of the active directory.
//       user - Name of the Active Directory user with the privilege to leave the domain.
//       password - Password for the Active Directory user.

//   Optional parameters:
//       force - Use this option when there is an error when leaving the domain.

func (svc *ActiveDirectoryMembershipService) RemoveActiveDirectoryMembership(id string, user string, password string, force *bool) error {

	if len(id) == 0 || len(user) == 0 || len(password) == 0 {
		return fmt.Errorf("RemoveActiveDirectoryMembership: invalid parameter specified id: %v, user: %v, ", id, user)
	}

	err := svc.objectSet.Remove(&id, &user, &password, force)
	return err
}

// ReportStatusActiveDirectoryMembership - reports the detail status of the Active Directory domain.
//   Required parameters:
//       id - ID of the active directory.

func (svc *ActiveDirectoryMembershipService) ReportStatusActiveDirectoryMembership(id string) (*nimbleos.NsADReportStatusReturn, error) {

	if len(id) == 0 {
		return nil, fmt.Errorf("ReportStatusActiveDirectoryMembership: invalid parameter specified id: %v ", id)
	}

	resp, err := svc.objectSet.ReportStatus(&id)
	return resp, err
}

// TestUserActiveDirectoryMembership - tests whether the user exist in the Active Directory. If the user is present, then the user's group and role information is reported.
//   Required parameters:
//       id - ID of the Active Directory.
//       name - Name of the Active Directory user.

func (svc *ActiveDirectoryMembershipService) TestUserActiveDirectoryMembership(id string, name string) (*nimbleos.NsADTestUserReturn, error) {

	if len(id) == 0 || len(name) == 0 {
		return nil, fmt.Errorf("TestUserActiveDirectoryMembership: invalid parameter specified id: %v, name: %v ", id, name)
	}

	resp, err := svc.objectSet.TestUser(&id, &name)
	return resp, err
}

// TestGroupActiveDirectoryMembership - tests whether the user group exist in the Active Directory.
//   Required parameters:
//       id - ID of the Active Directory.
//       name - Name of the Active Directory group.

func (svc *ActiveDirectoryMembershipService) TestGroupActiveDirectoryMembership(id string, name string) error {

	if len(id) == 0 || len(name) == 0 {
		return fmt.Errorf("TestGroupActiveDirectoryMembership: invalid parameter specified id: %v, name: %v ", id, name)
	}

	err := svc.objectSet.TestGroup(&id, &name)
	return err
}
