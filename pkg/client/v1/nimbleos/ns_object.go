// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObject - Arbitrary object.
// Export NsObjectFields for advance operations like search filter etc.
var NsObjectFields *NsObject

func init() {
	IDfield := "id"

	NsObjectFields = &NsObject{
		ID: &IDfield,
	}
}

type NsObject struct {
	// ID - ID of object.
	ID *string `json:"id,omitempty"`
}
