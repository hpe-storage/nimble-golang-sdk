/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVolumeSummaryWithAppCategory


// NsVolumeSummaryWithAppCategory :
type NsVolumeSummaryWithAppCategory struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // AppCategory
   AppCategory string `json:"app_category,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // Lun
   Lun float64 `json:"lun,omitempty"`
}
