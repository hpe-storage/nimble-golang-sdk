// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapshotFromVolumes - Snapshot as presented in volumes object set.

// Export NsSnapshotFromVolumesFields provides field names to use in filter parameters, for example.
var NsSnapshotFromVolumesFields *NsSnapshotFromVolumesStringFields

func init() {
	fieldID := "id"
	fieldSnapId := "snap_id"
	fieldName := "name"
	fieldSnapName := "snap_name"

	NsSnapshotFromVolumesFields = &NsSnapshotFromVolumesStringFields{
		ID:       &fieldID,
		SnapId:   &fieldSnapId,
		Name:     &fieldName,
		SnapName: &fieldSnapName,
	}
}

type NsSnapshotFromVolumes struct {
	// ID - Snapshot id.
	ID *string `json:"id,omitempty"`
	// SnapId - Snapshot id.
	SnapId *string `json:"snap_id,omitempty"`
	// Name - Snapshot name.
	Name *string `json:"name,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
}

// Struct for NsSnapshotFromVolumesFields
type NsSnapshotFromVolumesStringFields struct {
	ID       *string
	SnapId   *string
	Name     *string
	SnapName *string
}
