// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorData - A list of samples (in order of sample time) for a sensor.
// Export NsSensorDataFields for advance operations like search filter etc.
var NsSensorDataFields *NsSensorDataStringFields

func init() {
	Sensorfield := "sensor"
	Samplesfield := "samples"

	NsSensorDataFields = &NsSensorDataStringFields{
		Sensor:  &Sensorfield,
		Samples: &Samplesfield,
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
