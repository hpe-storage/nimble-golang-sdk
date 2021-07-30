// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsShelfSensorFields provides field names to use in filter parameters, for example.
var NsShelfSensorFields *NsShelfSensorFieldHandles

func init() {
	NsShelfSensorFields = &NsShelfSensorFieldHandles{
		Type:        "type",
		Name:        "name",
		DisplayName: "display_name",
		Location:    "location",
		Cid:         "cid",
		Value:       "value",
		Status:      "status",
	}
}

// NsShelfSensor - A shelf sensor data.
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

// NsShelfSensorFieldHandles provides a string representation for each AccessControlRecord field.
type NsShelfSensorFieldHandles struct {
	Type        string
	Name        string
	DisplayName string
	Location    string
	Cid         string
	Value       string
	Status      string
}
