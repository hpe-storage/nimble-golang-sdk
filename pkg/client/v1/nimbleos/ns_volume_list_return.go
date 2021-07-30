// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeListReturn - Object containing a list of volume names and IDs.
// Export NsVolumeListReturnFields for advance operations like search filter etc.
var NsVolumeListReturnFields *NsVolumeListReturnStringFields

func init() {
	VolListfield := "vol_list"

	NsVolumeListReturnFields = &NsVolumeListReturnStringFields{
		VolList: &VolListfield,
	}
}

type NsVolumeListReturn struct {
	// VolList - A list of volume names and IDs.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}

// Struct for NsVolumeListReturnFields
type NsVolumeListReturnStringFields struct {
	VolList *string
}
