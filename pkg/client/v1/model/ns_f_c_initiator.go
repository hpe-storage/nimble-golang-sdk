/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFCInitiator


// NsFCInitiator :
type NsFCInitiator struct {
   // ID
   ID string `json:"id,omitempty"`
   // InitiatorID
   InitiatorID string `json:"initiator_id,omitempty"`
   // Wwpn
   Wwpn string `json:"wwpn,omitempty"`
   // Alias
   Alias string `json:"alias,omitempty"`
}
