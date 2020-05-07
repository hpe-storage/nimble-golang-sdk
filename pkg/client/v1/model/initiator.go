/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Initiator


// Initiator :
type Initiator struct {
   // ID
   ID string `json:"id,omitempty"`
   // InitiatorGroupID
   InitiatorGroupID string `json:"initiator_group_id,omitempty"`
   // InitiatorGroupName
   InitiatorGroupName string `json:"initiator_group_name,omitempty"`
   // Label
   Label string `json:"label,omitempty"`
   // Iqn
   Iqn string `json:"iqn,omitempty"`
   // IpAddress
   IpAddress string `json:"ip_address,omitempty"`
   // Alias
   Alias string `json:"alias,omitempty"`
   // ChapuserID
   ChapuserID string `json:"chapuser_id,omitempty"`
   // Wwpn
   Wwpn string `json:"wwpn,omitempty"`
   // VpOverrIDe
   VpOverrIDe bool `json:"vp_override,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // OverrIDeExistingAlias
   OverrIDeExistingAlias bool `json:"override_existing_alias,omitempty"`
}
