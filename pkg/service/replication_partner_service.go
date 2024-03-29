// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// ReplicationPartner Service - Manage replication partner. Replication partners let one storage array talk to another for replication purposes. The two arrays must be able to communicate over a network, and
// use ports 4213 and 4214. Replication partners have the same name as the remote group. Replication partners can be reciprocal, upstream (the source of replicas), or downstream
// (the receiver of replicas) partners.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ReplicationPartnerService type
type ReplicationPartnerService struct {
	objectSet *client.ReplicationPartnerObjectSet
}

// NewReplicationPartnerService - method to initialize "ReplicationPartnerService"
func NewReplicationPartnerService(gs *NsGroupService) *ReplicationPartnerService {
	objectSet := gs.client.GetReplicationPartnerObjectSet()
	return &ReplicationPartnerService{objectSet: objectSet}
}

// GetReplicationPartners - method returns a array of pointers of type "ReplicationPartners"
func (svc *ReplicationPartnerService) GetReplicationPartners(params *param.GetParams) ([]*nimbleos.ReplicationPartner, error) {
	replicationPartnerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return replicationPartnerResp, nil
}

// CreateReplicationPartner - method creates a "ReplicationPartner"
func (svc *ReplicationPartnerService) CreateReplicationPartner(obj *nimbleos.ReplicationPartner) (*nimbleos.ReplicationPartner, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateReplicationPartner: invalid parameter specified, %v", obj)
	}

	replicationPartnerResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return replicationPartnerResp, nil
}

// UpdateReplicationPartner - method modifies  the "ReplicationPartner"
func (svc *ReplicationPartnerService) UpdateReplicationPartner(id string, obj *nimbleos.ReplicationPartner) (*nimbleos.ReplicationPartner, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateReplicationPartner: invalid parameter specified, %v", obj)
	}

	replicationPartnerResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return replicationPartnerResp, nil
}

// GetReplicationPartnerById - method returns a pointer to "ReplicationPartner"
func (svc *ReplicationPartnerService) GetReplicationPartnerById(id string) (*nimbleos.ReplicationPartner, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetReplicationPartnerById: invalid parameter specified, %v", id)
	}

	replicationPartnerResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return replicationPartnerResp, nil
}

// GetReplicationPartnerByName - method returns a pointer "ReplicationPartner"
func (svc *ReplicationPartnerService) GetReplicationPartnerByName(name string) (*nimbleos.ReplicationPartner, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: &nimbleos.ReplicationPartnerFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	replicationPartnerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(replicationPartnerResp) == 0 {
		return nil, nil
	}

	return replicationPartnerResp[0], nil
}

// DeleteReplicationPartner - deletes the "ReplicationPartner"
func (svc *ReplicationPartnerService) DeleteReplicationPartner(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteReplicationPartner: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// PauseReplicationPartner - pause replication for the specified partner.
//   Required parameters:
//       id - ID of the partner to pause.

func (svc *ReplicationPartnerService) PauseReplicationPartner(id string) error {

	if len(id) == 0 {
		return fmt.Errorf("PauseReplicationPartner: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.Pause(&id)
	return err
}

// ResumeReplicationPartner - resume replication for the specified partner.
//   Required parameters:
//       id - ID of the partner to resume.

func (svc *ReplicationPartnerService) ResumeReplicationPartner(id string) error {

	if len(id) == 0 {
		return fmt.Errorf("ResumeReplicationPartner: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.Resume(&id)
	return err
}

// TestReplicationPartner - test connectivity to the specified partner.
//   Required parameters:
//       id - ID of the partner to test.

func (svc *ReplicationPartnerService) TestReplicationPartner(id string) error {

	if len(id) == 0 {
		return fmt.Errorf("TestReplicationPartner: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.Test(&id)
	return err
}
