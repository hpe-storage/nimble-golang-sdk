/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// UserGroup - Represents Active Directory groups configured to manage the system.
// Export UserGroupFields for advance operations like search filter etc.
var UserGroupFields *UserGroup

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Descriptionfield:= "description"
	RoleIDfield:= "role_id"
	ExternalIDfield:= "external_id"
	DomainIDfield:= "domain_id"
	DomainNamefield:= "domain_name"
		
	UserGroupFields= &UserGroup{
		ID: &IDfield,
		Name: &Namefield,
		Description: &Descriptionfield,
		RoleID: &RoleIDfield,
		ExternalID: &ExternalIDfield,
		DomainID: &DomainIDfield,
		DomainName: &DomainNamefield,
		
	}
}

type UserGroup struct {
   
    // Identifier for the user group.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the user group.
    
 	Name *string `json:"name,omitempty"`
   
    // Description of the user group.
    
 	Description *string `json:"description,omitempty"`
   
    // Identifier for the user group's role.
    
 	RoleID *string `json:"role_id,omitempty"`
   
    // Role of the user.
    
   	Role *NsUserRoles `json:"role,omitempty"`
   
    // The amount of time that the user session is inactive before timing out. A value of 0 indicates that the timeout is taken from the group setting.
    
   	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`
   
    // Time when this user was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this user was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // User is currently disabled.
    
 	Disabled *bool `json:"disabled,omitempty"`
   
    // External ID of the user group. In Active Directory, it is the group's SID (Security Identifier).
    
 	ExternalID *string `json:"external_id,omitempty"`
   
    // Identifier of the domain this user group belongs to.
    
 	DomainID *string `json:"domain_id,omitempty"`
   
    // Role of the user.
    
 	DomainName *string `json:"domain_name,omitempty"`
}
