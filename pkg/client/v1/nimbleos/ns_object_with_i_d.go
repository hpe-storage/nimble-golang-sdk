// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsObjectWithIDFields provides field names to use in filter parameters, for example.
var NsObjectWithIDFields *NsObjectWithIDFieldHandles

func init() {
	NsObjectWithIDFields = &NsObjectWithIDFieldHandles{
		ID: "id",
	}
}

// NsObjectWithID - An object with an ID.
type NsObjectWithID struct {
	// ID - ID of object.
	ID *string `json:"id,omitempty"`
}

// NsObjectWithIDFieldHandles provides a string representation for each AccessControlRecord field.
type NsObjectWithIDFieldHandles struct {
	ID string
}
