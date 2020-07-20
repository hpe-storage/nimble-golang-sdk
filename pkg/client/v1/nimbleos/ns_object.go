// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObject - Arbitrary object.
// Export NsObjectFields for advance operations like search filter etc.
var NsObjectFields *NsObject

func init() {

	NsObjectFields = &NsObject{
		ID: "id",
	}
}

type NsObject struct {
	// ID - ID of object.
	ID string `json:"id,omitempty"`
}

// sdk internal struct
type nsObject struct {
	ID *string `json:"id,omitempty"`
}

// EncodeNsObject - Transform NsObject to nsObject type
func EncodeNsObject(request interface{}) (*nsObject, error) {
	reqNsObject := request.(*NsObject)
	byte, err := json.Marshal(reqNsObject)
	if err != nil {
		return nil, err
	}
	respNsObjectPtr := &nsObject{}
	err = json.Unmarshal(byte, respNsObjectPtr)
	return respNsObjectPtr, err
}

// DecodeNsObject Transform nsObject to NsObject type
func DecodeNsObject(request interface{}) (*NsObject, error) {
	reqNsObject := request.(*nsObject)
	byte, err := json.Marshal(reqNsObject)
	if err != nil {
		return nil, err
	}
	respNsObject := &NsObject{}
	err = json.Unmarshal(byte, respNsObject)
	return respNsObject, err
}
