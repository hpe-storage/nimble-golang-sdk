/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsPoolMergeReturn - Return values of pool merge.
// Export NsPoolMergeReturnFields for advance operations like search filter etc.
var NsPoolMergeReturnFields *NsPoolMergeReturn

func init(){
	Descriptionfield:= "description"
	FullNamefield:= "full_name"
	IDfield:= "id"
	Namefield:= "name"
	SearchNamefield:= "search_name"
		
	NsPoolMergeReturnFields= &NsPoolMergeReturn{
		Description: &Descriptionfield,
		FullName: &FullNamefield,
		ID: &IDfield,
		Name: &Namefield,
		SearchName: &SearchNamefield,
		
	}
}

type NsPoolMergeReturn struct {
   
    // Snapshot collection count.
    
   	SnapcollCount *int64 `json:"snapcoll_count,omitempty"`
   
    // List of arrays in the pool with detailed information. When create/update array list, only arrays' ID is required.
    
   	ArrayList []*NsArrayDetail `json:"array_list,omitempty"`
   
    // Number of arrays in the pool.
    
   	ArrayCount *int64 `json:"array_count,omitempty"`
   
    // List of arrays being unassigned from the pool.
    
   	UnassignedArrayList []*NsArraySummary `json:"unassigned_array_list,omitempty"`
   
    // Indicate whether the pool is an all_flash pool.
    
 	AllFlash *bool `json:"all_flash,omitempty"`
   
    // Total usable cache capacity of the pool in bytes.
    
   	CacheCapacity *int64 `json:"cache_capacity,omitempty"`
   
    // Total storage space of the pool in bytes.
    
   	Capacity *int64 `json:"capacity,omitempty"`
   
    // Compressed usage of snapshots in the pool.
    
   	CompressedSnapUsageBytes *int64 `json:"compressed_snap_usage_bytes,omitempty"`
   
    // Compressed usage of volumes in the pool.
    
   	CompressedVolUsageBytes *int64 `json:"compressed_vol_usage_bytes,omitempty"`
   
    // Compression savings for the pool expressed as ratio.
    
  	CompressionRatio *float32 `json:"compression_ratio,omitempty"`
   
    // Time when this pool was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Dedupe savings for the pool expressed as ratio.
    
  	DedupeRatio *float32 `json:"dedupe_ratio,omitempty"`
   
    // Text description of pool. Default: ''.
    
 	Description *string `json:"description,omitempty"`
   
    // Free space of the pool in bytes.
    
   	FreeSpace *int64 `json:"free_space,omitempty"`
   
    // Fully qualified name of pool.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Identifier for the pool.
    
 	ID *string `json:"id,omitempty"`
   
    // Indicate whether the pool is a dedupe enabled pool.
    
 	DedupeEnabled *bool `json:"dedupe_enabled,omitempty"`
   
    // Time when this pool was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Name of pool.
    
 	Name *string `json:"name,omitempty"`
   
    // Total pinnable cache capacity of the pool in bytes.
    
   	PinnableCacheCapacity *int64 `json:"pinnable_cache_capacity,omitempty"`
   
    // Total pinned cache capacity of the pool in bytes.
    
   	PinnedCacheCapacity *int64 `json:"pinned_cache_capacity,omitempty"`
   
    // The list of pinned volumes in the pool.
    
   	PinnedVolList []*NsVolumeSummary `json:"pinned_vol_list,omitempty"`
   
    // Space usage savings in the pool due to compression.
    
   	SavingsCompression *int64 `json:"savings_compression,omitempty"`
   
    // Space usage savings in the pool due to deduplication.
    
   	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
   
    // Name of pool used for object search.
    
 	SearchName *string `json:"search_name,omitempty"`
   
    // Compression ratio of snapshots in the pool.
    
  	SnapCompressionRatio *float32 `json:"snap_compression_ratio,omitempty"`
   
    // Snapshot count.
    
   	SnapCount *int64 `json:"snap_count,omitempty"`
   
    // Uncompressed usage of snapshots in the pool.
    
   	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   
    // Uncompressed usage of volumes in the pool.
    
   	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   
    // Unused reserve space of the pool in bytes.
    
   	UnusedReserve *int64 `json:"unused_reserve,omitempty"`
   
    // Used space of the pool in bytes.
    
   	Usage *int64 `json:"usage,omitempty"`
   
    // Indicate whether the usage of pool is valid.
    
 	UsageValID *bool `json:"usage_valid,omitempty"`
   
    // Compression ratio of volumes in the pool.
    
  	VolCompressionRatio *float32 `json:"vol_compression_ratio,omitempty"`
   
    // Number of volumes assigned to the pool.
    
   	VolCount *int64 `json:"vol_count,omitempty"`
   
    // The list of volumes in the pool.
    
   	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
}
