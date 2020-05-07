/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NetworkInterface


// NetworkInterface :
type NetworkInterface struct {
   // ID
   ID string `json:"id,omitempty"`
   // ArrayNameOrSerial
   ArrayNameOrSerial string `json:"array_name_or_serial,omitempty"`
   // PartialResponseOk
   PartialResponseOk bool `json:"partial_response_ok,omitempty"`
   // ArrayID
   ArrayID string `json:"array_id,omitempty"`
   // ControllerName
   ControllerName string `json:"controller_name,omitempty"`
   // ControllerID
   ControllerID string `json:"controller_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Mac
   Mac string `json:"mac,omitempty"`
   // IsPresent
   IsPresent bool `json:"is_present,omitempty"`
   // Mtu
   Mtu float64 `json:"mtu,omitempty"`
   // Port
   Port float64 `json:"port,omitempty"`
   // Slot
   Slot float64 `json:"slot,omitempty"`
}
