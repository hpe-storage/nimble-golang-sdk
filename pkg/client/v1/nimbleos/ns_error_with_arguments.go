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
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Severity - Severity of the error.
	Severity *NsApiSeverityLevel `json:"severity,omitempty"`
	// Text - Full error message with argument populated.
	Text *string `json:"text,omitempty"`
}

// EncodeNsErrorWithArguments - Transform NsErrorWithArguments to nsErrorWithArguments type
func EncodeNsErrorWithArguments(request interface{}) (*nsErrorWithArguments, error) {
	reqNsErrorWithArguments := request.(*NsErrorWithArguments)
	byte, err := json.Marshal(reqNsErrorWithArguments)
	objPtr := &nsErrorWithArguments{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsErrorWithArguments Transform nsErrorWithArguments to NsErrorWithArguments type
func DecodeNsErrorWithArguments(request interface{}) (*NsErrorWithArguments, error) {
	reqNsErrorWithArguments := request.(*nsErrorWithArguments)
	byte, err := json.Marshal(reqNsErrorWithArguments)
	obj := &NsErrorWithArguments{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
