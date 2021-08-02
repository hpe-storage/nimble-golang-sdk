// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrRaidInfoFields provides field names to use in filter parameters, for example.
var NsCtrlrRaidInfoFields *NsCtrlrRaidInfoFieldHandles

func init() {
	NsCtrlrRaidInfoFields = &NsCtrlrRaidInfoFieldHandles{
		RaidId:      "raid_id",
		RaidType:    "raid_type",
		MaxCopies:   "max_copies",
		CurCopies:   "cur_copies",
		IsResyncing: "is_resyncing",
	}
}

// NsCtrlrRaidInfo - Information about a controller's raid configuration.
type NsCtrlrRaidInfo struct {
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

// NsCtrlrRaidInfoFieldHandles provides a string representation for each NsCtrlrRaidInfo field.
type NsCtrlrRaidInfoFieldHandles struct {
	RaidId      string
	RaidType    string
	MaxCopies   string
	CurCopies   string
	IsResyncing string
}
