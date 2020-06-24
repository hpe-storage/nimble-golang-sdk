// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsWitnessTestResponse - Results of witness connection test.
// Export NsWitnessTestResponseFields for advance operations like search filter etc.
var NsWitnessTestResponseFields *NsWitnessTestResponse

func init(){
	ArrayNamefield:= "array_name"
	WitnessConnectivityStatefield:= "witness_connectivity_state"
	WitnessConnectivityMessagefield:= "witness_connectivity_message"
		
	NsWitnessTestResponseFields= &NsWitnessTestResponse{
	ArrayName: &ArrayNamefield,
	WitnessConnectivityState: &WitnessConnectivityStatefield,
	WitnessConnectivityMessage: &WitnessConnectivityMessagefield,
		
	}
}

type NsWitnessTestResponse struct {
	// ArrayName - Name of an array.
 	ArrayName *string `json:"array_name,omitempty"`
	// Role - Role of an array in the group.
   	Role *NsArrayRole `json:"role,omitempty"`
	// WitnessConnectivityState - Reachability status of the witness.
 	WitnessConnectivityState *string `json:"witness_connectivity_state,omitempty"`
	// WitnessConnectivityMessage - Reachability message of the witness.
 	WitnessConnectivityMessage *string `json:"witness_connectivity_message,omitempty"`
}
