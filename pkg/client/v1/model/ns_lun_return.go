// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsLunReturn - Return LU number.
// Export NsLunReturnFields for advance operations like search filter etc.
var NsLunReturnFields *NsLunReturn

func init(){
		
	NsLunReturnFields= &NsLunReturn{
	}
}

type NsLunReturn struct {
	// Lun - LU number in hexadecimal.
   	Lun *int64 `json:"lun,omitempty"`
	// LuNumber - LU number in decimal.
   	LuNumber *int64 `json:"lu_number,omitempty"`
}
