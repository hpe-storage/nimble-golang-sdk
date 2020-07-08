// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsEulaReturn - Return end-user license information.
// Export NsEulaReturnFields for advance operations like search filter etc.
var NsEulaReturnFields *NsEulaReturn

func init() {

	NsEulaReturnFields = &NsEulaReturn{
		Eula: "eula",
	}
}

type NsEulaReturn struct {
	// Eula - License information.
	Eula string `json:"eula,omitempty"`
}

// sdk internal struct
type nsEulaReturn struct {
	// Eula - License information.
	Eula *string `json:"eula,omitempty"`
}

// EncodeNsEulaReturn - Transform NsEulaReturn to nsEulaReturn type
func EncodeNsEulaReturn(request interface{}) (*nsEulaReturn, error) {
	reqNsEulaReturn := request.(*NsEulaReturn)
	byte, err := json.Marshal(reqNsEulaReturn)
	objPtr := &nsEulaReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsEulaReturn Transform nsEulaReturn to NsEulaReturn type
func DecodeNsEulaReturn(request interface{}) (*NsEulaReturn, error) {
	reqNsEulaReturn := request.(*nsEulaReturn)
	byte, err := json.Marshal(reqNsEulaReturn)
	obj := &NsEulaReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
