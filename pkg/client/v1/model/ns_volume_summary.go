/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumeSummary - Select fields containing volume info.
// Export NsVolumeSummaryFields for advance operations like search filter etc.
var NsVolumeSummaryFields *NsVolumeSummary

func init(){
	IDfield:= "id"
	VolIDfield:= "vol_id"
	Namefield:= "name"
	VolNamefield:= "vol_name"
		
	NsVolumeSummaryFields= &NsVolumeSummary{
		ID: &IDfield,
		VolID: &VolIDfield,
		Name: &Namefield,
		VolName: &VolNamefield,
		
	}
}

type NsVolumeSummary struct {
   
    // ID of volume.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of volume.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Name of volume.
    
 	Name *string `json:"name,omitempty"`
   
    // Name of volume.
    
 	VolName *string `json:"vol_name,omitempty"`
}
