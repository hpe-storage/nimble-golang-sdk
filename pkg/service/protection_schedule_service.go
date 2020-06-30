// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtectionSchedule Service - Manage protection schedules used in protection templates.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetProtectionSchedulesWithFields - method returns a array of pointers of type "ProtectionSchedule" 
func (svc *ProtectionScheduleService) GetProtectionSchedulesWithFields(fields []string) ([]*model.ProtectionSchedule, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateProtectionSchedule - method creates a "ProtectionSchedule"
func (svc *ProtectionScheduleService) CreateProtectionSchedule(obj *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditProtectionSchedule - method modifies  the "ProtectionSchedule" 
func (svc *ProtectionScheduleService) EditProtectionSchedule(id string, obj *model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyProtectionSchedule - private method for more than one element check. 
func onlyProtectionSchedule(objs []*model.ProtectionSchedule) (*model.ProtectionSchedule, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ProtectionSchedule found with the given filter")
	}

	return objs[0], nil
}

 
// GetProtectionSchedulesByID - method returns associative a array of pointers of type "ProtectionSchedule", filter by Id
func (svc *ProtectionScheduleService) GetProtectionSchedulesByID(pool *model.Pool, fields []string) (map[string]*model.ProtectionSchedule, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetProtectionSchedules(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ProtectionSchedule)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetProtectionScheduleById - method returns a pointer to "ProtectionSchedule"
func (svc *ProtectionScheduleService) GetProtectionScheduleById(id string) (*model.ProtectionSchedule, error) {
	return svc.objectSet.GetObject(id)
}

// GetProtectionSchedulesByName - method returns a associative array of pointers of type "ProtectionSchedule", filter by name 
func (svc *ProtectionScheduleService) GetProtectionSchedulesByName(pool *model.Pool, fields []string) (map[string]*model.ProtectionSchedule, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetProtectionSchedules(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ProtectionSchedule)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlyProtectionSchedule(objs)
}	

