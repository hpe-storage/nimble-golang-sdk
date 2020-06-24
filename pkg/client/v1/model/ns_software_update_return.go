// Copyright 2020 Hewlett Packard Enterprise Development LP

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
	// Error - Top level error.
 	Error *string `json:"error,omitempty"`
	// ArrayResponseList - Errors from all the arrays in the group.
   	ArrayResponseList []*NsArraySoftwareUpdateStatus `json:"array_response_list,omitempty"`
}
