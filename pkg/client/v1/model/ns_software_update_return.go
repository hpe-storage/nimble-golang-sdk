/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSoftwareUpdateReturn - The status returned by the software update precheck and start actions.
// Export NsSoftwareUpdateReturnFields for advance operations like search filter etc.
var NsSoftwareUpdateReturnFields *NsSoftwareUpdateReturn

func init(){
	Errorfield:= "error"
		
	NsSoftwareUpdateReturnFields= &NsSoftwareUpdateReturn{
		Error: &Errorfield,
		
	}
}

type NsSoftwareUpdateReturn struct {
   
    // Top level error.
    
 	Error *string `json:"error,omitempty"`
   
    // Errors from all the arrays in the group.
    
   	ArrayResponseList []*NsArraySoftwareUpdateStatus `json:"array_response_list,omitempty"`
}
