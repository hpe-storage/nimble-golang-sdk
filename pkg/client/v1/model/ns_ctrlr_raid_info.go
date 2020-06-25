// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsCtrlrRaidInfo - Information about a controller's raid configuration.
// Export NsCtrlrRaidInfoFields for advance operations like search filter etc.
var NsCtrlrRaidInfoFields *NsCtrlrRaidInfo

func init(){
	RaidTypefield:= "raid_type"
		
	NsCtrlrRaidInfoFields= &NsCtrlrRaidInfo{
		RaidType:    &RaidTypefield,
	}
}

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
