// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrHwSensorInfo - Information on a controller hardware sensor.
// Export NsCtrlrHwSensorInfoFields for advance operations like search filter etc.
var NsCtrlrHwSensorInfoFields *NsCtrlrHwSensorInfoStringFields

func init() {
	Namefield := "name"
	DisplayNamefield := "display_name"
	Locationfield := "location"
	CtrlrOwnerfield := "ctrlr_owner"
	Statefield := "state"
	CurrentReadingfield := "current_reading"

	NsCtrlrHwSensorInfoFields = &NsCtrlrHwSensorInfoStringFields{
		Name:           &Namefield,
		DisplayName:    &DisplayNamefield,
		Location:       &Locationfield,
		CtrlrOwner:     &CtrlrOwnerfield,
		State:          &Statefield,
		CurrentReading: &CurrentReadingfield,
	}
}

type NsCtrlrHwSensorInfo struct {
	// Name - A uniquely identifying name.
	Name *string `json:"name,omitempty"`
	// DisplayName - A human readable name for the sensor.
	DisplayName *string `json:"display_name,omitempty"`
	// Location - The location of this sensor.
	Location *string `json:"location,omitempty"`
	// CtrlrOwner - The controller owning this sensor.
	CtrlrOwner *NsControllerId `json:"ctrlr_owner,omitempty"`
	// State - The current state of this sensor.
	State *NsSensorState `json:"state,omitempty"`
	// CurrentReading - A sensor type specific reading (RPM for fans, degrees celcius for temperature).
	CurrentReading *int64 `json:"current_reading,omitempty"`
}

// Struct for NsCtrlrHwSensorInfoFields
type NsCtrlrHwSensorInfoStringFields struct {
	Name           *string
	DisplayName    *string
	Location       *string
	CtrlrOwner     *string
	State          *string
	CurrentReading *string
}
