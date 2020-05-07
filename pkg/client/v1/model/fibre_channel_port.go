/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/FibreChannelPort


// FibreChannelPort :
type FibreChannelPort struct {
   // ID
   ID string `json:"id,omitempty"`
   // ArrayNameOrSerial
   ArrayNameOrSerial string `json:"array_name_or_serial,omitempty"`
   // ControllerName
   ControllerName string `json:"controller_name,omitempty"`
   // FcPortName
   FcPortName string `json:"fc_port_name,omitempty"`
   // BusLocation
   BusLocation string `json:"bus_location,omitempty"`
   // Port
   Port float64 `json:"port,omitempty"`
   // Slot
   Slot float64 `json:"slot,omitempty"`
   // RxPower
   RxPower float64 `json:"rx_power,omitempty"`
   // TxPower
   TxPower float64 `json:"tx_power,omitempty"`
}
