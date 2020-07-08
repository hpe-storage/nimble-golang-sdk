// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsTimezonesReturn - Group timezone list attribute.
// Export NsTimezonesReturnFields for advance operations like search filter etc.
var NsTimezonesReturnFields *NsTimezonesReturn

func init() {

	NsTimezonesReturnFields = &NsTimezonesReturn{}
}

type NsTimezonesReturn struct {
	// Timezones - Group timezone list.
	Timezones []*string `json:"timezones,omitempty"`
}

// sdk internal struct
type nsTimezonesReturn struct {
	// Timezones - Group timezone list.
	Timezones []*string `json:"timezones,omitempty"`
}

// EncodeNsTimezonesReturn - Transform NsTimezonesReturn to nsTimezonesReturn type
func EncodeNsTimezonesReturn(request interface{}) (*nsTimezonesReturn, error) {
	reqNsTimezonesReturn := request.(*NsTimezonesReturn)
	byte, err := json.Marshal(reqNsTimezonesReturn)
	objPtr := &nsTimezonesReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsTimezonesReturn Transform nsTimezonesReturn to NsTimezonesReturn type
func DecodeNsTimezonesReturn(request interface{}) (*NsTimezonesReturn, error) {
	reqNsTimezonesReturn := request.(*nsTimezonesReturn)
	byte, err := json.Marshal(reqNsTimezonesReturn)
	obj := &NsTimezonesReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
