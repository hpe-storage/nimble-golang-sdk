// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeBulkUpdateAttr - Volume object used in bulk update.
// Export NsVolumeBulkUpdateAttrFields for advance operations like search filter etc.
var NsVolumeBulkUpdateAttrFields *NsVolumeBulkUpdateAttrStringFields

func init() {
	IDfield := "id"
	FolderIdfield := "folder_id"
	Onlinefield := "online"

	NsVolumeBulkUpdateAttrFields = &NsVolumeBulkUpdateAttrStringFields{
		ID:       &IDfield,
		FolderId: &FolderIdfield,
		Online:   &Onlinefield,
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

// Struct for NsVolumeBulkUpdateAttrFields
type NsVolumeBulkUpdateAttrStringFields struct {
	ID       *string
	FolderId *string
	Online   *string
}
