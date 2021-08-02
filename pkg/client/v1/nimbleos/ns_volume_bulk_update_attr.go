// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeBulkUpdateAttrFields provides field names to use in filter parameters, for example.
var NsVolumeBulkUpdateAttrFields *NsVolumeBulkUpdateAttrFieldHandles

func init() {
	NsVolumeBulkUpdateAttrFields = &NsVolumeBulkUpdateAttrFieldHandles{
		ID:       "id",
		FolderId: "folder_id",
		Online:   "online",
	}
}

// NsVolumeBulkUpdateAttr - Volume object used in bulk update.
type NsVolumeBulkUpdateAttr struct {
	// ID - ID of volume.
	ID *string `json:"id,omitempty"`
	// FolderId - ID of folder.
	FolderId *string `json:"folder_id,omitempty"`
	// Online - Online state of the volume.
	Online *bool `json:"online,omitempty"`
}

// NsVolumeBulkUpdateAttrFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolumeBulkUpdateAttrFieldHandles struct {
	ID       string
	FolderId string
	Online   string
}
