/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVolumePoolInfo


// NsVolumePoolInfo :
type NsVolumePoolInfo struct {
   // VolID
   VolID string `json:"vol_id,omitempty"`
   // VolName
   VolName string `json:"vol_name,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
}
