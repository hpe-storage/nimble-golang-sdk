// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAverageStatsFields provides field names to use in filter parameters, for example.
var NsAverageStatsFields *NsAverageStatsFieldHandles

func init() {
	NsAverageStatsFields = &NsAverageStatsFieldHandles{
		ReadIops:           "read_iops",
		ReadThroughput:     "read_throughput",
		ReadLatency:        "read_latency",
		WriteIops:          "write_iops",
		WriteThroughput:    "write_throughput",
		WriteLatency:       "write_latency",
		CombinedIops:       "combined_iops",
		CombinedThroughput: "combined_throughput",
		CombinedLatency:    "combined_latency",
	}
}

// NsAverageStats - Average statistics.
type NsAverageStats struct {
	// ReadIops - Average read iops.
	ReadIops *int64 `json:"read_iops,omitempty"`
	// ReadThroughput - Average read throughput.
	ReadThroughput *int64 `json:"read_throughput,omitempty"`
	// ReadLatency - Average read latency.
	ReadLatency *int64 `json:"read_latency,omitempty"`
	// WriteIops - Average write iops.
	WriteIops *int64 `json:"write_iops,omitempty"`
	// WriteThroughput - Average write throughput.
	WriteThroughput *int64 `json:"write_throughput,omitempty"`
	// WriteLatency - Average write latency.
	WriteLatency *int64 `json:"write_latency,omitempty"`
	// CombinedIops - Average combined iops.
	CombinedIops *int64 `json:"combined_iops,omitempty"`
	// CombinedThroughput - Average combined throughput.
	CombinedThroughput *int64 `json:"combined_throughput,omitempty"`
	// CombinedLatency - Average combined latency.
	CombinedLatency *int64 `json:"combined_latency,omitempty"`
}

// NsAverageStatsFieldHandles provides a string representation for each AccessControlRecord field.
type NsAverageStatsFieldHandles struct {
	ReadIops           string
	ReadThroughput     string
	ReadLatency        string
	WriteIops          string
	WriteThroughput    string
	WriteLatency       string
	CombinedIops       string
	CombinedThroughput string
	CombinedLatency    string
}
