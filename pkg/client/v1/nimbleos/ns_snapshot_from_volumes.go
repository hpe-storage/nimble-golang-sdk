// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapshotFromVolumesFields provides field names to use in filter parameters, for example.
var NsSnapshotFromVolumesFields *NsSnapshotFromVolumesFieldHandles

func init() {
	NsSnapshotFromVolumesFields = &NsSnapshotFromVolumesFieldHandles{
		ID:       "id",
		SnapId:   "snap_id",
		Name:     "name",
		SnapName: "snap_name",
	}
}

// NsSnapshotFromVolumes - Snapshot as presented in volumes object set.
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

// NsSnapshotFromVolumesFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapshotFromVolumesFieldHandles struct {
	ID       string
	SnapId   string
	Name     string
	SnapName string
}
