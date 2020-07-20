// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVssResp - Response from VSS app server.
// Export NsVssRespFields for advance operations like search filter etc.
var NsVssRespFields *NsVssResp

func init() {

	NsVssRespFields = &NsVssResp{
		VssError:        "vss_error",
		VssErrorMessage: "vss_error_message",
	}
}

type NsVssResp struct {
	// VssError - Error code from VSS app server.
	VssError string `json:"vss_error,omitempty"`
	// VssErrorMessage - Detailed error message from VSS app server.
	VssErrorMessage string `json:"vss_error_message,omitempty"`
}

// sdk internal struct
type nsVssResp struct {
	VssError        *string `json:"vss_error,omitempty"`
	VssErrorMessage *string `json:"vss_error_message,omitempty"`
}

// EncodeNsVssResp - Transform NsVssResp to nsVssResp type
func EncodeNsVssResp(request interface{}) (*nsVssResp, error) {
	reqNsVssResp := request.(*NsVssResp)
	byte, err := json.Marshal(reqNsVssResp)
	if err != nil {
		return nil, err
	}
	respNsVssRespPtr := &nsVssResp{}
	err = json.Unmarshal(byte, respNsVssRespPtr)
	return respNsVssRespPtr, err
}

// DecodeNsVssResp Transform nsVssResp to NsVssResp type
func DecodeNsVssResp(request interface{}) (*NsVssResp, error) {
	reqNsVssResp := request.(*nsVssResp)
	byte, err := json.Marshal(reqNsVssResp)
	if err != nil {
		return nil, err
	}
	respNsVssResp := &NsVssResp{}
	err = json.Unmarshal(byte, respNsVssResp)
	return respNsVssResp, err
}
