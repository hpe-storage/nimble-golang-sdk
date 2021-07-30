// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumePoolInfo - Volume information along with the pool to which it belongs to.

// Export NsVolumePoolInfoFields provides field names to use in filter parameters, for example.
var NsVolumePoolInfoFields *NsVolumePoolInfoStringFields

func init() {
	fieldVolId := "vol_id"
	fieldVolName := "vol_name"
	fieldPoolId := "pool_id"
	fieldPoolName := "pool_name"

	NsVolumePoolInfoFields = &NsVolumePoolInfoStringFields{
		VolId:    &fieldVolId,
		VolName:  &fieldVolName,
		PoolId:   &fieldPoolId,
		PoolName: &fieldPoolName,
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

// Struct for NsVolumePoolInfoFields
type NsVolumePoolInfoStringFields struct {
	VolId    *string
	VolName  *string
	PoolId   *string
	PoolName *string
}
