/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Subnet


// Subnet :
type Subnet struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Network
   Network string `json:"network,omitempty"`
   // Netmask
   Netmask string `json:"netmask,omitempty"`
   // AllowIscsi
   AllowIscsi bool `json:"allow_iscsi,omitempty"`
   // AllowGroup
   AllowGroup bool `json:"allow_group,omitempty"`
   // DiscoveryIp
   DiscoveryIp string `json:"discovery_ip,omitempty"`
   // Mtu
   Mtu float64 `json:"mtu,omitempty"`
   // VlanID
   VlanID float64 `json:"vlan_id,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
}
