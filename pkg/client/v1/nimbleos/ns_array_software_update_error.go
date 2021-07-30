// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArraySoftwareUpdateError - Software update error for specific controller.
// Export NsArraySoftwareUpdateErrorFields for advance operations like search filter etc.
var NsArraySoftwareUpdateErrorFields *NsArraySoftwareUpdateErrorStringFields

func init() {
	Errorfield := "error"

	NsArraySoftwareUpdateErrorFields = &NsArraySoftwareUpdateErrorStringFields{
		Error: &Errorfield,
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
