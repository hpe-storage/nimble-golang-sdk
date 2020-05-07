/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapshotFromSnapshotCollections


// NsSnapshotFromSnapshotCollections :
type NsSnapshotFromSnapshotCollections struct {
   // ID
   ID string `json:"id,omitempty"`
   // SnapcollID
   SnapcollID string `json:"snapcoll_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // SnapcollName
   SnapcollName string `json:"snapcoll_name,omitempty"`
   // VolID
   VolID string `json:"vol_id,omitempty"`
   // VolName
   VolName string `json:"vol_name,omitempty"`
   // SnapID
   SnapID string `json:"snap_id,omitempty"`
   // SnapName
   SnapName string `json:"snap_name,omitempty"`
   // ScheduleID
   ScheduleID string `json:"schedule_id,omitempty"`
   // ScheduleName
   ScheduleName string `json:"schedule_name,omitempty"`
   // ExpiryTime
   ExpiryTime float64 `json:"expiry_time,omitempty"`
}
