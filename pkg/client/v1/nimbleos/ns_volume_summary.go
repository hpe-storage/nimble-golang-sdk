// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolumeSummary - Select fields containing volume info.
// Export NsVolumeSummaryFields for advance operations like search filter etc.
var NsVolumeSummaryFields *NsVolumeSummary

func init() {

	NsVolumeSummaryFields = &NsVolumeSummary{
		ID:      "id",
		VolId:   "vol_id",
		Name:    "name",
		VolName: "vol_name",
	}
}

type NsVolumeSummary struct {
	// ID - ID of volume.
	ID string `json:"id,omitempty"`
	// VolId - ID of volume.
	VolId string `json:"vol_id,omitempty"`
	// Name - Name of volume.
	Name string `json:"name,omitempty"`
	// VolName - Name of volume.
	VolName string `json:"vol_name,omitempty"`
}

// sdk internal struct
type nsVolumeSummary struct {
	ID      *string `json:"id,omitempty"`
	VolId   *string `json:"vol_id,omitempty"`
	Name    *string `json:"name,omitempty"`
	VolName *string `json:"vol_name,omitempty"`
}

// EncodeNsVolumeSummary - Transform NsVolumeSummary to nsVolumeSummary type
func EncodeNsVolumeSummary(request interface{}) (*nsVolumeSummary, error) {
	reqNsVolumeSummary := request.(*NsVolumeSummary)
	byte, err := json.Marshal(reqNsVolumeSummary)
	if err != nil {
		return nil, err
	}
	respNsVolumeSummaryPtr := &nsVolumeSummary{}
	err = json.Unmarshal(byte, respNsVolumeSummaryPtr)
	return respNsVolumeSummaryPtr, err
}

// DecodeNsVolumeSummary Transform nsVolumeSummary to NsVolumeSummary type
func DecodeNsVolumeSummary(request interface{}) (*NsVolumeSummary, error) {
	reqNsVolumeSummary := request.(*nsVolumeSummary)
	byte, err := json.Marshal(reqNsVolumeSummary)
	if err != nil {
		return nil, err
	}
	respNsVolumeSummary := &NsVolumeSummary{}
	err = json.Unmarshal(byte, respNsVolumeSummary)
	return respNsVolumeSummary, err
}
