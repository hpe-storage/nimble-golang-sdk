// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObject - Arbitrary object.

// Export NsObjectFields provides field names to use in filter parameters, for example.
var NsObjectFields *NsObjectStringFields

func init() {
	fieldID := "id"

	NsObjectFields = &NsObjectStringFields{
		ID: &fieldID,
	}
}

type NsObject struct {
	// ID - ID of object.
	ID *string `json:"id,omitempty"`
}

// Struct for NsObjectFields
type NsObjectStringFields struct {
	ID *string
}
