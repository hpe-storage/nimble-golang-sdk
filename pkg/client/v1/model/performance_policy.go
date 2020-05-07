/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model

var (
	PerformancePolicyFields = &PerformancePolicy{
		ID:   "id",
		Name: "name",
		// TODO: generate this with remaining fields
	}
)

// PerformancePolicy :
type PerformancePolicy struct {
	// ID
	ID string `json:"id,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// FullName
	FullName string `json:"full_name,omitempty"`
	// SearchName
	SearchName string `json:"search_name,omitempty"`
	// Description
	Description string `json:"description,omitempty"`
	// BlockSize
	BlockSize float64 `json:"block_size,omitempty"`
	// Compress
	Compress bool `json:"compress,omitempty"`
	// Cache
	Cache bool `json:"cache,omitempty"`
	// AppCategory
	AppCategory string `json:"app_category,omitempty"`
	// DedupeEnabled
	DedupeEnabled bool `json:"dedupe_enabled,omitempty"`
	// Deprecated
	Deprecated bool `json:"deprecated,omitempty"`
	// Predefined
	Predefined bool `json:"predefined,omitempty"`
	// CreationTime
	CreationTime float64 `json:"creation_time,omitempty"`
	// LastModified
	LastModified float64 `json:"last_modified,omitempty"`
	// SampleRate
	SampleRate float64 `json:"sample_rate,omitempty"`
	// VolumeCount
	VolumeCount float64 `json:"volume_count,omitempty"`
}
