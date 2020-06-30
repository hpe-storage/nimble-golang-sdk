// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Version Service - Show the API version.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// VersionService type 
type VersionService struct {
	objectSet *client.VersionObjectSet
}

// NewVersionService - method to initialize "VersionService" 
func NewVersionService(gs *NsGroupService) (*VersionService) {
	objectSet := gs.client.GetVersionObjectSet()
	return &VersionService{objectSet: objectSet}
}

// GetVersions - method returns a array of pointers of type "Versions"
func (svc *VersionService) GetVersions(params *util.GetParams) ([]*model.Version, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetVersionsWithFields - method returns a array of pointers of type "Version" 
func (svc *VersionService) GetVersionsWithFields(fields []string) ([]*model.Version, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateVersion - method creates a "Version"
func (svc *VersionService) CreateVersion(obj *model.Version) (*model.Version, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditVersion - method modifies  the "Version" 
func (svc *VersionService) EditVersion(id string, obj *model.Version) (*model.Version, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyVersion - private method for more than one element check. 
func onlyVersion(objs []*model.Version) (*model.Version, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Version found with the given filter")
	}

	return objs[0], nil
}

// GetVersionsByName - method returns a associative array of pointers of type "Version", filter by name 
func (svc *VersionService) GetVersionsByName(pool *model.Pool, fields []string) (map[string]*model.Version, error) {
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
	objs, err := svc.GetVersions(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Version)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetVersionByName - method returns a pointer "Version" 
func (svc *VersionService) GetVersionByName(name string) (*model.Version, error) {
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
	return onlyVersion(objs)
}	

