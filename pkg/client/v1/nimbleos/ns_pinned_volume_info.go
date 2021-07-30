// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsPinnedVolumeInfoFields provides field names to use in filter parameters, for example.
var NsPinnedVolumeInfoFields *NsPinnedVolumeInfoFieldHandles

func init() {
	NsPinnedVolumeInfoFields = &NsPinnedVolumeInfoFieldHandles{
		ID:                   "id",
		Name:                 "name",
		SizePinnedCacheBytes: "size_pinned_cache_bytes",
	}
}

// NsPinnedVolumeInfo - Select fields containing pinned volume info.
type NsPinnedVolumeInfo struct {
	// ID - ID of volume.
	ID *string `json:"id,omitempty"`
	// Name - Name of volume.
	Name *string `json:"name,omitempty"`
	// SizePinnedCacheBytes - Total pinned cache size of the volume in bytes.
	SizePinnedCacheBytes *int64 `json:"size_pinned_cache_bytes,omitempty"`
}

// NsPinnedVolumeInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsPinnedVolumeInfoFieldHandles struct {
	ID                   string
	Name                 string
	SizePinnedCacheBytes string
}
