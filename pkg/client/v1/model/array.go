/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Array


// Array :
type Array struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Force
   Force bool `json:"force,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // Model
   Model string `json:"model,omitempty"`
   // Serial
   Serial string `json:"serial,omitempty"`
   // Version
   Version string `json:"version,omitempty"`
   // IsSfa
   IsSfa bool `json:"is_sfa,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // UsageValID
   UsageValID bool `json:"usage_valid,omitempty"`
   // UsableCapacityBytes
   UsableCapacityBytes float64 `json:"usable_capacity_bytes,omitempty"`
   // UsableCacheCapacityBytes
   UsableCacheCapacityBytes float64 `json:"usable_cache_capacity_bytes,omitempty"`
   // RawCapacityBytes
   RawCapacityBytes float64 `json:"raw_capacity_bytes,omitempty"`
   // VolUsageBytes
   VolUsageBytes float64 `json:"vol_usage_bytes,omitempty"`
   // VolUsageUncompressedBytes
   VolUsageUncompressedBytes float64 `json:"vol_usage_uncompressed_bytes,omitempty"`
   // VolCompression
   VolCompression float32 `json:"vol_compression,omitempty"`
   // VolSavedBytes
   VolSavedBytes float64 `json:"vol_saved_bytes,omitempty"`
   // SnapUsageBytes
   SnapUsageBytes float64 `json:"snap_usage_bytes,omitempty"`
   // SnapUsageUncompressedBytes
   SnapUsageUncompressedBytes float64 `json:"snap_usage_uncompressed_bytes,omitempty"`
   // SnapCompression
   SnapCompression float32 `json:"snap_compression,omitempty"`
   // SnapSpaceReduction
   SnapSpaceReduction float32 `json:"snap_space_reduction,omitempty"`
   // SnapSavedBytes
   SnapSavedBytes float64 `json:"snap_saved_bytes,omitempty"`
   // PendingDeleteBytes
   PendingDeleteBytes float64 `json:"pending_delete_bytes,omitempty"`
   // AvailableBytes
   AvailableBytes float64 `json:"available_bytes,omitempty"`
   // Usage
   Usage float64 `json:"usage,omitempty"`
   // AllFlash
   AllFlash bool `json:"all_flash,omitempty"`
   // DedupeCapacityBytes
   DedupeCapacityBytes float64 `json:"dedupe_capacity_bytes,omitempty"`
   // DedupeUsageBytes
   DedupeUsageBytes float64 `json:"dedupe_usage_bytes,omitempty"`
   // ExtendedModel
   ExtendedModel string `json:"extended_model,omitempty"`
   // IsSupportedHwConfig
   IsSupportedHwConfig bool `json:"is_supported_hw_config,omitempty"`
   // GigNicPortCount
   GigNicPortCount float64 `json:"gig_nic_port_count,omitempty"`
   // TenGigSfpNicPortCount
   TenGigSfpNicPortCount float64 `json:"ten_gig_sfp_nic_port_count,omitempty"`
   // TenGigTNicPortCount
   TenGigTNicPortCount float64 `json:"ten_gig_t_nic_port_count,omitempty"`
   // FcPortCount
   FcPortCount float64 `json:"fc_port_count,omitempty"`
   // CreatePool
   CreatePool bool `json:"create_pool,omitempty"`
   // PoolDescription
   PoolDescription string `json:"pool_description,omitempty"`
   // AllowLowerLimits
   AllowLowerLimits bool `json:"allow_lower_limits,omitempty"`
   // CtrlrASupportIp
   CtrlrASupportIp string `json:"ctrlr_a_support_ip,omitempty"`
   // CtrlrBSupportIp
   CtrlrBSupportIp string `json:"ctrlr_b_support_ip,omitempty"`
   // ModelSubType
   ModelSubType string `json:"model_sub_type,omitempty"`
   // SecondaryMgmtIp
   SecondaryMgmtIp string `json:"secondary_mgmt_ip,omitempty"`
}
