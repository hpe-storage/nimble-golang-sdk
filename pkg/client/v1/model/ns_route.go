/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsRoute


// NsRoute :
type NsRoute struct {
   // TgtNetwork
   TgtNetwork string `json:"tgt_network,omitempty"`
   // TgtNetmask
   TgtNetmask string `json:"tgt_netmask,omitempty"`
   // Gateway
   Gateway string `json:"gateway,omitempty"`
}
