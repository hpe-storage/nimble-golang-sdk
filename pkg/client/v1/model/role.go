/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Role - Retrieve roles and privileges for role-based access control.
// Export RoleFields for advance operations like search filter etc.
var RoleFields *Role

func init(){
	IDfield:= "id"
	Namefield:= "name"
	FullNamefield:= "full_name"
	Descriptionfield:= "description"
		
	RoleFields= &Role{
		ID: &IDfield,
		Name: &Namefield,
		FullName: &FullNamefield,
		Description: &Descriptionfield,
		
	}
}

type Role struct {
   
    // Identifier for role.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of role.
    
 	Name *string `json:"name,omitempty"`
   
    // Full name of role.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Description of role.
    
 	Description *string `json:"description,omitempty"`
   
    // List of privileges for this role.
    
   	PrivilegeList []*NsPrivilege `json:"privilege_list,omitempty"`
   
    // List of extended privileges for this role.
    
   	ExtendedPrivilegeList []*NsExtendedPrivilege `json:"extended_privilege_list,omitempty"`
   
    // Time when this role was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this role was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Indicate whether the role is hidden.
    
 	HIDden *bool `json:"hidden,omitempty"`
}
