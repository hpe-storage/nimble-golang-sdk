/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFcSessionInitiator


// NsFcSessionInitiator :
type NsFcSessionInitiator struct {
   // InitiatorAlias
   InitiatorAlias string `json:"initiator_alias,omitempty"`
   // InitiatorWwpn
   InitiatorWwpn string `json:"initiator_wwpn,omitempty"`
   // InitiatorWwnn
   InitiatorWwnn string `json:"initiator_wwnn,omitempty"`
   // InitiatorSwitchName
   InitiatorSwitchName string `json:"initiator_switch_name,omitempty"`
   // InitiatorSwitchPort
   InitiatorSwitchPort string `json:"initiator_switch_port,omitempty"`
   // InitiatorSymbolicPortname
   InitiatorSymbolicPortname string `json:"initiator_symbolic_portname,omitempty"`
   // InitiatorSymbolicNodename
   InitiatorSymbolicNodename string `json:"initiator_symbolic_nodename,omitempty"`
   // InitiatorFcID
   InitiatorFcID string `json:"initiator_fcid,omitempty"`
}
