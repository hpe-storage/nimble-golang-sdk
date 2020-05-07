/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArrayUnassignMigStatus


// NsArrayUnassignMigStatus :
type NsArrayUnassignMigStatus struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // BytesMigrated
   BytesMigrated float64 `json:"bytes_migrated,omitempty"`
   // BytesRemaining
   BytesRemaining float64 `json:"bytes_remaining,omitempty"`
   // StartTime
   StartTime float64 `json:"start_time,omitempty"`
   // EstimatedCompletionTime
   EstimatedCompletionTime float64 `json:"estimated_completion_time,omitempty"`
}
