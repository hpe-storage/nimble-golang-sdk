// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsVolumePoolInfo - Volume information along with the pool to which it belongs to.
// Export NsVolumePoolInfoFields for advance operations like search filter etc.
var NsVolumePoolInfoFields *NsVolumePoolInfo

func init(){
	VolIdfield:= "vol_id"
	VolNamefield:= "vol_name"
	PoolIdfield:= "pool_id"
	PoolNamefield:= "pool_name"
		
	NsVolumePoolInfoFields= &NsVolumePoolInfo{
	VolId: &VolIdfield,
	VolName: &VolNamefield,
	PoolId: &PoolIdfield,
	PoolName: &PoolNamefield,
		
	}
}

type NsVolumePoolInfo struct {
	// VolId - ID of the volume.
 	VolId *string `json:"vol_id,omitempty"`
	// VolName - Name of the volume.
 	VolName *string `json:"vol_name,omitempty"`
	// PoolId - ID of the pool to which the volume belongs to.
 	PoolId *string `json:"pool_id,omitempty"`
	// PoolName - Name of the pool to which volume belongs to.
 	PoolName *string `json:"pool_name,omitempty"`
}
