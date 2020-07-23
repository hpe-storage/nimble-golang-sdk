// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// PerformancePolicy - Manage performance policies. A performance policy is a set of optimizations including block size, compression, and caching, to ensure that the volume's performance is the best configuration for its intended use like databases or log files. By default, a volume uses the \\"default\\" performance policy, which is set to use 4096 byte blocks with full compression and caching enabled. For replicated volumes, the same performance policy must exist on each replication partner.
// Export PerformancePolicyFields for advance operations like search filter etc.
var PerformancePolicyFields *PerformancePolicy

func init() {
	IDfield := "id"
	Namefield := "name"
	FullNamefield := "full_name"
	SearchNamefield := "search_name"
	Descriptionfield := "description"
	AppCategoryfield := "app_category"

	PerformancePolicyFields = &PerformancePolicy{
		ID:          &IDfield,
		Name:        &Namefield,
		FullName:    &FullNamefield,
		SearchName:  &SearchNamefield,
		Description: &Descriptionfield,
		AppCategory: &AppCategoryfield,
	}
}

type PerformancePolicy struct {
	// ID - Unique Identifier for the Performance Policy.
	ID *string `json:"id,omitempty"`
	// Name - Name of the Performance Policy.
	Name *string `json:"name,omitempty"`
	// FullName - Fully qualified name of the Performance Policy.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - Name of the Performance Policy used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Description of a performance policy.
	Description *string `json:"description,omitempty"`
	// BlockSize - Block Size in bytes to be used by the volumes created with this specific performance policy. Supported block sizes are 4096 bytes (4 KB), 8192 bytes (8 KB), 16384 bytes(16 KB), and 32768 bytes (32 KB). Block size of a performance policy cannot be changed once the performance policy is created.
	BlockSize *int64 `json:"block_size,omitempty"`
	// Compress - Flag denoting if data in the associated volume should be compressed.
	Compress *bool `json:"compress,omitempty"`
	// Cache - Flag denoting if data in the associated volume should be cached.
	Cache *bool `json:"cache,omitempty"`
	// CachePolicy - Specifies how data of associated volume should be cached. Supports two policies, 'normal' and 'aggressive'. 'normal' policy caches data but skips in certain conditions such as sequential I/O. 'aggressive' policy will accelerate caching of all data belonging to this volume, regardless of sequentiality.
	CachePolicy *NsCachePolicy `json:"cache_policy,omitempty"`
	// SpacePolicy - Specifies the state of the volume upon space constraint violation such as volume limit violation or volumes above their volume reserve, if the pool free space is exhausted. Supports two policies, 'offline' and 'non_writable'.
	SpacePolicy *NsSpacePolicy `json:"space_policy,omitempty"`
	// AppCategory - Specifies the application category of the associated volume.
	AppCategory *string `json:"app_category,omitempty"`
	// DedupeEnabled - Specifies if dedupe is enabled for volumes created with this performance policy.
	DedupeEnabled *bool `json:"dedupe_enabled,omitempty"`
	// Deprecated - Specifies if this performance policy is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`
	// Predefined - Specifies if this performance policy is predefined (read-only).
	Predefined *bool `json:"predefined,omitempty"`
	// CreationTime - Time when the performance policy was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when the performance policy's configurations were last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// SampleRate - Sample rate value.
	SampleRate *int64 `json:"sample_rate,omitempty"`
	// VolumeCount - Number of volumes using this performance policy.
	VolumeCount *int64 `json:"volume_count,omitempty"`
	// DedupeOverridePools - List of pools that override performance policy's dedupe setting.
	DedupeOverridePools []*NsPoolSummary `json:"dedupe_override_pools,omitempty"`
}
