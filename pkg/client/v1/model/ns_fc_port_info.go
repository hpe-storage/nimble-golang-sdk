/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFcPortInfo


// NsFcPortInfo :
type NsFcPortInfo struct {
   // Name
   Name string `json:"name,omitempty"`
   // BusLocation
   BusLocation string `json:"bus_location,omitempty"`
   // Port
   Port float64 `json:"port,omitempty"`
   // Slot
   Slot float64 `json:"slot,omitempty"`
}
