// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsObjectFields provides field names to use in filter parameters, for example.
var NsObjectFields *NsObjectFieldHandles

func init() {
	fieldID := "id"

	NsObjectFields = &NsObjectFieldHandles{
		ID: &fieldID,
	}
}

// NsObject - Arbitrary object.
type NsObject struct {
	// ID - ID of object.
	ID *string `json:"id,omitempty"`
}

// NsObjectFieldHandles provides a string representation for each AccessControlRecord field.
type NsObjectFieldHandles struct {
	ID *string
}
