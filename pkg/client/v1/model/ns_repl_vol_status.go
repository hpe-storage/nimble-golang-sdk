/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsReplVolStatus


// NsReplVolStatus :
type NsReplVolStatus struct {
   // Name
   Name string `json:"name,omitempty"`
   // SnapName
   SnapName string `json:"snap_name,omitempty"`
   // ReplBytesDone
   ReplBytesDone float64 `json:"repl_bytes_done,omitempty"`
   // ReplBytesTotal
   ReplBytesTotal float64 `json:"repl_bytes_total,omitempty"`
}
