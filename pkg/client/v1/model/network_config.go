/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NetworkConfig


// NetworkConfig :
type NetworkConfig struct {
   // ID
   ID string `json:"id,omitempty"`
   // MgmtIp
   MgmtIp string `json:"mgmt_ip,omitempty"`
   // SecondaryMgmtIp
   SecondaryMgmtIp string `json:"secondary_mgmt_ip,omitempty"`
   // IscsiAutomaticConnectionMethod
   IscsiAutomaticConnectionMethod bool `json:"iscsi_automatic_connection_method,omitempty"`
   // IscsiConnectionRebalancing
   IscsiConnectionRebalancing bool `json:"iscsi_connection_rebalancing,omitempty"`
   // GroupLeaderArray
   GroupLeaderArray string `json:"group_leader_array,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // ActiveSince
   ActiveSince float64 `json:"active_since,omitempty"`
   // LastActive
   LastActive float64 `json:"last_active,omitempty"`
   // IgnoreValIDationMask
   IgnoreValIDationMask float64 `json:"ignore_validation_mask,omitempty"`
}
