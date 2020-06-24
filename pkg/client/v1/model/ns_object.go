/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsObject - Arbitrary object.
// Export NsObjectFields for advance operations like search filter etc.
var NsObjectFields *NsObject

func init(){
	IDfield:= "id"
		
	NsObjectFields= &NsObject{
		ID: &IDfield,
		
	}
}

type NsObject struct {
   
    // ID of object.
    
 	ID *string `json:"id,omitempty"`
}
