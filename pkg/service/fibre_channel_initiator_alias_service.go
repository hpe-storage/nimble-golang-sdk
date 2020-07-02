// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelInitiatorAlias Service - This API provides the alias information for Fibre Channel initiators.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// FibreChannelInitiatorAliasService type 
type FibreChannelInitiatorAliasService struct {
	objectSet *client.FibreChannelInitiatorAliasObjectSet
}

// NewFibreChannelInitiatorAliasService - method to initialize "FibreChannelInitiatorAliasService" 
func NewFibreChannelInitiatorAliasService(gs *NsGroupService) (*FibreChannelInitiatorAliasService) {
	objectSet := gs.client.GetFibreChannelInitiatorAliasObjectSet()
	return &FibreChannelInitiatorAliasService{objectSet: objectSet}
}

// GetFibreChannelInitiatorAliass - method returns a array of pointers of type "FibreChannelInitiatorAliass"
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliass(params *util.GetParams) ([]*model.FibreChannelInitiatorAlias, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelInitiatorAlias - method creates a "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) CreateFibreChannelInitiatorAlias(obj *model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateFibreChannelInitiatorAlias - method modifies  the "FibreChannelInitiatorAlias" 
func (svc *FibreChannelInitiatorAliasService) UpdateFibreChannelInitiatorAlias(id string, obj *model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetFibreChannelInitiatorAliasById - method returns a pointer to "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliasById(id string) (*model.FibreChannelInitiatorAlias, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteFibreChannelInitiatorAlias - deletes the "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) DeleteFibreChannelInitiatorAlias(id string) error {
	return svc.objectSet.DeleteObject(id)
}
