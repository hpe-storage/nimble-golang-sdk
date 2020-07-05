// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtocolEndpoint Service - Protocol endpoints are administrative logical units (LUs) in an LU conglomerate to be used with VMware Virtual Volumes.

import (
	"fmt"
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
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	protocolEndpointResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return protocolEndpointResp,nil
}

// CreateProtocolEndpoint - method creates a "ProtocolEndpoint"
func (svc *ProtocolEndpointService) CreateProtocolEndpoint(obj *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	protocolEndpointResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return protocolEndpointResp,nil
}

// UpdateProtocolEndpoint - method modifies  the "ProtocolEndpoint" 
func (svc *ProtocolEndpointService) UpdateProtocolEndpoint(id string, obj *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	protocolEndpointResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return protocolEndpointResp,nil
}

// GetProtocolEndpointById - method returns a pointer to "ProtocolEndpoint"
func (svc *ProtocolEndpointService) GetProtocolEndpointById(id string) (*model.ProtocolEndpoint, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	protocolEndpointResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return protocolEndpointResp,nil
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
	protocolEndpointResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(protocolEndpointResp) == 0 {
    	return nil, nil
    }
    
	return protocolEndpointResp[0],nil
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
	protocolEndpointResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	if len(protocolEndpointResp) == 0 {
    	return nil, nil
    }
    
	return protocolEndpointResp[0],nil
}

// DeleteProtocolEndpoint - deletes the "ProtocolEndpoint"
func (svc *ProtocolEndpointService) DeleteProtocolEndpoint(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
