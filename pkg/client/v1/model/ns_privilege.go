/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsPrivilege - Privilege info.
// Export NsPrivilegeFields for advance operations like search filter etc.
var NsPrivilegeFields *NsPrivilege

func init(){
	ObjectTypefield:= "object_type"
		
	NsPrivilegeFields= &NsPrivilege{
		ObjectType: &ObjectTypefield,
		
	}
}

type NsPrivilege struct {
   
    // Object type name associated with this privilege.
    
 	ObjectType *string `json:"object_type,omitempty"`
   
    // List of operations associated with the above object for this privilege.
    
	Operations []*string `json:"operations,omitempty"`
}
