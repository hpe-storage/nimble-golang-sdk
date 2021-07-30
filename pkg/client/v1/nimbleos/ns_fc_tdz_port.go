// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFcTdzPortFields provides field names to use in filter parameters, for example.
var NsFcTdzPortFields *NsFcTdzPortFieldHandles

func init() {
	fieldArrayName := "array_name"
	fieldFcName := "fc_name"

	NsFcTdzPortFields = &NsFcTdzPortFieldHandles{
		ArrayName: &fieldArrayName,
		FcName:    &fieldFcName,
	}
}

// NsFcTdzPort - Fibre Channel Target port.
type NsFcTdzPort struct {
	// ArrayName - Unique name of the array.
	ArrayName *string `json:"array_name,omitempty"`
	// FcName - Target port interface name.
	FcName *string `json:"fc_name,omitempty"`
}

// NsFcTdzPortFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcTdzPortFieldHandles struct {
	ArrayName *string
	FcName    *string
}
