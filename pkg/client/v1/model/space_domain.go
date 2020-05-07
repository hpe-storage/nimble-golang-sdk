/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/SpaceDomain


// SpaceDomain :
type SpaceDomain struct {
   // ID
   ID string `json:"id,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
   // AppCategoryID
   AppCategoryID string `json:"app_category_id,omitempty"`
   // AppCategoryName
   AppCategoryName string `json:"app_category_name,omitempty"`
   // SampleRate
   SampleRate float64 `json:"sample_rate,omitempty"`
   // VolumeCount
   VolumeCount float64 `json:"volume_count,omitempty"`
   // DedupedVolumeCount
   DedupedVolumeCount float64 `json:"deduped_volume_count,omitempty"`
   // BlockSize
   BlockSize float64 `json:"block_size,omitempty"`
   // Deduped
   Deduped bool `json:"deduped,omitempty"`
   // Encrypted
   Encrypted bool `json:"encrypted,omitempty"`
   // Usage
   Usage float64 `json:"usage,omitempty"`
   // VolLogicalUsage
   VolLogicalUsage float64 `json:"vol_logical_usage,omitempty"`
   // SnapLogicalUsage
   SnapLogicalUsage float64 `json:"snap_logical_usage,omitempty"`
   // VolMappedUsage
   VolMappedUsage float64 `json:"vol_mapped_usage,omitempty"`
   // LogicalDedupeUsage
   LogicalDedupeUsage float64 `json:"logical_dedupe_usage,omitempty"`
   // PhysicalDedupeUsage
   PhysicalDedupeUsage float64 `json:"physical_dedupe_usage,omitempty"`
   // SavingsCompression
   SavingsCompression float64 `json:"savings_compression,omitempty"`
   // SavingsDedupe
   SavingsDedupe float64 `json:"savings_dedupe,omitempty"`
   // SavingsClone
   SavingsClone float64 `json:"savings_clone,omitempty"`
   // CompressedUsageBytes
   CompressedUsageBytes float64 `json:"compressed_usage_bytes,omitempty"`
   // UncompressedUsageBytes
   UncompressedUsageBytes float64 `json:"uncompressed_usage_bytes,omitempty"`
   // CompressionRatio
   CompressionRatio float32 `json:"compression_ratio,omitempty"`
   // DedupeRatio
   DedupeRatio float32 `json:"dedupe_ratio,omitempty"`
   // CloneRatio
   CloneRatio float32 `json:"clone_ratio,omitempty"`
}
