// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsCtrlrHwSensorInfo - Information on a controller hardware sensor.
// Export NsCtrlrHwSensorInfoFields for advance operations like search filter etc.
var NsCtrlrHwSensorInfoFields *NsCtrlrHwSensorInfo

func init() {

	NsCtrlrHwSensorInfoFields = &NsCtrlrHwSensorInfo{
		Name:        "name",
		DisplayName: "display_name",
		Location:    "location",
	}
}

type NsCtrlrHwSensorInfo struct {
	// Name - A uniquely identifying name.
	Name string `json:"name,omitempty"`
	// DisplayName - A human readable name for the sensor.
	DisplayName string `json:"display_name,omitempty"`
	// Location - The location of this sensor.
	Location string `json:"location,omitempty"`
	// CtrlrOwner - The controller owning this sensor.
	CtrlrOwner *NsControllerId `json:"ctrlr_owner,omitempty"`
	// State - The current state of this sensor.
	State *NsSensorState `json:"state,omitempty"`
	// CurrentReading - A sensor type specific reading (RPM for fans, degrees celcius for temperature).
	CurrentReading int64 `json:"current_reading,omitempty"`
}

// sdk internal struct
type nsCtrlrHwSensorInfo struct {
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

// EncodeNsCtrlrHwSensorInfo - Transform NsCtrlrHwSensorInfo to nsCtrlrHwSensorInfo type
func EncodeNsCtrlrHwSensorInfo(request interface{}) (*nsCtrlrHwSensorInfo, error) {
	reqNsCtrlrHwSensorInfo := request.(*NsCtrlrHwSensorInfo)
	byte, err := json.Marshal(reqNsCtrlrHwSensorInfo)
	objPtr := &nsCtrlrHwSensorInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsCtrlrHwSensorInfo Transform nsCtrlrHwSensorInfo to NsCtrlrHwSensorInfo type
func DecodeNsCtrlrHwSensorInfo(request interface{}) (*NsCtrlrHwSensorInfo, error) {
	reqNsCtrlrHwSensorInfo := request.(*nsCtrlrHwSensorInfo)
	byte, err := json.Marshal(reqNsCtrlrHwSensorInfo)
	obj := &NsCtrlrHwSensorInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
