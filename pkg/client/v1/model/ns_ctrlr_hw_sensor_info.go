/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsCtrlrHwSensorInfo - Information on a controller hardware sensor.
// Export NsCtrlrHwSensorInfoFields for advance operations like search filter etc.
var NsCtrlrHwSensorInfoFields *NsCtrlrHwSensorInfo

func init(){
	Namefield:= "name"
	DisplayNamefield:= "display_name"
	Locationfield:= "location"
		
	NsCtrlrHwSensorInfoFields= &NsCtrlrHwSensorInfo{
		Name: &Namefield,
		DisplayName: &DisplayNamefield,
		Location: &Locationfield,
		
	}
}

type NsCtrlrHwSensorInfo struct {
   
    // A uniquely identifying name.
    
 	Name *string `json:"name,omitempty"`
   
    // A human readable name for the sensor.
    
 	DisplayName *string `json:"display_name,omitempty"`
   
    // The location of this sensor.
    
 	Location *string `json:"location,omitempty"`
   
    // The controller owning this sensor.
    
   	CtrlrOwner *NsControllerID `json:"ctrlr_owner,omitempty"`
   
    // The current state of this sensor.
    
   	State *NsSensorState `json:"state,omitempty"`
   
    // A sensor type specific reading (RPM for fans, degrees celcius for temperature).
    
  	CurrentReading  *int64 `json:"current_reading,omitempty"`
}
