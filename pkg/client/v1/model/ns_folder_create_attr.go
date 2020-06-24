/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFolderCreateAttr - Attributes for folder creation.
// Export NsFolderCreateAttrFields for advance operations like search filter etc.
var NsFolderCreateAttrFields *NsFolderCreateAttr

func init(){
	Namefield:= "name"
	PoolIDfield:= "pool_id"
		
	NsFolderCreateAttrFields= &NsFolderCreateAttr{
		Name: &Namefield,
		PoolID: &PoolIDfield,
		
	}
}

type NsFolderCreateAttr struct {
   
    // Name of folder.
    
 	Name *string `json:"name,omitempty"`
   
    // ID of pool to create the folder in.
    
 	PoolID *string `json:"pool_id,omitempty"`
}
