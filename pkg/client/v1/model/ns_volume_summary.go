// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsVolumeSummary - Select fields containing volume info.
// Export NsVolumeSummaryFields for advance operations like search filter etc.
var NsVolumeSummaryFields *NsVolumeSummary

func init(){
	IDfield:= "id"
	VolIdfield:= "vol_id"
	Namefield:= "name"
	VolNamefield:= "vol_name"
		
	NsVolumeSummaryFields= &NsVolumeSummary{
		ID:      &IDfield,
		VolId:   &VolIdfield,
		Name:    &Namefield,
		VolName: &VolNamefield,
	}
}

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
