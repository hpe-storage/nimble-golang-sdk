/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFcInterfaceInfo


// NsFcInterfaceInfo :
type NsFcInterfaceInfo struct {
   // Name
   Name string `json:"name,omitempty"`
   // Wwnn
   Wwnn string `json:"wwnn,omitempty"`
   // Wwpn
   Wwpn string `json:"wwpn,omitempty"`
   // Online
   Online bool `json:"online,omitempty"`
   // BusLocation
   BusLocation string `json:"bus_location,omitempty"`
   // Slot
   Slot float64 `json:"slot,omitempty"`
   // Port
   Port float64 `json:"port,omitempty"`
}
