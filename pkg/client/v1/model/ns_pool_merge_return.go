/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsPoolMergeReturn


// NsPoolMergeReturn :
type NsPoolMergeReturn struct {
   // SnapcollCount
   SnapcollCount float64 `json:"snapcoll_count,omitempty"`
   // ArrayCount
   ArrayCount float64 `json:"array_count,omitempty"`
   // AllFlash
   AllFlash bool `json:"all_flash,omitempty"`
   // CacheCapacity
   CacheCapacity float64 `json:"cache_capacity,omitempty"`
   // Capacity
   Capacity float64 `json:"capacity,omitempty"`
   // CompressedSnapUsageBytes
   CompressedSnapUsageBytes float64 `json:"compressed_snap_usage_bytes,omitempty"`
   // CompressedVolUsageBytes
   CompressedVolUsageBytes float64 `json:"compressed_vol_usage_bytes,omitempty"`
   // CompressionRatio
   CompressionRatio float32 `json:"compression_ratio,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // DedupeRatio
   DedupeRatio float32 `json:"dedupe_ratio,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // FreeSpace
   FreeSpace float64 `json:"free_space,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // ID
   ID string `json:"id,omitempty"`
   // DedupeEnabled
   DedupeEnabled bool `json:"dedupe_enabled,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // PinnableCacheCapacity
   PinnableCacheCapacity float64 `json:"pinnable_cache_capacity,omitempty"`
   // PinnedCacheCapacity
   PinnedCacheCapacity float64 `json:"pinned_cache_capacity,omitempty"`
   // SavingsCompression
   SavingsCompression float64 `json:"savings_compression,omitempty"`
   // SavingsDedupe
   SavingsDedupe float64 `json:"savings_dedupe,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // SnapCompressionRatio
   SnapCompressionRatio float32 `json:"snap_compression_ratio,omitempty"`
   // SnapCount
   SnapCount float64 `json:"snap_count,omitempty"`
   // UncompressedSnapUsageBytes
   UncompressedSnapUsageBytes float64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   // UncompressedVolUsageBytes
   UncompressedVolUsageBytes float64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   // UnusedReserve
   UnusedReserve float64 `json:"unused_reserve,omitempty"`
   // Usage
   Usage float64 `json:"usage,omitempty"`
   // UsageValID
   UsageValID bool `json:"usage_valid,omitempty"`
   // VolCompressionRatio
   VolCompressionRatio float32 `json:"vol_compression_ratio,omitempty"`
   // VolCount
   VolCount float64 `json:"vol_count,omitempty"`
}
