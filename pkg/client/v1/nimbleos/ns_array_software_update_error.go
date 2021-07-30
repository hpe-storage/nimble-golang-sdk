// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArraySoftwareUpdateError - Software update error for specific controller.

// Export NsArraySoftwareUpdateErrorFields provides field names to use in filter parameters, for example.
var NsArraySoftwareUpdateErrorFields *NsArraySoftwareUpdateErrorStringFields

func init() {
	fieldError := "error"

	NsArraySoftwareUpdateErrorFields = &NsArraySoftwareUpdateErrorStringFields{
		Error: &fieldError,
	}
}

type NsArraySoftwareUpdateError struct {
	// Error - Error code from software update.
	Error *string `json:"error,omitempty"`
}

// Struct for NsArraySoftwareUpdateErrorFields
type NsArraySoftwareUpdateErrorStringFields struct {
	Error *string
}
