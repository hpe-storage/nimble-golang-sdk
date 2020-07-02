// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtectionSchedule Service - Manage protection schedules used in protection templates.

import (
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
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateProtectionSchedule - method creates a "ProtectionSchedule"
func (svc *ProtectionScheduleService) CreateProtectionSchedule(obj *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateProtectionSchedule - method modifies  the "ProtectionSchedule" 
func (svc *ProtectionScheduleService) UpdateProtectionSchedule(id string, obj *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetProtectionScheduleById - method returns a pointer to "ProtectionSchedule"
func (svc *ProtectionScheduleService) GetProtectionScheduleById(id string) (*model.ProtectionSchedule, error) {
	return svc.objectSet.GetObject(id)
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
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteProtectionSchedule - deletes the "ProtectionSchedule"
func (svc *ProtectionScheduleService) DeleteProtectionSchedule(id string) error {
	return svc.objectSet.DeleteObject(id)
}
