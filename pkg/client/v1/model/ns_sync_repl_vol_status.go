/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSyncReplVolStatus


// NsSyncReplVolStatus :
type NsSyncReplVolStatus struct {
   // ResyncActive
   ResyncActive bool `json:"resync_active,omitempty"`
   // ResyncBytesDone
   ResyncBytesDone float64 `json:"resync_bytes_done,omitempty"`
   // ResyncBytesTotal
   ResyncBytesTotal float64 `json:"resync_bytes_total,omitempty"`
}
