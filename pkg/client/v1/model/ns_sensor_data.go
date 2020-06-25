// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsSensorData - A list of samples (in order of sample time) for a sensor.
// Export NsSensorDataFields for advance operations like search filter etc.
var NsSensorDataFields *NsSensorData

func init(){
	Sensorfield:= "sensor"
		
	NsSensorDataFields= &NsSensorData{
		Sensor: &Sensorfield,
	}
}

type NsSensorData struct {
	// Sensor - Sensor name.
 	Sensor *string `json:"sensor,omitempty"`
	// Samples - A list of samples for the sensor.
	Samples []*float64 `json:"samples,omitempty"`
}
