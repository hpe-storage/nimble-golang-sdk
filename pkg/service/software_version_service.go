// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SoftwareVersion Service - Show the software version.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// SoftwareVersionService type 
type SoftwareVersionService struct {
	objectSet *client.SoftwareVersionObjectSet
}

// NewSoftwareVersionService - method to initialize "SoftwareVersionService" 
func NewSoftwareVersionService(gs *NsGroupService) (*SoftwareVersionService) {
	objectSet := gs.client.GetSoftwareVersionObjectSet()
	return &SoftwareVersionService{objectSet: objectSet}
}

// GetSoftwareVersions - method returns a array of pointers of type "SoftwareVersions"
func (svc *SoftwareVersionService) GetSoftwareVersions(params *util.GetParams) ([]*model.SoftwareVersion, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetSoftwareVersionsWithFields - method returns a array of pointers of type "SoftwareVersion" 
func (svc *SoftwareVersionService) GetSoftwareVersionsWithFields(fields []string) ([]*model.SoftwareVersion, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSoftwareVersion - method creates a "SoftwareVersion"
func (svc *SoftwareVersionService) CreateSoftwareVersion(obj *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSoftwareVersion - method modifies  the "SoftwareVersion" 
func (svc *SoftwareVersionService) EditSoftwareVersion(id string, obj *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySoftwareVersion - private method for more than one element check. 
func onlySoftwareVersion(objs []*model.SoftwareVersion) (*model.SoftwareVersion, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one SoftwareVersion found with the given filter")
	}

	return objs[0], nil
}

// GetSoftwareVersionsByName - method returns a associative array of pointers of type "SoftwareVersion", filter by name 
func (svc *SoftwareVersionService) GetSoftwareVersionsByName(pool *model.Pool, fields []string) (map[string]*model.SoftwareVersion, error) {
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
	objs, err := svc.GetSoftwareVersions(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.SoftwareVersion)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetSoftwareVersionByName - method returns a pointer "SoftwareVersion" 
func (svc *SoftwareVersionService) GetSoftwareVersionByName(name string) (*model.SoftwareVersion, error) {
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
	return onlySoftwareVersion(objs)
}	

