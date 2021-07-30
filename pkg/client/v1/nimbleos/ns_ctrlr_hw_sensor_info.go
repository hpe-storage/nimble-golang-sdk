// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrHwSensorInfo - Information on a controller hardware sensor.

// Export NsCtrlrHwSensorInfoFields provides field names to use in filter parameters, for example.
var NsCtrlrHwSensorInfoFields *NsCtrlrHwSensorInfoStringFields

func init() {
	fieldName := "name"
	fieldDisplayName := "display_name"
	fieldLocation := "location"
	fieldCtrlrOwner := "ctrlr_owner"
	fieldState := "state"
	fieldCurrentReading := "current_reading"

	NsCtrlrHwSensorInfoFields = &NsCtrlrHwSensorInfoStringFields{
		Name:           &fieldName,
		DisplayName:    &fieldDisplayName,
		Location:       &fieldLocation,
		CtrlrOwner:     &fieldCtrlrOwner,
		State:          &fieldState,
		CurrentReading: &fieldCurrentReading,
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
