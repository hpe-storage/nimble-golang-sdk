// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSnapVolListReturn - Object returned after creating snapshot collection.
// Export NsSnapVolListReturnFields for advance operations like search filter etc.
var NsSnapVolListReturnFields *NsSnapVolListReturn

func init() {

	NsSnapVolListReturnFields = &NsSnapVolListReturn{}
}

type NsSnapVolListReturn struct {
	// SnapIds - A list of snapshot ids.
	SnapIds []*NsObjectIDKV `json:"snap_ids,omitempty"`
}

// sdk internal struct
type nsSnapVolListReturn struct {
	SnapIds []*NsObjectIDKV `json:"snap_ids,omitempty"`
}

// EncodeNsSnapVolListReturn - Transform NsSnapVolListReturn to nsSnapVolListReturn type
func EncodeNsSnapVolListReturn(request interface{}) (*nsSnapVolListReturn, error) {
	reqNsSnapVolListReturn := request.(*NsSnapVolListReturn)
	byte, err := json.Marshal(reqNsSnapVolListReturn)
	if err != nil {
		return nil, err
	}
	respNsSnapVolListReturnPtr := &nsSnapVolListReturn{}
	err = json.Unmarshal(byte, respNsSnapVolListReturnPtr)
	return respNsSnapVolListReturnPtr, err
}

// DecodeNsSnapVolListReturn Transform nsSnapVolListReturn to NsSnapVolListReturn type
func DecodeNsSnapVolListReturn(request interface{}) (*NsSnapVolListReturn, error) {
	reqNsSnapVolListReturn := request.(*nsSnapVolListReturn)
	byte, err := json.Marshal(reqNsSnapVolListReturn)
	if err != nil {
		return nil, err
	}
	respNsSnapVolListReturn := &NsSnapVolListReturn{}
	err = json.Unmarshal(byte, respNsSnapVolListReturn)
	return respNsSnapVolListReturn, err
}
