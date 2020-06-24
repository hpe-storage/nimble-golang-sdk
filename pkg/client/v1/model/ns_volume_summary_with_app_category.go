/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumeSummaryWithAppCategory - Select fields containing volume info.
// Export NsVolumeSummaryWithAppCategoryFields for advance operations like search filter etc.
var NsVolumeSummaryWithAppCategoryFields *NsVolumeSummaryWithAppCategory

func init(){
	IDfield:= "id"
	Namefield:= "name"
	AppCategoryfield:= "app_category"
	FullNamefield:= "full_name"
		
	NsVolumeSummaryWithAppCategoryFields= &NsVolumeSummaryWithAppCategory{
		ID: &IDfield,
		Name: &Namefield,
		AppCategory: &AppCategoryfield,
		FullName: &FullNamefield,
		
	}
}

type NsVolumeSummaryWithAppCategory struct {
   
    // ID of volume.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of volume.
    
 	Name *string `json:"name,omitempty"`
   
    // Application category that the volume belongs to.
    
 	AppCategory *string `json:"app_category,omitempty"`
   
    // Fully qualified name of volume.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // LUN of volume. Secondary LUN if this is Virtual Volume.
    
   	Lun *int64 `json:"lun,omitempty"`
}
