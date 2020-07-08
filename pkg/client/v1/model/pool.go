// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// Pool - Manage pools. Pools are an aggregation of arrays.
// Export PoolFields for advance operations like search filter etc.
var PoolFields *Pool

func init() {

	PoolFields = &Pool{
		ID:          "id",
		Name:        "name",
		FullName:    "full_name",
		SearchName:  "search_name",
		Description: "description",
	}
}

type Pool struct {
	// ID - Identifier for the pool.
	ID string `json:"id,omitempty"`
	// Name - Name of pool.
	Name string `json:"name,omitempty"`
	// FullName - Fully qualified name of pool.
	FullName string `json:"full_name,omitempty"`
	// SearchName - Name of pool used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description - Text description of pool.
	Description string `json:"description,omitempty"`
	// CreationTime - Time when this pool was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this pool was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// Capacity - Total storage space of the pool in bytes.
	Capacity int64 `json:"capacity,omitempty"`
	// Usage - Used space of the pool in bytes.
	Usage int64 `json:"usage,omitempty"`
	// Savings - Overall space usage savings in the pool.
	Savings int64 `json:"savings,omitempty"`
	// SavingsDataReduction - Space usage savings in the pool that does not include thin-provisioning savings.
	SavingsDataReduction int64 `json:"savings_data_reduction,omitempty"`
	// SavingsCompression - Space usage savings in the pool due to compression.
	SavingsCompression int64 `json:"savings_compression,omitempty"`
	// SavingsDedupe - Space usage savings in the pool due to deduplication.
	SavingsDedupe int64 `json:"savings_dedupe,omitempty"`
	// SavingsClone - Space usage savings in the pool due to cloning of volumes.
	SavingsClone int64 `json:"savings_clone,omitempty"`
	// SavingsVolThinProvisioning - Space usage savings in the pool due to thin provisioning of volumes.
	SavingsVolThinProvisioning int64 `json:"savings_vol_thin_provisioning,omitempty"`
	// Reserve - Reserved space of the pool in bytes. Sum of volume reserve in the pool.
	Reserve int64 `json:"reserve,omitempty"`
	// UnusedReserve - Unused reserve space of the pool in bytes.
	UnusedReserve int64 `json:"unused_reserve,omitempty"`
	// FreeSpace - Free space of the pool in bytes.
	FreeSpace int64 `json:"free_space,omitempty"`
	// CacheCapacity - Total usable cache capacity of the pool in bytes.
	CacheCapacity int64 `json:"cache_capacity,omitempty"`
	// PinnableCacheCapacity - Total pinnable cache capacity of the pool in bytes.
	PinnableCacheCapacity int64 `json:"pinnable_cache_capacity,omitempty"`
	// PinnedCacheCapacity - Total pinned cache capacity of the pool in bytes.
	PinnedCacheCapacity int64 `json:"pinned_cache_capacity,omitempty"`
	// DedupeCapacityBytes - The dedupe capacity of a hybrid pool. Does not apply to all-flash pools.
	DedupeCapacityBytes int64 `json:"dedupe_capacity_bytes,omitempty"`
	// DedupeUsageBytes - The dedupe usage of a hybrid pool. Does not apply to all-flash pools.
	DedupeUsageBytes int64 `json:"dedupe_usage_bytes,omitempty"`
	// SavingsRatio - Overall space usage savings in the pool expressed as ratio.
	SavingsRatio float32 `json:"savings_ratio,omitempty"`
	// DataReductionRatio - Space usage savings in the pool expressed as ratio that does not include thin-provisioning savings.
	DataReductionRatio float32 `json:"data_reduction_ratio,omitempty"`
	// CompressionRatio - Compression savings for the pool expressed as ratio.
	CompressionRatio float32 `json:"compression_ratio,omitempty"`
	// DedupeRatio - Dedupe savings for the pool expressed as ratio.
	DedupeRatio float32 `json:"dedupe_ratio,omitempty"`
	// CloneRatio - Clone savings for the pool expressed as ratio.
	CloneRatio float32 `json:"clone_ratio,omitempty"`
	// VolThinProvisioningRatio - Thin provisioning savings for volumes in the pool expressed as ratio.
	VolThinProvisioningRatio float32 `json:"vol_thin_provisioning_ratio,omitempty"`
	// SnapcollCount - Snapshot collection count.
	SnapcollCount int64 `json:"snapcoll_count,omitempty"`
	// SnapCount - Snapshot count.
	SnapCount int64 `json:"snap_count,omitempty"`
	// ArrayCount - Number of arrays in the pool.
	ArrayCount int64 `json:"array_count,omitempty"`
	// VolCount - Number of volumes assigned to the pool.
	VolCount int64 `json:"vol_count,omitempty"`
	// ArrayList - List of arrays in the pool with detailed information. When create/update array list, only arrays' ID is required.
	ArrayList []*NsArrayDetail `json:"array_list,omitempty"`
	// UnassignedArrayList - List of arrays being unassigned from the pool.
	UnassignedArrayList []*NsArraySummary `json:"unassigned_array_list,omitempty"`
	// VolList - The list of volumes in the pool.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
	// PinnedVolList - The list of pinned volumes in the pool.
	PinnedVolList []*NsPinnedVolumeInfo `json:"pinned_vol_list,omitempty"`
	// FolderList - The list of fully qualified names of folders in the pool.
	FolderList []*NsFolderSummary `json:"folder_list,omitempty"`
	// Force - Forcibly delete the specified pool even if it contains deleted volumes whose space is being reclaimed. Forcibly remove an array from array_list via an update operation even if the array is not reachable. There should no volumes currently in the pool for the forced update operation to succeed.
	Force *bool `json:"force,omitempty"`
	// UsageValid - Indicates whether the usage of pool is valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// UncompressedVolUsageBytes - Uncompressed usage of volumes in the pool.
	UncompressedVolUsageBytes int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
	// UncompressedSnapUsageBytes - Uncompressed usage of snapshots in the pool.
	UncompressedSnapUsageBytes int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
	// AllFlash - Indicate whether the pool is an all_flash pool.
	AllFlash *bool `json:"all_flash,omitempty"`
	// DedupeCapable - Indicates whether the pool is capable of hosting deduped volumes.
	DedupeCapable *bool `json:"dedupe_capable,omitempty"`
	// DedupeAllVolumesCapable - Indicates whether the pool can enable dedupe by default.
	DedupeAllVolumesCapable *bool `json:"dedupe_all_volumes_capable,omitempty"`
	// DedupeAllVolumes - Indicates if dedupe is enabled by default for new volumes on this pool.
	DedupeAllVolumes *bool `json:"dedupe_all_volumes,omitempty"`
	// IsDefault - Indicates if this is the default pool.
	IsDefault *bool `json:"is_default,omitempty"`
}

