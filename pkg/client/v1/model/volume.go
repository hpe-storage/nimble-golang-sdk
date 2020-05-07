/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model

var (
	VolumeFields = &Volume{
		ID:           "id",
		Name:         "name",
		FullName:     "full_name",
		SerialNumber: "serial_number",
		// TODO: generate this with remaining fields
	}
)

// Volume :
type Volume struct {
	// ID
	ID string `json:"id,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// FullName
	FullName string `json:"full_name,omitempty"`
	// SearchName
	SearchName string `json:"search_name,omitempty"`
	// Size
	Size float64 `json:"size,omitempty"`
	// Description
	Description string `json:"description,omitempty"`
	// PerfpolicyName
	PerfpolicyName string `json:"perfpolicy_name,omitempty"`
	// PerfpolicyID
	PerfpolicyID string `json:"perfpolicy_id,omitempty"`
	// Reserve
	Reserve float64 `json:"reserve,omitempty"`
	// WarnLevel
	WarnLevel float64 `json:"warn_level,omitempty"`
	// Limit
	Limit float64 `json:"limit,omitempty"`
	// SnapReserve
	SnapReserve float64 `json:"snap_reserve,omitempty"`
	// SnapWarnLevel
	SnapWarnLevel float64 `json:"snap_warn_level,omitempty"`
	// SnapLimit
	SnapLimit float64 `json:"snap_limit,omitempty"`
	// SnapLimitPercent
	SnapLimitPercent int32 `json:"snap_limit_percent,omitempty"`
	// NumSnaps
	NumSnaps float64 `json:"num_snaps,omitempty"`
	// ProjectedNumSnaps
	ProjectedNumSnaps float64 `json:"projected_num_snaps,omitempty"`
	// Online
	Online *bool `json:"online,omitempty"`
	// OwnedByGroup
	OwnedByGroup string `json:"owned_by_group,omitempty"`
	// OwnedByGroupID
	OwnedByGroupID string `json:"owned_by_group_id,omitempty"`
	// MultiInitiator
	MultiInitiator bool `json:"multi_initiator,omitempty"`
	// PoolName
	PoolName string `json:"pool_name,omitempty"`
	// PoolID
	PoolID string `json:"pool_id,omitempty"`
	// ReadOnly
	ReadOnly bool `json:"read_only,omitempty"`
	// SerialNumber
	SerialNumber string `json:"serial_number,omitempty"`
	// SecondarySerialNumber
	SecondarySerialNumber string `json:"secondary_serial_number,omitempty"`
	// TargetName
	TargetName string `json:"target_name,omitempty"`
	// BlockSize
	BlockSize float64 `json:"block_size,omitempty"`
	// Clone
	Clone *bool `json:"clone,omitempty"`
	// ParentVolName
	ParentVolName string `json:"parent_vol_name,omitempty"`
	// ParentVolID
	ParentVolID string `json:"parent_vol_id,omitempty"`
	// BaseSnapName
	BaseSnapName string `json:"base_snap_name,omitempty"`
	// BaseSnapID
	BaseSnapID string `json:"base_snap_id,omitempty"`
	// VolcollName
	VolcollName string `json:"volcoll_name,omitempty"`
	// VolcollID
	VolcollID string `json:"volcoll_id,omitempty"`
	// Force
	Force *bool `json:"force,omitempty"`
	// CreationTime
	CreationTime float64 `json:"creation_time,omitempty"`
	// LastModified
	LastModified float64 `json:"last_modified,omitempty"`
	// DestPoolName
	DestPoolName string `json:"dest_pool_name,omitempty"`
	// DestPoolID
	DestPoolID string `json:"dest_pool_id,omitempty"`
	// MoveStartTime
	MoveStartTime float64 `json:"move_start_time,omitempty"`
	// MoveAborting
	MoveAborting *bool `json:"move_aborting,omitempty"`
	// MoveBytesMigrated
	MoveBytesMigrated float64 `json:"move_bytes_migrated,omitempty"`
	// MoveBytesRemaining
	MoveBytesRemaining float64 `json:"move_bytes_remaining,omitempty"`
	// MoveEstComplTime
	MoveEstComplTime float64 `json:"move_est_compl_time,omitempty"`
	// UsageValID
	UsageValID *bool `json:"usage_valid,omitempty"`
	// TotalUsageBytes
	TotalUsageBytes float64 `json:"total_usage_bytes,omitempty"`
	// VolUsageCompressedBytes
	VolUsageCompressedBytes float64 `json:"vol_usage_compressed_bytes,omitempty"`
	// VolUsageUncompressedBytes
	VolUsageUncompressedBytes float64 `json:"vol_usage_uncompressed_bytes,omitempty"`
	// VolUsageMappedBytes
	VolUsageMappedBytes float64 `json:"vol_usage_mapped_bytes,omitempty"`
	// SnapUsageCompressedBytes
	SnapUsageCompressedBytes float64 `json:"snap_usage_compressed_bytes,omitempty"`
	// SnapUsageUncompressedBytes
	SnapUsageUncompressedBytes float64 `json:"snap_usage_uncompressed_bytes,omitempty"`
	// SnapUsagePopulatedBytes
	SnapUsagePopulatedBytes float64 `json:"snap_usage_populated_bytes,omitempty"`
	// CachePinned
	CachePinned *bool `json:"cache_pinned,omitempty"`
	// PinnedCacheSize
	PinnedCacheSize float64 `json:"pinned_cache_size,omitempty"`
	// CacheNeededForPin
	CacheNeededForPin float64 `json:"cache_needed_for_pin,omitempty"`
	// UpstreamCachePinned
	UpstreamCachePinned bool `json:"upstream_cache_pinned,omitempty"`
	// ThinlyProvisioned
	ThinlyProvisioned bool `json:"thinly_provisioned,omitempty"`
	// NumConnections
	NumConnections float64 `json:"num_connections,omitempty"`
	// NumIscsiConnections
	NumIscsiConnections float64 `json:"num_iscsi_connections,omitempty"`
	// NumFcConnections
	NumFcConnections float64 `json:"num_fc_connections,omitempty"`
	// InheritAcl
	InheritAcl bool `json:"inherit_acl,omitempty"`
	// AppUuID
	AppUuID string `json:"app_uuid,omitempty"`
	// FolderID
	FolderID string `json:"folder_id,omitempty"`
	// FolderName
	FolderName string `json:"folder_name,omitempty"`
	// CachingEnabled
	CachingEnabled bool `json:"caching_enabled,omitempty"`
	// PreviouslyDeduped
	PreviouslyDeduped bool `json:"previously_deduped,omitempty"`
	// DedupeEnabled
	DedupeEnabled bool `json:"dedupe_enabled,omitempty"`
	// VpdT10
	VpdT10 string `json:"vpd_t10,omitempty"`
	// VpdIeee0
	VpdIeee0 string `json:"vpd_ieee0,omitempty"`
	// VpdIeee1
	VpdIeee1 string `json:"vpd_ieee1,omitempty"`
	// AppCategory
	AppCategory string `json:"app_category,omitempty"`
	// LimitIops
	LimitIops int32 `json:"limit_iops,omitempty"`
	// LimitMbps
	LimitMbps int32 `json:"limit_mbps,omitempty"`
	// NeedsContentRepl
	NeedsContentRepl bool `json:"needs_content_repl,omitempty"`
	// ContentReplErrorsFound
	ContentReplErrorsFound bool `json:"content_repl_errors_found,omitempty"`
	// LastContentSnapBrCgUID
	LastContentSnapBrCgUID float64 `json:"last_content_snap_br_cg_uid,omitempty"`
	// LastContentSnapBrGID
	LastContentSnapBrGID float64 `json:"last_content_snap_br_gid,omitempty"`
	// LastContentSnapID
	LastContentSnapID float64 `json:"last_content_snap_id,omitempty"`
	// CksumLastVerified
	CksumLastVerified float64 `json:"cksum_last_verified,omitempty"`
	// PreFilter
	PreFilter string `json:"pre_filter,omitempty"`
	// SrepLastSync
	SrepLastSync float64 `json:"srep_last_sync,omitempty"`
	// SrepResyncPercent
	SrepResyncPercent float64 `json:"srep_resync_percent,omitempty"`
}
