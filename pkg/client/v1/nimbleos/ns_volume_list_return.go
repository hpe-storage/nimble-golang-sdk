// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsVolumeListReturnFields provides field names to use in filter parameters, for example.
var NsVolumeListReturnFields *NsVolumeListReturnFieldHandles

func init() {
	fieldVolList := "vol_list"

	NsVolumeListReturnFields = &NsVolumeListReturnFieldHandles{
		VolList: &fieldVolList,
	}
}

// NsVolumeListReturn - Object containing a list of volume names and IDs.
type NsVolumeListReturn struct {
	// VolList - A list of volume names and IDs.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}

// NsVolumeListReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolumeListReturnFieldHandles struct {
	VolList *string
}
