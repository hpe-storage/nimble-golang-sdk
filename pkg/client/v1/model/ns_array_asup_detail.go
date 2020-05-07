/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArrayAsupDetail


// NsArrayAsupDetail :
type NsArrayAsupDetail struct {
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // AsupValIDate
   AsupValIDate bool `json:"asup_validate,omitempty"`
   // NameResolution
   NameResolution bool `json:"name_resolution,omitempty"`
   // PingfromMgmtip
   PingfromMgmtip bool `json:"pingfrom_mgmtip,omitempty"`
   // PingfromCtrlra
   PingfromCtrlra bool `json:"pingfrom_ctrlra,omitempty"`
   // PingfromCtrlrb
   PingfromCtrlrb bool `json:"pingfrom_ctrlrb,omitempty"`
   // Heartbeat
   Heartbeat bool `json:"heartbeat,omitempty"`
}
