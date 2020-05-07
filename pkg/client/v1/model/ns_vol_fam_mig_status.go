/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVolFamMigStatus


// NsVolFamMigStatus :
type NsVolFamMigStatus struct {
   // RootVolID
   RootVolID string `json:"root_vol_id,omitempty"`
   // RootVolName
   RootVolName string `json:"root_vol_name,omitempty"`
   // SourcePoolID
   SourcePoolID string `json:"source_pool_id,omitempty"`
   // SourcePoolName
   SourcePoolName string `json:"source_pool_name,omitempty"`
   // DestPoolID
   DestPoolID string `json:"dest_pool_id,omitempty"`
   // DestPoolName
   DestPoolName string `json:"dest_pool_name,omitempty"`
   // MoveBytesMigrated
   MoveBytesMigrated float64 `json:"move_bytes_migrated,omitempty"`
   // MoveBytesRemaining
   MoveBytesRemaining float64 `json:"move_bytes_remaining,omitempty"`
   // MoveStartTime
   MoveStartTime float64 `json:"move_start_time,omitempty"`
   // MoveEstComplTime
   MoveEstComplTime float64 `json:"move_est_compl_time,omitempty"`
}
