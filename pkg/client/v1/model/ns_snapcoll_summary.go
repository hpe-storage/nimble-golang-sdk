/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapcollSummary


// NsSnapcollSummary :
type NsSnapcollSummary struct {
   // SnapcollID
   SnapcollID string `json:"snapcoll_id,omitempty"`
   // SnapcollName
   SnapcollName string `json:"snapcoll_name,omitempty"`
   // SnapcollCreationTime
   SnapcollCreationTime float64 `json:"snapcoll_creation_time,omitempty"`
}
