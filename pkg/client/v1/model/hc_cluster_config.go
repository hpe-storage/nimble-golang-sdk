/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/HcClusterConfig


// HcClusterConfig :
type HcClusterConfig struct {
   // ID
   ID string `json:"id,omitempty"`
   // UniqueID
   UniqueID string `json:"unique_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // Username
   Username string `json:"username,omitempty"`
   // Password
   Password string `json:"password,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
}
