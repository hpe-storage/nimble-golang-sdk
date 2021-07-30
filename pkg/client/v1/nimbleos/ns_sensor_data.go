// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorData - A list of samples (in order of sample time) for a sensor.

// Export NsSensorDataFields provides field names to use in filter parameters, for example.
var NsSensorDataFields *NsSensorDataStringFields

func init() {
	fieldSensor := "sensor"
	fieldSamples := "samples"

	NsSensorDataFields = &NsSensorDataStringFields{
		Sensor:  &fieldSensor,
		Samples: &fieldSamples,
	}
}

type NsSensorData struct {
	// Sensor - Sensor name.
	Sensor *string `json:"sensor,omitempty"`
	// Samples - A list of samples for the sensor.
	Samples []*uint64 `json:"samples,omitempty"`
}

// Struct for NsSensorDataFields
type NsSensorDataStringFields struct {
	Sensor  *string
	Samples *string
}
