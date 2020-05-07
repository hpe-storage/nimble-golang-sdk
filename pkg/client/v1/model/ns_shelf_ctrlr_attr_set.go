/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsShelfCtrlrAttrSet


// NsShelfCtrlrAttrSet :
type NsShelfCtrlrAttrSet struct {
   // SessionSerial
   SessionSerial string `json:"session_serial,omitempty"`
   // CachedSerial
   CachedSerial string `json:"cached_serial,omitempty"`
   // DiskSerials
   DiskSerials string `json:"disk_serials,omitempty"`
   // DiskTypes
   DiskTypes string `json:"disk_types,omitempty"`
}
