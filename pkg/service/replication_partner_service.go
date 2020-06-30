// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ReplicationPartner Service - Manage replication partner. Replication partners let one storage array talk to another for replication purposes. The two arrays must be able to communicate over a network, and
// use ports 4213 and 4214. Replication partners have the same name as the remote group. Replication partners can be reciprocal, upstream (the source of replicas), or downstream
// (the receiver of replicas) partners.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// ReplicationPartnerService type 
type ReplicationPartnerService struct {
	objectSet *client.ReplicationPartnerObjectSet
}

// NewReplicationPartnerService - method to initialize "ReplicationPartnerService" 
func NewReplicationPartnerService(gs *NsGroupService) (*ReplicationPartnerService) {
	objectSet := gs.client.GetReplicationPartnerObjectSet()
	return &ReplicationPartnerService{objectSet: objectSet}
}

// GetReplicationPartners - method returns a array of pointers of type "ReplicationPartners"
func (svc *ReplicationPartnerService) GetReplicationPartners(params *util.GetParams) ([]*model.ReplicationPartner, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetReplicationPartnersWithFields - method returns a array of pointers of type "ReplicationPartner" 
func (svc *ReplicationPartnerService) GetReplicationPartnersWithFields(fields []string) ([]*model.ReplicationPartner, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateReplicationPartner - method creates a "ReplicationPartner"
func (svc *ReplicationPartnerService) CreateReplicationPartner(obj *model.ReplicationPartner) (*model.ReplicationPartner, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditReplicationPartner - method modifies  the "ReplicationPartner" 
func (svc *ReplicationPartnerService) EditReplicationPartner(id string, obj *model.ReplicationPartner) (*model.ReplicationPartner, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyReplicationPartner - private method for more than one element check. 
func onlyReplicationPartner(objs []*model.ReplicationPartner) (*model.ReplicationPartner, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ReplicationPartner found with the given filter")
	}

	return objs[0], nil
}

 
// GetReplicationPartnersByID - method returns associative a array of pointers of type "ReplicationPartner", filter by Id
func (svc *ReplicationPartnerService) GetReplicationPartnersByID(pool *model.Pool, fields []string) (map[string]*model.ReplicationPartner, error) {
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
	objs, err := svc.GetReplicationPartners(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ReplicationPartner)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetReplicationPartnerById - method returns a pointer to "ReplicationPartner"
func (svc *ReplicationPartnerService) GetReplicationPartnerById(id string) (*model.ReplicationPartner, error) {
	return svc.objectSet.GetObject(id)
}

// GetReplicationPartnersByName - method returns a associative array of pointers of type "ReplicationPartner", filter by name 
func (svc *ReplicationPartnerService) GetReplicationPartnersByName(pool *model.Pool, fields []string) (map[string]*model.ReplicationPartner, error) {
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
	objs, err := svc.GetReplicationPartners(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ReplicationPartner)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetReplicationPartnerByName - method returns a pointer "ReplicationPartner" 
func (svc *ReplicationPartnerService) GetReplicationPartnerByName(name string) (*model.ReplicationPartner, error) {
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
	return onlyReplicationPartner(objs)
}	

