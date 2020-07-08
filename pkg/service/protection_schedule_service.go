// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtectionSchedule Service - Manage protection schedules used in protection templates.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ProtectionScheduleService type 
type ProtectionScheduleService struct {
	objectSet *client.ProtectionScheduleObjectSet
}

// NewProtectionScheduleService - method to initialize "ProtectionScheduleService" 
func NewProtectionScheduleService(gs *NsGroupService) (*ProtectionScheduleService) {
	objectSet := gs.client.GetProtectionScheduleObjectSet()
	return &ProtectionScheduleService{objectSet: objectSet}
}

// GetProtectionSchedules - method returns a array of pointers of type "ProtectionSchedules"
func (svc *ProtectionScheduleService) GetProtectionSchedules(params *util.GetParams) ([]*model.ProtectionSchedule, error) {
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	protectionScheduleResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return protectionScheduleResp,nil
}

// CreateProtectionSchedule - method creates a "ProtectionSchedule"
func (svc *ProtectionScheduleService) CreateProtectionSchedule(obj *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	protectionScheduleResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return protectionScheduleResp,nil
}

// UpdateProtectionSchedule - method modifies  the "ProtectionSchedule" 
func (svc *ProtectionScheduleService) UpdateProtectionSchedule(id string, obj *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	protectionScheduleResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return protectionScheduleResp,nil
}

// GetProtectionScheduleById - method returns a pointer to "ProtectionSchedule"
func (svc *ProtectionScheduleService) GetProtectionScheduleById(id string) (*model.ProtectionSchedule, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	protectionScheduleResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return protectionScheduleResp,nil
}

// GetProtectionScheduleByName - method returns a pointer "ProtectionSchedule" 
func (svc *ProtectionScheduleService) GetProtectionScheduleByName(name string) (*model.ProtectionSchedule, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	protectionScheduleResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(protectionScheduleResp) == 0 {
    	return nil, nil
    }
    
	return protectionScheduleResp[0],nil
}	

// DeleteProtectionSchedule - deletes the "ProtectionSchedule"
func (svc *ProtectionScheduleService) DeleteProtectionSchedule(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
