// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsLunReturn - Return LU number.
// Export NsLunReturnFields for advance operations like search filter etc.
var NsLunReturnFields *NsLunReturn

func init() {

	NsLunReturnFields = &NsLunReturn{}
}

type NsLunReturn struct {
	// Lun - LU number in hexadecimal.
	Lun int64 `json:"lun,omitempty"`
	// LuNumber - LU number in decimal.
	LuNumber int64 `json:"lu_number,omitempty"`
}

// sdk internal struct
type nsLunReturn struct {
	Lun      *int64 `json:"lun,omitempty"`
	LuNumber *int64 `json:"lu_number,omitempty"`
}

// EncodeNsLunReturn - Transform NsLunReturn to nsLunReturn type
func EncodeNsLunReturn(request interface{}) (*nsLunReturn, error) {
	reqNsLunReturn := request.(*NsLunReturn)
	byte, err := json.Marshal(reqNsLunReturn)
	if err != nil {
		return nil, err
	}
	respNsLunReturnPtr := &nsLunReturn{}
	err = json.Unmarshal(byte, respNsLunReturnPtr)
	return respNsLunReturnPtr, err
}

// DecodeNsLunReturn Transform nsLunReturn to NsLunReturn type
func DecodeNsLunReturn(request interface{}) (*NsLunReturn, error) {
	reqNsLunReturn := request.(*nsLunReturn)
	byte, err := json.Marshal(reqNsLunReturn)
	if err != nil {
		return nil, err
	}
	respNsLunReturn := &NsLunReturn{}
	err = json.Unmarshal(byte, respNsLunReturn)
	return respNsLunReturn, err
}
