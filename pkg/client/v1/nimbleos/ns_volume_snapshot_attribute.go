// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolumeSnapshotAttribute - Snapshot attributes that could be specified for individual snapshots during snapshot collection creation.
// Export NsVolumeSnapshotAttributeFields for advance operations like search filter etc.
var NsVolumeSnapshotAttributeFields *NsVolumeSnapshotAttribute

func init() {

	NsVolumeSnapshotAttributeFields = &NsVolumeSnapshotAttribute{
		VolId:   "vol_id",
		AppUuid: "app_uuid",
	}
}

type NsVolumeSnapshotAttribute struct {
	// VolId - ID of the volume on which snapshot will be created.
	VolId string `json:"vol_id,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// AppUuid - Application identifier of snapshot.
	AppUuid string `json:"app_uuid,omitempty"`
}

// sdk internal struct
type nsVolumeSnapshotAttribute struct {
	VolId    *string       `json:"vol_id,omitempty"`
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	AppUuid  *string       `json:"app_uuid,omitempty"`
}

// EncodeNsVolumeSnapshotAttribute - Transform NsVolumeSnapshotAttribute to nsVolumeSnapshotAttribute type
func EncodeNsVolumeSnapshotAttribute(request interface{}) (*nsVolumeSnapshotAttribute, error) {
	reqNsVolumeSnapshotAttribute := request.(*NsVolumeSnapshotAttribute)
	byte, err := json.Marshal(reqNsVolumeSnapshotAttribute)
	if err != nil {
		return nil, err
	}
	respNsVolumeSnapshotAttributePtr := &nsVolumeSnapshotAttribute{}
	err = json.Unmarshal(byte, respNsVolumeSnapshotAttributePtr)
	return respNsVolumeSnapshotAttributePtr, err
}

// DecodeNsVolumeSnapshotAttribute Transform nsVolumeSnapshotAttribute to NsVolumeSnapshotAttribute type
func DecodeNsVolumeSnapshotAttribute(request interface{}) (*NsVolumeSnapshotAttribute, error) {
	reqNsVolumeSnapshotAttribute := request.(*nsVolumeSnapshotAttribute)
	byte, err := json.Marshal(reqNsVolumeSnapshotAttribute)
	if err != nil {
		return nil, err
	}
	respNsVolumeSnapshotAttribute := &NsVolumeSnapshotAttribute{}
	err = json.Unmarshal(byte, respNsVolumeSnapshotAttribute)
	return respNsVolumeSnapshotAttribute, err
}
