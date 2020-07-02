// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ReplicationPartner Service - Manage replication partner. Replication partners let one storage array talk to another for replication purposes. The two arrays must be able to communicate over a network, and
// use ports 4213 and 4214. Replication partners have the same name as the remote group. Replication partners can be reciprocal, upstream (the source of replicas), or downstream
// (the receiver of replicas) partners.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateReplicationPartner - method creates a "ReplicationPartner"
func (svc *ReplicationPartnerService) CreateReplicationPartner(obj *model.ReplicationPartner) (*model.ReplicationPartner, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateReplicationPartner - method modifies  the "ReplicationPartner" 
func (svc *ReplicationPartnerService) UpdateReplicationPartner(id string, obj *model.ReplicationPartner) (*model.ReplicationPartner, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetReplicationPartnerById - method returns a pointer to "ReplicationPartner"
func (svc *ReplicationPartnerService) GetReplicationPartnerById(id string) (*model.ReplicationPartner, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteReplicationPartner - deletes the "ReplicationPartner"
func (svc *ReplicationPartnerService) DeleteReplicationPartner(id string) error {
	return svc.objectSet.DeleteObject(id)
}
