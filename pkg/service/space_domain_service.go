// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// SpaceDomain Service - A space domain is created for each application category and block size for each each pool.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetSpaceDomainsWithFields - method returns a array of pointers of type "SpaceDomain" 
func (svc *SpaceDomainService) GetSpaceDomainsWithFields(fields []string) ([]*model.SpaceDomain, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateSpaceDomain - method creates a "SpaceDomain"
func (svc *SpaceDomainService) CreateSpaceDomain(obj *model.SpaceDomain) (*model.SpaceDomain, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditSpaceDomain - method modifies  the "SpaceDomain" 
func (svc *SpaceDomainService) EditSpaceDomain(id string, obj *model.SpaceDomain) (*model.SpaceDomain, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlySpaceDomain - private method for more than one element check. 
func onlySpaceDomain(objs []*model.SpaceDomain) (*model.SpaceDomain, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one SpaceDomain found with the given filter")
	}

	return objs[0], nil
}

 
// GetSpaceDomainsByID - method returns associative a array of pointers of type "SpaceDomain", filter by Id
func (svc *SpaceDomainService) GetSpaceDomainsByID(pool *model.Pool, fields []string) (map[string]*model.SpaceDomain, error) {
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
	objs, err := svc.GetSpaceDomains(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.SpaceDomain)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetSpaceDomainById - method returns a pointer to "SpaceDomain"
func (svc *SpaceDomainService) GetSpaceDomainById(id string) (*model.SpaceDomain, error) {
	return svc.objectSet.GetObject(id)
}


