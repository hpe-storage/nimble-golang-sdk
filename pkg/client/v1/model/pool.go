/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Pool - Manage pools. Pools are an aggregation of arrays.
// Export PoolFields for advance operations like search filter etc.
var PoolFields *Pool

func init(){
	IDfield:= "id"
	Namefield:= "name"
	FullNamefield:= "full_name"
	SearchNamefield:= "search_name"
	Descriptionfield:= "description"
		
	PoolFields= &Pool{
		ID: &IDfield,
		Name: &Namefield,
		FullName: &FullNamefield,
		SearchName: &SearchNamefield,
		Description: &Descriptionfield,
		
	}
}

type Pool struct {
   
    // Identifier for the pool.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of pool.
    
 	Name *string `json:"name,omitempty"`
   
    // Fully qualified name of pool.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Name of pool used for object search.
    
 	SearchName *string `json:"search_name,omitempty"`
   
    // Text description of pool.
    
 	Description *string `json:"description,omitempty"`
   
    // Time when this pool was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this pool was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Total storage space of the pool in bytes.
    
   	Capacity *int64 `json:"capacity,omitempty"`
   
    // Used space of the pool in bytes.
    
   	Usage *int64 `json:"usage,omitempty"`
   
    // Overall space usage savings in the pool.
    
   	Savings *int64 `json:"savings,omitempty"`
   
    // Space usage savings in the pool that does not include thin-provisioning savings.
    
   	SavingsDataReduction *int64 `json:"savings_data_reduction,omitempty"`
   
    // Space usage savings in the pool due to compression.
    
   	SavingsCompression *int64 `json:"savings_compression,omitempty"`
   
    // Space usage savings in the pool due to deduplication.
    
   	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
   
    // Space usage savings in the pool due to cloning of volumes.
    
   	SavingsClone *int64 `json:"savings_clone,omitempty"`
   
    // Space usage savings in the pool due to thin provisioning of volumes.
    
   	SavingsVolThinProvisioning *int64 `json:"savings_vol_thin_provisioning,omitempty"`
   
    // Reserved space of the pool in bytes. Sum of volume reserve in the pool.
    
   	Reserve *int64 `json:"reserve,omitempty"`
   
    // Unused reserve space of the pool in bytes.
    
   	UnusedReserve *int64 `json:"unused_reserve,omitempty"`
   
    // Free space of the pool in bytes.
    
   	FreeSpace *int64 `json:"free_space,omitempty"`
   
    // Total usable cache capacity of the pool in bytes.
    
   	CacheCapacity *int64 `json:"cache_capacity,omitempty"`
   
    // Total pinnable cache capacity of the pool in bytes.
    
   	PinnableCacheCapacity *int64 `json:"pinnable_cache_capacity,omitempty"`
   
    // Total pinned cache capacity of the pool in bytes.
    
   	PinnedCacheCapacity *int64 `json:"pinned_cache_capacity,omitempty"`
   
    // The dedupe capacity of a hybrid pool. Does not apply to all-flash pools.
    
   	DedupeCapacityBytes *int64 `json:"dedupe_capacity_bytes,omitempty"`
   
    // The dedupe usage of a hybrid pool. Does not apply to all-flash pools.
    
   	DedupeUsageBytes *int64 `json:"dedupe_usage_bytes,omitempty"`
   
    // Overall space usage savings in the pool expressed as ratio.
    
  	SavingsRatio *float32 `json:"savings_ratio,omitempty"`
   
    // Space usage savings in the pool expressed as ratio that does not include thin-provisioning savings.
    
  	DataReductionRatio *float32 `json:"data_reduction_ratio,omitempty"`
   
    // Compression savings for the pool expressed as ratio.
    
  	CompressionRatio *float32 `json:"compression_ratio,omitempty"`
   
    // Dedupe savings for the pool expressed as ratio.
    
  	DedupeRatio *float32 `json:"dedupe_ratio,omitempty"`
   
    // Clone savings for the pool expressed as ratio.
    
  	CloneRatio *float32 `json:"clone_ratio,omitempty"`
   
    // Thin provisioning savings for volumes in the pool expressed as ratio.
    
  	VolThinProvisioningRatio *float32 `json:"vol_thin_provisioning_ratio,omitempty"`
   
    // Snapshot collection count.
    
   	SnapcollCount *int64 `json:"snapcoll_count,omitempty"`
   
    // Snapshot count.
    
   	SnapCount *int64 `json:"snap_count,omitempty"`
   
    // Number of arrays in the pool.
    
   	ArrayCount *int64 `json:"array_count,omitempty"`
   
    // Number of volumes assigned to the pool.
    
   	VolCount *int64 `json:"vol_count,omitempty"`
   
    // List of arrays in the pool with detailed information. When create/update array list, only arrays' ID is required.
    
   	ArrayList []*NsArrayDetail `json:"array_list,omitempty"`
   
    // List of arrays being unassigned from the pool.
    
   	UnassignedArrayList []*NsArraySummary `json:"unassigned_array_list,omitempty"`
   
    // The list of volumes in the pool.
    
   	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
   
    // The list of pinned volumes in the pool.
    
   	PinnedVolList []*NsPinnedVolumeInfo `json:"pinned_vol_list,omitempty"`
   
    // The list of fully qualified names of folders in the pool.
    
   	FolderList []*NsFolderSummary `json:"folder_list,omitempty"`
   
    // Forcibly delete the specified pool even if it contains deleted volumes whose space is being reclaimed. Forcibly remove an array from array_list via an update operation even if the array is not reachable. There should no volumes currently in the pool for the forced update operation to succeed.
    
 	Force *bool `json:"force,omitempty"`
   
    // Indicates whether the usage of pool is valid.
    
 	UsageValID *bool `json:"usage_valid,omitempty"`
   
    // Uncompressed usage of volumes in the pool.
    
   	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   
    // Uncompressed usage of snapshots in the pool.
    
   	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   
    // Indicate whether the pool is an all_flash pool.
    
 	AllFlash *bool `json:"all_flash,omitempty"`
   
    // Indicates whether the pool is capable of hosting deduped volumes.
    
 	DedupeCapable *bool `json:"dedupe_capable,omitempty"`
   
    // Indicates whether the pool can enable dedupe by default.
    
 	DedupeAllVolumesCapable *bool `json:"dedupe_all_volumes_capable,omitempty"`
   
    // Indicates if dedupe is enabled by default for new volumes on this pool.
    
 	DedupeAllVolumes *bool `json:"dedupe_all_volumes,omitempty"`
   
    // Indicates if this is the default pool.
    
 	IsDefault *bool `json:"is_default,omitempty"`
}
