// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsIopsMbpsStatsFields provides field names to use in filter parameters, for example.
var NsIopsMbpsStatsFields *NsIopsMbpsStatsFieldHandles

func init() {
	fieldAvgIops := "avg_iops"
	fieldMaxIops := "max_iops"
	fieldAvgMbps := "avg_mbps"
	fieldMaxMbps := "max_mbps"

	NsIopsMbpsStatsFields = &NsIopsMbpsStatsFieldHandles{
		AvgIops: &fieldAvgIops,
		MaxIops: &fieldMaxIops,
		AvgMbps: &fieldAvgMbps,
		MaxMbps: &fieldMaxMbps,
	}
}

// NsIopsMbpsStats - Average and maximum iops and mbps statistics in last 24 hours.
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

// NsIopsMbpsStatsFieldHandles provides a string representation for each AccessControlRecord field.
type NsIopsMbpsStatsFieldHandles struct {
	AvgIops *string
	MaxIops *string
	AvgMbps *string
	MaxMbps *string
}
