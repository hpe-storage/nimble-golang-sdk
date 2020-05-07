/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/FibreChannelInitiatorAlias


// FibreChannelInitiatorAlias :
type FibreChannelInitiatorAlias struct {
   // ID
   ID string `json:"id,omitempty"`
   // Alias
   Alias string `json:"alias,omitempty"`
   // Wwpn
   Wwpn string `json:"wwpn,omitempty"`
}
