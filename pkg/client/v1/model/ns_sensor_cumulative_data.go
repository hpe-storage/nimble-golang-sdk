// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSensorCumulativeData - Stat sensor cumulative data.
// Export NsSensorCumulativeDataFields for advance operations like search filter etc.
var NsSensorCumulativeDataFields *NsSensorCumulativeData

func init() {

	NsSensorCumulativeDataFields = &NsSensorCumulativeData{
		Name: "name",
	}
}

type NsSensorCumulativeData struct {
	// Name - Name of sensor.
	Name string `json:"name,omitempty"`
	// Index - Index of sensor in stat for versioning.
	Index int64 `json:"index,omitempty"`
	// Msec - Creation time of stat sensor in millisecs.
	Msec int64 `json:"msec,omitempty"`
	// PrevUnavail - Previous stat value is not valid.
	PrevUnavail *bool `json:"prev_unavail,omitempty"`
	// CurrUnavail - Current stat value is not valid.
	CurrUnavail *bool `json:"curr_unavail,omitempty"`
	// Curr - Current stat value.
	Curr int64 `json:"curr,omitempty"`
	// Prev - Previous stat value from last sample period.
	Prev int64 `json:"prev,omitempty"`
}

// sdk internal struct
type nsSensorCumulativeData struct {
	// Name - Name of sensor.
	Name *string `json:"name,omitempty"`
	// Index - Index of sensor in stat for versioning.
	Index *int64 `json:"index,omitempty"`
	// Msec - Creation time of stat sensor in millisecs.
	Msec *int64 `json:"msec,omitempty"`
	// PrevUnavail - Previous stat value is not valid.
	PrevUnavail *bool `json:"prev_unavail,omitempty"`
	// CurrUnavail - Current stat value is not valid.
	CurrUnavail *bool `json:"curr_unavail,omitempty"`
	// Curr - Current stat value.
	Curr *int64 `json:"curr,omitempty"`
	// Prev - Previous stat value from last sample period.
	Prev *int64 `json:"prev,omitempty"`
}

// EncodeNsSensorCumulativeData - Transform NsSensorCumulativeData to nsSensorCumulativeData type
func EncodeNsSensorCumulativeData(request interface{}) (*nsSensorCumulativeData, error) {
	reqNsSensorCumulativeData := request.(*NsSensorCumulativeData)
	byte, err := json.Marshal(reqNsSensorCumulativeData)
	objPtr := &nsSensorCumulativeData{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSensorCumulativeData Transform nsSensorCumulativeData to NsSensorCumulativeData type
func DecodeNsSensorCumulativeData(request interface{}) (*NsSensorCumulativeData, error) {
	reqNsSensorCumulativeData := request.(*nsSensorCumulativeData)
	byte, err := json.Marshal(reqNsSensorCumulativeData)
	obj := &NsSensorCumulativeData{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
