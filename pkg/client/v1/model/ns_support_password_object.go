// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSupportPasswordObject - Support password blob for a user.
// Export NsSupportPasswordObjectFields for advance operations like search filter etc.
var NsSupportPasswordObjectFields *NsSupportPasswordObject

func init() {

	NsSupportPasswordObjectFields = &NsSupportPasswordObject{
		Username: "username",
		Blob:     "blob",
	}
}

type NsSupportPasswordObject struct {
	// Username - The username for the account the password blob relates to.
	Username string `json:"username,omitempty"`
	// Blob - The ciphertext blob holding the randomly produced password.
	Blob string `json:"blob,omitempty"`
}

// sdk internal struct
type nsSupportPasswordObject struct {
	// Username - The username for the account the password blob relates to.
	Username *string `json:"username,omitempty"`
	// Blob - The ciphertext blob holding the randomly produced password.
	Blob *string `json:"blob,omitempty"`
}

// EncodeNsSupportPasswordObject - Transform NsSupportPasswordObject to nsSupportPasswordObject type
func EncodeNsSupportPasswordObject(request interface{}) (*nsSupportPasswordObject, error) {
	reqNsSupportPasswordObject := request.(*NsSupportPasswordObject)
	byte, err := json.Marshal(reqNsSupportPasswordObject)
	objPtr := &nsSupportPasswordObject{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSupportPasswordObject Transform nsSupportPasswordObject to NsSupportPasswordObject type
func DecodeNsSupportPasswordObject(request interface{}) (*NsSupportPasswordObject, error) {
	reqNsSupportPasswordObject := request.(*nsSupportPasswordObject)
	byte, err := json.Marshal(reqNsSupportPasswordObject)
	obj := &NsSupportPasswordObject{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
