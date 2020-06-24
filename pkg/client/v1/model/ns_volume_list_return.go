// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsVolumeListReturn - Object containing a list of volume names and IDs.
// Export NsVolumeListReturnFields for advance operations like search filter etc.
var NsVolumeListReturnFields *NsVolumeListReturn

func init(){
		
	NsVolumeListReturnFields= &NsVolumeListReturn{
		
	}
}

type NsVolumeListReturn struct {
	// VolList - A list of volume names and IDs.
   	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}
