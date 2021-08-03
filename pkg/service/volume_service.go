// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// Volume Service - Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// VolumeService type
type VolumeService struct {
	objectSet *client.VolumeObjectSet
}

// NewVolumeService - method to initialize "VolumeService"
func NewVolumeService(gs *NsGroupService) *VolumeService {
	objectSet := gs.client.GetVolumeObjectSet()
	return &VolumeService{objectSet: objectSet}
}

// GetVolumes - method returns a array of pointers of type "Volumes"
func (svc *VolumeService) GetVolumes(params *param.GetParams) ([]*nimbleos.Volume, error) {
	volumeResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// CreateVolume - method creates a "Volume"
func (svc *VolumeService) CreateVolume(obj *nimbleos.Volume) (*nimbleos.Volume, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateVolume: invalid parameter specified, %v", obj)
	}

	volumeResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// UpdateVolume - method modifies  the "Volume"
func (svc *VolumeService) UpdateVolume(id string, obj *nimbleos.Volume) (*nimbleos.Volume, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateVolume: invalid parameter specified, %v", obj)
	}

	volumeResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// GetVolumeById - method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeById(id string) (*nimbleos.Volume, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetVolumeById: invalid parameter specified, %v", id)
	}

	volumeResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// GetVolumeByName - method returns a pointer "Volume"
func (svc *VolumeService) GetVolumeByName(name string) (*nimbleos.Volume, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: &nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	volumeResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(volumeResp) == 0 {
		return nil, nil
	}

	return volumeResp[0], nil
}

// GetVolumeBySerialNumber method returns a pointer to "Volume"
func (svc *VolumeService) GetVolumeBySerialNumber(serialNumber string) (*nimbleos.Volume, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: &nimbleos.VolumeFields.SerialNumber,
			Operator:  param.EQUALS.String(),
			Value:     serialNumber,
		},
	}
	volumeResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	if len(volumeResp) == 0 {
		return nil, nil
	}

	return volumeResp[0], nil
}

//OnlineVolume - method makes the volume online
func (svc *VolumeService) OnlineVolume(id string, force bool) (*nimbleos.Volume, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("OnlineVolume: invalid parameter specified, %s", id)
	}

	volumeResp, err := svc.UpdateVolume(id, &nimbleos.Volume{
		Online: param.NewBool(true),
		Force:  param.NewBool(true),
	})
	if err != nil {
		return nil, err
	}
	return volumeResp, nil
}

// OfflineVolume - makes the volume offline
func (svc *VolumeService) OfflineVolume(id string, force bool) (*nimbleos.Volume, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("OfflineVolume: invalid parameter specified, %s", id)
	}

	volumeResp, err := svc.UpdateVolume(id, &nimbleos.Volume{
		Online: param.NewBool(false),
		Force:  param.NewBool(false),
	})
	if err != nil {
		return nil, err
	}
	return volumeResp, nil

}

// DestroyVolume - makes the volume offline and delete it
func (svc *VolumeService) DestroyVolume(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DestroyVolume: invalid parameter specified, %s", id)
	}
	_, err := svc.OfflineVolume(id, false)
	if err != nil {
		return err
	}

	err = svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// AssociateVolume - add a volume to a volume collection
func (svc *VolumeService) AssociateVolume(volId string, volcollId string) error {
	if len(volId) == 0 || len(volcollId) == 0 {
		return fmt.Errorf("AssociateVolume: invalid parameter specified, %s", volId)
	}

	_, err := svc.UpdateVolume(volId, &nimbleos.Volume{
		VolcollId: param.NewString(volcollId),
	})
	if err != nil {
		return err
	}
	return nil
}

// DisassociateVolume - remove a volume from a volume collection
func (svc *VolumeService) DisassociateVolume(volId string) error {
	if len(volId) == 0 {
		return fmt.Errorf("DisassociateVolume: invalid parameter specified, %s", volId)
	}

	_, err := svc.UpdateVolume(volId, &nimbleos.Volume{
		VolcollId: param.NewString(""),
	})
	if err != nil {
		return err
	}
	return nil

}

