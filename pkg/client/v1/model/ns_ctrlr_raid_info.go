/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsCtrlrRaidInfo


// NsCtrlrRaidInfo :
type NsCtrlrRaidInfo struct {
   // RaIDID
   RaIDID float64 `json:"raid_id,omitempty"`
   // RaIDType
   RaIDType string `json:"raid_type,omitempty"`
   // MaxCopies
   MaxCopies float64 `json:"max_copies,omitempty"`
   // CurCopies
   CurCopies float64 `json:"cur_copies,omitempty"`
   // IsResyncing
   IsResyncing bool `json:"is_resyncing,omitempty"`
}
