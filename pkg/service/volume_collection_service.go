// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// VolumeCollection Service - Manage volume collections. Volume collections are logical groups of volumes that share protection characteristics such as snapshot and replication schedules. Volume collections
// can be created from scratch or based on predefined protection templates.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// VolumeCollectionService type
type VolumeCollectionService struct {
	objectSet *client.VolumeCollectionObjectSet
}

// NewVolumeCollectionService - method to initialize "VolumeCollectionService"
func NewVolumeCollectionService(gs *NsGroupService) *VolumeCollectionService {
	objectSet := gs.client.GetVolumeCollectionObjectSet()
	return &VolumeCollectionService{objectSet: objectSet}
}

// GetVolumeCollections - method returns a array of pointers of type "VolumeCollections"
func (svc *VolumeCollectionService) GetVolumeCollections(params *param.GetParams) ([]*nimbleos.VolumeCollection, error) {
	if params == nil {
		return nil, fmt.Errorf("GetVolumeCollections: invalid parameter specified, %v", params)
	}

	volumeCollectionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return volumeCollectionResp, nil
}

// CreateVolumeCollection - method creates a "VolumeCollection"
func (svc *VolumeCollectionService) CreateVolumeCollection(obj *nimbleos.VolumeCollection) (*nimbleos.VolumeCollection, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateVolumeCollection: invalid parameter specified, %v", obj)
	}

	volumeCollectionResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return volumeCollectionResp, nil
}

// UpdateVolumeCollection - method modifies  the "VolumeCollection"
func (svc *VolumeCollectionService) UpdateVolumeCollection(id string, obj *nimbleos.VolumeCollection) (*nimbleos.VolumeCollection, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateVolumeCollection: invalid parameter specified, %v", obj)
	}

	volumeCollectionResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return volumeCollectionResp, nil
}

// GetVolumeCollectionById - method returns a pointer to "VolumeCollection"
func (svc *VolumeCollectionService) GetVolumeCollectionById(id string) (*nimbleos.VolumeCollection, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetVolumeCollectionById: invalid parameter specified, %v", id)
	}

	volumeCollectionResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return volumeCollectionResp, nil
}

// GetVolumeCollectionByName - method returns a pointer "VolumeCollection"
func (svc *VolumeCollectionService) GetVolumeCollectionByName(name string) (*nimbleos.VolumeCollection, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	volumeCollectionResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(volumeCollectionResp) == 0 {
		return nil, nil
	}

	return volumeCollectionResp[0], nil
}

// DeleteVolumeCollection - deletes the "VolumeCollection"
func (svc *VolumeCollectionService) DeleteVolumeCollection(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteVolumeCollection: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// PromoteVolumeCollection - take ownership of the specified volume collection. The volumes associated with the volume collection will be set to online and be available for reading and writing. Replication will be disabled on the affected schedules and must be re-configured if desired. Snapshot retention for the affected schedules will be set to the greater of the current local or replica retention values. This operation is not supported for synchronous replication volume collections.
//   Required parameters:
//       id - ID of the promoted volume collection.

func (svc *VolumeCollectionService) PromoteVolumeCollection(id string) error {

	if len(id) == 0 {
		return fmt.Errorf("PromoteVolumeCollection: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.Promote(&id)
	return err
}

// DemoteVolumeCollection - release ownership of the specified volume collection. The volumes associated with the volume collection will set to offline and a snapshot will be created, then full control over the volume collection will be transferred to the new owner. This option can be used following a promote to revert the volume collection back to its prior configured state. This operation does not alter the configuration on the new owner itself, but does require the new owner to be running in order to obtain its identity information. This operation is not supported for synchronous replication volume collections.
//   Required parameters:
//       id - ID of the demoted volume collection.
//       replicationPartnerId - ID of the new owner. If invoke_on_upstream_partner is provided, utilize the ID of the current owner i.e. upstream replication partner.

//   Optional parameters:
//       invokeOnUpstreamPartner - Invoke demote request on upstream partner. Default: 'false'. This operation is not supported for synchronous replication volume vollections.

func (svc *VolumeCollectionService) DemoteVolumeCollection(id string, replicationPartnerId string, invokeOnUpstreamPartner *bool) error {

	if len(id) == 0 || len(replicationPartnerId) == 0 {
		return fmt.Errorf("DemoteVolumeCollection: invalid parameter specified id: %v, replicationPartnerId: %v ", id, replicationPartnerId)
	}

	err := svc.objectSet.Demote(&id, &replicationPartnerId, invokeOnUpstreamPartner)
	return err
}

// HandoverVolumeCollection - gracefully transfer ownership of the specified volume collection. This action can be used to pass control of the volume collection to the downstream replication partner. Ownership and full control over the volume collection will be given to the downstream replication partner. The volumes associated with the volume collection will be set to offline prior to the final snapshot being taken and replicated, thus ensuring full data synchronization as part of the transfer. By default, the new owner will automatically begin replicating the volume collection back to this node when the handover completes.
//   Required parameters:
//       id - ID of the volume collection be handed over to the downstream replication partner.
//       replicationPartnerId - ID of the new owner.

//   Optional parameters:
//       noReverse - Do not automatically reverse direction of replication. Using this argument will prevent the new owner from automatically replicating the volume collection to this node when the handover completes. The default behavior is to enable replication back to this node. Default: 'false'.
//       invokeOnUpstreamPartner - Invoke handover request on upstream partner. Default: 'false'. This operation is not supported for synchronous replication volume vollections.
//       overrideUpstreamDown - Allow the handover request to proceed even if upstream array is down. The default behavior is to return an error when upstream is down. This option is applicable for synchronous replication only. Default: 'false'.

func (svc *VolumeCollectionService) HandoverVolumeCollection(id string, replicationPartnerId string, noReverse *bool, invokeOnUpstreamPartner *bool, overrideUpstreamDown *bool) error {

	if len(id) == 0 || len(replicationPartnerId) == 0 {
		return fmt.Errorf("HandoverVolumeCollection: invalid parameter specified id: %v, replicationPartnerId: %v ", id, replicationPartnerId)
	}

	err := svc.objectSet.Handover(&id, &replicationPartnerId, noReverse, invokeOnUpstreamPartner, overrideUpstreamDown)
	return err
}

// AbortHandoverVolumeCollection - abort in-progress handover. If for some reason a previously invoked handover request is unable to complete, this action can be used to cancel it. This operation is not supported for synchronous replication volume collections.
//   Required parameters:
//       id - ID of the volume collection on which to abort handover.

func (svc *VolumeCollectionService) AbortHandoverVolumeCollection(id string) error {

	if len(id) == 0 {
		return fmt.Errorf("AbortHandoverVolumeCollection: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.AbortHandover(&id)
	return err
}

// ValidateVolumeCollection - validate a volume collection with either Microsoft VSS or VMware application synchronization.
//   Required parameters:
//       id - ID of the volume collection that is to be validated.

func (svc *VolumeCollectionService) ValidateVolumeCollection(id string) (*nimbleos.NsAppServerResp, error) {

	if len(id) == 0 {
		return nil, fmt.Errorf("ValidateVolumeCollection: invalid parameter specified id: %v ", id)
	}

	resp, err := svc.objectSet.Validate(&id)
	return resp, err
}
