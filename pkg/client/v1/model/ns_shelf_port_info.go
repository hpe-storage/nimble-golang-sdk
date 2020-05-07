/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsShelfPortInfo


// NsShelfPortInfo :
type NsShelfPortInfo struct {
   // PortIDx
   PortIDx float64 `json:"port_idx,omitempty"`
   // PortName
   PortName string `json:"port_name,omitempty"`
   // PortErrors
   PortErrors string `json:"port_errors,omitempty"`
   // RemoteSasAddr
   RemoteSasAddr string `json:"remote_sas_addr,omitempty"`
   // RemoteSasPhyID
   RemoteSasPhyID string `json:"remote_sas_phy_id,omitempty"`
   // RemoteLocID
   RemoteLocID  int32 `json:"remote_loc_id,omitempty"`
   // RemotePortID
   RemotePortID  int32 `json:"remote_port_id,omitempty"`
}
