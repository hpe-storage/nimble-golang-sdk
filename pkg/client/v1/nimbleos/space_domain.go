// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// SpaceDomain - A space domain is created for each application category and block size for each each pool.
// Export SpaceDomainFields for advance operations like search filter etc.
var SpaceDomainFields *SpaceDomain

func init() {

	SpaceDomainFields = &SpaceDomain{
		ID:              "id",
		PoolId:          "pool_id",
		PoolName:        "pool_name",
		AppCategoryId:   "app_category_id",
		AppCategoryName: "app_category_name",
	}
}

type SpaceDomain struct {
	// ID - Identifier for the space domain.
	ID string `json:"id,omitempty"`
	// PoolId - Identifier associated with the pool in the storage pool table.
	PoolId string `json:"pool_id,omitempty"`
	// PoolName - Name of the pool containing the space domain.
	PoolName string `json:"pool_name,omitempty"`
	// AppCategoryId - Identifier of the application category associated with the space domain.
	AppCategoryId string `json:"app_category_id,omitempty"`
	// AppCategoryName - Name of the application category associated with the space domain.
	AppCategoryName string `json:"app_category_name,omitempty"`
	// PerfPolicyNames - Name of the performance policies associated with the space domain.
	PerfPolicyNames []*NsPerfPolicySummary `json:"perf_policy_names,omitempty"`
	// SampleRate - Sample rate value.
	SampleRate int64 `json:"sample_rate,omitempty"`
	// VolumeCount - Number of volumes belonging to the space domain.
	VolumeCount int64 `json:"volume_count,omitempty"`
	// DedupedVolumeCount - Number of deduplicated volumes belonging to the space domain.
	DedupedVolumeCount int64 `json:"deduped_volume_count,omitempty"`
	// Volumes - Volumes belonging to the space domain.
	Volumes []*NsVolumeSummary `json:"volumes,omitempty"`
	// BlockSize - Block size in bytes of volumes belonging to the space domain.
	BlockSize int64 `json:"block_size,omitempty"`
	// Deduped - Volumes in space domain are deduplicated by default.
	Deduped *bool `json:"deduped,omitempty"`
	// Encrypted - Volumes in space domain are encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
	// Usage - Physical space usage of volumes in the space domain.
	Usage int64 `json:"usage,omitempty"`
	// VolLogicalUsage - Logical usage of volumes in the space domain.
	VolLogicalUsage int64 `json:"vol_logical_usage,omitempty"`
	// SnapLogicalUsage - Logical usage of snapshots in the space domain.
	SnapLogicalUsage int64 `json:"snap_logical_usage,omitempty"`
	// VolMappedUsage - Mapped usage of volumes in the space domain, useful for computing clone savings.
	VolMappedUsage int64 `json:"vol_mapped_usage,omitempty"`
	// LogicalDedupeUsage - Logical space usage of volumes when deduped.
	LogicalDedupeUsage int64 `json:"logical_dedupe_usage,omitempty"`
	// PhysicalDedupeUsage - Physical space usage of volumes including snapshots when deduped.
	PhysicalDedupeUsage int64 `json:"physical_dedupe_usage,omitempty"`
	// SavingsCompression - Space usage savings in the space domain due to compression.
	SavingsCompression int64 `json:"savings_compression,omitempty"`
	// SavingsDedupe - Space usage savings in the space domain due to deduplication.
	SavingsDedupe int64 `json:"savings_dedupe,omitempty"`
	// SavingsClone - Space usage savings in the space domain due to cloning of volumes.
	SavingsClone int64 `json:"savings_clone,omitempty"`
	// CompressedUsageBytes - Compressed usage of volumes and snapshots in the space domain.
	CompressedUsageBytes int64 `json:"compressed_usage_bytes,omitempty"`
	// UncompressedUsageBytes - Uncompressed usage of volumes and snapshots in the space domain.
	UncompressedUsageBytes int64 `json:"uncompressed_usage_bytes,omitempty"`
	// CompressionRatio - Compression savings for the space domain expressed as ratio.
	CompressionRatio float32 `json:"compression_ratio,omitempty"`
	// DedupeRatio - Deduplication savings for the space domain expressed as ratio.
	DedupeRatio float32 `json:"dedupe_ratio,omitempty"`
	// CloneRatio - Clone savings for the space domain expressed as ratio.
	CloneRatio float32 `json:"clone_ratio,omitempty"`
}

// sdk internal struct
type spaceDomain struct {
	ID                     *string                `json:"id,omitempty"`
	PoolId                 *string                `json:"pool_id,omitempty"`
	PoolName               *string                `json:"pool_name,omitempty"`
	AppCategoryId          *string                `json:"app_category_id,omitempty"`
	AppCategoryName        *string                `json:"app_category_name,omitempty"`
	PerfPolicyNames        []*NsPerfPolicySummary `json:"perf_policy_names,omitempty"`
	SampleRate             *int64                 `json:"sample_rate,omitempty"`
	VolumeCount            *int64                 `json:"volume_count,omitempty"`
	DedupedVolumeCount     *int64                 `json:"deduped_volume_count,omitempty"`
	Volumes                []*NsVolumeSummary     `json:"volumes,omitempty"`
	BlockSize              *int64                 `json:"block_size,omitempty"`
	Deduped                *bool                  `json:"deduped,omitempty"`
	Encrypted              *bool                  `json:"encrypted,omitempty"`
	Usage                  *int64                 `json:"usage,omitempty"`
	VolLogicalUsage        *int64                 `json:"vol_logical_usage,omitempty"`
	SnapLogicalUsage       *int64                 `json:"snap_logical_usage,omitempty"`
	VolMappedUsage         *int64                 `json:"vol_mapped_usage,omitempty"`
	LogicalDedupeUsage     *int64                 `json:"logical_dedupe_usage,omitempty"`
	PhysicalDedupeUsage    *int64                 `json:"physical_dedupe_usage,omitempty"`
	SavingsCompression     *int64                 `json:"savings_compression,omitempty"`
	SavingsDedupe          *int64                 `json:"savings_dedupe,omitempty"`
	SavingsClone           *int64                 `json:"savings_clone,omitempty"`
	CompressedUsageBytes   *int64                 `json:"compressed_usage_bytes,omitempty"`
	UncompressedUsageBytes *int64                 `json:"uncompressed_usage_bytes,omitempty"`
	CompressionRatio       *float32               `json:"compression_ratio,omitempty"`
	DedupeRatio            *float32               `json:"dedupe_ratio,omitempty"`
	CloneRatio             *float32               `json:"clone_ratio,omitempty"`
}

// EncodeSpaceDomain - Transform SpaceDomain to spaceDomain type
func EncodeSpaceDomain(request interface{}) (*spaceDomain, error) {
	reqSpaceDomain := request.(*SpaceDomain)
	byte, err := json.Marshal(reqSpaceDomain)
	if err != nil {
		return nil, err
	}
	respSpaceDomainPtr := &spaceDomain{}
	err = json.Unmarshal(byte, respSpaceDomainPtr)
	return respSpaceDomainPtr, err
}

// DecodeSpaceDomain Transform spaceDomain to SpaceDomain type
func DecodeSpaceDomain(request interface{}) (*SpaceDomain, error) {
	reqSpaceDomain := request.(*spaceDomain)
	byte, err := json.Marshal(reqSpaceDomain)
	if err != nil {
		return nil, err
	}
	respSpaceDomain := &SpaceDomain{}
	err = json.Unmarshal(byte, respSpaceDomain)
	return respSpaceDomain, err
}
