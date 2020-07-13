// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Subnet Service - Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets
// let you create logical addressing for selective routing.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// SubnetService type
type SubnetService struct {
	objectSet *client.SubnetObjectSet
}

// NewSubnetService - method to initialize "SubnetService"
func NewSubnetService(gs *NsGroupService) *SubnetService {
	objectSet := gs.client.GetSubnetObjectSet()
	return &SubnetService{objectSet: objectSet}
}

// GetSubnets - method returns a array of pointers of type "Subnets"
func (svc *SubnetService) GetSubnets(params *param.GetParams) ([]*nimbleos.Subnet, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	subnetResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return subnetResp, nil
}

// CreateSubnet - method creates a "Subnet"
func (svc *SubnetService) CreateSubnet(obj *nimbleos.Subnet) (*nimbleos.Subnet, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	subnetResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return subnetResp, nil
}

// UpdateSubnet - method modifies  the "Subnet"
func (svc *SubnetService) UpdateSubnet(id string, obj *nimbleos.Subnet) (*nimbleos.Subnet, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	subnetResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return subnetResp, nil
}

// GetSubnetById - method returns a pointer to "Subnet"
func (svc *SubnetService) GetSubnetById(id string) (*nimbleos.Subnet, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	subnetResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return subnetResp, nil
}

// GetSubnetByName - method returns a pointer "Subnet"
func (svc *SubnetService) GetSubnetByName(name string) (*nimbleos.Subnet, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: &nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	subnetResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(subnetResp) == 0 {
		return nil, nil
	}

	return subnetResp[0], nil
}

// DeleteSubnet - deletes the "Subnet"
func (svc *SubnetService) DeleteSubnet(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
