// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsRequest - May contain anything that any REST API request can contain.
// Export NsRequestFields for advance operations like search filter etc.
var NsRequestFields *NsRequest

func init() {

	NsRequestFields = &NsRequest{
		Path: "path",
	}
}

type NsRequest struct {
	// Data - Request data.
	Data *NsObject `json:"data,omitempty"`
	// Method - HTTP method.
	Method int64 `json:"method,omitempty"`
	// Path - Path which identifies the target resource.
	Path string `json:"path,omitempty"`
}

// sdk internal struct
type nsRequest struct {
	Data   *NsObject `json:"data,omitempty"`
	Method *int64    `json:"method,omitempty"`
	Path   *string   `json:"path,omitempty"`
}

// EncodeNsRequest - Transform NsRequest to nsRequest type
func EncodeNsRequest(request interface{}) (*nsRequest, error) {
	reqNsRequest := request.(*NsRequest)
	byte, err := json.Marshal(reqNsRequest)
	if err != nil {
		return nil, err
	}
	respNsRequestPtr := &nsRequest{}
	err = json.Unmarshal(byte, respNsRequestPtr)
	return respNsRequestPtr, err
}

// DecodeNsRequest Transform nsRequest to NsRequest type
func DecodeNsRequest(request interface{}) (*NsRequest, error) {
	reqNsRequest := request.(*nsRequest)
	byte, err := json.Marshal(reqNsRequest)
	if err != nil {
		return nil, err
	}
	respNsRequest := &NsRequest{}
	err = json.Unmarshal(byte, respNsRequest)
	return respNsRequest, err
}