// DeleteVolume - deletes the "Volume"
func (svc *VolumeService) DeleteVolume(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteVolume: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// RestoreVolume - restore volume data from a previous snapshot.
//   Required parameters:
//       id - ID of the restored volume.
//       baseSnapId - ID of the snapshot to which this the volume is restored.

func (svc *VolumeService) RestoreVolume(id string, baseSnapId string) error {

	if len(id) == 0 || len(baseSnapId) == 0 {
		return fmt.Errorf("RestoreVolume: invalid parameter specified id: %v, baseSnapId: %v ", id, baseSnapId)
	}

	err := svc.objectSet.Restore(&id, &baseSnapId)
	return err
}

// MoveVolume - move a volume and its related volumes to another pool. To change a single volume's folder assignment (while remaining in the same pool), use a volume update operation to change the folder_id attribute.
//   Required parameters:
//       id - ID of the volume to move.
//       destPoolId - ID of the destination pool or folder. Specify a pool ID when the volumes should not be in a folder; otherwise, specify a folder ID and the pool will be derived from the folder.

//   Optional parameters:
//       forceVvol - Forcibly move a Virtual Volume. Moving Virtual Volume is disruptive to the vCenter, hence it should only be done by the VASA Provider (VP). This flag should only be set by the VP when it calls this API.

func (svc *VolumeService) MoveVolume(id string, destPoolId string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error) {

	if len(id) == 0 || len(destPoolId) == 0 {
		return nil, fmt.Errorf("MoveVolume: invalid parameter specified id: %v, destPoolId: %v ", id, destPoolId)
	}

	resp, err := svc.objectSet.Move(&id, &destPoolId, forceVvol)
	return resp, err
}

// BulkMoveVolumes - move volumes and their related volumes to another pool. To change a single volume's folder assignment (while remaining in the same pool), use a volume update operation to change the folder_id attribute.
//   Required parameters:
//       volIds - IDs of the volumes to move.
//       destPoolId - ID of the destination pool or folder. Specify a pool ID when the volumes should not be in a folder; otherwise, specify a folder ID and the pool will be derived from the folder.

//   Optional parameters:
//       forceVvol - Forcibly move a Virtual Volume. Moving Virtual Volume is disruptive to the vCenter, hence it should only be done by the VASA Provider (VP). This flag should only be set by the VP when it calls this API.

func (svc *VolumeService) BulkMoveVolumes(volIds []*string, destPoolId string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error) {

	if len(volIds) == 0 || len(destPoolId) == 0 {
		return nil, fmt.Errorf("BulkMoveVolume: invalid parameter specified volIds: %v, destPoolId: %v ", volIds, destPoolId)
	}

	resp, err := svc.objectSet.BulkMove(volIds, &destPoolId, forceVvol)
	return resp, err
}

// AbortMoveVolume - abort the in-progress move of the specified volume to another pool.
//   Required parameters:
//       id - ID of the volume to stop moving to a different pool.

func (svc *VolumeService) AbortMoveVolume(id string) error {

	if len(id) == 0 {
		return fmt.Errorf("AbortMoveVolume: invalid parameter specified id: %v ", id)
	}

	err := svc.objectSet.AbortMove(&id)
	return err
}

// BulkSetDedupeVolumes - enable or disable dedupe on a list of volumes. If the volumes are not dedupe capable, the operation will fail for the specified volume.
//   Required parameters:
//       volIds - IDs of the volumes to enable/disable dedupe.
//       dedupeEnabled - Dedupe property to be applied to the list of volumes.

func (svc *VolumeService) BulkSetDedupeVolumes(volIds []*string, dedupeEnabled bool) error {

	if len(volIds) == 0 {
		return fmt.Errorf("BulkSetDedupeVolume: invalid parameter specified volIds: %v ", volIds)
	}

	err := svc.objectSet.BulkSetDedupe(volIds, &dedupeEnabled)
	return err
}

// BulkSetOnlineAndOfflineVolumes - bring a list of volumes online or offline.
//   Required parameters:
//       volIds - IDs of the volumes to set online or offline.
//       online - Desired state of the volumes. "true" for online, "false" for offline.

func (svc *VolumeService) BulkSetOnlineAndOfflineVolumes(volIds []*string, online bool) error {

	if len(volIds) == 0 {
		return fmt.Errorf("BulkSetOnlineAndOfflineVolume: invalid parameter specified volIds: %v ", volIds)
	}

	err := svc.objectSet.BulkSetOnlineAndOffline(volIds, &online)
	return err
}
