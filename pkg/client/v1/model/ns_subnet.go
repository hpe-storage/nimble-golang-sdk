/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSubnet


// NsSubnet :
type NsSubnet struct {
   // Label
   Label string `json:"label,omitempty"`
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
}
