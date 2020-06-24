/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSensorRate - Rate stats for a sensor.
// Export NsSensorRateFields for advance operations like search filter etc.
var NsSensorRateFields *NsSensorRate

func init(){
	Namefield:= "name"
		
	NsSensorRateFields= &NsSensorRate{
		Name: &Namefield,
		
	}
}

type NsSensorRate struct {
   
    // Sensor name.
    
 	Name *string `json:"name,omitempty"`
   
    // Sensor value.
    
  	Rate *float32 `json:"rate,omitempty"`
}
