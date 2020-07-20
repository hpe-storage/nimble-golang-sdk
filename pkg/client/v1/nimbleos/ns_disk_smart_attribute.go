// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsDiskSmartAttribute - One Smart attribute of a disk.
// Export NsDiskSmartAttributeFields for advance operations like search filter etc.
var NsDiskSmartAttributeFields *NsDiskSmartAttribute

func init() {

	NsDiskSmartAttributeFields = &NsDiskSmartAttribute{
		Name: "name",
	}
}

type NsDiskSmartAttribute struct {
	// Name - Name of the Smart attribute.
	Name string `json:"name,omitempty"`
	// SmartId - Smart attribute ID.
	SmartId int64 `json:"smart_id,omitempty"`
	// CurValue - Current value.
	CurValue int64 `json:"cur_value,omitempty"`
	// WorstValue - Worst value.
	WorstValue int64 `json:"worst_value,omitempty"`
	// Threshold - Smart threshold.
	Threshold int64 `json:"threshold,omitempty"`
	// Flags - Smart flags.
	Flags int64 `json:"flags,omitempty"`
	// RawValue - Raw value.
	RawValue int64 `json:"raw_value,omitempty"`
	// LastUpdatedEpochSecs - Last update time in epoch seconds.
	LastUpdatedEpochSecs int64 `json:"last_updated_epoch_secs,omitempty"`
}

// sdk internal struct
type nsDiskSmartAttribute struct {
	Name                 *string `json:"name,omitempty"`
	SmartId              *int64  `json:"smart_id,omitempty"`
	CurValue             *int64  `json:"cur_value,omitempty"`
	WorstValue           *int64  `json:"worst_value,omitempty"`
	Threshold            *int64  `json:"threshold,omitempty"`
	Flags                *int64  `json:"flags,omitempty"`
	RawValue             *int64  `json:"raw_value,omitempty"`
	LastUpdatedEpochSecs *int64  `json:"last_updated_epoch_secs,omitempty"`
}

// EncodeNsDiskSmartAttribute - Transform NsDiskSmartAttribute to nsDiskSmartAttribute type
func EncodeNsDiskSmartAttribute(request interface{}) (*nsDiskSmartAttribute, error) {
	reqNsDiskSmartAttribute := request.(*NsDiskSmartAttribute)
	byte, err := json.Marshal(reqNsDiskSmartAttribute)
	if err != nil {
		return nil, err
	}
	respNsDiskSmartAttributePtr := &nsDiskSmartAttribute{}
	err = json.Unmarshal(byte, respNsDiskSmartAttributePtr)
	return respNsDiskSmartAttributePtr, err
}

// DecodeNsDiskSmartAttribute Transform nsDiskSmartAttribute to NsDiskSmartAttribute type
func DecodeNsDiskSmartAttribute(request interface{}) (*NsDiskSmartAttribute, error) {
	reqNsDiskSmartAttribute := request.(*nsDiskSmartAttribute)
	byte, err := json.Marshal(reqNsDiskSmartAttribute)
	if err != nil {
		return nil, err
	}
	respNsDiskSmartAttribute := &NsDiskSmartAttribute{}
	err = json.Unmarshal(byte, respNsDiskSmartAttribute)
	return respNsDiskSmartAttribute, err
}
