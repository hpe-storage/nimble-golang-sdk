/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsTokenReportUserDetailsReturn


// NsTokenReportUserDetailsReturn :
type NsTokenReportUserDetailsReturn struct {
   // UserName
   UserName string `json:"user_name,omitempty"`
   // PrimaryGroupID
   PrimaryGroupID string `json:"primary_group_id,omitempty"`
   // PrimaryGroupName
   PrimaryGroupName string `json:"primary_group_name,omitempty"`
   // GroupCount
   GroupCount float64 `json:"group_count,omitempty"`
   // InactivityTimeout
   InactivityTimeout float64 `json:"inactivity_timeout,omitempty"`
   // UserID
   UserID string `json:"user_id,omitempty"`
}
