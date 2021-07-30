// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapshotFromVolumes - Snapshot as presented in volumes object set.
// Export NsSnapshotFromVolumesFields for advance operations like search filter etc.
var NsSnapshotFromVolumesFields *NsSnapshotFromVolumesStringFields

func init() {
	IDfield := "id"
	SnapIdfield := "snap_id"
	Namefield := "name"
	SnapNamefield := "snap_name"

	NsSnapshotFromVolumesFields = &NsSnapshotFromVolumesStringFields{
		ID:       &IDfield,
		SnapId:   &SnapIdfield,
		Name:     &Namefield,
		SnapName: &SnapNamefield,
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
