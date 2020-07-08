// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SpaceDomain Service - A space domain is created for each application category and block size for each each pool.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// SpaceDomainService type
type SpaceDomainService struct {
	objectSet *client.SpaceDomainObjectSet
}

// NewSpaceDomainService - method to initialize "SpaceDomainService"
func NewSpaceDomainService(gs *NsGroupService) *SpaceDomainService {
	objectSet := gs.client.GetSpaceDomainObjectSet()
	return &SpaceDomainService{objectSet: objectSet}
}

// GetSpaceDomains - method returns a array of pointers of type "SpaceDomains"
func (svc *SpaceDomainService) GetSpaceDomains(params *util.GetParams) ([]*model.SpaceDomain, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	spaceDomainResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return spaceDomainResp, nil
}

// CreateSpaceDomain - method creates a "SpaceDomain"
func (svc *SpaceDomainService) CreateSpaceDomain(obj *model.SpaceDomain) (*model.SpaceDomain, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	spaceDomainResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return spaceDomainResp, nil
}

// UpdateSpaceDomain - method modifies  the "SpaceDomain"
func (svc *SpaceDomainService) UpdateSpaceDomain(id string, obj *model.SpaceDomain) (*model.SpaceDomain, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	spaceDomainResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return spaceDomainResp, nil
}

// GetSpaceDomainById - method returns a pointer to "SpaceDomain"
func (svc *SpaceDomainService) GetSpaceDomainById(id string) (*model.SpaceDomain, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	spaceDomainResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return spaceDomainResp, nil
}

// DeleteSpaceDomain - deletes the "SpaceDomain"
func (svc *SpaceDomainService) DeleteSpaceDomain(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
