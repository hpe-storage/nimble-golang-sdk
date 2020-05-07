/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsADTestUserReturn


// NsADTestUserReturn :
type NsADTestUserReturn struct {
   // Username
   Username string `json:"username,omitempty"`
   // PrimaryGroupName
   PrimaryGroupName string `json:"primary_group_name,omitempty"`
   // PrimaryGroupID
   PrimaryGroupID string `json:"primary_group_id,omitempty"`
   // GroupCount
   GroupCount float64 `json:"group_count,omitempty"`
}
