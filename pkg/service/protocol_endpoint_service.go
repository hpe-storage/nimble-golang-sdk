// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtocolEndpoint Service - Protocol endpoints are administrative logical units (LUs) in an LU conglomerate to be used with VMware Virtual Volumes.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ProtocolEndpointService type 
type ProtocolEndpointService struct {
	objectSet *client.ProtocolEndpointObjectSet
}

// NewProtocolEndpointService - method to initialize "ProtocolEndpointService" 
func NewProtocolEndpointService(gs *NsGroupService) (*ProtocolEndpointService) {
	objectSet := gs.client.GetProtocolEndpointObjectSet()
	return &ProtocolEndpointService{objectSet: objectSet}
}

// GetProtocolEndpoints - method returns a array of pointers of type "ProtocolEndpoints"
func (svc *ProtocolEndpointService) GetProtocolEndpoints(params *util.GetParams) ([]*model.ProtocolEndpoint, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateProtocolEndpoint - method creates a "ProtocolEndpoint"
func (svc *ProtocolEndpointService) CreateProtocolEndpoint(obj *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateProtocolEndpoint - method modifies  the "ProtocolEndpoint" 
func (svc *ProtocolEndpointService) UpdateProtocolEndpoint(id string, obj *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetProtocolEndpointById - method returns a pointer to "ProtocolEndpoint"
func (svc *ProtocolEndpointService) GetProtocolEndpointById(id string) (*model.ProtocolEndpoint, error) {
	return svc.objectSet.GetObject(id)
}

// GetProtocolEndpointByName - method returns a pointer "ProtocolEndpoint" 
func (svc *ProtocolEndpointService) GetProtocolEndpointByName(name string) (*model.ProtocolEndpoint, error) {
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
// GetProtocolEndpointBySerialNumber method returns a pointer to "ProtocolEndpoint"
func (svc *ProtocolEndpointService) GetProtocolEndpointBySerialNumber(serialNumber string) (*model.ProtocolEndpoint, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.ProtocolEndpointFields.SerialNumber,
			Operator:  util.EQUALS.String(),
			Value:     serialNumber,
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

// DeleteProtocolEndpoint - deletes the "ProtocolEndpoint"
func (svc *ProtocolEndpointService) DeleteProtocolEndpoint(id string) error {
	return svc.objectSet.DeleteObject(id)
}
