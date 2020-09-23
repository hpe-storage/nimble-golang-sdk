// Copyright 2020 Hewlett Packard Enterprise Development LP

// create true

// update true
package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
	"strings"
)

// Manage volume collections. Volume collections are logical groups of volumes that share protection characteristics such as snapshot and replication schedules. Volume collections
// can be created from scratch or based on predefined protection templates.
const (
	volumeCollectionPath = "volume_collections"
)

// VolumeCollectionObjectSet
type VolumeCollectionObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new VolumeCollection object
func (objectSet *VolumeCollectionObjectSet) CreateObject(payload *nimbleos.VolumeCollection) (*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.Post(volumeCollectionPath, payload, &nimbleos.VolumeCollection{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)
			}
		}
		return nil, err
	}

	return resp.(*nimbleos.VolumeCollection), err
}

// UpdateObject Modify existing VolumeCollection object
func (objectSet *VolumeCollectionObjectSet) UpdateObject(id string, payload *nimbleos.VolumeCollection) (*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.Put(volumeCollectionPath, id, payload, &nimbleos.VolumeCollection{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)

			}
		}
		return nil, err
	}

	return resp.(*nimbleos.VolumeCollection), err
}

// DeleteObject deletes the VolumeCollection object with the specified ID
func (objectSet *VolumeCollectionObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(volumeCollectionPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a VolumeCollection object with the given ID
func (objectSet *VolumeCollectionObjectSet) GetObject(id string) (*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.Get(volumeCollectionPath, id, &nimbleos.VolumeCollection{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.VolumeCollection), err
}

// GetObjectList returns the list of VolumeCollection objects
func (objectSet *VolumeCollectionObjectSet) GetObjectList() ([]*nimbleos.VolumeCollection, error) {
	resp, err := objectSet.Client.List(volumeCollectionPath)
	if err != nil {
		return nil, err
	}
	return buildVolumeCollectionObjectSet(resp), err
}

// GetObjectListFromParams returns the list of VolumeCollection objects using the given params query info
func (objectSet *VolumeCollectionObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.VolumeCollection, error) {
	volumeCollectionObjectSetResp, err := objectSet.Client.ListFromParams(volumeCollectionPath, params)
	if err != nil {
		return nil, err
	}
	return buildVolumeCollectionObjectSet(volumeCollectionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildVolumeCollectionObjectSet(response interface{}) []*nimbleos.VolumeCollection {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.VolumeCollection, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.VolumeCollection{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Promote - Take ownership of the specified volume collection. The volumes associated with the volume collection will be set to online and be available for reading and writing. Replication will be disabled on the affected schedules and must be re-configured if desired. Snapshot retention for the affected schedules will be set to the greater of the current local or replica retention values. This operation is not supported for synchronous replication volume collections.
func (objectSet *VolumeCollectionObjectSet) Promote(id *string) error {

	promoteUri := volumeCollectionPath
	promoteUri = promoteUri + "/" + *id
	promoteUri = promoteUri + "/actions/" + "promote"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(promoteUri, payload, &emptyStruct)
	return err
}

// Demote - Release ownership of the specified volume collection. The volumes associated with the volume collection will set to offline and a snapshot will be created, then full control over the volume collection will be transferred to the new owner. This option can be used following a promote to revert the volume collection back to its prior configured state. This operation does not alter the configuration on the new owner itself, but does require the new owner to be running in order to obtain its identity information. This operation is not supported for synchronous replication volume collections.
func (objectSet *VolumeCollectionObjectSet) Demote(id *string, replicationPartnerId *string, invokeOnUpstreamPartner *bool) error {

	demoteUri := volumeCollectionPath
	demoteUri = demoteUri + "/" + *id
	demoteUri = demoteUri + "/actions/" + "demote"

	payload := &struct {
		Id                      *string `json:"id,omitempty"`
		ReplicationPartnerId    *string `json:"replication_partner_id,omitempty"`
		InvokeOnUpstreamPartner *bool   `json:"invoke_on_upstream_partner,omitempty"`
	}{
		id,
		replicationPartnerId,
		invokeOnUpstreamPartner,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(demoteUri, payload, &emptyStruct)
	return err
}

// Handover - Gracefully transfer ownership of the specified volume collection. This action can be used to pass control of the volume collection to the downstream replication partner. Ownership and full control over the volume collection will be given to the downstream replication partner. The volumes associated with the volume collection will be set to offline prior to the final snapshot being taken and replicated, thus ensuring full data synchronization as part of the transfer. By default, the new owner will automatically begin replicating the volume collection back to this node when the handover completes.
func (objectSet *VolumeCollectionObjectSet) Handover(id *string, replicationPartnerId *string, noReverse *bool, invokeOnUpstreamPartner *bool, overrideUpstreamDown *bool) error {

	handoverUri := volumeCollectionPath
	handoverUri = handoverUri + "/" + *id
	handoverUri = handoverUri + "/actions/" + "handover"

	payload := &struct {
		Id                      *string `json:"id,omitempty"`
		ReplicationPartnerId    *string `json:"replication_partner_id,omitempty"`
		NoReverse               *bool   `json:"no_reverse,omitempty"`
		InvokeOnUpstreamPartner *bool   `json:"invoke_on_upstream_partner,omitempty"`
		OverrideUpstreamDown    *bool   `json:"override_upstream_down,omitempty"`
	}{
		id,
		replicationPartnerId,
		noReverse,
		invokeOnUpstreamPartner,
		overrideUpstreamDown,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(handoverUri, payload, &emptyStruct)
	return err
}

// AbortHandover - Abort in-progress handover. If for some reason a previously invoked handover request is unable to complete, this action can be used to cancel it. This operation is not supported for synchronous replication volume collections.
func (objectSet *VolumeCollectionObjectSet) AbortHandover(id *string) error {

	abortHandoverUri := volumeCollectionPath
	abortHandoverUri = abortHandoverUri + "/" + *id
	abortHandoverUri = abortHandoverUri + "/actions/" + "abort_handover"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(abortHandoverUri, payload, &emptyStruct)
	return err
}

// Validate - Validate a volume collection with either Microsoft VSS or VMware application synchronization.
func (objectSet *VolumeCollectionObjectSet) Validate(id *string) (*nimbleos.NsAppServerResp, error) {

	validateUri := volumeCollectionPath
	validateUri = validateUri + "/" + *id
	validateUri = validateUri + "/actions/" + "validate"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	resp, err := objectSet.Client.Post(validateUri, payload, &nimbleos.NsAppServerResp{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsAppServerResp), err
}
