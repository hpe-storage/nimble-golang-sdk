// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsCtrlrRaidInfo - Information about a controller's raid configuration.
// Export NsCtrlrRaidInfoFields for advance operations like search filter etc.
var NsCtrlrRaidInfoFields *NsCtrlrRaidInfo

func init() {

	NsCtrlrRaidInfoFields = &NsCtrlrRaidInfo{
		RaidType: "raid_type",
	}
}

type NsCtrlrRaidInfo struct {
	// RaidId - Raid ID for this raid array.
	RaidId int64 `json:"raid_id,omitempty"`
	// RaidType - Type of raid for this array.
	RaidType string `json:"raid_type,omitempty"`
	// MaxCopies - Maximum number of copies.
	MaxCopies int64 `json:"max_copies,omitempty"`
	// CurCopies - Current number of copies.
	CurCopies int64 `json:"cur_copies,omitempty"`
	// IsResyncing - Is this raid array resynchronizing.
	IsResyncing *bool `json:"is_resyncing,omitempty"`
}

// sdk internal struct
type nsCtrlrRaidInfo struct {
	// RaidId - Raid ID for this raid array.
	RaidId *int64 `json:"raid_id,omitempty"`
	// RaidType - Type of raid for this array.
	RaidType *string `json:"raid_type,omitempty"`
	// MaxCopies - Maximum number of copies.
	MaxCopies *int64 `json:"max_copies,omitempty"`
	// CurCopies - Current number of copies.
	CurCopies *int64 `json:"cur_copies,omitempty"`
	// IsResyncing - Is this raid array resynchronizing.
	IsResyncing *bool `json:"is_resyncing,omitempty"`
}

// EncodeNsCtrlrRaidInfo - Transform NsCtrlrRaidInfo to nsCtrlrRaidInfo type
func EncodeNsCtrlrRaidInfo(request interface{}) (*nsCtrlrRaidInfo, error) {
	reqNsCtrlrRaidInfo := request.(*NsCtrlrRaidInfo)
	byte, err := json.Marshal(reqNsCtrlrRaidInfo)
	objPtr := &nsCtrlrRaidInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsCtrlrRaidInfo Transform nsCtrlrRaidInfo to NsCtrlrRaidInfo type
func DecodeNsCtrlrRaidInfo(request interface{}) (*NsCtrlrRaidInfo, error) {
	reqNsCtrlrRaidInfo := request.(*nsCtrlrRaidInfo)
	byte, err := json.Marshal(reqNsCtrlrRaidInfo)
	obj := &NsCtrlrRaidInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
