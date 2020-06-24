/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// VolumeFamily - A volume family contains all the volumes, snapshots, and clones derived from and including a root volume.
// Export VolumeFamilyFields for advance operations like search filter etc.
var VolumeFamilyFields *VolumeFamily

func init(){
	IDfield:= "id"
	PoolIDfield:= "pool_id"
	PoolNamefield:= "pool_name"
	RootVolNamefield:= "root_vol_name"
		
	VolumeFamilyFields= &VolumeFamily{
		ID: &IDfield,
		PoolID: &PoolIDfield,
		PoolName: &PoolNamefield,
		RootVolName: &RootVolNamefield,
		
	}
}

type VolumeFamily struct {
   
    // Identifier for the volume family.
    
 	ID *string `json:"id,omitempty"`
   
    // Identifier associated with the pool in the pools table.
    
 	PoolID *string `json:"pool_id,omitempty"`
   
    // Name of the pool where the volume family resides.
    
 	PoolName *string `json:"pool_name,omitempty"`
   
    // Size of blocks in the volume.
    
   	Blocksize *int64 `json:"blocksize,omitempty"`
   
    // Volume family root volume name.
    
 	RootVolName *string `json:"root_vol_name,omitempty"`
   
    // List of volumes.
    
   	Volumes []*NsVolumeSummary `json:"volumes,omitempty"`
   
    // Sum of compressed bytes stored in the volumes of this family.
    
   	VolUsageCompressedBytes *int64 `json:"vol_usage_compressed_bytes,omitempty"`
   
    // Sum of compressed bytes stored in the snapshots of this family.
    
   	SnapUsageCompressedBytes *int64 `json:"snap_usage_compressed_bytes,omitempty"`
}
