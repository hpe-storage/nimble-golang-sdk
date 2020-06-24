/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

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
   
    // Sensor name.
    
 	Sensor *string `json:"sensor,omitempty"`
   
    // A list of samples for the sensor.
    
	Samples []*float64 `json:"samples,omitempty"`
}
