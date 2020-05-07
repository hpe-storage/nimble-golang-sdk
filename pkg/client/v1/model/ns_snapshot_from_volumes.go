/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapshotFromVolumes


// NsSnapshotFromVolumes :
type NsSnapshotFromVolumes struct {
   // ID
   ID string `json:"id,omitempty"`
   // SnapID
   SnapID string `json:"snap_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // SnapName
   SnapName string `json:"snap_name,omitempty"`
}
