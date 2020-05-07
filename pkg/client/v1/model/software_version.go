/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/SoftwareVersion


// SoftwareVersion :
type SoftwareVersion struct {
   // Version
   Version string `json:"version,omitempty"`
   // Signature
   Signature string `json:"signature,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Status
   Status string `json:"status,omitempty"`
   // TotalBytes
   TotalBytes float64 `json:"total_bytes,omitempty"`
   // DownloadedBytes
   DownloadedBytes float64 `json:"downloaded_bytes,omitempty"`
   // BlacklistReason
   BlacklistReason string `json:"blacklist_reason,omitempty"`
   // ReleaseDate
   ReleaseDate float64 `json:"release_date,omitempty"`
   // IsManuallyDownloaded
   IsManuallyDownloaded bool `json:"is_manually_downloaded,omitempty"`
   // NoPartialResponse
   NoPartialResponse bool `json:"no_partial_response,omitempty"`
}
