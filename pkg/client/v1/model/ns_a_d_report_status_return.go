/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsADReportStatusReturn


// NsADReportStatusReturn :
type NsADReportStatusReturn struct {
   // Joined
   Joined bool `json:"joined,omitempty"`
   // Enabled
   Enabled bool `json:"enabled,omitempty"`
   // LocalServiceStatus
   LocalServiceStatus bool `json:"local_service_status,omitempty"`
   // RemoteServiceStatus
   RemoteServiceStatus bool `json:"remote_service_status,omitempty"`
   // TrustValID
   TrustValID bool `json:"trust_valid,omitempty"`
}
