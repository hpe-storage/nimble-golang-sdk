// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSensorRateFields provides field names to use in filter parameters, for example.
var NsSensorRateFields *NsSensorRateFieldHandles

func init() {
	fieldName := "name"
	fieldRate := "rate"

	NsSensorRateFields = &NsSensorRateFieldHandles{
		Name: &fieldName,
		Rate: &fieldRate,
	}
}

// NsSensorRate - Rate stats for a sensor.
type NsSensorRate struct {
	// Name - Sensor name.
	Name *string `json:"name,omitempty"`
	// Rate - Sensor value.
	Rate *float64 `json:"rate,omitempty"`
}

// NsSensorRateFieldHandles provides a string representation for each AccessControlRecord field.
type NsSensorRateFieldHandles struct {
	Name *string
	Rate *string
}
