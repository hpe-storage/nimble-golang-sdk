/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Pool


// Pool :
type Pool struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // Capacity
   Capacity float64 `json:"capacity,omitempty"`
   // Usage
   Usage float64 `json:"usage,omitempty"`
   // Savings
   Savings float64 `json:"savings,omitempty"`
   // SavingsDataReduction
   SavingsDataReduction float64 `json:"savings_data_reduction,omitempty"`
   // SavingsCompression
   SavingsCompression float64 `json:"savings_compression,omitempty"`
   // SavingsDedupe
   SavingsDedupe float64 `json:"savings_dedupe,omitempty"`
   // SavingsClone
   SavingsClone float64 `json:"savings_clone,omitempty"`
   // SavingsVolThinProvisioning
   SavingsVolThinProvisioning float64 `json:"savings_vol_thin_provisioning,omitempty"`
   // Reserve
   Reserve float64 `json:"reserve,omitempty"`
   // UnusedReserve
   UnusedReserve float64 `json:"unused_reserve,omitempty"`
   // FreeSpace
   FreeSpace float64 `json:"free_space,omitempty"`
   // CacheCapacity
   CacheCapacity float64 `json:"cache_capacity,omitempty"`
   // PinnableCacheCapacity
   PinnableCacheCapacity float64 `json:"pinnable_cache_capacity,omitempty"`
   // PinnedCacheCapacity
   PinnedCacheCapacity float64 `json:"pinned_cache_capacity,omitempty"`
   // DedupeCapacityBytes
   DedupeCapacityBytes float64 `json:"dedupe_capacity_bytes,omitempty"`
   // DedupeUsageBytes
   DedupeUsageBytes float64 `json:"dedupe_usage_bytes,omitempty"`
   // SavingsRatio
   SavingsRatio float32 `json:"savings_ratio,omitempty"`
   // DataReductionRatio
   DataReductionRatio float32 `json:"data_reduction_ratio,omitempty"`
   // CompressionRatio
   CompressionRatio float32 `json:"compression_ratio,omitempty"`
   // DedupeRatio
   DedupeRatio float32 `json:"dedupe_ratio,omitempty"`
   // CloneRatio
   CloneRatio float32 `json:"clone_ratio,omitempty"`
   // VolThinProvisioningRatio
   VolThinProvisioningRatio float32 `json:"vol_thin_provisioning_ratio,omitempty"`
   // SnapcollCount
   SnapcollCount float64 `json:"snapcoll_count,omitempty"`
   // SnapCount
   SnapCount float64 `json:"snap_count,omitempty"`
   // ArrayCount
   ArrayCount float64 `json:"array_count,omitempty"`
   // VolCount
   VolCount float64 `json:"vol_count,omitempty"`
   // Force
   Force bool `json:"force,omitempty"`
   // UsageValID
   UsageValID bool `json:"usage_valid,omitempty"`
   // UncompressedVolUsageBytes
   UncompressedVolUsageBytes float64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   // UncompressedSnapUsageBytes
   UncompressedSnapUsageBytes float64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   // AllFlash
   AllFlash bool `json:"all_flash,omitempty"`
   // DedupeCapable
   DedupeCapable bool `json:"dedupe_capable,omitempty"`
   // DedupeAllVolumesCapable
   DedupeAllVolumesCapable bool `json:"dedupe_all_volumes_capable,omitempty"`
   // DedupeAllVolumes
   DedupeAllVolumes bool `json:"dedupe_all_volumes,omitempty"`
   // IsDefault
   IsDefault bool `json:"is_default,omitempty"`
}
