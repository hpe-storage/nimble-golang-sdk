/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsTargetSubnet - List of subnet labels.
// Export NsTargetSubnetFields for advance operations like search filter etc.
var NsTargetSubnetFields *NsTargetSubnet

func init(){
	IDfield:= "id"
	Labelfield:= "label"
		
	NsTargetSubnetFields= &NsTargetSubnet{
		ID: &IDfield,
		Label: &Labelfield,
		
	}
}

type NsTargetSubnet struct {
   
    // Subnet ID.
    
 	ID *string `json:"id,omitempty"`
   
    // Subnet label.
    
 	Label *string `json:"label,omitempty"`
}
