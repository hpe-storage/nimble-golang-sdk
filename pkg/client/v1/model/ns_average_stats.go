/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAverageStats - Average statistics.
// Export NsAverageStatsFields for advance operations like search filter etc.
var NsAverageStatsFields *NsAverageStats

func init(){
		
	NsAverageStatsFields= &NsAverageStats{
		
	}
}

type NsAverageStats struct {
   
    // Average read iops.
    
   	ReadIops *int64 `json:"read_iops,omitempty"`
   
    // Average read throughput.
    
   	ReadThroughput *int64 `json:"read_throughput,omitempty"`
   
    // Average read latency.
    
   	ReadLatency *int64 `json:"read_latency,omitempty"`
   
    // Average write iops.
    
   	WriteIops *int64 `json:"write_iops,omitempty"`
   
    // Average write throughput.
    
   	WriteThroughput *int64 `json:"write_throughput,omitempty"`
   
    // Average write latency.
    
   	WriteLatency *int64 `json:"write_latency,omitempty"`
   
    // Average combined iops.
    
   	CombinedIops *int64 `json:"combined_iops,omitempty"`
   
    // Average combined throughput.
    
   	CombinedThroughput *int64 `json:"combined_throughput,omitempty"`
   
    // Average combined latency.
    
   	CombinedLatency *int64 `json:"combined_latency,omitempty"`
}
