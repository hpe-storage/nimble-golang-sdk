// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Initiator Service - Manage initiators in initiator groups. An initiator group has a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// InitiatorService type
type InitiatorService struct {
	objectSet *client.InitiatorObjectSet
}

// NewInitiatorService - method to initialize "InitiatorService"
func NewInitiatorService(gs *NsGroupService) *InitiatorService {
	objectSet := gs.client.GetInitiatorObjectSet()
	return &InitiatorService{objectSet: objectSet}
}

// GetInitiators - method returns a array of pointers of type "Initiators"
func (svc *InitiatorService) GetInitiators(params *util.GetParams) ([]*model.Initiator, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	initiatorResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return initiatorResp, nil
}

// CreateInitiator - method creates a "Initiator"
func (svc *InitiatorService) CreateInitiator(obj *model.Initiator) (*model.Initiator, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	initiatorResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return initiatorResp, nil
}

// UpdateInitiator - method modifies  the "Initiator"
func (svc *InitiatorService) UpdateInitiator(id string, obj *model.Initiator) (*model.Initiator, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	initiatorResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return initiatorResp, nil
}

// GetInitiatorById - method returns a pointer to "Initiator"
func (svc *InitiatorService) GetInitiatorById(id string) (*model.Initiator, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	initiatorResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return initiatorResp, nil
}

// DeleteInitiator - deletes the "Initiator"
func (svc *InitiatorService) DeleteInitiator(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
