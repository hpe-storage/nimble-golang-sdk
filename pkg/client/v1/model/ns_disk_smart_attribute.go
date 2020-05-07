/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsDiskSmartAttribute


// NsDiskSmartAttribute :
type NsDiskSmartAttribute struct {
   // Name
   Name string `json:"name,omitempty"`
   // SmartID
   SmartID float64 `json:"smart_id,omitempty"`
   // CurValue
   CurValue float64 `json:"cur_value,omitempty"`
   // WorstValue
   WorstValue float64 `json:"worst_value,omitempty"`
   // Threshold
   Threshold float64 `json:"threshold,omitempty"`
   // Flags
   Flags float64 `json:"flags,omitempty"`
   // RawValue
   RawValue float64 `json:"raw_value,omitempty"`
   // LastUpdatedEpochSecs
   LastUpdatedEpochSecs float64 `json:"last_updated_epoch_secs,omitempty"`
}
