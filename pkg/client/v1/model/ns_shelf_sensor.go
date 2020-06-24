/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsShelfSensor - A shelf sensor data.
// Export NsShelfSensorFields for advance operations like search filter etc.
var NsShelfSensorFields *NsShelfSensor

func init(){
	Namefield:= "name"
	DisplayNamefield:= "display_name"
	Locationfield:= "location"
		
	NsShelfSensorFields= &NsShelfSensor{
		Name: &Namefield,
		DisplayName: &DisplayNamefield,
		Location: &Locationfield,
		
	}
}

type NsShelfSensor struct {
   
    // Type of the sensor.
    
   	Type *NsShelfSensorType `json:"type,omitempty"`
   
    // Internal name of the sensor.
    
 	Name *string `json:"name,omitempty"`
   
    // Name for display purpose.
    
 	DisplayName *string `json:"display_name,omitempty"`
   
    // Location of the sensor.
    
 	Location *string `json:"location,omitempty"`
   
    // Which controller this sensor applies to.
    
   	CID *NsShelfCtrlrSIDe `json:"cid,omitempty"`
   
    // Value of the sensor reading.
    
   	Value *int64 `json:"value,omitempty"`
   
    // Sensor status.
    
   	Status *NsShelfSensorState `json:"status,omitempty"`
}
