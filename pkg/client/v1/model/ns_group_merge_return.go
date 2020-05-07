/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsGroupMergeReturn


// NsGroupMergeReturn :
type NsGroupMergeReturn struct {
   // SrcSID
   SrcSID string `json:"src_sid,omitempty"`
   // SrcGroupName
   SrcGroupName string `json:"src_group_name,omitempty"`
   // DstGroupName
   DstGroupName string `json:"dst_group_name,omitempty"`
   // DstGroupSwversion
   DstGroupSwversion string `json:"dst_group_swversion,omitempty"`
   // ValIDationErrorMsg
   ValIDationErrorMsg string `json:"validation_error_msg,omitempty"`
}
