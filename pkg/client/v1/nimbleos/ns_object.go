// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectFields provides field names to use in filter parameters, for example.
var NsObjectFields *NsObjectFieldHandles

func init() {
	NsObjectFields = &NsObjectFieldHandles{
		ID: "id",
	}
}

// NsObject - Arbitrary object.
type NsObject struct {
	// ID - ID of object.
	ID *string `json:"id,omitempty"`
}

// NsObjectFieldHandles provides a string representation for each NsObject field.
type NsObjectFieldHandles struct {
	ID string
}
