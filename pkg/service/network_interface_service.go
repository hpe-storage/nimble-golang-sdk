// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// NetworkInterface Service - Manage per array network interface configuration.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// NetworkInterfaceService type 
type NetworkInterfaceService struct {
	objectSet *client.NetworkInterfaceObjectSet
}

// NewNetworkInterfaceService - method to initialize "NetworkInterfaceService" 
func NewNetworkInterfaceService(gs *NsGroupService) (*NetworkInterfaceService) {
	objectSet := gs.client.GetNetworkInterfaceObjectSet()
	return &NetworkInterfaceService{objectSet: objectSet}
}

// GetNetworkInterfaces - method returns a array of pointers of type "NetworkInterfaces"
func (svc *NetworkInterfaceService) GetNetworkInterfaces(params *util.GetParams) ([]*model.NetworkInterface, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetNetworkInterfacesWithFields - method returns a array of pointers of type "NetworkInterface" 
func (svc *NetworkInterfaceService) GetNetworkInterfacesWithFields(fields []string) ([]*model.NetworkInterface, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateNetworkInterface - method creates a "NetworkInterface"
func (svc *NetworkInterfaceService) CreateNetworkInterface(obj *model.NetworkInterface) (*model.NetworkInterface, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditNetworkInterface - method modifies  the "NetworkInterface" 
func (svc *NetworkInterfaceService) EditNetworkInterface(id string, obj *model.NetworkInterface) (*model.NetworkInterface, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyNetworkInterface - private method for more than one element check. 
func onlyNetworkInterface(objs []*model.NetworkInterface) (*model.NetworkInterface, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one NetworkInterface found with the given filter")
	}

	return objs[0], nil
}

 
// GetNetworkInterfacesByID - method returns associative a array of pointers of type "NetworkInterface", filter by Id
func (svc *NetworkInterfaceService) GetNetworkInterfacesByID(pool *model.Pool, fields []string) (map[string]*model.NetworkInterface, error) {
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
	objs, err := svc.GetNetworkInterfaces(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.NetworkInterface)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetNetworkInterfaceById - method returns a pointer to "NetworkInterface"
func (svc *NetworkInterfaceService) GetNetworkInterfaceById(id string) (*model.NetworkInterface, error) {
	return svc.objectSet.GetObject(id)
}

// GetNetworkInterfacesByName - method returns a associative array of pointers of type "NetworkInterface", filter by name 
func (svc *NetworkInterfaceService) GetNetworkInterfacesByName(pool *model.Pool, fields []string) (map[string]*model.NetworkInterface, error) {
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
	objs, err := svc.GetNetworkInterfaces(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.NetworkInterface)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetNetworkInterfaceByName - method returns a pointer "NetworkInterface" 
func (svc *NetworkInterfaceService) GetNetworkInterfaceByName(name string) (*model.NetworkInterface, error) {
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
	return onlyNetworkInterface(objs)
}	

