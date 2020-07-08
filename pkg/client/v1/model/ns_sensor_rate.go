// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSensorRate - Rate stats for a sensor.
// Export NsSensorRateFields for advance operations like search filter etc.
var NsSensorRateFields *NsSensorRate

func init() {

	NsSensorRateFields = &NsSensorRate{
		Name: "name",
	}
}

type NsSensorRate struct {
	// Name - Sensor name.
	Name string `json:"name,omitempty"`
	// Rate - Sensor value.
	Rate float32 `json:"rate,omitempty"`
}

// sdk internal struct
type nsSensorRate struct {
	// Name - Sensor name.
	Name *string `json:"name,omitempty"`
	// Rate - Sensor value.
	Rate *float32 `json:"rate,omitempty"`
}

// EncodeNsSensorRate - Transform NsSensorRate to nsSensorRate type
func EncodeNsSensorRate(request interface{}) (*nsSensorRate, error) {
	reqNsSensorRate := request.(*NsSensorRate)
	byte, err := json.Marshal(reqNsSensorRate)
	objPtr := &nsSensorRate{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSensorRate Transform nsSensorRate to NsSensorRate type
func DecodeNsSensorRate(request interface{}) (*NsSensorRate, error) {
	reqNsSensorRate := request.(*nsSensorRate)
	byte, err := json.Marshal(reqNsSensorRate)
	obj := &NsSensorRate{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
