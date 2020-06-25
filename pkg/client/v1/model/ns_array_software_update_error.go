// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsArraySoftwareUpdateError - Software update error for specific controller.
// Export NsArraySoftwareUpdateErrorFields for advance operations like search filter etc.
var NsArraySoftwareUpdateErrorFields *NsArraySoftwareUpdateError

func init(){
	Errorfield:= "error"
		
	NsArraySoftwareUpdateErrorFields= &NsArraySoftwareUpdateError{
		Error:&Errorfield,
	}
}

type NsArraySoftwareUpdateError struct {
	// Error - Error code from software update.
 	Error *string `json:"error,omitempty"`
}
