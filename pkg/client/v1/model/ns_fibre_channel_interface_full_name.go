/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFibreChannelInterfaceFullName


// NsFibreChannelInterfaceFullName :
type NsFibreChannelInterfaceFullName struct {
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // CtrlrName
   CtrlrName string `json:"ctrlr_name,omitempty"`
   // IntfName
   IntfName string `json:"intf_name,omitempty"`
}
