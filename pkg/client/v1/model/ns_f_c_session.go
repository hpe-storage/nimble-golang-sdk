/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFCSession


// NsFCSession :
type NsFCSession struct {
   // ID
   ID string `json:"id,omitempty"`
   // SessionID
   SessionID string `json:"session_id,omitempty"`
   // PrKey
   PrKey  int32 `json:"pr_key,omitempty"`
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
   InitiatorFcID float64 `json:"initiator_fcid,omitempty"`
   // TargetPortArrayName
   TargetPortArrayName string `json:"target_port_array_name,omitempty"`
   // TargetPortCtrlrID
   TargetPortCtrlrID float64 `json:"target_port_ctrlr_id,omitempty"`
   // TargetPortInterfaceName
   TargetPortInterfaceName string `json:"target_port_interface_name,omitempty"`
   // TargetWwnn
   TargetWwnn string `json:"target_wwnn,omitempty"`
   // TargetWwpn
   TargetWwpn string `json:"target_wwpn,omitempty"`
   // TargetFcID
   TargetFcID float64 `json:"target_fcid,omitempty"`
}
