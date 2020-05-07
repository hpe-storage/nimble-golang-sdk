/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapLunInfo


// NsSnapLunInfo :
type NsSnapLunInfo struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Lun
   Lun float64 `json:"lun,omitempty"`
}
