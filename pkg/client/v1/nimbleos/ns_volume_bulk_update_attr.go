// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeBulkUpdateAttr - Volume object used in bulk update.

// Export NsVolumeBulkUpdateAttrFields provides field names to use in filter parameters, for example.
var NsVolumeBulkUpdateAttrFields *NsVolumeBulkUpdateAttrStringFields

func init() {
	fieldID := "id"
	fieldFolderId := "folder_id"
	fieldOnline := "online"

	NsVolumeBulkUpdateAttrFields = &NsVolumeBulkUpdateAttrStringFields{
		ID:       &fieldID,
		FolderId: &fieldFolderId,
		Online:   &fieldOnline,
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
