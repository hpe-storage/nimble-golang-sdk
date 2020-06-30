// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ObjectLimit Service - List the maximum limits and warning thresholds for number of objects in the storage group.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// ObjectLimitService type 
type ObjectLimitService struct {
	objectSet *client.ObjectLimitObjectSet
}

// NewObjectLimitService - method to initialize "ObjectLimitService" 
func NewObjectLimitService(gs *NsGroupService) (*ObjectLimitService) {
	objectSet := gs.client.GetObjectLimitObjectSet()
	return &ObjectLimitService{objectSet: objectSet}
}

// GetObjectLimits - method returns a array of pointers of type "ObjectLimits"
func (svc *ObjectLimitService) GetObjectLimits(params *util.GetParams) ([]*model.ObjectLimit, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetObjectLimitsWithFields - method returns a array of pointers of type "ObjectLimit" 
func (svc *ObjectLimitService) GetObjectLimitsWithFields(fields []string) ([]*model.ObjectLimit, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateObjectLimit - method creates a "ObjectLimit"
func (svc *ObjectLimitService) CreateObjectLimit(obj *model.ObjectLimit) (*model.ObjectLimit, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditObjectLimit - method modifies  the "ObjectLimit" 
func (svc *ObjectLimitService) EditObjectLimit(id string, obj *model.ObjectLimit) (*model.ObjectLimit, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyObjectLimit - private method for more than one element check. 
func onlyObjectLimit(objs []*model.ObjectLimit) (*model.ObjectLimit, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ObjectLimit found with the given filter")
	}

	return objs[0], nil
}

 
// GetObjectLimitsByID - method returns associative a array of pointers of type "ObjectLimit", filter by Id
func (svc *ObjectLimitService) GetObjectLimitsByID(pool *model.Pool, fields []string) (map[string]*model.ObjectLimit, error) {
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
	objs, err := svc.GetObjectLimits(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ObjectLimit)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetObjectLimitById - method returns a pointer to "ObjectLimit"
func (svc *ObjectLimitService) GetObjectLimitById(id string) (*model.ObjectLimit, error) {
	return svc.objectSet.GetObject(id)
}


