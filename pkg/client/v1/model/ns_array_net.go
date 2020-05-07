/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArrayNet


// NsArrayNet :
type NsArrayNet struct {
   // Name
   Name string `json:"name,omitempty"`
   // MemberGID
   MemberGID float64 `json:"member_gid,omitempty"`
   // CtrlrASupportIp
   CtrlrASupportIp string `json:"ctrlr_a_support_ip,omitempty"`
   // CtrlrBSupportIp
   CtrlrBSupportIp string `json:"ctrlr_b_support_ip,omitempty"`
}
