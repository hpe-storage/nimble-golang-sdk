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
	// VolList - A list of volume names and IDs.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}

// EncodeNsVolumeListReturn - Transform NsVolumeListReturn to nsVolumeListReturn type
func EncodeNsVolumeListReturn(request interface{}) (*nsVolumeListReturn, error) {
	reqNsVolumeListReturn := request.(*NsVolumeListReturn)
	byte, err := json.Marshal(reqNsVolumeListReturn)
	objPtr := &nsVolumeListReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsVolumeListReturn Transform nsVolumeListReturn to NsVolumeListReturn type
func DecodeNsVolumeListReturn(request interface{}) (*NsVolumeListReturn, error) {
	reqNsVolumeListReturn := request.(*nsVolumeListReturn)
	byte, err := json.Marshal(reqNsVolumeListReturn)
	obj := &NsVolumeListReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
