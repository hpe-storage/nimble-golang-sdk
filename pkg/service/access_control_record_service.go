// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// AccessControlRecord Service - Manage access control records for volumes.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateAccessControlRecord - method creates a "AccessControlRecord"
func (svc *AccessControlRecordService) CreateAccessControlRecord(obj *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateAccessControlRecord - method modifies  the "AccessControlRecord" 
func (svc *AccessControlRecordService) UpdateAccessControlRecord(id string, obj *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetAccessControlRecordById - method returns a pointer to "AccessControlRecord"
func (svc *AccessControlRecordService) GetAccessControlRecordById(id string) (*model.AccessControlRecord, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteAccessControlRecord - deletes the "AccessControlRecord"
func (svc *AccessControlRecordService) DeleteAccessControlRecord(id string) error {
	return svc.objectSet.DeleteObject(id)
}
