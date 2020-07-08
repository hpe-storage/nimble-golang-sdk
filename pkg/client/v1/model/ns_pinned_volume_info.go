// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsPinnedVolumeInfo - Select fields containing pinned volume info.
// Export NsPinnedVolumeInfoFields for advance operations like search filter etc.
var NsPinnedVolumeInfoFields *NsPinnedVolumeInfo

func init() {

	NsPinnedVolumeInfoFields = &NsPinnedVolumeInfo{
		ID:   "id",
		Name: "name",
	}
}

type NsPinnedVolumeInfo struct {
	// ID - ID of volume.
	ID string `json:"id,omitempty"`
	// Name - Name of volume.
	Name string `json:"name,omitempty"`
	// SizePinnedCacheBytes - Total pinned cache size of the volume in bytes.
	SizePinnedCacheBytes int64 `json:"size_pinned_cache_bytes,omitempty"`
}

// sdk internal struct
type nsPinnedVolumeInfo struct {
	// ID - ID of volume.
	ID *string `json:"id,omitempty"`
	// Name - Name of volume.
	Name *string `json:"name,omitempty"`
	// SizePinnedCacheBytes - Total pinned cache size of the volume in bytes.
	SizePinnedCacheBytes *int64 `json:"size_pinned_cache_bytes,omitempty"`
}

// EncodeNsPinnedVolumeInfo - Transform NsPinnedVolumeInfo to nsPinnedVolumeInfo type
func EncodeNsPinnedVolumeInfo(request interface{}) (*nsPinnedVolumeInfo, error) {
	reqNsPinnedVolumeInfo := request.(*NsPinnedVolumeInfo)
	byte, err := json.Marshal(reqNsPinnedVolumeInfo)
	objPtr := &nsPinnedVolumeInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsPinnedVolumeInfo Transform nsPinnedVolumeInfo to NsPinnedVolumeInfo type
func DecodeNsPinnedVolumeInfo(request interface{}) (*NsPinnedVolumeInfo, error) {
	reqNsPinnedVolumeInfo := request.(*nsPinnedVolumeInfo)
	byte, err := json.Marshal(reqNsPinnedVolumeInfo)
	obj := &NsPinnedVolumeInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
