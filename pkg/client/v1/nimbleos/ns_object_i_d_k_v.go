// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObjectIDKV - A key value pair containing an object ID as a value.
// Export NsObjectIDKVFields for advance operations like search filter etc.
var NsObjectIDKVFields *NsObjectIDKV

func init() {

	NsObjectIDKVFields = &NsObjectIDKV{
		Key: "key",
		ID:  "id",
	}
}

type NsObjectIDKV struct {
	// Key - Identifier of key-value pair.
	Key string `json:"key,omitempty"`
	// ID - An object ID.
	ID string `json:"id,omitempty"`
}

// sdk internal struct
type nsObjectIDKV struct {
	// Key - Identifier of key-value pair.
	Key *string `json:"key,omitempty"`
	// ID - An object ID.
	ID *string `json:"id,omitempty"`
}

// EncodeNsObjectIDKV - Transform NsObjectIDKV to nsObjectIDKV type
func EncodeNsObjectIDKV(request interface{}) (*nsObjectIDKV, error) {
	reqNsObjectIDKV := request.(*NsObjectIDKV)
	byte, err := json.Marshal(reqNsObjectIDKV)
	objPtr := &nsObjectIDKV{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsObjectIDKV Transform nsObjectIDKV to NsObjectIDKV type
func DecodeNsObjectIDKV(request interface{}) (*NsObjectIDKV, error) {
	reqNsObjectIDKV := request.(*nsObjectIDKV)
	byte, err := json.Marshal(reqNsObjectIDKV)
	obj := &NsObjectIDKV{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
