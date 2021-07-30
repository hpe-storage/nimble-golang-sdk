// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorRate - Rate stats for a sensor.

// Export NsSensorRateFields provides field names to use in filter parameters, for example.
var NsSensorRateFields *NsSensorRateStringFields

func init() {
	fieldName := "name"
	fieldRate := "rate"

	NsSensorRateFields = &NsSensorRateStringFields{
		Name: &fieldName,
		Rate: &fieldRate,
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
