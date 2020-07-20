// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	Name *string  `json:"name,omitempty"`
	Rate *float32 `json:"rate,omitempty"`
}

// EncodeNsSensorRate - Transform NsSensorRate to nsSensorRate type
func EncodeNsSensorRate(request interface{}) (*nsSensorRate, error) {
	reqNsSensorRate := request.(*NsSensorRate)
	byte, err := json.Marshal(reqNsSensorRate)
	if err != nil {
		return nil, err
	}
	respNsSensorRatePtr := &nsSensorRate{}
	err = json.Unmarshal(byte, respNsSensorRatePtr)
	return respNsSensorRatePtr, err
}

// DecodeNsSensorRate Transform nsSensorRate to NsSensorRate type
func DecodeNsSensorRate(request interface{}) (*NsSensorRate, error) {
	reqNsSensorRate := request.(*nsSensorRate)
	byte, err := json.Marshal(reqNsSensorRate)
	if err != nil {
		return nil, err
	}
	respNsSensorRate := &NsSensorRate{}
	err = json.Unmarshal(byte, respNsSensorRate)
	return respNsSensorRate, err
}
