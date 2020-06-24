/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsADTestUserReturn - Active Directory user details.
// Export NsADTestUserReturnFields for advance operations like search filter etc.
var NsADTestUserReturnFields *NsADTestUserReturn

func init(){
	Usernamefield:= "username"
	PrimaryGroupNamefield:= "primary_group_name"
	PrimaryGroupIDfield:= "primary_group_id"
		
	NsADTestUserReturnFields= &NsADTestUserReturn{
		Username: &Usernamefield,
		PrimaryGroupName: &PrimaryGroupNamefield,
		PrimaryGroupID: &PrimaryGroupIDfield,
		
	}
}

type NsADTestUserReturn struct {
   
    // Name of the Active Directory user.
    
 	Username *string `json:"username,omitempty"`
   
    // Name of the Active Directory group the user belongs to. RBAC is based on this primary group.
    
 	PrimaryGroupName *string `json:"primary_group_name,omitempty"`
   
    // ID of the Active Directory group the user belongs to. RBAC is based on this primary group.
    
 	PrimaryGroupID *string `json:"primary_group_id,omitempty"`
   
    // Number of Active Directory groups the user belongs to.
    
   	GroupCount *int64 `json:"group_count,omitempty"`
   
    // List of Active Directory groups the user belongs to.
    
	Groups []*string `json:"groups,omitempty"`
   
    // The role the user belongs to.
    
   	Role *NsUserRoles `json:"role,omitempty"`
}
