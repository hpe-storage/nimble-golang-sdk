// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelInterface Service - Represent information of specified Fibre Channel interfaces. Fibre Channel interfaces are hosted on Fibre Channel ports to provide data access.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// FibreChannelInterfaceService type 
type FibreChannelInterfaceService struct {
	objectSet *client.FibreChannelInterfaceObjectSet
}

// NewFibreChannelInterfaceService - method to initialize "FibreChannelInterfaceService" 
func NewFibreChannelInterfaceService(gs *NsGroupService) (*FibreChannelInterfaceService) {
	objectSet := gs.client.GetFibreChannelInterfaceObjectSet()
	return &FibreChannelInterfaceService{objectSet: objectSet}
}

// GetFibreChannelInterfaces - method returns a array of pointers of type "FibreChannelInterfaces"
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaces(params *util.GetParams) ([]*model.FibreChannelInterface, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateFibreChannelInterface - method creates a "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) CreateFibreChannelInterface(obj *model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateFibreChannelInterface - method modifies  the "FibreChannelInterface" 
func (svc *FibreChannelInterfaceService) UpdateFibreChannelInterface(id string, obj *model.FibreChannelInterface) (*model.FibreChannelInterface, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetFibreChannelInterfaceById - method returns a pointer to "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaceById(id string) (*model.FibreChannelInterface, error) {
	return svc.objectSet.GetObject(id)
}

// GetFibreChannelInterfaceByName - method returns a pointer "FibreChannelInterface" 
func (svc *FibreChannelInterfaceService) GetFibreChannelInterfaceByName(name string) (*model.FibreChannelInterface, error) {
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

// DeleteFibreChannelInterface - deletes the "FibreChannelInterface"
func (svc *FibreChannelInterfaceService) DeleteFibreChannelInterface(id string) error {
	return svc.objectSet.DeleteObject(id)
}
