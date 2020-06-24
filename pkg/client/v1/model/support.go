/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Support - View and alter support-based parameters.
// Export SupportFields for advance operations like search filter etc.
var SupportFields *Support

func init(){
	IDfield:= "id"
		
	SupportFields= &Support{
		ID: &IDfield,
		
	}
}

type Support struct {
   
    // Object identifier for the group.
    
 	ID *string `json:"id,omitempty"`
   
    // Mode for support password.
    
   	PasswordMode *NsSupportPasswordMode `json:"password_mode,omitempty"`
   
    // Count of arrays for support password.
    
   	ArrayCount *int64 `json:"array_count,omitempty"`
   
    // Details of support passwords for arrays.
    
   	ArrayList []*NsSupportPasswordArray `json:"array_list,omitempty"`
}
