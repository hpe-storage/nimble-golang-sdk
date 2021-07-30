// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsVssRespFields provides field names to use in filter parameters, for example.
var NsVssRespFields *NsVssRespFieldHandles

func init() {
	fieldVssError := "vss_error"
	fieldVssErrorMessage := "vss_error_message"

	NsVssRespFields = &NsVssRespFieldHandles{
		VssError:        &fieldVssError,
		VssErrorMessage: &fieldVssErrorMessage,
	}
}

// NsVssResp - Response from VSS app server.
type NsVssResp struct {
	// VssError - Error code from VSS app server.
	VssError *string `json:"vss_error,omitempty"`
	// VssErrorMessage - Detailed error message from VSS app server.
	VssErrorMessage *string `json:"vss_error_message,omitempty"`
}

// NsVssRespFieldHandles provides a string representation for each AccessControlRecord field.
type NsVssRespFieldHandles struct {
	VssError        *string
	VssErrorMessage *string
}
