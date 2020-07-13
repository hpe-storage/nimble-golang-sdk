// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSshKey - SSH key.
// Export NsSshKeyFields for advance operations like search filter etc.
var NsSshKeyFields *NsSshKey

func init() {

	NsSshKeyFields = &NsSshKey{
		KeyName: "key_name",
		KeyType: "key_type",
		Key:     "key",
	}
}

type NsSshKey struct {
	// KeyName - The user that owns the key.
	KeyName string `json:"key_name,omitempty"`
	// KeyType - The key type.
	KeyType string `json:"key_type,omitempty"`
	// Key - The key.
	Key string `json:"key,omitempty"`
}

// sdk internal struct
type nsSshKey struct {
	// KeyName - The user that owns the key.
	KeyName *string `json:"key_name,omitempty"`
	// KeyType - The key type.
	KeyType *string `json:"key_type,omitempty"`
	// Key - The key.
	Key *string `json:"key,omitempty"`
}

// EncodeNsSshKey - Transform NsSshKey to nsSshKey type
func EncodeNsSshKey(request interface{}) (*nsSshKey, error) {
	reqNsSshKey := request.(*NsSshKey)
	byte, err := json.Marshal(reqNsSshKey)
	objPtr := &nsSshKey{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSshKey Transform nsSshKey to NsSshKey type
func DecodeNsSshKey(request interface{}) (*NsSshKey, error) {
	reqNsSshKey := request.(*nsSshKey)
	byte, err := json.Marshal(reqNsSshKey)
	obj := &NsSshKey{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
