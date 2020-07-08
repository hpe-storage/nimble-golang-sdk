// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// FibreChannelSession Service - Fibre Channel session is created when Fibre Channel initiator connects to this group.

import (
	"fmt"
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
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	fibreChannelSessionResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return fibreChannelSessionResp,nil
}

// CreateFibreChannelSession - method creates a "FibreChannelSession"
func (svc *FibreChannelSessionService) CreateFibreChannelSession(obj *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	fibreChannelSessionResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return fibreChannelSessionResp,nil
}

// UpdateFibreChannelSession - method modifies  the "FibreChannelSession" 
func (svc *FibreChannelSessionService) UpdateFibreChannelSession(id string, obj *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	fibreChannelSessionResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return fibreChannelSessionResp,nil
}

// GetFibreChannelSessionById - method returns a pointer to "FibreChannelSession"
func (svc *FibreChannelSessionService) GetFibreChannelSessionById(id string) (*model.FibreChannelSession, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	fibreChannelSessionResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return fibreChannelSessionResp,nil
}


// DeleteFibreChannelSession - deletes the "FibreChannelSession"
func (svc *FibreChannelSessionService) DeleteFibreChannelSession(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
