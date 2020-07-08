// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsKeyValue - Key-value pair.
// Export NsKeyValueFields for advance operations like search filter etc.
var NsKeyValueFields *NsKeyValue

func init() {

	NsKeyValueFields = &NsKeyValue{
		Key:   "key",
		Value: "value",
	}
}

type NsKeyValue struct {
	// Key - Identifier of key-value pair.
	Key string `json:"key,omitempty"`
	// Value - Value of key-value pair.
	Value string `json:"value,omitempty"`
}

// sdk internal struct
type nsKeyValue struct {
	// Key - Identifier of key-value pair.
	Key *string `json:"key,omitempty"`
	// Value - Value of key-value pair.
	Value *string `json:"value,omitempty"`
}

// EncodeNsKeyValue - Transform NsKeyValue to nsKeyValue type
func EncodeNsKeyValue(request interface{}) (*nsKeyValue, error) {
	reqNsKeyValue := request.(*NsKeyValue)
	byte, err := json.Marshal(reqNsKeyValue)
	objPtr := &nsKeyValue{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsKeyValue Transform nsKeyValue to NsKeyValue type
func DecodeNsKeyValue(request interface{}) (*NsKeyValue, error) {
	reqNsKeyValue := request.(*nsKeyValue)
	byte, err := json.Marshal(reqNsKeyValue)
	obj := &NsKeyValue{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
