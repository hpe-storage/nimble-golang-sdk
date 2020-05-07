/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArrayDetail


// NsArrayDetail :
type NsArrayDetail struct {
   // ID
   ID string `json:"id,omitempty"`
   // ArrayID
   ArrayID string `json:"array_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // EvacTime
   EvacTime float64 `json:"evac_time,omitempty"`
   // EvacUsage
   EvacUsage float64 `json:"evac_usage,omitempty"`
   // UsableCapacity
   UsableCapacity float64 `json:"usable_capacity,omitempty"`
   // Usage
   Usage float64 `json:"usage,omitempty"`
   // VolUsageCompressedBytes
   VolUsageCompressedBytes float64 `json:"vol_usage_compressed_bytes,omitempty"`
   // SnapUsageCompressedBytes
   SnapUsageCompressedBytes float64 `json:"snap_usage_compressed_bytes,omitempty"`
   // UsageValID
   UsageValID bool `json:"usage_valid,omitempty"`
}
