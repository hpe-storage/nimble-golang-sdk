// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manage replication partner. Replication partners let one storage array talk to another for replication purposes. The two arrays must be able to communicate over a network, and
// use ports 4213 and 4214. Replication partners have the same name as the remote group. Replication partners can be reciprocal, upstream (the source of replicas), or downstream
// (the receiver of replicas) partners.
const (
    replicationPartnerPath = "replication_partners"
)

// ReplicationPartnerObjectSet
type ReplicationPartnerObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ReplicationPartner object
func (objectSet *ReplicationPartnerObjectSet) CreateObject(payload *model.ReplicationPartner) (*model.ReplicationPartner, error) {
	replicationPartnerObjectSetResp, err := objectSet.Client.Post(replicationPartnerPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if replicationPartnerObjectSetResp == nil {
		return nil,nil
	}
	return replicationPartnerObjectSetResp.(*model.ReplicationPartner), err
}

// UpdateObject Modify existing ReplicationPartner object
func (objectSet *ReplicationPartnerObjectSet) UpdateObject(id string, payload *model.ReplicationPartner) (*model.ReplicationPartner, error) {
	replicationPartnerObjectSetResp, err := objectSet.Client.Put(replicationPartnerPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if replicationPartnerObjectSetResp == nil {
		return nil,nil
	}
	return replicationPartnerObjectSetResp.(*model.ReplicationPartner), err
}

// DeleteObject deletes the ReplicationPartner object with the specified ID
func (objectSet *ReplicationPartnerObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(replicationPartnerPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a ReplicationPartner object with the given ID
func (objectSet *ReplicationPartnerObjectSet) GetObject(id string) (*model.ReplicationPartner, error) {
	replicationPartnerObjectSetResp, err := objectSet.Client.Get(replicationPartnerPath, id, model.ReplicationPartner{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if replicationPartnerObjectSetResp == nil {
		return nil,nil
	}
	return replicationPartnerObjectSetResp.(*model.ReplicationPartner), err
}

// GetObjectList returns the list of ReplicationPartner objects
func (objectSet *ReplicationPartnerObjectSet) GetObjectList() ([]*model.ReplicationPartner, error) {
	replicationPartnerObjectSetResp, err := objectSet.Client.List(replicationPartnerPath)
	if err != nil {
		return nil, err
	}
	return buildReplicationPartnerObjectSet(replicationPartnerObjectSetResp), err
}

// GetObjectListFromParams returns the list of ReplicationPartner objects using the given params query info
func (objectSet *ReplicationPartnerObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ReplicationPartner, error) {
	replicationPartnerObjectSetResp, err := objectSet.Client.ListFromParams(replicationPartnerPath, params)
	if err != nil {
		return nil, err
	}
	return buildReplicationPartnerObjectSet(replicationPartnerObjectSetResp), err
}

// generated function to build the appropriate response types
func buildReplicationPartnerObjectSet(response interface{}) ([]*model.ReplicationPartner) {
	values := reflect.ValueOf(response)
	results := make([]*model.ReplicationPartner, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ReplicationPartner{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}