// sdk internal struct
type pool struct {
	// ID - Identifier for the pool.
	ID *string `json:"id,omitempty"`
	// Name - Name of pool.
	Name *string `json:"name,omitempty"`
	// FullName - Fully qualified name of pool.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - Name of pool used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Text description of pool.
	Description *string `json:"description,omitempty"`
	// CreationTime - Time when this pool was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this pool was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// Capacity - Total storage space of the pool in bytes.
	Capacity *int64 `json:"capacity,omitempty"`
	// Usage - Used space of the pool in bytes.
	Usage *int64 `json:"usage,omitempty"`
	// Savings - Overall space usage savings in the pool.
	Savings *int64 `json:"savings,omitempty"`
	// SavingsDataReduction - Space usage savings in the pool that does not include thin-provisioning savings.
	SavingsDataReduction *int64 `json:"savings_data_reduction,omitempty"`
	// SavingsCompression - Space usage savings in the pool due to compression.
	SavingsCompression *int64 `json:"savings_compression,omitempty"`
	// SavingsDedupe - Space usage savings in the pool due to deduplication.
	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
	// SavingsClone - Space usage savings in the pool due to cloning of volumes.
	SavingsClone *int64 `json:"savings_clone,omitempty"`
	// SavingsVolThinProvisioning - Space usage savings in the pool due to thin provisioning of volumes.
	SavingsVolThinProvisioning *int64 `json:"savings_vol_thin_provisioning,omitempty"`
	// Reserve - Reserved space of the pool in bytes. Sum of volume reserve in the pool.
	Reserve *int64 `json:"reserve,omitempty"`
	// UnusedReserve - Unused reserve space of the pool in bytes.
	UnusedReserve *int64 `json:"unused_reserve,omitempty"`
	// FreeSpace - Free space of the pool in bytes.
	FreeSpace *int64 `json:"free_space,omitempty"`
	// CacheCapacity - Total usable cache capacity of the pool in bytes.
	CacheCapacity *int64 `json:"cache_capacity,omitempty"`
	// PinnableCacheCapacity - Total pinnable cache capacity of the pool in bytes.
	PinnableCacheCapacity *int64 `json:"pinnable_cache_capacity,omitempty"`
	// PinnedCacheCapacity - Total pinned cache capacity of the pool in bytes.
	PinnedCacheCapacity *int64 `json:"pinned_cache_capacity,omitempty"`
	// DedupeCapacityBytes - The dedupe capacity of a hybrid pool. Does not apply to all-flash pools.
	DedupeCapacityBytes *int64 `json:"dedupe_capacity_bytes,omitempty"`
	// DedupeUsageBytes - The dedupe usage of a hybrid pool. Does not apply to all-flash pools.
	DedupeUsageBytes *int64 `json:"dedupe_usage_bytes,omitempty"`
	// SavingsRatio - Overall space usage savings in the pool expressed as ratio.
	SavingsRatio *float32 `json:"savings_ratio,omitempty"`
	// DataReductionRatio - Space usage savings in the pool expressed as ratio that does not include thin-provisioning savings.
	DataReductionRatio *float32 `json:"data_reduction_ratio,omitempty"`
	// CompressionRatio - Compression savings for the pool expressed as ratio.
	CompressionRatio *float32 `json:"compression_ratio,omitempty"`
	// DedupeRatio - Dedupe savings for the pool expressed as ratio.
	DedupeRatio *float32 `json:"dedupe_ratio,omitempty"`
	// CloneRatio - Clone savings for the pool expressed as ratio.
	CloneRatio *float32 `json:"clone_ratio,omitempty"`
	// VolThinProvisioningRatio - Thin provisioning savings for volumes in the pool expressed as ratio.
	VolThinProvisioningRatio *float32 `json:"vol_thin_provisioning_ratio,omitempty"`
	// SnapcollCount - Snapshot collection count.
	SnapcollCount *int64 `json:"snapcoll_count,omitempty"`
	// SnapCount - Snapshot count.
	SnapCount *int64 `json:"snap_count,omitempty"`
	// ArrayCount - Number of arrays in the pool.
	ArrayCount *int64 `json:"array_count,omitempty"`
	// VolCount - Number of volumes assigned to the pool.
	VolCount *int64 `json:"vol_count,omitempty"`
	// ArrayList - List of arrays in the pool with detailed information. When create/update array list, only arrays' ID is required.
	ArrayList []*NsArrayDetail `json:"array_list,omitempty"`
	// UnassignedArrayList - List of arrays being unassigned from the pool.
	UnassignedArrayList []*NsArraySummary `json:"unassigned_array_list,omitempty"`
	// VolList - The list of volumes in the pool.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
	// PinnedVolList - The list of pinned volumes in the pool.
	PinnedVolList []*NsPinnedVolumeInfo `json:"pinned_vol_list,omitempty"`
	// FolderList - The list of fully qualified names of folders in the pool.
	FolderList []*NsFolderSummary `json:"folder_list,omitempty"`
	// Force - Forcibly delete the specified pool even if it contains deleted volumes whose space is being reclaimed. Forcibly remove an array from array_list via an update operation even if the array is not reachable. There should no volumes currently in the pool for the forced update operation to succeed.
	Force *bool `json:"force,omitempty"`
	// UsageValid - Indicates whether the usage of pool is valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// UncompressedVolUsageBytes - Uncompressed usage of volumes in the pool.
	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
	// UncompressedSnapUsageBytes - Uncompressed usage of snapshots in the pool.
	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
	// AllFlash - Indicate whether the pool is an all_flash pool.
	AllFlash *bool `json:"all_flash,omitempty"`
	// DedupeCapable - Indicates whether the pool is capable of hosting deduped volumes.
	DedupeCapable *bool `json:"dedupe_capable,omitempty"`
	// DedupeAllVolumesCapable - Indicates whether the pool can enable dedupe by default.
	DedupeAllVolumesCapable *bool `json:"dedupe_all_volumes_capable,omitempty"`
	// DedupeAllVolumes - Indicates if dedupe is enabled by default for new volumes on this pool.
	DedupeAllVolumes *bool `json:"dedupe_all_volumes,omitempty"`
	// IsDefault - Indicates if this is the default pool.
	IsDefault *bool `json:"is_default,omitempty"`
}

// EncodePool - Transform Pool to pool type
func EncodePool(request interface{}) (*pool, error) {
	reqPool := request.(*Pool)
	byte, err := json.Marshal(reqPool)
	objPtr := &pool{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodePool Transform pool to Pool type
func DecodePool(request interface{}) (*Pool, error) {
	reqPool := request.(*pool)
	byte, err := json.Marshal(reqPool)
	obj := &Pool{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
