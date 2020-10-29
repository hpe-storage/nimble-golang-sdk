// Copyright 2020 Hewlett Packard Enterprise Development LP

package fakeservice

// Volume Service - Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// VolumeService type
type VolumeService struct {
	// create a map
}

// NewVolumeService - method to initialize "VolumeService"
func NewVolumeService(gs *NsGroupService) *VolumeService {

	return &VolumeService{}
}

// GetVolumes - method returns a array of pointers of type "Volumes"
func (svc *VolumeService) GetVolumes(params *param.GetParams) ([]*nimbleos.Volume, error) {
	return nil, nil
}

// CreateVolume - method creates a "Volume"
func (svc *VolumeService) CreateVolume(obj *nimbleos.Volume) (*nimbleos.Volume, error) {
	return nil, nil
}

// UpdateVolume - method modifies  the "Volume"
func (svc *VolumeService) UpdateVolume(id string, obj *nimbleos.Volume) (*nimbleos.Volume, error) {
	return nil, nil
}

// GetVolumeById - method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeById(id string) (*nimbleos.Volume, error) {
	return nil, nil
}

// GetVolumeByName - method returns a pointer "Volume"
func (svc *VolumeService) GetVolumeByName(name string) (*nimbleos.Volume, error) {
	return nil, nil
}

// GetVolumeBySerialNumber method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeBySerialNumber(serialNumber string) (*nimbleos.Volume, error) {
	return nil, nil
}

//OnlineVolume - method makes the volume online
func (svc *VolumeService) OnlineVolume(id string, force bool) (*nimbleos.Volume, error) {
	return nil, nil
}

// OfflineVolume - makes the volume offline
func (svc *VolumeService) OfflineVolume(id string, force bool) (*nimbleos.Volume, error) {
	return nil, nil

}

// DeleteVolume - deletes the volume
func (svc *VolumeService) DeleteVolume(id string) error {
	return nil
}

// AssociateVolume - add a volume to a volume collection
func (svc *VolumeService) AssociateVolume(volId string, volcollId string) error {

	return nil
}

// DisassociateVolume - remove a volume from a volume collection
func (svc *VolumeService) DisassociateVolume(volId string) error {

	return nil

}

// RestoreVolume - restore volume data from a previous snapshot.
//   Required parameters:
//       id - ID of the restored volume.
//       baseSnapId - ID of the snapshot to which this the volume is restored.

func (svc *VolumeService) RestoreVolume(id string, baseSnapId string) error {

	return nil
}

// MoveVolume - move a volume and its related volumes to another pool. To change a single volume's folder assignment (while remaining in the same pool), use a volume update operation to change the folder_id attribute.
//   Required parameters:
//       id - ID of the volume to move.
//       destPoolId - ID of the destination pool or folder. Specify a pool ID when the volumes should not be in a folder; otherwise, specify a folder ID and the pool will be derived from the folder.

//   Optional parameters:
//       forceVvol - Forcibly move a Virtual Volume. Moving Virtual Volume is disruptive to the vCenter, hence it should only be done by the VASA Provider (VP). This flag should only be set by the VP when it calls this API.

func (svc *VolumeService) MoveVolume(id string, destPoolId string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error) {

	return nil, nil
}

// BulkMoveVolumes - move volumes and their related volumes to another pool. To change a single volume's folder assignment (while remaining in the same pool), use a volume update operation to change the folder_id attribute.
//   Required parameters:
//       volIds - IDs of the volumes to move.
//       destPoolId - ID of the destination pool or folder. Specify a pool ID when the volumes should not be in a folder; otherwise, specify a folder ID and the pool will be derived from the folder.

//   Optional parameters:
//       forceVvol - Forcibly move a Virtual Volume. Moving Virtual Volume is disruptive to the vCenter, hence it should only be done by the VASA Provider (VP). This flag should only be set by the VP when it calls this API.

func (svc *VolumeService) BulkMoveVolumes(volIds []*string, destPoolId string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error) {
	return nil, nil
}

// AbortMoveVolume - abort the in-progress move of the specified volume to another pool.
//   Required parameters:
//       id - ID of the volume to stop moving to a different pool.

func (svc *VolumeService) AbortMoveVolume(id string) error {

	return nil
}

// BulkSetDedupeVolumes - enable or disable dedupe on a list of volumes. If the volumes are not dedupe capable, the operation will fail for the specified volume.
//   Required parameters:
//       volIds - IDs of the volumes to enable/disable dedupe.
//       dedupeEnabled - Dedupe property to be applied to the list of volumes.

func (svc *VolumeService) BulkSetDedupeVolumes(volIds []*string, dedupeEnabled bool) error {

	return nil
}

// BulkSetOnlineAndOfflineVolumes - bring a list of volumes online or offline.
//   Required parameters:
//       volIds - IDs of the volumes to set online or offline.
//       online - Desired state of the volumes. "true" for online, "false" for offline.

func (svc *VolumeService) BulkSetOnlineAndOfflineVolumes(volIds []*string, online bool) error {

	return nil
}
