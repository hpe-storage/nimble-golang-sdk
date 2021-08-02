// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeListReturnFields provides field names to use in filter parameters, for example.
var NsVolumeListReturnFields *NsVolumeListReturnFieldHandles

func init() {
	NsVolumeListReturnFields = &NsVolumeListReturnFieldHandles{
		VolList: "vol_list",
	}
}

// NsVolumeListReturn - Object containing a list of volume names and IDs.
type NsVolumeListReturn struct {
	// VolList - A list of volume names and IDs.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}

// NsVolumeListReturnFieldHandles provides a string representation for each NsVolumeListReturn field.
type NsVolumeListReturnFieldHandles struct {
	VolList string
}
