// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsFcTdzPort - Fibre Channel Target port.
// Export NsFcTdzPortFields for advance operations like search filter etc.
var NsFcTdzPortFields *NsFcTdzPort

func init(){
	ArrayNamefield:= "array_name"
	FcNamefield:= "fc_name"
		
	NsFcTdzPortFields= &NsFcTdzPort{
		ArrayName: &ArrayNamefield,
		FcName:    &FcNamefield,
	}
}

type NsFcTdzPort struct {
	// ArrayName - Unique name of the array.
 	ArrayName *string `json:"array_name,omitempty"`
	// FcName - Target port interface name.
 	FcName *string `json:"fc_name,omitempty"`
}
