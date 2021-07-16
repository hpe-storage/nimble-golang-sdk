// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *ReplicationPartnerObjectSet) CreateObject(payload *nimbleos.ReplicationPartner) (*nimbleos.ReplicationPartner, error) {
	resp, err := objectSet.Client.Post(replicationPartnerPath, payload, &nimbleos.ReplicationPartner{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.ReplicationPartner), err
}

// UpdateObject Modify existing ReplicationPartner object
func (objectSet *ReplicationPartnerObjectSet) UpdateObject(id string, payload *nimbleos.ReplicationPartner) (*nimbleos.ReplicationPartner, error) {
	resp, err := objectSet.Client.Put(replicationPartnerPath, id, payload, &nimbleos.ReplicationPartner{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.ReplicationPartner), err
}

// DeleteObject deletes the ReplicationPartner object with the specified ID
func (objectSet *ReplicationPartnerObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(replicationPartnerPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a ReplicationPartner object with the given ID
func (objectSet *ReplicationPartnerObjectSet) GetObject(id string) (*nimbleos.ReplicationPartner, error) {
	resp, err := objectSet.Client.Get(replicationPartnerPath, id, &nimbleos.ReplicationPartner{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.ReplicationPartner), err
}

// GetObjectList returns the list of ReplicationPartner objects
func (objectSet *ReplicationPartnerObjectSet) GetObjectList() ([]*nimbleos.ReplicationPartner, error) {
	resp, err := objectSet.Client.List(replicationPartnerPath)
	if err != nil {
		return nil, err
	}
	return buildReplicationPartnerObjectSet(resp), err
}

// GetObjectListFromParams returns the list of ReplicationPartner objects using the given params query info
func (objectSet *ReplicationPartnerObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ReplicationPartner, error) {
	replicationPartnerObjectSetResp, err := objectSet.Client.ListFromParams(replicationPartnerPath, params)
	if err != nil {
		return nil, err
	}
	return buildReplicationPartnerObjectSet(replicationPartnerObjectSetResp), err
}

// generated function to build the appropriate response types
func buildReplicationPartnerObjectSet(response interface{}) []*nimbleos.ReplicationPartner {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.ReplicationPartner, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.ReplicationPartner{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Pause - Pause replication for the specified partner.
func (objectSet *ReplicationPartnerObjectSet) Pause(id *string) error {
	pauseUri := replicationPartnerPath
	pauseUri = pauseUri + "/" + *id
	pauseUri = pauseUri + "/actions/" + "pause"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	_, err := objectSet.Client.Post(pauseUri, payload, &nimbleos.ReplicationPartner{})
	return err
}

// Resume - Resume replication for the specified partner.
func (objectSet *ReplicationPartnerObjectSet) Resume(id *string) error {
	resumeUri := replicationPartnerPath
	resumeUri = resumeUri + "/" + *id
	resumeUri = resumeUri + "/actions/" + "resume"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	_, err := objectSet.Client.Post(resumeUri, payload, &nimbleos.ReplicationPartner{})
	return err
}

// Test - Test connectivity to the specified partner.
func (objectSet *ReplicationPartnerObjectSet) Test(id *string) error {
	testUri := replicationPartnerPath
	testUri = testUri + "/" + *id
	testUri = testUri + "/actions/" + "test"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	_, err := objectSet.Client.Post(testUri, payload, &nimbleos.ReplicationPartner{})
	return err
}
