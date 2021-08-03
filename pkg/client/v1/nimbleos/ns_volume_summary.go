// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeSummaryFields provides field names to use in filter parameters, for example.
var NsVolumeSummaryFields *NsVolumeSummaryFieldHandles

func init() {
	NsVolumeSummaryFields = &NsVolumeSummaryFieldHandles{
		ID:      "id",
		VolId:   "vol_id",
		Name:    "name",
		VolName: "vol_name",
	}
}

// NsVolumeSummary - Select fields containing volume info.
type NsVolumeSummary struct {
	// ID - ID of volume.
	ID *string `json:"id,omitempty"`
	// VolId - ID of volume.
	VolId *string `json:"vol_id,omitempty"`
	// Name - Name of volume.
	Name *string `json:"name,omitempty"`
	// VolName - Name of volume.
	VolName *string `json:"vol_name,omitempty"`
}

// NsVolumeSummaryFieldHandles provides a string representation for each NsVolumeSummary field.
type NsVolumeSummaryFieldHandles struct {
	ID      string
	VolId   string
	Name    string
	VolName string
}
