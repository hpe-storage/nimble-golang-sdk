/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumeBulkUpdateAttr - Volume object used in bulk update.
// Export NsVolumeBulkUpdateAttrFields for advance operations like search filter etc.
var NsVolumeBulkUpdateAttrFields *NsVolumeBulkUpdateAttr

func init(){
	IDfield:= "id"
	FolderIDfield:= "folder_id"
		
	NsVolumeBulkUpdateAttrFields= &NsVolumeBulkUpdateAttr{
		ID: &IDfield,
		FolderID: &FolderIDfield,
		
	}
}

type NsVolumeBulkUpdateAttr struct {
   
    // ID of volume.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of folder.
    
 	FolderID *string `json:"folder_id,omitempty"`
   
    // Online state of the volume.
    
 	Online *bool `json:"online,omitempty"`
}
