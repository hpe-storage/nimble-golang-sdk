/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Stat


// Stat :
type Stat struct {
   // Scope
   Scope string `json:"scope,omitempty"`
   // DomainID
   DomainID string `json:"domain_id,omitempty"`
   // SetID
   SetID string `json:"set_id,omitempty"`
   // VolIDs
   VolIDs string `json:"vol_ids,omitempty"`
   // Sensors
   Sensors string `json:"sensors,omitempty"`
   // Starttime
   Starttime float64 `json:"starttime,omitempty"`
   // Endtime
   Endtime float64 `json:"endtime,omitempty"`
   // Interval
   Interval float64 `json:"interval,omitempty"`
   // Cumulative
   Cumulative bool `json:"cumulative,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // VolID
   VolID string `json:"vol_id,omitempty"`
}
