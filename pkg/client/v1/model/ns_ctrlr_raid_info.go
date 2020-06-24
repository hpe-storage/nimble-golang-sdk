/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsCtrlrRaIDInfo - Information about a controller's raid configuration.
// Export NsCtrlrRaIDInfoFields for advance operations like search filter etc.
var NsCtrlrRaIDInfoFields *NsCtrlrRaIDInfo

func init(){
	RaIDTypefield:= "raid_type"
		
	NsCtrlrRaIDInfoFields= &NsCtrlrRaIDInfo{
		RaIDType: &RaIDTypefield,
		
	}
}

type NsCtrlrRaIDInfo struct {
   
    // Raid ID for this raid array.
    
   	RaIDID *int64 `json:"raid_id,omitempty"`
   
    // Type of raid for this array.
    
 	RaIDType *string `json:"raid_type,omitempty"`
   
    // Maximum number of copies.
    
   	MaxCopies *int64 `json:"max_copies,omitempty"`
   
    // Current number of copies.
    
   	CurCopies *int64 `json:"cur_copies,omitempty"`
   
    // Is this raid array resynchronizing.
    
 	IsResyncing *bool `json:"is_resyncing,omitempty"`
}
