// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// AccessControlRecord Service - Manage access control records for volumes.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// AccessControlRecordService type 
type AccessControlRecordService struct {
	objectSet *client.AccessControlRecordObjectSet
}

// NewAccessControlRecordService - method to initialize "AccessControlRecordService" 
func NewAccessControlRecordService(gs *NsGroupService) (*AccessControlRecordService) {
	objectSet := gs.client.GetAccessControlRecordObjectSet()
	return &AccessControlRecordService{objectSet: objectSet}
}

// GetAccessControlRecords - method returns a array of pointers of type "AccessControlRecords"
func (svc *AccessControlRecordService) GetAccessControlRecords(params *util.GetParams) ([]*model.AccessControlRecord, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetAccessControlRecordsWithFields - method returns a array of pointers of type "AccessControlRecord" 
func (svc *AccessControlRecordService) GetAccessControlRecordsWithFields(fields []string) ([]*model.AccessControlRecord, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateAccessControlRecord - method creates a "AccessControlRecord"
func (svc *AccessControlRecordService) CreateAccessControlRecord(obj *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditAccessControlRecord - method modifies  the "AccessControlRecord" 
func (svc *AccessControlRecordService) EditAccessControlRecord(id string, obj *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyAccessControlRecord - private method for more than one element check. 
func onlyAccessControlRecord(objs []*model.AccessControlRecord) (*model.AccessControlRecord, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one AccessControlRecord found with the given filter")
	}

	return objs[0], nil
}

 
// GetAccessControlRecordsByID - method returns associative a array of pointers of type "AccessControlRecord", filter by Id
func (svc *AccessControlRecordService) GetAccessControlRecordsByID(pool *model.Pool, fields []string) (map[string]*model.AccessControlRecord, error) {
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
	objs, err := svc.GetAccessControlRecords(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.AccessControlRecord)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetAccessControlRecordById - method returns a pointer to "AccessControlRecord"
func (svc *AccessControlRecordService) GetAccessControlRecordById(id string) (*model.AccessControlRecord, error) {
	return svc.objectSet.GetObject(id)
}


