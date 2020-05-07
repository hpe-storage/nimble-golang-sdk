/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsPinnedVolumeInfo


// NsPinnedVolumeInfo :
type NsPinnedVolumeInfo struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // SizePinnedCacheBytes
   SizePinnedCacheBytes float64 `json:"size_pinned_cache_bytes,omitempty"`
}
