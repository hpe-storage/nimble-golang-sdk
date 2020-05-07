/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVolAndSnapName


// NsVolAndSnapName :
type NsVolAndSnapName struct {
   // VolName
   VolName string `json:"vol_name,omitempty"`
   // SnapName
   SnapName string `json:"snap_name,omitempty"`
}
