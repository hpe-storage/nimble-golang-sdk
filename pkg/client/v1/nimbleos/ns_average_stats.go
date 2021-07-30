// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAverageStats - Average statistics.
// Export NsAverageStatsFields for advance operations like search filter etc.
var NsAverageStatsFields *NsAverageStatsStringFields

func init() {
	ReadIopsfield := "read_iops"
	ReadThroughputfield := "read_throughput"
	ReadLatencyfield := "read_latency"
	WriteIopsfield := "write_iops"
	WriteThroughputfield := "write_throughput"
	WriteLatencyfield := "write_latency"
	CombinedIopsfield := "combined_iops"
	CombinedThroughputfield := "combined_throughput"
	CombinedLatencyfield := "combined_latency"

	NsAverageStatsFields = &NsAverageStatsStringFields{
		ReadIops:           &ReadIopsfield,
		ReadThroughput:     &ReadThroughputfield,
		ReadLatency:        &ReadLatencyfield,
		WriteIops:          &WriteIopsfield,
		WriteThroughput:    &WriteThroughputfield,
		WriteLatency:       &WriteLatencyfield,
		CombinedIops:       &CombinedIopsfield,
		CombinedThroughput: &CombinedThroughputfield,
		CombinedLatency:    &CombinedLatencyfield,
	}
}

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

// Struct for NsAverageStatsFields
type NsAverageStatsStringFields struct {
	ReadIops           *string
	ReadThroughput     *string
	ReadLatency        *string
	WriteIops          *string
	WriteThroughput    *string
	WriteLatency       *string
	CombinedIops       *string
	CombinedThroughput *string
	CombinedLatency    *string
}
