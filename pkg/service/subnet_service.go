// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Subnet Service - Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets
// let you create logical addressing for selective routing.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// SubnetService type 
type SubnetService struct {
	objectSet *client.SubnetObjectSet
}

// NewSubnetService - method to initialize "SubnetService" 
func NewSubnetService(gs *NsGroupService) (*SubnetService) {
	objectSet := gs.client.GetSubnetObjectSet()
	return &SubnetService{objectSet: objectSet}
}

// GetSubnets - method returns a array of pointers of type "Subnets"
func (svc *SubnetService) GetSubnets(params *util.GetParams) ([]*model.Subnet, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetSubnetsWithFields - method returns a array of pointers of type "Subnet" 
func (svc *SubnetService) GetSubnetsWithFields(fields []string) ([]*model.Subnet, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSubnet - method creates a "Subnet"
func (svc *SubnetService) CreateSubnet(obj *model.Subnet) (*model.Subnet, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSubnet - method modifies  the "Subnet" 
func (svc *SubnetService) EditSubnet(id string, obj *model.Subnet) (*model.Subnet, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySubnet - private method for more than one element check. 
func onlySubnet(objs []*model.Subnet) (*model.Subnet, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Subnet found with the given filter")
	}

	return objs[0], nil
}

 
// GetSubnetsByID - method returns associative a array of pointers of type "Subnet", filter by Id
func (svc *SubnetService) GetSubnetsByID(pool *model.Pool, fields []string) (map[string]*model.Subnet, error) {
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
	objs, err := svc.GetSubnets(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Subnet)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetSubnetById - method returns a pointer to "Subnet"
func (svc *SubnetService) GetSubnetById(id string) (*model.Subnet, error) {
	return svc.objectSet.GetObject(id)
}

// GetSubnetsByName - method returns a associative array of pointers of type "Subnet", filter by name 
func (svc *SubnetService) GetSubnetsByName(pool *model.Pool, fields []string) (map[string]*model.Subnet, error) {
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
	objs, err := svc.GetSubnets(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Subnet)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetSubnetByName - method returns a pointer "Subnet" 
func (svc *SubnetService) GetSubnetByName(name string) (*model.Subnet, error) {
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
	return onlySubnet(objs)
}	

