// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsResponse - May contain anything that any REST API response can contain.
// Export NsResponseFields for advance operations like search filter etc.
var NsResponseFields *NsResponse

func init() {

	NsResponseFields = &NsResponse{}
}

type NsResponse struct {
	// Data - Response data.
	Data *NsObject `json:"data,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// sdk internal struct
type nsResponse struct {
	// Data - Response data.
	Data *NsObject `json:"data,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// EncodeNsResponse - Transform NsResponse to nsResponse type
func EncodeNsResponse(request interface{}) (*nsResponse, error) {
	reqNsResponse := request.(*NsResponse)
	byte, err := json.Marshal(reqNsResponse)
	objPtr := &nsResponse{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsResponse Transform nsResponse to NsResponse type
func DecodeNsResponse(request interface{}) (*NsResponse, error) {
	reqNsResponse := request.(*nsResponse)
	byte, err := json.Marshal(reqNsResponse)
	obj := &NsResponse{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
