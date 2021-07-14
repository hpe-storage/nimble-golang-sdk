// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsAssignedIP - IP address assignment details to network interface.
// Export NsAssignedIPFields for advance operations like search filter etc.
var NsAssignedIPFields *NsAssignedIP

func init(){
 Ipfield:= "ip"

 NsAssignedIPFields= &NsAssignedIP{
  Ip:     &Ipfield,
 }
}


type NsAssignedIP struct {
 // VlanId - VLAN id of network interface.
    VlanId *int64 `json:"vlan_id,omitempty"`
 // Ip - Assigned IP address to network interface.
  Ip *string `json:"ip,omitempty"`
}
