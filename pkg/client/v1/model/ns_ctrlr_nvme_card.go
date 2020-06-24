// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsCtrlrNvmeCard - NVMe accelerator card.
// Export NsCtrlrNvmeCardFields for advance operations like search filter etc.
var NsCtrlrNvmeCardFields *NsCtrlrNvmeCard

func init(){
	SerialNumberfield:= "serial_number"
		
	NsCtrlrNvmeCardFields= &NsCtrlrNvmeCard{
	SerialNumber: &SerialNumberfield,
		
	}
}

type NsCtrlrNvmeCard struct {
	// SerialNumber - Serial number.
 	SerialNumber *string `json:"serial_number,omitempty"`
	// Size - NVMe card cache size in bytes.
   	Size *int64 `json:"size,omitempty"`
	// State - Online state.
   	State *NsCtrlrNvmeCardState `json:"state,omitempty"`
}
