/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// SpaceDomain - A space domain is created for each application category and block size for each each pool.
// Export SpaceDomainFields for advance operations like search filter etc.
var SpaceDomainFields *SpaceDomain

func init(){
	IDfield:= "id"
	PoolIDfield:= "pool_id"
	PoolNamefield:= "pool_name"
	AppCategoryIDfield:= "app_category_id"
	AppCategoryNamefield:= "app_category_name"
		
	SpaceDomainFields= &SpaceDomain{
		ID: &IDfield,
		PoolID: &PoolIDfield,
		PoolName: &PoolNamefield,
		AppCategoryID: &AppCategoryIDfield,
		AppCategoryName: &AppCategoryNamefield,
		
	}
}

type SpaceDomain struct {
   
    // Identifier for the space domain.
    
 	ID *string `json:"id,omitempty"`
   
    // Identifier associated with the pool in the storage pool table.
    
 	PoolID *string `json:"pool_id,omitempty"`
   
    // Name of the pool containing the space domain.
    
 	PoolName *string `json:"pool_name,omitempty"`
   
    // Identifier of the application category associated with the space domain.
    
 	AppCategoryID *string `json:"app_category_id,omitempty"`
   
    // Name of the application category associated with the space domain.
    
 	AppCategoryName *string `json:"app_category_name,omitempty"`
   
    // Name of the performance policies associated with the space domain.
    
   	PerfPolicyNames []*NsPerfPolicySummary `json:"perf_policy_names,omitempty"`
   
    // Sample rate value.
    
   	SampleRate *int64 `json:"sample_rate,omitempty"`
   
    // Number of volumes belonging to the space domain.
    
   	VolumeCount *int64 `json:"volume_count,omitempty"`
   
    // Number of deduplicated volumes belonging to the space domain.
    
   	DedupedVolumeCount *int64 `json:"deduped_volume_count,omitempty"`
   
    // Volumes belonging to the space domain.
    
   	Volumes []*NsVolumeSummary `json:"volumes,omitempty"`
   
    // Block size in bytes of volumes belonging to the space domain.
    
   	BlockSize *int64 `json:"block_size,omitempty"`
   
    // Volumes in space domain are deduplicated by default.
    
 	Deduped *bool `json:"deduped,omitempty"`
   
    // Volumes in space domain are encrypted.
    
 	Encrypted *bool `json:"encrypted,omitempty"`
   
    // Physical space usage of volumes in the space domain.
    
   	Usage *int64 `json:"usage,omitempty"`
   
    // Logical usage of volumes in the space domain.
    
   	VolLogicalUsage *int64 `json:"vol_logical_usage,omitempty"`
   
    // Logical usage of snapshots in the space domain.
    
   	SnapLogicalUsage *int64 `json:"snap_logical_usage,omitempty"`
   
    // Mapped usage of volumes in the space domain, useful for computing clone savings.
    
   	VolMappedUsage *int64 `json:"vol_mapped_usage,omitempty"`
   
    // Logical space usage of volumes when deduped.
    
   	LogicalDedupeUsage *int64 `json:"logical_dedupe_usage,omitempty"`
   
    // Physical space usage of volumes including snapshots when deduped.
    
   	PhysicalDedupeUsage *int64 `json:"physical_dedupe_usage,omitempty"`
   
    // Space usage savings in the space domain due to compression.
    
   	SavingsCompression *int64 `json:"savings_compression,omitempty"`
   
    // Space usage savings in the space domain due to deduplication.
    
   	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
   
    // Space usage savings in the space domain due to cloning of volumes.
    
   	SavingsClone *int64 `json:"savings_clone,omitempty"`
   
    // Compressed usage of volumes and snapshots in the space domain.
    
   	CompressedUsageBytes *int64 `json:"compressed_usage_bytes,omitempty"`
   
    // Uncompressed usage of volumes and snapshots in the space domain.
    
   	UncompressedUsageBytes *int64 `json:"uncompressed_usage_bytes,omitempty"`
   
    // Compression savings for the space domain expressed as ratio.
    
  	CompressionRatio *float32 `json:"compression_ratio,omitempty"`
   
    // Deduplication savings for the space domain expressed as ratio.
    
  	DedupeRatio *float32 `json:"dedupe_ratio,omitempty"`
   
    // Clone savings for the space domain expressed as ratio.
    
  	CloneRatio *float32 `json:"clone_ratio,omitempty"`
}
