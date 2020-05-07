/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapSummary


// NsSnapSummary :
type NsSnapSummary struct {
   // SnapID
   SnapID string `json:"snap_id,omitempty"`
   // SnapName
   SnapName string `json:"snap_name,omitempty"`
   // SnapCreationTime
   SnapCreationTime float64 `json:"snap_creation_time,omitempty"`
}
