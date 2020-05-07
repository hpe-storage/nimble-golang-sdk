/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsISCSIInitiator


// NsISCSIInitiator :
type NsISCSIInitiator struct {
   // ID
   ID string `json:"id,omitempty"`
   // InitiatorID
   InitiatorID string `json:"initiator_id,omitempty"`
   // Label
   Label string `json:"label,omitempty"`
   // Iqn
   Iqn string `json:"iqn,omitempty"`
   // IpAddress
   IpAddress string `json:"ip_address,omitempty"`
}
