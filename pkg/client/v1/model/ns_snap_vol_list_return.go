// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsSnapVolListReturn - Object returned after creating snapshot collection.
// Export NsSnapVolListReturnFields for advance operations like search filter etc.
var NsSnapVolListReturnFields *NsSnapVolListReturn

func init(){
		
	NsSnapVolListReturnFields= &NsSnapVolListReturn{
		
	}
}

type NsSnapVolListReturn struct {
	// SnapIds - A list of snapshot ids.
   	SnapIds []*NsObjectIDKV `json:"snap_ids,omitempty"`
}
