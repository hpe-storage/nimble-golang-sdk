// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorDataFields provides field names to use in filter parameters, for example.
var NsSensorDataFields *NsSensorDataFieldHandles

func init() {
	NsSensorDataFields = &NsSensorDataFieldHandles{
		Sensor:  "sensor",
		Samples: "samples",
	}
}

// NsSensorData - A list of samples (in order of sample time) for a sensor.
type NsSensorData struct {
	// Sensor - Sensor name.
	Sensor *string `json:"sensor,omitempty"`
	// Samples - A list of samples for the sensor.
	Samples []*uint64 `json:"samples,omitempty"`
}

// NsSensorDataFieldHandles provides a string representation for each NsSensorData field.
type NsSensorDataFieldHandles struct {
	Sensor  string
	Samples string
}
