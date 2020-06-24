/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumePoolInfo - Volume information along with the pool to which it belongs to.
// Export NsVolumePoolInfoFields for advance operations like search filter etc.
var NsVolumePoolInfoFields *NsVolumePoolInfo

func init(){
	VolIDfield:= "vol_id"
	VolNamefield:= "vol_name"
	PoolIDfield:= "pool_id"
	PoolNamefield:= "pool_name"
		
	NsVolumePoolInfoFields= &NsVolumePoolInfo{
		VolID: &VolIDfield,
		VolName: &VolNamefield,
		PoolID: &PoolIDfield,
		PoolName: &PoolNamefield,
		
	}
}

type NsVolumePoolInfo struct {
   
    // ID of the volume.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Name of the volume.
    
 	VolName *string `json:"vol_name,omitempty"`
   
    // ID of the pool to which the volume belongs to.
    
 	PoolID *string `json:"pool_id,omitempty"`
   
    // Name of the pool to which volume belongs to.
    
 	PoolName *string `json:"pool_name,omitempty"`
}
