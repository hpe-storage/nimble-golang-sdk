/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Snapshot


// Snapshot :
type Snapshot struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // Size
   Size float64 `json:"size,omitempty"`
   // VolName
   VolName string `json:"vol_name,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
   // VolID
   VolID string `json:"vol_id,omitempty"`
   // SnapCollectionName
   SnapCollectionName string `json:"snap_collection_name,omitempty"`
   // SnapCollectionID
   SnapCollectionID string `json:"snap_collection_id,omitempty"`
   // Online
   Online bool `json:"online,omitempty"`
   // Writable
   Writable bool `json:"writable,omitempty"`
   // ExpiryTime
   ExpiryTime float64 `json:"expiry_time,omitempty"`
   // ExpiryAfter
   ExpiryAfter  int32 `json:"expiry_after,omitempty"`
   // OriginName
   OriginName string `json:"origin_name,omitempty"`
   // IsReplica
   IsReplica bool `json:"is_replica,omitempty"`
   // IsUnmanaged
   IsUnmanaged bool `json:"is_unmanaged,omitempty"`
   // IsManuallyManaged
   IsManuallyManaged bool `json:"is_manually_managed,omitempty"`
   // SerialNumber
   SerialNumber string `json:"serial_number,omitempty"`
   // TargetName
   TargetName string `json:"target_name,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // ScheduleName
   ScheduleName string `json:"schedule_name,omitempty"`
   // ScheduleID
   ScheduleID string `json:"schedule_id,omitempty"`
   // AppUuID
   AppUuID string `json:"app_uuid,omitempty"`
   // NewDataValID
   NewDataValID bool `json:"new_data_valid,omitempty"`
   // NewDataCompressedBytes
   NewDataCompressedBytes float64 `json:"new_data_compressed_bytes,omitempty"`
   // NewDataUncompressedBytes
   NewDataUncompressedBytes float64 `json:"new_data_uncompressed_bytes,omitempty"`
   // VpdT10
   VpdT10 string `json:"vpd_t10,omitempty"`
   // VpdIeee0
   VpdIeee0 string `json:"vpd_ieee0,omitempty"`
   // VpdIeee1
   VpdIeee1 string `json:"vpd_ieee1,omitempty"`
}
