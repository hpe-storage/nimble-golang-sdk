// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ProtocolEndpoint Service - Protocol endpoints are administrative logical units (LUs) in an LU conglomerate to be used with VMware Virtual Volumes.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetProtocolEndpointsWithFields - method returns a array of pointers of type "ProtocolEndpoint" 
func (svc *ProtocolEndpointService) GetProtocolEndpointsWithFields(fields []string) ([]*model.ProtocolEndpoint, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateProtocolEndpoint - method creates a "ProtocolEndpoint"
func (svc *ProtocolEndpointService) CreateProtocolEndpoint(obj *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditProtocolEndpoint - method modifies  the "ProtocolEndpoint" 
func (svc *ProtocolEndpointService) EditProtocolEndpoint(id string, obj *model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyProtocolEndpoint - private method for more than one element check. 
func onlyProtocolEndpoint(objs []*model.ProtocolEndpoint) (*model.ProtocolEndpoint, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ProtocolEndpoint found with the given filter")
	}

	return objs[0], nil
}

 
// GetProtocolEndpointsByID - method returns associative a array of pointers of type "ProtocolEndpoint", filter by Id
func (svc *ProtocolEndpointService) GetProtocolEndpointsByID(pool *model.Pool, fields []string) (map[string]*model.ProtocolEndpoint, error) {
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
	objs, err := svc.GetProtocolEndpoints(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ProtocolEndpoint)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetProtocolEndpointById - method returns a pointer to "ProtocolEndpoint"
func (svc *ProtocolEndpointService) GetProtocolEndpointById(id string) (*model.ProtocolEndpoint, error) {
	return svc.objectSet.GetObject(id)
}

// GetProtocolEndpointsByName - method returns a associative array of pointers of type "ProtocolEndpoint", filter by name 
func (svc *ProtocolEndpointService) GetProtocolEndpointsByName(pool *model.Pool, fields []string) (map[string]*model.ProtocolEndpoint, error) {
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
	objs, err := svc.GetProtocolEndpoints(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ProtocolEndpoint)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return onlyProtocolEndpoint(objs)
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
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return onlyProtocolEndpoint(objs)
}

