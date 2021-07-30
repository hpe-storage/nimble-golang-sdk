// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export SpaceDomainFields provides field names to use in filter parameters, for example.
var SpaceDomainFields *SpaceDomainFieldHandles

func init() {
	fieldID := "id"
	fieldPoolId := "pool_id"
	fieldPoolName := "pool_name"
	fieldAppCategoryId := "app_category_id"
	fieldAppCategoryName := "app_category_name"
	fieldPerfPolicyNames := "perf_policy_names"
	fieldSampleRate := "sample_rate"
	fieldVolumeCount := "volume_count"
	fieldDedupedVolumeCount := "deduped_volume_count"
	fieldVolumes := "volumes"
	fieldBlockSize := "block_size"
	fieldDeduped := "deduped"
	fieldEncrypted := "encrypted"
	fieldUsage := "usage"
	fieldVolLogicalUsage := "vol_logical_usage"
	fieldSnapLogicalUsage := "snap_logical_usage"
	fieldVolMappedUsage := "vol_mapped_usage"
	fieldLogicalDedupeUsage := "logical_dedupe_usage"
	fieldPhysicalDedupeUsage := "physical_dedupe_usage"
	fieldSavingsCompression := "savings_compression"
	fieldSavingsDedupe := "savings_dedupe"
	fieldSavingsClone := "savings_clone"
	fieldCompressedUsageBytes := "compressed_usage_bytes"
	fieldUncompressedUsageBytes := "uncompressed_usage_bytes"
	fieldCompressionRatio := "compression_ratio"
	fieldDedupeRatio := "dedupe_ratio"
	fieldCloneRatio := "clone_ratio"

	SpaceDomainFields = &SpaceDomainFieldHandles{
		ID:                     &fieldID,
		PoolId:                 &fieldPoolId,
		PoolName:               &fieldPoolName,
		AppCategoryId:          &fieldAppCategoryId,
		AppCategoryName:        &fieldAppCategoryName,
		PerfPolicyNames:        &fieldPerfPolicyNames,
		SampleRate:             &fieldSampleRate,
		VolumeCount:            &fieldVolumeCount,
		DedupedVolumeCount:     &fieldDedupedVolumeCount,
		Volumes:                &fieldVolumes,
		BlockSize:              &fieldBlockSize,
		Deduped:                &fieldDeduped,
		Encrypted:              &fieldEncrypted,
		Usage:                  &fieldUsage,
		VolLogicalUsage:        &fieldVolLogicalUsage,
		SnapLogicalUsage:       &fieldSnapLogicalUsage,
		VolMappedUsage:         &fieldVolMappedUsage,
		LogicalDedupeUsage:     &fieldLogicalDedupeUsage,
		PhysicalDedupeUsage:    &fieldPhysicalDedupeUsage,
		SavingsCompression:     &fieldSavingsCompression,
		SavingsDedupe:          &fieldSavingsDedupe,
		SavingsClone:           &fieldSavingsClone,
		CompressedUsageBytes:   &fieldCompressedUsageBytes,
		UncompressedUsageBytes: &fieldUncompressedUsageBytes,
		CompressionRatio:       &fieldCompressionRatio,
		DedupeRatio:            &fieldDedupeRatio,
		CloneRatio:             &fieldCloneRatio,
	}
}

