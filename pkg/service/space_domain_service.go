// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SpaceDomain Service - A space domain is created for each application category and block size for each each pool.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// SpaceDomainService type 
type SpaceDomainService struct {
	objectSet *client.SpaceDomainObjectSet
}

// NewSpaceDomainService - method to initialize "SpaceDomainService" 
func NewSpaceDomainService(gs *NsGroupService) (*SpaceDomainService) {
	objectSet := gs.client.GetSpaceDomainObjectSet()
	return &SpaceDomainService{objectSet: objectSet}
}

// GetSpaceDomains - method returns a array of pointers of type "SpaceDomains"
func (svc *SpaceDomainService) GetSpaceDomains(params *util.GetParams) ([]*model.SpaceDomain, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSpaceDomain - method creates a "SpaceDomain"
func (svc *SpaceDomainService) CreateSpaceDomain(obj *model.SpaceDomain) (*model.SpaceDomain, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateSpaceDomain - method modifies  the "SpaceDomain" 
func (svc *SpaceDomainService) UpdateSpaceDomain(id string, obj *model.SpaceDomain) (*model.SpaceDomain, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetSpaceDomainById - method returns a pointer to "SpaceDomain"
func (svc *SpaceDomainService) GetSpaceDomainById(id string) (*model.SpaceDomain, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteSpaceDomain - deletes the "SpaceDomain"
func (svc *SpaceDomainService) DeleteSpaceDomain(id string) error {
	return svc.objectSet.DeleteObject(id)
}
