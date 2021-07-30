// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorRate - Rate stats for a sensor.
// Export NsSensorRateFields for advance operations like search filter etc.
var NsSensorRateFields *NsSensorRateStringFields

func init() {
	Namefield := "name"
	Ratefield := "rate"

	NsSensorRateFields = &NsSensorRateStringFields{
		Name: &Namefield,
		Rate: &Ratefield,
	}
}

type NsSensorRate struct {
	// Name - Sensor name.
	Name *string `json:"name,omitempty"`
	// Rate - Sensor value.
	Rate *float64 `json:"rate,omitempty"`
}

// Struct for NsSensorRateFields
type NsSensorRateStringFields struct {
	Name *string
	Rate *string
}
