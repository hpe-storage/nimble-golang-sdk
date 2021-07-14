// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelInitiatorAlias Service - This API provides the alias information for Fibre Channel initiators.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliass(params *param.GetParams) ([]*nimbleos.FibreChannelInitiatorAlias, error) {
    fibreChannelInitiatorAliasResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return fibreChannelInitiatorAliasResp,nil
}

// CreateFibreChannelInitiatorAlias - method creates a "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) CreateFibreChannelInitiatorAlias(obj *nimbleos.FibreChannelInitiatorAlias) (*nimbleos.FibreChannelInitiatorAlias, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateFibreChannelInitiatorAlias: invalid parameter specified, %v",obj)
    }

    fibreChannelInitiatorAliasResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return fibreChannelInitiatorAliasResp,nil
}

// UpdateFibreChannelInitiatorAlias - method modifies  the "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) UpdateFibreChannelInitiatorAlias(id string, obj *nimbleos.FibreChannelInitiatorAlias) (*nimbleos.FibreChannelInitiatorAlias, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateFibreChannelInitiatorAlias: invalid parameter specified, %v",obj)
    }

    fibreChannelInitiatorAliasResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return fibreChannelInitiatorAliasResp,nil
}

// GetFibreChannelInitiatorAliasById - method returns a pointer to "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) GetFibreChannelInitiatorAliasById(id string) (*nimbleos.FibreChannelInitiatorAlias, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetFibreChannelInitiatorAliasById: invalid parameter specified, %v",id)
    }

    fibreChannelInitiatorAliasResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return fibreChannelInitiatorAliasResp,nil
}

// DeleteFibreChannelInitiatorAlias - deletes the "FibreChannelInitiatorAlias"
func (svc *FibreChannelInitiatorAliasService) DeleteFibreChannelInitiatorAlias(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteFibreChannelInitiatorAlias: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

