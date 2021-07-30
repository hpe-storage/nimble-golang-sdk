// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsVolumeSnapshotAttributeFields provides field names to use in filter parameters, for example.
var NsVolumeSnapshotAttributeFields *NsVolumeSnapshotAttributeFieldHandles

func init() {
	fieldVolId := "vol_id"
	fieldMetadata := "metadata"
	fieldAppUuid := "app_uuid"

	NsVolumeSnapshotAttributeFields = &NsVolumeSnapshotAttributeFieldHandles{
		VolId:    &fieldVolId,
		Metadata: &fieldMetadata,
		AppUuid:  &fieldAppUuid,
	}
}

// NsVolumeSnapshotAttribute - Snapshot attributes that could be specified for individual snapshots during snapshot collection creation.
type NsVolumeSnapshotAttribute struct {
	// VolId - ID of the volume on which snapshot will be created.
	VolId *string `json:"vol_id,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// AppUuid - Application identifier of snapshot.
	AppUuid *string `json:"app_uuid,omitempty"`
}

// NsVolumeSnapshotAttributeFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolumeSnapshotAttributeFieldHandles struct {
	VolId    *string
	Metadata *string
	AppUuid  *string
}
