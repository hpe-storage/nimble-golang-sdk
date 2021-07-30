// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectWithID - An object with an ID.

// Export NsObjectWithIDFields provides field names to use in filter parameters, for example.
var NsObjectWithIDFields *NsObjectWithIDStringFields

func init() {
	fieldID := "id"

	NsObjectWithIDFields = &NsObjectWithIDStringFields{
		ID: &fieldID,
	}
}

type NsObjectWithID struct {
	// ID - ID of object.
	ID *string `json:"id,omitempty"`
}

// Struct for NsObjectWithIDFields
type NsObjectWithIDStringFields struct {
	ID *string
}
