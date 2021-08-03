// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsGenericRespFields provides field names to use in filter parameters, for example.
var NsGenericRespFields *NsGenericRespFieldHandles

func init() {
	NsGenericRespFields = &NsGenericRespFieldHandles{
		GenericError:        "generic_error",
		GenericErrorMessage: "generic_error_message",
		ConnStatusOk:        "conn_status_ok",
		ConnMessage:         "conn_message",
	}
}

// NsGenericResp - Response from generic app server.
type NsGenericResp struct {
	// GenericError - Error code from generic app server.
	GenericError *string `json:"generic_error,omitempty"`
	// GenericErrorMessage - Detailed error message from generic app server.
	GenericErrorMessage *string `json:"generic_error_message,omitempty"`
	// ConnStatusOk - Is the connection status OK.
	ConnStatusOk *bool `json:"conn_status_ok,omitempty"`
	// ConnMessage - Detailed connection message.
	ConnMessage *string `json:"conn_message,omitempty"`
}

// NsGenericRespFieldHandles provides a string representation for each NsGenericResp field.
type NsGenericRespFieldHandles struct {
	GenericError        string
	GenericErrorMessage string
	ConnStatusOk        string
	ConnMessage         string
}
