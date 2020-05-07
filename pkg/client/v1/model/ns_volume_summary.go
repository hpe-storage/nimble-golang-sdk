/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVolumeSummary


// NsVolumeSummary :
type NsVolumeSummary struct {
   // ID
   ID string `json:"id,omitempty"`
   // VolID
   VolID string `json:"vol_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // VolName
   VolName string `json:"vol_name,omitempty"`
}
