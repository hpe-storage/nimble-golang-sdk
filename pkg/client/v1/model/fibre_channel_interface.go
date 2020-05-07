/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/FibreChannelInterface


// FibreChannelInterface :
type FibreChannelInterface struct {
   // ID
   ID string `json:"id,omitempty"`
   // ArrayNameOrSerial
   ArrayNameOrSerial string `json:"array_name_or_serial,omitempty"`
   // PartialResponseOk
   PartialResponseOk bool `json:"partial_response_ok,omitempty"`
   // ControllerName
   ControllerName string `json:"controller_name,omitempty"`
   // FcPortID
   FcPortID string `json:"fc_port_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Wwnn
   Wwnn string `json:"wwnn,omitempty"`
   // Wwpn
   Wwpn string `json:"wwpn,omitempty"`
   // Online
   Online bool `json:"online,omitempty"`
   // FirmwareVersion
   FirmwareVersion string `json:"firmware_version,omitempty"`
   // LogicalPortNumber
   LogicalPortNumber float64 `json:"logical_port_number,omitempty"`
   // FcPortName
   FcPortName string `json:"fc_port_name,omitempty"`
   // BusLocation
   BusLocation string `json:"bus_location,omitempty"`
   // Slot
   Slot float64 `json:"slot,omitempty"`
   // Port
   Port float64 `json:"port,omitempty"`
}
