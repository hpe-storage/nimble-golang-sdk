/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapshotCreateAttr


// NsSnapshotCreateAttr :
type NsSnapshotCreateAttr struct {
   // VolID
   VolID string `json:"vol_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // Online
   Online bool `json:"online,omitempty"`
   // Writable
   Writable bool `json:"writable,omitempty"`
}
