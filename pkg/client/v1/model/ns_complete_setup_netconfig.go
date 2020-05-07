/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsCompleteSetupNetconfig


// NsCompleteSetupNetconfig :
type NsCompleteSetupNetconfig struct {
   // DiscoveryIp
   DiscoveryIp string `json:"discovery_ip,omitempty"`
   // DiscoveryNetmask
   DiscoveryNetmask string `json:"discovery_netmask,omitempty"`
   // MgmtIp
   MgmtIp string `json:"mgmt_ip,omitempty"`
   // SecondaryMgmtIp
   SecondaryMgmtIp string `json:"secondary_mgmt_ip,omitempty"`
   // MgmtNetmask
   MgmtNetmask string `json:"mgmt_netmask,omitempty"`
   // IscsiAutomaticConnectionMethod
   IscsiAutomaticConnectionMethod bool `json:"iscsi_automatic_connection_method,omitempty"`
   // IscsiConnectionRebalancing
   IscsiConnectionRebalancing bool `json:"iscsi_connection_rebalancing,omitempty"`
}
