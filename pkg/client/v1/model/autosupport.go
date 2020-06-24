/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Autosupport - Get status of autosupport.
// Export AutosupportFields for advance operations like search filter etc.
var AutosupportFields *Autosupport

func init(){
	IDfield:= "id"
	GroupIDfield:= "group_id"
	GroupNamefield:= "group_name"
		
	AutosupportFields= &Autosupport{
		ID: &IDfield,
		GroupID: &GroupIDfield,
		GroupName: &GroupNamefield,
		
	}
}

type Autosupport struct {
   
    // Identifier of the autosupport.
    
 	ID *string `json:"id,omitempty"`
   
    // List of arrays in the group with autosupport information.
    
   	ArrayList []*NsArrayAsupDetail `json:"array_list,omitempty"`
   
    // Number of arrays in the group.
    
   	ArrayCount *int64 `json:"array_count,omitempty"`
   
    // Identifier for the group.
    
 	GroupID *string `json:"group_id,omitempty"`
   
    // Name of the group.
    
 	GroupName *string `json:"group_name,omitempty"`
}
