// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcTdzPortFields provides field names to use in filter parameters, for example.
var NsFcTdzPortFields *NsFcTdzPortFieldHandles

func init() {
	NsFcTdzPortFields = &NsFcTdzPortFieldHandles{
		ArrayName: "array_name",
		FcName:    "fc_name",
	}
}

// NsFcTdzPort - Fibre Channel Target port.
type NsFcTdzPort struct {
	// ArrayName - Unique name of the array.
	ArrayName *string `json:"array_name,omitempty"`
	// FcName - Target port interface name.
	FcName *string `json:"fc_name,omitempty"`
}

// NsFcTdzPortFieldHandles provides a string representation for each NsFcTdzPort field.
type NsFcTdzPortFieldHandles struct {
	ArrayName string
	FcName    string
}
