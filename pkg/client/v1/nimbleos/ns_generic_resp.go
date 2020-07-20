// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsGenericResp - Response from generic app server.
// Export NsGenericRespFields for advance operations like search filter etc.
var NsGenericRespFields *NsGenericResp

func init() {

	NsGenericRespFields = &NsGenericResp{
		GenericError:        "generic_error",
		GenericErrorMessage: "generic_error_message",
		ConnMessage:         "conn_message",
	}
}

type NsGenericResp struct {
	// GenericError - Error code from generic app server.
	GenericError string `json:"generic_error,omitempty"`
	// GenericErrorMessage - Detailed error message from generic app server.
	GenericErrorMessage string `json:"generic_error_message,omitempty"`
	// ConnStatusOk - Is the connection status OK.
	ConnStatusOk *bool `json:"conn_status_ok,omitempty"`
	// ConnMessage - Detailed connection message.
	ConnMessage string `json:"conn_message,omitempty"`
}

// sdk internal struct
type nsGenericResp struct {
	GenericError        *string `json:"generic_error,omitempty"`
	GenericErrorMessage *string `json:"generic_error_message,omitempty"`
	ConnStatusOk        *bool   `json:"conn_status_ok,omitempty"`
	ConnMessage         *string `json:"conn_message,omitempty"`
}

// EncodeNsGenericResp - Transform NsGenericResp to nsGenericResp type
func EncodeNsGenericResp(request interface{}) (*nsGenericResp, error) {
	reqNsGenericResp := request.(*NsGenericResp)
	byte, err := json.Marshal(reqNsGenericResp)
	if err != nil {
		return nil, err
	}
	respNsGenericRespPtr := &nsGenericResp{}
	err = json.Unmarshal(byte, respNsGenericRespPtr)
	return respNsGenericRespPtr, err
}

// DecodeNsGenericResp Transform nsGenericResp to NsGenericResp type
func DecodeNsGenericResp(request interface{}) (*NsGenericResp, error) {
	reqNsGenericResp := request.(*nsGenericResp)
	byte, err := json.Marshal(reqNsGenericResp)
	if err != nil {
		return nil, err
	}
	respNsGenericResp := &NsGenericResp{}
	err = json.Unmarshal(byte, respNsGenericResp)
	return respNsGenericResp, err
}
