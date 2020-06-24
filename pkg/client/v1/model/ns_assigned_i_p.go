/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAssignedIP - IP address assignment details to network interface.
// Export NsAssignedIPFields for advance operations like search filter etc.
var NsAssignedIPFields *NsAssignedIP

func init(){
	Ipfield:= "ip"
		
	NsAssignedIPFields= &NsAssignedIP{
		Ip: &Ipfield,
		
	}
}

type NsAssignedIP struct {
   
    // VLAN id of network interface.
    
   	VlanID *int64 `json:"vlan_id,omitempty"`
   
    // Assigned IP address to network interface.
    
 	Ip *string `json:"ip,omitempty"`
}
