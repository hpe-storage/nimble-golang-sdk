/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsDiskSetAttr


// NsDiskSetAttr :
type NsDiskSetAttr struct {
   // Driveset
   Driveset float64 `json:"driveset,omitempty"`
   // IsFlashShelf
   IsFlashShelf bool `json:"is_flash_shelf,omitempty"`
   // IsCapacityValID
   IsCapacityValID bool `json:"is_capacity_valid,omitempty"`
   // UsableCapacity
   UsableCapacity float64 `json:"usable_capacity,omitempty"`
   // RawCapacity
   RawCapacity float64 `json:"raw_capacity,omitempty"`
   // UsableCacheCapacity
   UsableCacheCapacity float64 `json:"usable_cache_capacity,omitempty"`
   // RawCacheCapacity
   RawCacheCapacity float64 `json:"raw_cache_capacity,omitempty"`
}
