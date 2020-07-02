// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelSession Service - Fibre Channel session is created when Fibre Channel initiator connects to this group.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// FibreChannelSessionService type 
type FibreChannelSessionService struct {
	objectSet *client.FibreChannelSessionObjectSet
}

// NewFibreChannelSessionService - method to initialize "FibreChannelSessionService" 
func NewFibreChannelSessionService(gs *NsGroupService) (*FibreChannelSessionService) {
	objectSet := gs.client.GetFibreChannelSessionObjectSet()
	return &FibreChannelSessionService{objectSet: objectSet}
}

// GetFibreChannelSessions - method returns a array of pointers of type "FibreChannelSessions"
func (svc *FibreChannelSessionService) GetFibreChannelSessions(params *util.GetParams) ([]*model.FibreChannelSession, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelSession - method creates a "FibreChannelSession"
func (svc *FibreChannelSessionService) CreateFibreChannelSession(obj *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateFibreChannelSession - method modifies  the "FibreChannelSession" 
func (svc *FibreChannelSessionService) UpdateFibreChannelSession(id string, obj *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetFibreChannelSessionById - method returns a pointer to "FibreChannelSession"
func (svc *FibreChannelSessionService) GetFibreChannelSessionById(id string) (*model.FibreChannelSession, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteFibreChannelSession - deletes the "FibreChannelSession"
func (svc *FibreChannelSessionService) DeleteFibreChannelSession(id string) error {
	return svc.objectSet.DeleteObject(id)
}
