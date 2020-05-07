/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsDiscoveredGroupInfo


// NsDiscoveredGroupInfo :
type NsDiscoveredGroupInfo struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // VersionCurrent
   VersionCurrent string `json:"version_current,omitempty"`
   // CountOfMembers
   CountOfMembers float64 `json:"count_of_members,omitempty"`
   // ManagementIp
   ManagementIp string `json:"management_ip,omitempty"`
   // DiscoveryIp
   DiscoveryIp string `json:"discovery_ip,omitempty"`
   // IscsiEnabled
   IscsiEnabled bool `json:"iscsi_enabled,omitempty"`
   // FcEnabled
   FcEnabled bool `json:"fc_enabled,omitempty"`
}
