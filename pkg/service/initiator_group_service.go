// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// InitiatorGroup Service - Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through
// group membership.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *InitiatorGroupService) GetInitiatorGroups(params *param.GetParams) ([]*nimbleos.InitiatorGroup, error) {
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
func (svc *InitiatorGroupService) CreateInitiatorGroup(obj *nimbleos.InitiatorGroup) (*nimbleos.InitiatorGroup, error) {
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
func (svc *InitiatorGroupService) UpdateInitiatorGroup(id string, obj *nimbleos.InitiatorGroup) (*nimbleos.InitiatorGroup, error) {
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
func (svc *InitiatorGroupService) GetInitiatorGroupById(id string) (*nimbleos.InitiatorGroup, error) {
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
func (svc *InitiatorGroupService) GetInitiatorGroupByName(name string) (*nimbleos.InitiatorGroup, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
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

// SuggestLunInitiatorGroup - suggest an LU number for the volume and initiator group combination.
//   Required parameters:
//       id - ID of the initiator group.

//   Optional parameters:
//       volId - ID of the volume.

func (svc *InitiatorGroupService) SuggestLunInitiatorGroup(id string, volId *string) (*nimbleos.NsLunReturn, error) {

	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified id:%v, volId:%v ", id, volId)
	}

	resp, err := svc.objectSet.SuggestLun(&id, volId)
	return resp, err

}

// ValidateLunInitiatorGroup - validate an LU number for the volume and initiator group combination.
//   Required parameters:
//       id - ID of the initiator group.
//       lun - LU number to validate in decimal.

//   Optional parameters:
//       volId - ID of the volume.

func (svc *InitiatorGroupService) ValidateLunInitiatorGroup(id string, volId *string, lun uint64) error {

	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified id:%v, volId:%v, lun:%v ", id, volId, lun)
	}

	err := svc.objectSet.ValidateLun(&id, volId, &lun)
	return err
}
