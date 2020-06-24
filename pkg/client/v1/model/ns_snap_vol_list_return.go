/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSnapVolListReturn - Object returned after creating snapshot collection.
// Export NsSnapVolListReturnFields for advance operations like search filter etc.
var NsSnapVolListReturnFields *NsSnapVolListReturn

func init(){
		
	NsSnapVolListReturnFields= &NsSnapVolListReturn{
		
	}
}

type NsSnapVolListReturn struct {
   
    // A list of snapshot ids.
    
   	SnapIDs []*NsObjectIDKV `json:"snap_ids,omitempty"`
}
