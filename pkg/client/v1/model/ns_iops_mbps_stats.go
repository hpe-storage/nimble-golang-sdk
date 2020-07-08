// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsIopsMbpsStats - Average and maximum iops and mbps statistics in last 24 hours.
// Export NsIopsMbpsStatsFields for advance operations like search filter etc.
var NsIopsMbpsStatsFields *NsIopsMbpsStats

func init() {

	NsIopsMbpsStatsFields = &NsIopsMbpsStats{}
}

type NsIopsMbpsStats struct {
	// AvgIops - Average combined read and write iops.
	AvgIops int64 `json:"avg_iops,omitempty"`
	// MaxIops - Maximum combined read and write iops.
	MaxIops int64 `json:"max_iops,omitempty"`
	// AvgMbps - Average combined read and write throughput.
	AvgMbps int64 `json:"avg_mbps,omitempty"`
	// MaxMbps - Maximum combined read and write throughput.
	MaxMbps int64 `json:"max_mbps,omitempty"`
}

// sdk internal struct
type nsIopsMbpsStats struct {
	// AvgIops - Average combined read and write iops.
	AvgIops *int64 `json:"avg_iops,omitempty"`
	// MaxIops - Maximum combined read and write iops.
	MaxIops *int64 `json:"max_iops,omitempty"`
	// AvgMbps - Average combined read and write throughput.
	AvgMbps *int64 `json:"avg_mbps,omitempty"`
	// MaxMbps - Maximum combined read and write throughput.
	MaxMbps *int64 `json:"max_mbps,omitempty"`
}

// EncodeNsIopsMbpsStats - Transform NsIopsMbpsStats to nsIopsMbpsStats type
func EncodeNsIopsMbpsStats(request interface{}) (*nsIopsMbpsStats, error) {
	reqNsIopsMbpsStats := request.(*NsIopsMbpsStats)
	byte, err := json.Marshal(reqNsIopsMbpsStats)
	objPtr := &nsIopsMbpsStats{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsIopsMbpsStats Transform nsIopsMbpsStats to NsIopsMbpsStats type
func DecodeNsIopsMbpsStats(request interface{}) (*NsIopsMbpsStats, error) {
	reqNsIopsMbpsStats := request.(*nsIopsMbpsStats)
	byte, err := json.Marshal(reqNsIopsMbpsStats)
	obj := &NsIopsMbpsStats{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
