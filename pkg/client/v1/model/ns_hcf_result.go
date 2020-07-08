// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsHcfResult - Results from health check of a single element.
// Export NsHcfResultFields for advance operations like search filter etc.
var NsHcfResultFields *NsHcfResult

func init() {

	NsHcfResultFields = &NsHcfResult{
		ElementName: "element_name",
	}
}

type NsHcfResult struct {
	// ElementName - Name of the element.
	ElementName string `json:"element_name,omitempty"`
	// ErrorList - List of health check errors for this element.
	ErrorList []*string `json:"error_list,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// sdk internal struct
type nsHcfResult struct {
	// ElementName - Name of the element.
	ElementName *string `json:"element_name,omitempty"`
	// ErrorList - List of health check errors for this element.
	ErrorList []*string `json:"error_list,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// EncodeNsHcfResult - Transform NsHcfResult to nsHcfResult type
func EncodeNsHcfResult(request interface{}) (*nsHcfResult, error) {
	reqNsHcfResult := request.(*NsHcfResult)
	byte, err := json.Marshal(reqNsHcfResult)
	objPtr := &nsHcfResult{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsHcfResult Transform nsHcfResult to NsHcfResult type
func DecodeNsHcfResult(request interface{}) (*NsHcfResult, error) {
	reqNsHcfResult := request.(*nsHcfResult)
	byte, err := json.Marshal(reqNsHcfResult)
	obj := &NsHcfResult{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
