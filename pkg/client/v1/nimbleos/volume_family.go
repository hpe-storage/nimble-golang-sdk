// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// VolumeFamily - A volume family contains all the volumes, snapshots, and clones derived from and including a root volume.
// Export VolumeFamilyFields for advance operations like search filter etc.
var VolumeFamilyFields *VolumeFamily

func init() {
	IDfield := "id"
	PoolIdfield := "pool_id"
	PoolNamefield := "pool_name"
	RootVolNamefield := "root_vol_name"

	VolumeFamilyFields = &VolumeFamily{
		ID:          &IDfield,
		PoolId:      &PoolIdfield,
		PoolName:    &PoolNamefield,
		RootVolName: &RootVolNamefield,
	}
}

type VolumeFamily struct {
	// ID - Identifier for the volume family.
	ID *string `json:"id,omitempty"`
	// PoolId - Identifier associated with the pool in the pools table.
	PoolId *string `json:"pool_id,omitempty"`
	// PoolName - Name of the pool where the volume family resides.
	PoolName *string `json:"pool_name,omitempty"`
	// Blocksize - Size of blocks in the volume.
	Blocksize *int64 `json:"blocksize,omitempty"`
	// RootVolName - Volume family root volume name.
	RootVolName *string `json:"root_vol_name,omitempty"`
	// Volumes - List of volumes.
	Volumes []*NsVolumeSummary `json:"volumes,omitempty"`
	// VolUsageCompressedBytes - Sum of compressed bytes stored in the volumes of this family.
	VolUsageCompressedBytes *int64 `json:"vol_usage_compressed_bytes,omitempty"`
	// SnapUsageCompressedBytes - Sum of compressed bytes stored in the snapshots of this family.
	SnapUsageCompressedBytes *int64 `json:"snap_usage_compressed_bytes,omitempty"`
}
