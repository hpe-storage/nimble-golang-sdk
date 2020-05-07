/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSensorCumulativeData


// NsSensorCumulativeData :
type NsSensorCumulativeData struct {
   // Name
   Name string `json:"name,omitempty"`
   // Index
   Index float64 `json:"index,omitempty"`
   // Msec
   Msec float64 `json:"msec,omitempty"`
   // PrevUnavail
   PrevUnavail bool `json:"prev_unavail,omitempty"`
   // CurrUnavail
   CurrUnavail bool `json:"curr_unavail,omitempty"`
   // Curr
   Curr float64 `json:"curr,omitempty"`
   // Prev
   Prev float64 `json:"prev,omitempty"`
}
