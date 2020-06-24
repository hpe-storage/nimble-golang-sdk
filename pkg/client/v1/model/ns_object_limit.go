/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsObjectLimit - Limits (Max total number of objects) for object of a given type.
// Export NsObjectLimitFields for advance operations like search filter etc.
var NsObjectLimitFields *NsObjectLimit

func init(){
		
	NsObjectLimitFields= &NsObjectLimit{
		
	}
}

type NsObjectLimit struct {
   
    // Type of the object.
    
   	ObjType *NsObjectType `json:"obj_type,omitempty"`
   
    // Limit of the objects.
    
   	ObjLimit *int64 `json:"obj_limit,omitempty"`
   
    // Number of objects after group merge.
    
   	ObjNum *int64 `json:"obj_num,omitempty"`
}
