/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsExtendedPrivilege - Extended privilege info.
// Export NsExtendedPrivilegeFields for advance operations like search filter etc.
var NsExtendedPrivilegeFields *NsExtendedPrivilege

func init(){
	ObjectTypefield:= "object_type"
	Operationfield:= "operation"
		
	NsExtendedPrivilegeFields= &NsExtendedPrivilege{
		ObjectType: &ObjectTypefield,
		Operation: &Operationfield,
		
	}
}

type NsExtendedPrivilege struct {
   
    // Object type name associated with this privilege.
    
 	ObjectType *string `json:"object_type,omitempty"`
   
    // Operation associated with the above object for this privilege.
    
 	Operation *string `json:"operation,omitempty"`
   
    // Indicate whether the above operation is allowed for this privilege.
    
 	Allow *bool `json:"allow,omitempty"`
}
