/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapVol


// NsSnapVol :
type NsSnapVol struct {
   // VolID
   VolID string `json:"vol_id,omitempty"`
   // SnapName
   SnapName string `json:"snap_name,omitempty"`
   // SnapDescription
   SnapDescription string `json:"snap_description,omitempty"`
   // Cookie
   Cookie string `json:"cookie,omitempty"`
   // Online
   Online bool `json:"online,omitempty"`
   // Writable
   Writable bool `json:"writable,omitempty"`
   // AppUuID
   AppUuID string `json:"app_uuid,omitempty"`
}
