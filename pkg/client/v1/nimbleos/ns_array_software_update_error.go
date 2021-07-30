// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsArraySoftwareUpdateErrorFields provides field names to use in filter parameters, for example.
var NsArraySoftwareUpdateErrorFields *NsArraySoftwareUpdateErrorFieldHandles

func init() {
	NsArraySoftwareUpdateErrorFields = &NsArraySoftwareUpdateErrorFieldHandles{
		Error: "error",
	}
}

// NsArraySoftwareUpdateError - Software update error for specific controller.
type NsArraySoftwareUpdateError struct {
	// Error - Error code from software update.
	Error *string `json:"error,omitempty"`
}

// NsArraySoftwareUpdateErrorFieldHandles provides a string representation for each AccessControlRecord field.
type NsArraySoftwareUpdateErrorFieldHandles struct {
	Error string
}
