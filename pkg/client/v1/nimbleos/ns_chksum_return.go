// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsChksumReturn - Return computed checksum.
// Export NsChksumReturnFields for advance operations like search filter etc.
var NsChksumReturnFields *NsChksumReturn

func init() {

	NsChksumReturnFields = &NsChksumReturn{
		Cksum: "cksum",
	}
}

type NsChksumReturn struct {
	// Cksum - Computed checksum.
	Cksum string `json:"cksum,omitempty"`
}

// sdk internal struct
type nsChksumReturn struct {
	Cksum *string `json:"cksum,omitempty"`
}

// EncodeNsChksumReturn - Transform NsChksumReturn to nsChksumReturn type
func EncodeNsChksumReturn(request interface{}) (*nsChksumReturn, error) {
	reqNsChksumReturn := request.(*NsChksumReturn)
	byte, err := json.Marshal(reqNsChksumReturn)
	if err != nil {
		return nil, err
	}
	respNsChksumReturnPtr := &nsChksumReturn{}
	err = json.Unmarshal(byte, respNsChksumReturnPtr)
	return respNsChksumReturnPtr, err
}

// DecodeNsChksumReturn Transform nsChksumReturn to NsChksumReturn type
func DecodeNsChksumReturn(request interface{}) (*NsChksumReturn, error) {
	reqNsChksumReturn := request.(*nsChksumReturn)
	byte, err := json.Marshal(reqNsChksumReturn)
	if err != nil {
		return nil, err
	}
	respNsChksumReturn := &NsChksumReturn{}
	err = json.Unmarshal(byte, respNsChksumReturn)
	return respNsChksumReturn, err
}
