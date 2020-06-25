// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsVolumeBulkUpdateAttr - Volume object used in bulk update.
// Export NsVolumeBulkUpdateAttrFields for advance operations like search filter etc.
var NsVolumeBulkUpdateAttrFields *NsVolumeBulkUpdateAttr

func init(){
	IDfield:= "id"
	FolderIdfield:= "folder_id"
		
	NsVolumeBulkUpdateAttrFields= &NsVolumeBulkUpdateAttr{
		ID:       &IDfield,
		FolderId: &FolderIdfield,
	}
}

type NsVolumeBulkUpdateAttr struct {
	// ID - ID of volume.
 	ID *string `json:"id,omitempty"`
	// FolderId - ID of folder.
 	FolderId *string `json:"folder_id,omitempty"`
	// Online - Online state of the volume.
    Online *bool `json:"online,omitempty"`
}
