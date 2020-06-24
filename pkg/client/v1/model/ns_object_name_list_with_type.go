/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsObjectNameListWithType - List of objects of a given type.
// Export NsObjectNameListWithTypeFields for advance operations like search filter etc.
var NsObjectNameListWithTypeFields *NsObjectNameListWithType

func init(){
		
	NsObjectNameListWithTypeFields= &NsObjectNameListWithType{
		
	}
}

type NsObjectNameListWithType struct {
   
    // Type of the object.
    
   	ObjType *NsObjectType `json:"obj_type,omitempty"`
   
    // List of object names.
    
	ObjNameList []*string `json:"obj_name_list,omitempty"`
}
