// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Subnet Service - Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets
// let you create logical addressing for selective routing.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateSubnet - method creates a "Subnet"
func (svc *SubnetService) CreateSubnet(obj *model.Subnet) (*model.Subnet, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateSubnet - method modifies  the "Subnet" 
func (svc *SubnetService) UpdateSubnet(id string, obj *model.Subnet) (*model.Subnet, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetSubnetById - method returns a pointer to "Subnet"
func (svc *SubnetService) GetSubnetById(id string) (*model.Subnet, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteSubnet - deletes the "Subnet"
func (svc *SubnetService) DeleteSubnet(id string) error {
	return svc.objectSet.DeleteObject(id)
}
