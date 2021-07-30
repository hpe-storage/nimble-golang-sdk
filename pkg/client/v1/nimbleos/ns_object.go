// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObject - Arbitrary object.
// Export NsObjectFields for advance operations like search filter etc.
var NsObjectFields *NsObjectStringFields

func init() {
	IDfield := "id"

	NsObjectFields = &NsObjectStringFields{
		ID: &IDfield,
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
