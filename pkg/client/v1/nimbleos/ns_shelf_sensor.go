// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfSensor - A shelf sensor data.
// Export NsShelfSensorFields for advance operations like search filter etc.
var NsShelfSensorFields *NsShelfSensor

func init() {
	Namefield := "name"
	DisplayNamefield := "display_name"
	Locationfield := "location"

	NsShelfSensorFields = &NsShelfSensor{
		Name:        &Namefield,
		DisplayName: &DisplayNamefield,
		Location:    &Locationfield,
	}
}

type NsShelfSensor struct {
	// Type - Type of the sensor.
	Type *NsShelfSensorType `json:"type,omitempty"`
	// Name - Internal name of the sensor.
	Name *string `json:"name,omitempty"`
	// DisplayName - Name for display purpose.
	DisplayName *string `json:"display_name,omitempty"`
	// Location - Location of the sensor.
	Location *string `json:"location,omitempty"`
	// Cid - Which controller this sensor applies to.
	Cid *NsShelfCtrlrSide `json:"cid,omitempty"`
	// Value - Value of the sensor reading.
	Value *int64 `json:"value,omitempty"`
	// Status - Sensor status.
	Status *NsShelfSensorState `json:"status,omitempty"`
}
