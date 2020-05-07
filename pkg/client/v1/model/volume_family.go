/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/VolumeFamily


// VolumeFamily :
type VolumeFamily struct {
   // ID
   ID string `json:"id,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
   // Blocksize
   Blocksize float64 `json:"blocksize,omitempty"`
   // RootVolName
   RootVolName string `json:"root_vol_name,omitempty"`
   // VolUsageCompressedBytes
   VolUsageCompressedBytes float64 `json:"vol_usage_compressed_bytes,omitempty"`
   // SnapUsageCompressedBytes
   SnapUsageCompressedBytes float64 `json:"snap_usage_compressed_bytes,omitempty"`
}
