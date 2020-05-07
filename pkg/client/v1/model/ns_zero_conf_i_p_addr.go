/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsZeroConfIPAddr


// NsZeroConfIPAddr :
type NsZeroConfIPAddr struct {
   // Nic
   Nic string `json:"nic,omitempty"`
   // LocalIpaddr
   LocalIpaddr string `json:"local_ipaddr,omitempty"`
   // RemoteIpaddr
   RemoteIpaddr string `json:"remote_ipaddr,omitempty"`
}
