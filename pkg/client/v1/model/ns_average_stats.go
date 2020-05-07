/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsAverageStats


// NsAverageStats :
type NsAverageStats struct {
   // ReadIops
   ReadIops float64 `json:"read_iops,omitempty"`
   // ReadThroughput
   ReadThroughput float64 `json:"read_throughput,omitempty"`
   // ReadLatency
   ReadLatency float64 `json:"read_latency,omitempty"`
   // WriteIops
   WriteIops float64 `json:"write_iops,omitempty"`
   // WriteThroughput
   WriteThroughput float64 `json:"write_throughput,omitempty"`
   // WriteLatency
   WriteLatency float64 `json:"write_latency,omitempty"`
   // CombinedIops
   CombinedIops float64 `json:"combined_iops,omitempty"`
   // CombinedThroughput
   CombinedThroughput float64 `json:"combined_throughput,omitempty"`
   // CombinedLatency
   CombinedLatency float64 `json:"combined_latency,omitempty"`
}
