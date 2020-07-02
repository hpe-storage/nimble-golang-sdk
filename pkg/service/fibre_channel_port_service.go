// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelPort Service - Fibre Channel ports provide data access. This API provides the list of all Fibre Channel ports configured on the arrays.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// FibreChannelPortService type 
type FibreChannelPortService struct {
	objectSet *client.FibreChannelPortObjectSet
}

// NewFibreChannelPortService - method to initialize "FibreChannelPortService" 
func NewFibreChannelPortService(gs *NsGroupService) (*FibreChannelPortService) {
	objectSet := gs.client.GetFibreChannelPortObjectSet()
	return &FibreChannelPortService{objectSet: objectSet}
}

// GetFibreChannelPorts - method returns a array of pointers of type "FibreChannelPorts"
func (svc *FibreChannelPortService) GetFibreChannelPorts(params *util.GetParams) ([]*model.FibreChannelPort, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelPort - method creates a "FibreChannelPort"
func (svc *FibreChannelPortService) CreateFibreChannelPort(obj *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateFibreChannelPort - method modifies  the "FibreChannelPort" 
func (svc *FibreChannelPortService) UpdateFibreChannelPort(id string, obj *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetFibreChannelPortById - method returns a pointer to "FibreChannelPort"
func (svc *FibreChannelPortService) GetFibreChannelPortById(id string) (*model.FibreChannelPort, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteFibreChannelPort - deletes the "FibreChannelPort"
func (svc *FibreChannelPortService) DeleteFibreChannelPort(id string) error {
	return svc.objectSet.DeleteObject(id)
}
