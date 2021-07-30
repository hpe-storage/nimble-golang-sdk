// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPinnedVolumeInfo - Select fields containing pinned volume info.

// Export NsPinnedVolumeInfoFields provides field names to use in filter parameters, for example.
var NsPinnedVolumeInfoFields *NsPinnedVolumeInfoStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldSizePinnedCacheBytes := "size_pinned_cache_bytes"

	NsPinnedVolumeInfoFields = &NsPinnedVolumeInfoStringFields{
		ID:                   &fieldID,
		Name:                 &fieldName,
		SizePinnedCacheBytes: &fieldSizePinnedCacheBytes,
	}
}

type NsPinnedVolumeInfo struct {
	// ID - ID of volume.
	ID *string `json:"id,omitempty"`
	// Name - Name of volume.
	Name *string `json:"name,omitempty"`
	// SizePinnedCacheBytes - Total pinned cache size of the volume in bytes.
	SizePinnedCacheBytes *int64 `json:"size_pinned_cache_bytes,omitempty"`
}

// Struct for NsPinnedVolumeInfoFields
type NsPinnedVolumeInfoStringFields struct {
	ID                   *string
	Name                 *string
	SizePinnedCacheBytes *string
}
