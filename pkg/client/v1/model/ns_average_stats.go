// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsAverageStats - Average statistics.
// Export NsAverageStatsFields for advance operations like search filter etc.
var NsAverageStatsFields *NsAverageStats

func init() {

	NsAverageStatsFields = &NsAverageStats{}
}

type NsAverageStats struct {
	// ReadIops - Average read iops.
	ReadIops int64 `json:"read_iops,omitempty"`
	// ReadThroughput - Average read throughput.
	ReadThroughput int64 `json:"read_throughput,omitempty"`
	// ReadLatency - Average read latency.
	ReadLatency int64 `json:"read_latency,omitempty"`
	// WriteIops - Average write iops.
	WriteIops int64 `json:"write_iops,omitempty"`
	// WriteThroughput - Average write throughput.
	WriteThroughput int64 `json:"write_throughput,omitempty"`
	// WriteLatency - Average write latency.
	WriteLatency int64 `json:"write_latency,omitempty"`
	// CombinedIops - Average combined iops.
	CombinedIops int64 `json:"combined_iops,omitempty"`
	// CombinedThroughput - Average combined throughput.
	CombinedThroughput int64 `json:"combined_throughput,omitempty"`
	// CombinedLatency - Average combined latency.
	CombinedLatency int64 `json:"combined_latency,omitempty"`
}

// sdk internal struct
type nsAverageStats struct {
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

// EncodeNsAverageStats - Transform NsAverageStats to nsAverageStats type
func EncodeNsAverageStats(request interface{}) (*nsAverageStats, error) {
	reqNsAverageStats := request.(*NsAverageStats)
	byte, err := json.Marshal(reqNsAverageStats)
	objPtr := &nsAverageStats{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsAverageStats Transform nsAverageStats to NsAverageStats type
func DecodeNsAverageStats(request interface{}) (*NsAverageStats, error) {
	reqNsAverageStats := request.(*nsAverageStats)
	byte, err := json.Marshal(reqNsAverageStats)
	obj := &NsAverageStats{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
