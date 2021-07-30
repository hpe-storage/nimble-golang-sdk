// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsPoolMergeReturnFields provides field names to use in filter parameters, for example.
var NsPoolMergeReturnFields *NsPoolMergeReturnFieldHandles

func init() {
	NsPoolMergeReturnFields = &NsPoolMergeReturnFieldHandles{
		SnapcollCount:              "snapcoll_count",
		ArrayList:                  "array_list",
		ArrayCount:                 "array_count",
		UnassignedArrayList:        "unassigned_array_list",
		AllFlash:                   "all_flash",
		CacheCapacity:              "cache_capacity",
		Capacity:                   "capacity",
		CompressedSnapUsageBytes:   "compressed_snap_usage_bytes",
		CompressedVolUsageBytes:    "compressed_vol_usage_bytes",
		CompressionRatio:           "compression_ratio",
		CreationTime:               "creation_time",
		DedupeRatio:                "dedupe_ratio",
		Description:                "description",
		FreeSpace:                  "free_space",
		FullName:                   "full_name",
		ID:                         "id",
		DedupeEnabled:              "dedupe_enabled",
		LastModified:               "last_modified",
		Name:                       "name",
		PinnableCacheCapacity:      "pinnable_cache_capacity",
		PinnedCacheCapacity:        "pinned_cache_capacity",
		PinnedVolList:              "pinned_vol_list",
		SavingsCompression:         "savings_compression",
		SavingsDedupe:              "savings_dedupe",
		SearchName:                 "search_name",
		SnapCompressionRatio:       "snap_compression_ratio",
		SnapCount:                  "snap_count",
		UncompressedSnapUsageBytes: "uncompressed_snap_usage_bytes",
		UncompressedVolUsageBytes:  "uncompressed_vol_usage_bytes",
		UnusedReserve:              "unused_reserve",
		Usage:                      "usage",
		UsageValid:                 "usage_valid",
		VolCompressionRatio:        "vol_compression_ratio",
		VolCount:                   "vol_count",
		VolList:                    "vol_list",
	}
}

// NsPoolMergeReturn - Return values of pool merge.
type NsPoolMergeReturn struct {
	// SnapcollCount - Snapshot collection count.
	SnapcollCount *int64 `json:"snapcoll_count,omitempty"`
	// ArrayList - List of arrays in the pool with detailed information. When create/update array list, only arrays' ID is required.
	ArrayList []*NsArrayDetail `json:"array_list,omitempty"`
	// ArrayCount - Number of arrays in the pool.
	ArrayCount *int64 `json:"array_count,omitempty"`
	// UnassignedArrayList - List of arrays being unassigned from the pool.
	UnassignedArrayList []*NsArraySummary `json:"unassigned_array_list,omitempty"`
	// AllFlash - Indicate whether the pool is an all_flash pool.
	AllFlash *bool `json:"all_flash,omitempty"`
	// CacheCapacity - Total usable cache capacity of the pool in bytes.
	CacheCapacity *int64 `json:"cache_capacity,omitempty"`
	// Capacity - Total storage space of the pool in bytes.
	Capacity *int64 `json:"capacity,omitempty"`
	// CompressedSnapUsageBytes - Compressed usage of snapshots in the pool.
	CompressedSnapUsageBytes *int64 `json:"compressed_snap_usage_bytes,omitempty"`
	// CompressedVolUsageBytes - Compressed usage of volumes in the pool.
	CompressedVolUsageBytes *int64 `json:"compressed_vol_usage_bytes,omitempty"`
	// CompressionRatio - Compression savings for the pool expressed as ratio.
	CompressionRatio *float64 `json:"compression_ratio,omitempty"`
	// CreationTime - Time when this pool was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// DedupeRatio - Dedupe savings for the pool expressed as ratio.
	DedupeRatio *float64 `json:"dedupe_ratio,omitempty"`
	// Description - Text description of pool. Default: ''.
	Description *string `json:"description,omitempty"`
	// FreeSpace - Free space of the pool in bytes.
	FreeSpace *int64 `json:"free_space,omitempty"`
	// FullName - Fully qualified name of pool.
	FullName *string `json:"full_name,omitempty"`
	// ID - Identifier for the pool.
	ID *string `json:"id,omitempty"`
	// DedupeEnabled - Indicate whether the pool is a dedupe enabled pool.
	DedupeEnabled *bool `json:"dedupe_enabled,omitempty"`
	// LastModified - Time when this pool was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// Name - Name of pool.
	Name *string `json:"name,omitempty"`
	// PinnableCacheCapacity - Total pinnable cache capacity of the pool in bytes.
	PinnableCacheCapacity *int64 `json:"pinnable_cache_capacity,omitempty"`
	// PinnedCacheCapacity - Total pinned cache capacity of the pool in bytes.
	PinnedCacheCapacity *int64 `json:"pinned_cache_capacity,omitempty"`
	// PinnedVolList - The list of pinned volumes in the pool.
	PinnedVolList []*NsVolumeSummary `json:"pinned_vol_list,omitempty"`
	// SavingsCompression - Space usage savings in the pool due to compression.
	SavingsCompression *int64 `json:"savings_compression,omitempty"`
	// SavingsDedupe - Space usage savings in the pool due to deduplication.
	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
	// SearchName - Name of pool used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// SnapCompressionRatio - Compression ratio of snapshots in the pool.
	SnapCompressionRatio *float64 `json:"snap_compression_ratio,omitempty"`
	// SnapCount - Snapshot count.
	SnapCount *int64 `json:"snap_count,omitempty"`
	// UncompressedSnapUsageBytes - Uncompressed usage of snapshots in the pool.
	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
	// UncompressedVolUsageBytes - Uncompressed usage of volumes in the pool.
	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
	// UnusedReserve - Unused reserve space of the pool in bytes.
	UnusedReserve *int64 `json:"unused_reserve,omitempty"`
	// Usage - Used space of the pool in bytes.
	Usage *int64 `json:"usage,omitempty"`
	// UsageValid - Indicate whether the usage of pool is valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// VolCompressionRatio - Compression ratio of volumes in the pool.
	VolCompressionRatio *float64 `json:"vol_compression_ratio,omitempty"`
	// VolCount - Number of volumes assigned to the pool.
	VolCount *int64 `json:"vol_count,omitempty"`
	// VolList - The list of volumes in the pool.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}

// NsPoolMergeReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsPoolMergeReturnFieldHandles struct {
	SnapcollCount              string
	ArrayList                  string
	ArrayCount                 string
	UnassignedArrayList        string
	AllFlash                   string
	CacheCapacity              string
	Capacity                   string
	CompressedSnapUsageBytes   string
	CompressedVolUsageBytes    string
	CompressionRatio           string
	CreationTime               string
	DedupeRatio                string
	Description                string
	FreeSpace                  string
	FullName                   string
	ID                         string
	DedupeEnabled              string
	LastModified               string
	Name                       string
	PinnableCacheCapacity      string
	PinnedCacheCapacity        string
	PinnedVolList              string
	SavingsCompression         string
	SavingsDedupe              string
	SearchName                 string
	SnapCompressionRatio       string
	SnapCount                  string
	UncompressedSnapUsageBytes string
	UncompressedVolUsageBytes  string
	UnusedReserve              string
	Usage                      string
	UsageValid                 string
	VolCompressionRatio        string
	VolCount                   string
	VolList                    string
}
