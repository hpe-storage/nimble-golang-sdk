// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolAndSnapNameFields provides field names to use in filter parameters, for example.
var NsVolAndSnapNameFields *NsVolAndSnapNameFieldHandles

func init() {
	NsVolAndSnapNameFields = &NsVolAndSnapNameFieldHandles{
		VolName:  "vol_name",
		SnapName: "snap_name",
	}
}

// NsVolAndSnapName - Snapshot name and the belonging volume name.
type NsVolAndSnapName struct {
	// VolName - The name of the volume that the snapshot belongs to.
	VolName *string `json:"vol_name,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
}

// NsVolAndSnapNameFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolAndSnapNameFieldHandles struct {
	VolName  string
	SnapName string
}
