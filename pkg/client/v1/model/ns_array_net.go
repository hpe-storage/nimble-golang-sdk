/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArrayNet - Array network config.
// Export NsArrayNetFields for advance operations like search filter etc.
var NsArrayNetFields *NsArrayNet

func init(){
	Namefield:= "name"
	CtrlrASupportIpfield:= "ctrlr_a_support_ip"
	CtrlrBSupportIpfield:= "ctrlr_b_support_ip"
		
	NsArrayNetFields= &NsArrayNet{
		Name: &Namefield,
		CtrlrASupportIp: &CtrlrASupportIpfield,
		CtrlrBSupportIp: &CtrlrBSupportIpfield,
		
	}
}

type NsArrayNet struct {
   
    // Name of the array.
    
 	Name *string `json:"name,omitempty"`
   
    // GID of member. This field cannot be updated.
    
   	MemberGID *int64 `json:"member_gid,omitempty"`
   
    // IP address of controller A.
    
 	CtrlrASupportIp *string `json:"ctrlr_a_support_ip,omitempty"`
   
    // IP address of controller B.
    
 	CtrlrBSupportIp *string `json:"ctrlr_b_support_ip,omitempty"`
   
    // List of NICs.
    
   	NicList []*NsNIC `json:"nic_list,omitempty"`
}
