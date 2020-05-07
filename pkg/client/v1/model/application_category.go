/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ApplicationCategory


// ApplicationCategory :
type ApplicationCategory struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // DedupeEnabled
   DedupeEnabled bool `json:"dedupe_enabled,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
}
