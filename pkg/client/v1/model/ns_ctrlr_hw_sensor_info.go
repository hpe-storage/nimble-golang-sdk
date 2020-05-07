/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsCtrlrHwSensorInfo


// NsCtrlrHwSensorInfo :
type NsCtrlrHwSensorInfo struct {
   // Name
   Name string `json:"name,omitempty"`
   // DisplayName
   DisplayName string `json:"display_name,omitempty"`
   // Location
   Location string `json:"location,omitempty"`
   // CurrentReading
   CurrentReading float64 `json:"current_reading,omitempty"`
}
