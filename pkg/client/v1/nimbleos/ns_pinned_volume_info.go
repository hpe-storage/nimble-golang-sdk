// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	ID                   *string `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	SizePinnedCacheBytes *int64  `json:"size_pinned_cache_bytes,omitempty"`
}

// EncodeNsPinnedVolumeInfo - Transform NsPinnedVolumeInfo to nsPinnedVolumeInfo type
func EncodeNsPinnedVolumeInfo(request interface{}) (*nsPinnedVolumeInfo, error) {
	reqNsPinnedVolumeInfo := request.(*NsPinnedVolumeInfo)
	byte, err := json.Marshal(reqNsPinnedVolumeInfo)
	if err != nil {
		return nil, err
	}
	respNsPinnedVolumeInfoPtr := &nsPinnedVolumeInfo{}
	err = json.Unmarshal(byte, respNsPinnedVolumeInfoPtr)
	return respNsPinnedVolumeInfoPtr, err
}

// DecodeNsPinnedVolumeInfo Transform nsPinnedVolumeInfo to NsPinnedVolumeInfo type
func DecodeNsPinnedVolumeInfo(request interface{}) (*NsPinnedVolumeInfo, error) {
	reqNsPinnedVolumeInfo := request.(*nsPinnedVolumeInfo)
	byte, err := json.Marshal(reqNsPinnedVolumeInfo)
	if err != nil {
		return nil, err
	}
	respNsPinnedVolumeInfo := &NsPinnedVolumeInfo{}
	err = json.Unmarshal(byte, respNsPinnedVolumeInfo)
	return respNsPinnedVolumeInfo, err
}
