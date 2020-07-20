// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObjectWithID - An object with an ID.
// Export NsObjectWithIDFields for advance operations like search filter etc.
var NsObjectWithIDFields *NsObjectWithID

func init() {

	NsObjectWithIDFields = &NsObjectWithID{
		ID: "id",
	}
}

type NsObjectWithID struct {
	// ID - ID of object.
	ID string `json:"id,omitempty"`
}

// sdk internal struct
type nsObjectWithID struct {
	ID *string `json:"id,omitempty"`
}

// EncodeNsObjectWithID - Transform NsObjectWithID to nsObjectWithID type
func EncodeNsObjectWithID(request interface{}) (*nsObjectWithID, error) {
	reqNsObjectWithID := request.(*NsObjectWithID)
	byte, err := json.Marshal(reqNsObjectWithID)
	if err != nil {
		return nil, err
	}
	respNsObjectWithIDPtr := &nsObjectWithID{}
	err = json.Unmarshal(byte, respNsObjectWithIDPtr)
	return respNsObjectWithIDPtr, err
}

// DecodeNsObjectWithID Transform nsObjectWithID to NsObjectWithID type
func DecodeNsObjectWithID(request interface{}) (*NsObjectWithID, error) {
	reqNsObjectWithID := request.(*nsObjectWithID)
	byte, err := json.Marshal(reqNsObjectWithID)
	if err != nil {
		return nil, err
	}
	respNsObjectWithID := &NsObjectWithID{}
	err = json.Unmarshal(byte, respNsObjectWithID)
	return respNsObjectWithID, err
}
