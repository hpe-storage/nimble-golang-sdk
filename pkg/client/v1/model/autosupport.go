/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Autosupport


// Autosupport :
type Autosupport struct {
   // ID
   ID string `json:"id,omitempty"`
   // ArrayCount
   ArrayCount float64 `json:"array_count,omitempty"`
   // GroupID
   GroupID string `json:"group_id,omitempty"`
   // GroupName
   GroupName string `json:"group_name,omitempty"`
}
