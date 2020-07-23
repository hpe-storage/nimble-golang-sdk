// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeSnapshotAttribute - Snapshot attributes that could be specified for individual snapshots during snapshot collection creation.
// Export NsVolumeSnapshotAttributeFields for advance operations like search filter etc.
var NsVolumeSnapshotAttributeFields *NsVolumeSnapshotAttribute

func init() {
	VolIdfield := "vol_id"
	AppUuidfield := "app_uuid"

	NsVolumeSnapshotAttributeFields = &NsVolumeSnapshotAttribute{
		VolId:   &VolIdfield,
		AppUuid: &AppUuidfield,
	}
}

type NsVolumeSnapshotAttribute struct {
	// VolId - ID of the volume on which snapshot will be created.
	VolId *string `json:"vol_id,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// AppUuid - Application identifier of snapshot.
	AppUuid *string `json:"app_uuid,omitempty"`
}
