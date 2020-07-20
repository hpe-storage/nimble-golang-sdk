// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolumeListReturn - Object containing a list of volume names and IDs.
// Export NsVolumeListReturnFields for advance operations like search filter etc.
var NsVolumeListReturnFields *NsVolumeListReturn

func init() {

	NsVolumeListReturnFields = &NsVolumeListReturn{}
}

type NsVolumeListReturn struct {
	// VolList - A list of volume names and IDs.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}

// sdk internal struct
type nsVolumeListReturn struct {
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}

// EncodeNsVolumeListReturn - Transform NsVolumeListReturn to nsVolumeListReturn type
func EncodeNsVolumeListReturn(request interface{}) (*nsVolumeListReturn, error) {
	reqNsVolumeListReturn := request.(*NsVolumeListReturn)
	byte, err := json.Marshal(reqNsVolumeListReturn)
	if err != nil {
		return nil, err
	}
	respNsVolumeListReturnPtr := &nsVolumeListReturn{}
	err = json.Unmarshal(byte, respNsVolumeListReturnPtr)
	return respNsVolumeListReturnPtr, err
}

// DecodeNsVolumeListReturn Transform nsVolumeListReturn to NsVolumeListReturn type
func DecodeNsVolumeListReturn(request interface{}) (*NsVolumeListReturn, error) {
	reqNsVolumeListReturn := request.(*nsVolumeListReturn)
	byte, err := json.Marshal(reqNsVolumeListReturn)
	if err != nil {
		return nil, err
	}
	respNsVolumeListReturn := &NsVolumeListReturn{}
	err = json.Unmarshal(byte, respNsVolumeListReturn)
	return respNsVolumeListReturn, err
}
