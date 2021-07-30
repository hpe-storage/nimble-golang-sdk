// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVssResp - Response from VSS app server.

// Export NsVssRespFields provides field names to use in filter parameters, for example.
var NsVssRespFields *NsVssRespStringFields

func init() {
	fieldVssError := "vss_error"
	fieldVssErrorMessage := "vss_error_message"

	NsVssRespFields = &NsVssRespStringFields{
		VssError:        &fieldVssError,
		VssErrorMessage: &fieldVssErrorMessage,
	}
}

type NsVssResp struct {
	// VssError - Error code from VSS app server.
	VssError *string `json:"vss_error,omitempty"`
	// VssErrorMessage - Detailed error message from VSS app server.
	VssErrorMessage *string `json:"vss_error_message,omitempty"`
}

// Struct for NsVssRespFields
type NsVssRespStringFields struct {
	VssError        *string
	VssErrorMessage *string
}
