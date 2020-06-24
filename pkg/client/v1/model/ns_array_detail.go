/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArrayDetail - Detailed array information.
// Export NsArrayDetailFields for advance operations like search filter etc.
var NsArrayDetailFields *NsArrayDetail

func init(){
	IDfield:= "id"
	ArrayIDfield:= "array_id"
	Namefield:= "name"
	ArrayNamefield:= "array_name"
		
	NsArrayDetailFields= &NsArrayDetail{
		ID: &IDfield,
		ArrayID: &ArrayIDfield,
		Name: &Namefield,
		ArrayName: &ArrayNamefield,
		
	}
}

type NsArrayDetail struct {
   
    // Array API ID.
    
 	ID *string `json:"id,omitempty"`
   
    // Array API ID.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Unique name of array.
    
 	Name *string `json:"name,omitempty"`
   
    // Unique name of array.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // Start time of array evacuation.
    
   	EvacTime *int64 `json:"evac_time,omitempty"`
   
    // Initial data in the array.
    
   	EvacUsage *int64 `json:"evac_usage,omitempty"`
   
    // Usable capacity of the array.
    
   	UsableCapacity *int64 `json:"usable_capacity,omitempty"`
   
    // Usage of the array.
    
   	Usage *int64 `json:"usage,omitempty"`
   
    // Usage of volumes in the array.
    
   	VolUsageCompressedBytes *int64 `json:"vol_usage_compressed_bytes,omitempty"`
   
    // Usage of snapshots in the array.
    
   	SnapUsageCompressedBytes *int64 `json:"snap_usage_compressed_bytes,omitempty"`
   
    // Indicate whether usage of the array is valid.
    
 	UsageValID *bool `json:"usage_valid,omitempty"`
   
    // Migrate status of array.
    
   	Migrate *NsPoolMigrate `json:"migrate,omitempty"`
}
