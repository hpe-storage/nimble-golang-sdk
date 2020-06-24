/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsTokenReportUserDetailsReturn - Return values of token reporting user details.
// Export NsTokenReportUserDetailsReturnFields for advance operations like search filter etc.
var NsTokenReportUserDetailsReturnFields *NsTokenReportUserDetailsReturn

func init(){
	UserNamefield:= "user_name"
	PrimaryGroupIDfield:= "primary_group_id"
	PrimaryGroupNamefield:= "primary_group_name"
	UserIDfield:= "user_id"
		
	NsTokenReportUserDetailsReturnFields= &NsTokenReportUserDetailsReturn{
		UserName: &UserNamefield,
		PrimaryGroupID: &PrimaryGroupIDfield,
		PrimaryGroupName: &PrimaryGroupNamefield,
		UserID: &UserIDfield,
		
	}
}

type NsTokenReportUserDetailsReturn struct {
   
    // User name for the session.
    
 	UserName *string `json:"user_name,omitempty"`
   
    // The ID of the primary Active Directory user group the user belongs to. RBAC is granted based on this group.
    
 	PrimaryGroupID *string `json:"primary_group_id,omitempty"`
   
    // The primary Active Directory user group the user belongs to. RBAC is granted based on this group.
    
 	PrimaryGroupName *string `json:"primary_group_name,omitempty"`
   
    // The number of Active Directory user groups the user belongs to.
    
   	GroupCount *int64 `json:"group_count,omitempty"`
   
    // Role of the user.
    
   	Role *NsUserRoles `json:"role,omitempty"`
   
    // The amount of time that the user session is inactive before timing out.
    
   	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`
   
    // Global ID of the user.
    
 	UserID *string `json:"user_id,omitempty"`
   
    // The list of Active Directory groups the user belongs to.
    
	Groups []*string `json:"groups,omitempty"`
}
