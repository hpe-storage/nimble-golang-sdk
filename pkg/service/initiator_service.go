// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Initiator Service - Manage initiators in initiator groups. An initiator group has a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// InitiatorService type 
type InitiatorService struct {
	objectSet *client.InitiatorObjectSet
}

// NewInitiatorService - method to initialize "InitiatorService" 
func NewInitiatorService(gs *NsGroupService) (*InitiatorService) {
	objectSet := gs.client.GetInitiatorObjectSet()
	return &InitiatorService{objectSet: objectSet}
}

// GetInitiators - method returns a array of pointers of type "Initiators"
func (svc *InitiatorService) GetInitiators(params *util.GetParams) ([]*model.Initiator, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateInitiator - method creates a "Initiator"
func (svc *InitiatorService) CreateInitiator(obj *model.Initiator) (*model.Initiator, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateInitiator - method modifies  the "Initiator" 
func (svc *InitiatorService) UpdateInitiator(id string, obj *model.Initiator) (*model.Initiator, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetInitiatorById - method returns a pointer to "Initiator"
func (svc *InitiatorService) GetInitiatorById(id string) (*model.Initiator, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteInitiator - deletes the "Initiator"
func (svc *InitiatorService) DeleteInitiator(id string) error {
	return svc.objectSet.DeleteObject(id)
}
