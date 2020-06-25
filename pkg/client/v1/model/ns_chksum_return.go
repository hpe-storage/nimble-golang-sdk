// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsChksumReturn - Return computed checksum.
// Export NsChksumReturnFields for advance operations like search filter etc.
var NsChksumReturnFields *NsChksumReturn

func init(){
	Cksumfield:= "cksum"
		
	NsChksumReturnFields= &NsChksumReturn{
		Cksum:&Cksumfield,
	}
}

type NsChksumReturn struct {
	// Cksum - Computed checksum.
 	Cksum *string `json:"cksum,omitempty"`
}