// SpaceDomain - A space domain is created for each application category and block size for each each pool.
type SpaceDomain struct {
	// ID - Identifier for the space domain.
	ID *string `json:"id,omitempty"`
	// PoolId - Identifier associated with the pool in the storage pool table.
	PoolId *string `json:"pool_id,omitempty"`
	// PoolName - Name of the pool containing the space domain.
	PoolName *string `json:"pool_name,omitempty"`
	// AppCategoryId - Identifier of the application category associated with the space domain.
	AppCategoryId *string `json:"app_category_id,omitempty"`
	// AppCategoryName - Name of the application category associated with the space domain.
	AppCategoryName *string `json:"app_category_name,omitempty"`
	// PerfPolicyNames - Name of the performance policies associated with the space domain.
	PerfPolicyNames []*NsPerfPolicySummary `json:"perf_policy_names,omitempty"`
	// SampleRate - Sample rate value.
	SampleRate *int64 `json:"sample_rate,omitempty"`
	// VolumeCount - Number of volumes belonging to the space domain.
	VolumeCount *int64 `json:"volume_count,omitempty"`
	// DedupedVolumeCount - Number of deduplicated volumes belonging to the space domain.
	DedupedVolumeCount *int64 `json:"deduped_volume_count,omitempty"`
	// Volumes - Volumes belonging to the space domain.
	Volumes []*NsVolumeSummary `json:"volumes,omitempty"`
	// BlockSize - Block size in bytes of volumes belonging to the space domain.
	BlockSize *int64 `json:"block_size,omitempty"`
	// Deduped - Volumes in space domain are deduplicated by default.
	Deduped *bool `json:"deduped,omitempty"`
	// Encrypted - Volumes in space domain are encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
	// Usage - Physical space usage of volumes in the space domain.
	Usage *int64 `json:"usage,omitempty"`
	// VolLogicalUsage - Logical usage of volumes in the space domain.
	VolLogicalUsage *int64 `json:"vol_logical_usage,omitempty"`
	// SnapLogicalUsage - Logical usage of snapshots in the space domain.
	SnapLogicalUsage *int64 `json:"snap_logical_usage,omitempty"`
	// VolMappedUsage - Mapped usage of volumes in the space domain, useful for computing clone savings.
	VolMappedUsage *int64 `json:"vol_mapped_usage,omitempty"`
	// LogicalDedupeUsage - Logical space usage of volumes when deduped.
	LogicalDedupeUsage *int64 `json:"logical_dedupe_usage,omitempty"`
	// PhysicalDedupeUsage - Physical space usage of volumes including snapshots when deduped.
	PhysicalDedupeUsage *int64 `json:"physical_dedupe_usage,omitempty"`
	// SavingsCompression - Space usage savings in the space domain due to compression.
	SavingsCompression *int64 `json:"savings_compression,omitempty"`
	// SavingsDedupe - Space usage savings in the space domain due to deduplication.
	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
	// SavingsClone - Space usage savings in the space domain due to cloning of volumes.
	SavingsClone *int64 `json:"savings_clone,omitempty"`
	// CompressedUsageBytes - Compressed usage of volumes and snapshots in the space domain.
	CompressedUsageBytes *int64 `json:"compressed_usage_bytes,omitempty"`
	// UncompressedUsageBytes - Uncompressed usage of volumes and snapshots in the space domain.
	UncompressedUsageBytes *int64 `json:"uncompressed_usage_bytes,omitempty"`
	// CompressionRatio - Compression savings for the space domain expressed as ratio.
	CompressionRatio *float64 `json:"compression_ratio,omitempty"`
	// DedupeRatio - Deduplication savings for the space domain expressed as ratio.
	DedupeRatio *float64 `json:"dedupe_ratio,omitempty"`
	// CloneRatio - Clone savings for the space domain expressed as ratio.
	CloneRatio *float64 `json:"clone_ratio,omitempty"`
}

// SpaceDomainFieldHandles provides a string representation for each AccessControlRecord field.
type SpaceDomainFieldHandles struct {
	ID                     *string
	PoolId                 *string
	PoolName               *string
	AppCategoryId          *string
	AppCategoryName        *string
	PerfPolicyNames        *string
	SampleRate             *string
	VolumeCount            *string
	DedupedVolumeCount     *string
	Volumes                *string
	BlockSize              *string
	Deduped                *string
	Encrypted              *string
	Usage                  *string
	VolLogicalUsage        *string
	SnapLogicalUsage       *string
	VolMappedUsage         *string
	LogicalDedupeUsage     *string
	PhysicalDedupeUsage    *string
	SavingsCompression     *string
	SavingsDedupe          *string
	SavingsClone           *string
	CompressedUsageBytes   *string
	UncompressedUsageBytes *string
	CompressionRatio       *string
	DedupeRatio            *string
	CloneRatio             *string
}
