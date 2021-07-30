// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsGenericResp - Response from generic app server.
// Export NsGenericRespFields for advance operations like search filter etc.
var NsGenericRespFields *NsGenericRespStringFields

func init() {
	GenericErrorfield := "generic_error"
	GenericErrorMessagefield := "generic_error_message"
	ConnStatusOkfield := "conn_status_ok"
	ConnMessagefield := "conn_message"

	NsGenericRespFields = &NsGenericRespStringFields{
		GenericError:        &GenericErrorfield,
		GenericErrorMessage: &GenericErrorMessagefield,
		ConnStatusOk:        &ConnStatusOkfield,
		ConnMessage:         &ConnMessagefield,
	}
}

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

// Struct for NsGenericRespFields
type NsGenericRespStringFields struct {
	GenericError        *string
	GenericErrorMessage *string
	ConnStatusOk        *string
	ConnMessage         *string
}
