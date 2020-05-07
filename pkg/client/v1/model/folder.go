/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Folder


// Folder :
type Folder struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Fqn
   Fqn string `json:"fqn,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // LimitBytesSpecified
   LimitBytesSpecified bool `json:"limit_bytes_specified,omitempty"`
   // LimitBytes
   LimitBytes float64 `json:"limit_bytes,omitempty"`
   // LimitSizeBytes
   LimitSizeBytes  int32 `json:"limit_size_bytes,omitempty"`
   // ProvisionedLimitSizeBytes
   ProvisionedLimitSizeBytes  int32 `json:"provisioned_limit_size_bytes,omitempty"`
   // OverdraftLimitPct
   OverdraftLimitPct float64 `json:"overdraft_limit_pct,omitempty"`
   // CapacityBytes
   CapacityBytes float64 `json:"capacity_bytes,omitempty"`
   // FreeSpaceBytes
   FreeSpaceBytes float64 `json:"free_space_bytes,omitempty"`
   // ProvisionedBytes
   ProvisionedBytes float64 `json:"provisioned_bytes,omitempty"`
   // UsageBytes
   UsageBytes float64 `json:"usage_bytes,omitempty"`
   // VolumeMappedBytes
   VolumeMappedBytes float64 `json:"volume_mapped_bytes,omitempty"`
   // UsageValID
   UsageValID bool `json:"usage_valid,omitempty"`
   // InheritedVolPerfpolID
   InheritedVolPerfpolID string `json:"inherited_vol_perfpol_id,omitempty"`
   // InheritedVolPerfpolName
   InheritedVolPerfpolName string `json:"inherited_vol_perfpol_name,omitempty"`
   // UnusedReserveBytes
   UnusedReserveBytes float64 `json:"unused_reserve_bytes,omitempty"`
   // UnusedSnapReserveBytes
   UnusedSnapReserveBytes float64 `json:"unused_snap_reserve_bytes,omitempty"`
   // CompressedVolUsageBytes
   CompressedVolUsageBytes float64 `json:"compressed_vol_usage_bytes,omitempty"`
   // CompressedSnapUsageBytes
   CompressedSnapUsageBytes float64 `json:"compressed_snap_usage_bytes,omitempty"`
   // UncompressedVolUsageBytes
   UncompressedVolUsageBytes float64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   // UncompressedSnapUsageBytes
   UncompressedSnapUsageBytes float64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   // VolCompressionRatio
   VolCompressionRatio float32 `json:"vol_compression_ratio,omitempty"`
   // SnapCompressionRatio
   SnapCompressionRatio float32 `json:"snap_compression_ratio,omitempty"`
   // CompressionRatio
   CompressionRatio float32 `json:"compression_ratio,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // NumSnaps
   NumSnaps float64 `json:"num_snaps,omitempty"`
   // NumSnapcolls
   NumSnapcolls float64 `json:"num_snapcolls,omitempty"`
   // AppUuID
   AppUuID string `json:"app_uuid,omitempty"`
   // AppserverID
   AppserverID string `json:"appserver_id,omitempty"`
   // AppserverName
   AppserverName string `json:"appserver_name,omitempty"`
   // FolsetID
   FolsetID string `json:"folset_id,omitempty"`
   // FolsetName
   FolsetName string `json:"folset_name,omitempty"`
   // LimitIops
   LimitIops  int32 `json:"limit_iops,omitempty"`
   // LimitMbps
   LimitMbps  int32 `json:"limit_mbps,omitempty"`
   // TenantID
   TenantID string `json:"tenant_id,omitempty"`
}
