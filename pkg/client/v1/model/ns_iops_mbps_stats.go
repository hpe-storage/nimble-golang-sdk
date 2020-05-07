/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsIopsMbpsStats


// NsIopsMbpsStats :
type NsIopsMbpsStats struct {
   // AvgIops
   AvgIops float64 `json:"avg_iops,omitempty"`
   // MaxIops
   MaxIops float64 `json:"max_iops,omitempty"`
   // AvgMbps
   AvgMbps float64 `json:"avg_mbps,omitempty"`
   // MaxMbps
   MaxMbps float64 `json:"max_mbps,omitempty"`
}
