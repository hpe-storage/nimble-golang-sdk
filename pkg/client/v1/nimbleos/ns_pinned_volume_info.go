// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPinnedVolumeInfo - Select fields containing pinned volume info.
// Export NsPinnedVolumeInfoFields for advance operations like search filter etc.
var NsPinnedVolumeInfoFields *NsPinnedVolumeInfoStringFields

func init() {
	IDfield := "id"
	Namefield := "name"
	SizePinnedCacheBytesfield := "size_pinned_cache_bytes"

	NsPinnedVolumeInfoFields = &NsPinnedVolumeInfoStringFields{
		ID:                   &IDfield,
		Name:                 &Namefield,
		SizePinnedCacheBytes: &SizePinnedCacheBytesfield,
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
