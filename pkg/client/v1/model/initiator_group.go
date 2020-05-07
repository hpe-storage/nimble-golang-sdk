/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/InitiatorGroup


// InitiatorGroup :
type InitiatorGroup struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // HostType
   HostType string `json:"host_type,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // VpOverrIDe
   VpOverrIDe bool `json:"vp_override,omitempty"`
   // AppUuID
   AppUuID string `json:"app_uuid,omitempty"`
   // VolumeCount
   VolumeCount float64 `json:"volume_count,omitempty"`
   // NumConnections
   NumConnections float64 `json:"num_connections,omitempty"`
}
