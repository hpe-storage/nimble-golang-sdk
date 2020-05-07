/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFcFabricInfo


// NsFcFabricInfo :
type NsFcFabricInfo struct {
   // FabricSwitchName
   FabricSwitchName string `json:"fabric_switch_name,omitempty"`
   // FabricSwitchPort
   FabricSwitchPort string `json:"fabric_switch_port,omitempty"`
   // FabricSwitchWwnn
   FabricSwitchWwnn string `json:"fabric_switch_wwnn,omitempty"`
   // FabricSwitchWwpn
   FabricSwitchWwpn string `json:"fabric_switch_wwpn,omitempty"`
   // FabricWwn
   FabricWwn string `json:"fabric_wwn,omitempty"`
   // FcID
   FcID string `json:"fc_id,omitempty"`
   // LoggedIn
   LoggedIn bool `json:"logged_in,omitempty"`
}
