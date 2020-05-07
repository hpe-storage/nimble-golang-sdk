/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArraySummaryInfo


// NsArraySummaryInfo :
type NsArraySummaryInfo struct {
   // Name
   Name string `json:"name,omitempty"`
   // Version
   Version string `json:"version,omitempty"`
   // Serial
   Serial string `json:"serial,omitempty"`
   // Model
   Model string `json:"model,omitempty"`
   // CountOfFcPorts
   CountOfFcPorts float64 `json:"count_of_fc_ports,omitempty"`
   // AllFlash
   AllFlash bool `json:"all_flash,omitempty"`
}
