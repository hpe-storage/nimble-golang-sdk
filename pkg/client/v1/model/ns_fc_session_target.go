/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFcSessionTarget


// NsFcSessionTarget :
type NsFcSessionTarget struct {
   // TargetPortArrayName
   TargetPortArrayName string `json:"target_port_array_name,omitempty"`
   // TargetPortCtrlrName
   TargetPortCtrlrName string `json:"target_port_ctrlr_name,omitempty"`
   // TargetPortInterfaceName
   TargetPortInterfaceName string `json:"target_port_interface_name,omitempty"`
   // TargetWwnn
   TargetWwnn string `json:"target_wwnn,omitempty"`
   // TargetWwpn
   TargetWwpn string `json:"target_wwpn,omitempty"`
   // TargetFcID
   TargetFcID string `json:"target_fcid,omitempty"`
}
