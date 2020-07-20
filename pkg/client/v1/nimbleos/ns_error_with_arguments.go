// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsErrorWithArguments - Non-negative integer in range [0,9000].
// Export NsErrorWithArgumentsFields for advance operations like search filter etc.
var NsErrorWithArgumentsFields *NsErrorWithArguments

func init() {

	NsErrorWithArgumentsFields = &NsErrorWithArguments{
		Code: "code",
		Text: "text",
	}
}

type NsErrorWithArguments struct {
	// Code - Error code.
	Code string `json:"code,omitempty"`
	// Severity - Severity of the error.
	Severity *NsApiSeverityLevel `json:"severity,omitempty"`
	// Text - Full error message with argument populated.
	Text string `json:"text,omitempty"`
}

// sdk internal struct
type nsErrorWithArguments struct {
	Code     *string             `json:"code,omitempty"`
	Severity *NsApiSeverityLevel `json:"severity,omitempty"`
	Text     *string             `json:"text,omitempty"`
}

// EncodeNsErrorWithArguments - Transform NsErrorWithArguments to nsErrorWithArguments type
func EncodeNsErrorWithArguments(request interface{}) (*nsErrorWithArguments, error) {
	reqNsErrorWithArguments := request.(*NsErrorWithArguments)
	byte, err := json.Marshal(reqNsErrorWithArguments)
	if err != nil {
		return nil, err
	}
	respNsErrorWithArgumentsPtr := &nsErrorWithArguments{}
	err = json.Unmarshal(byte, respNsErrorWithArgumentsPtr)
	return respNsErrorWithArgumentsPtr, err
}

// DecodeNsErrorWithArguments Transform nsErrorWithArguments to NsErrorWithArguments type
func DecodeNsErrorWithArguments(request interface{}) (*NsErrorWithArguments, error) {
	reqNsErrorWithArguments := request.(*nsErrorWithArguments)
	byte, err := json.Marshal(reqNsErrorWithArguments)
	if err != nil {
		return nil, err
	}
	respNsErrorWithArguments := &NsErrorWithArguments{}
	err = json.Unmarshal(byte, respNsErrorWithArguments)
	return respNsErrorWithArguments, err
}
