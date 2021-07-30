// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolAndSnapName - Snapshot name and the belonging volume name.
// Export NsVolAndSnapNameFields for advance operations like search filter etc.
var NsVolAndSnapNameFields *NsVolAndSnapNameStringFields

func init() {
	VolNamefield := "vol_name"
	SnapNamefield := "snap_name"

	NsVolAndSnapNameFields = &NsVolAndSnapNameStringFields{
		VolName:  &VolNamefield,
		SnapName: &SnapNamefield,
	}
}

type NsVolAndSnapName struct {
	// VolName - The name of the volume that the snapshot belongs to.
	VolName *string `json:"vol_name,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
}

// Struct for NsVolAndSnapNameFields
type NsVolAndSnapNameStringFields struct {
	VolName  *string
	SnapName *string
}
