// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsIopsMbpsStats - Average and maximum iops and mbps statistics in last 24 hours.
// Export NsIopsMbpsStatsFields for advance operations like search filter etc.
var NsIopsMbpsStatsFields *NsIopsMbpsStats

func init(){
		
	NsIopsMbpsStatsFields= &NsIopsMbpsStats{
		
	}
}

type NsIopsMbpsStats struct {
	// AvgIops - Average combined read and write iops.
   	AvgIops *int64 `json:"avg_iops,omitempty"`
	// MaxIops - Maximum combined read and write iops.
   	MaxIops *int64 `json:"max_iops,omitempty"`
	// AvgMbps - Average combined read and write throughput.
   	AvgMbps *int64 `json:"avg_mbps,omitempty"`
	// MaxMbps - Maximum combined read and write throughput.
   	MaxMbps *int64 `json:"max_mbps,omitempty"`
}
