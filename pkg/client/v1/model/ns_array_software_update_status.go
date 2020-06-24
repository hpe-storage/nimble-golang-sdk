/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArraySoftwareUpdateStatus - List of software update errors for a specific array.
// Export NsArraySoftwareUpdateStatusFields for advance operations like search filter etc.
var NsArraySoftwareUpdateStatusFields *NsArraySoftwareUpdateStatus

func init(){
	ArrayNamefield:= "array_name"
	Errorfield:= "error"
		
	NsArraySoftwareUpdateStatusFields= &NsArraySoftwareUpdateStatus{
		ArrayName: &ArrayNamefield,
		Error: &Errorfield,
		
	}
}

type NsArraySoftwareUpdateStatus struct {
   
    // Name of the array.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // Top level software update error for the array.
    
 	Error *string `json:"error,omitempty"`
   
    // Error mask for the controller.
    
   	CtrlrErrorMask *int64 `json:"ctrlr_error_mask,omitempty"`
   
    // List of software update errors for controller A. The key is always "error" and the value is the actual error code string.
    
   	CtrlrAErrs []*NsArraySoftwareUpdateError `json:"ctrlr_a_errs,omitempty"`
   
    // List of software update errors for controller B. The key is always "error" and the value is the actual error code string.
    
   	CtrlrBErrs []*NsArraySoftwareUpdateError `json:"ctrlr_b_errs,omitempty"`
}
