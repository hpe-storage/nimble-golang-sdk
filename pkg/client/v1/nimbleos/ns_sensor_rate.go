// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorRate - Rate stats for a sensor.
// Export NsSensorRateFields for advance operations like search filter etc.
var NsSensorRateFields *NsSensorRate

func init() {
	Namefield := "name"

	NsSensorRateFields = &NsSensorRate{
		Name: &Namefield,
	}
}

type NsSensorRate struct {
	// Name - Sensor name.
	Name *string `json:"name,omitempty"`
	// Rate - Sensor value.
	Rate *float64 `json:"rate,omitempty"`
}
