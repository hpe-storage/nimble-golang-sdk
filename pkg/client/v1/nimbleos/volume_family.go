// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export VolumeFamilyFields provides field names to use in filter parameters, for example.
var VolumeFamilyFields *VolumeFamilyFieldHandles

func init() {
	VolumeFamilyFields = &VolumeFamilyFieldHandles{
		ID:                       "id",
		PoolId:                   "pool_id",
		PoolName:                 "pool_name",
		Blocksize:                "blocksize",
		RootVolName:              "root_vol_name",
		Volumes:                  "volumes",
		VolUsageCompressedBytes:  "vol_usage_compressed_bytes",
		SnapUsageCompressedBytes: "snap_usage_compressed_bytes",
	}
}

// VolumeFamily - A volume family contains all the volumes, snapshots, and clones derived from and including a root volume.
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

// VolumeFamilyFieldHandles provides a string representation for each AccessControlRecord field.
type VolumeFamilyFieldHandles struct {
	ID                       string
	PoolId                   string
	PoolName                 string
	Blocksize                string
	RootVolName              string
	Volumes                  string
	VolUsageCompressedBytes  string
	SnapUsageCompressedBytes string
}
