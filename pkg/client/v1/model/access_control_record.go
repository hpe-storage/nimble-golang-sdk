/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/AccessControlRecord


// AccessControlRecord :
type AccessControlRecord struct {
   // ID
   ID string `json:"id,omitempty"`
   // ChapUserID
   ChapUserID string `json:"chap_user_id,omitempty"`
   // ChapUserName
   ChapUserName string `json:"chap_user_name,omitempty"`
   // InitiatorGroupID
   InitiatorGroupID string `json:"initiator_group_id,omitempty"`
   // InitiatorGroupName
   InitiatorGroupName string `json:"initiator_group_name,omitempty"`
   // Lun
   Lun float64 `json:"lun,omitempty"`
   // VolID
   VolID string `json:"vol_id,omitempty"`
   // VolName
   VolName string `json:"vol_name,omitempty"`
   // PeID
   PeID string `json:"pe_id,omitempty"`
   // PeName
   PeName string `json:"pe_name,omitempty"`
   // PeLun
   PeLun float64 `json:"pe_lun,omitempty"`
   // SnapID
   SnapID string `json:"snap_id,omitempty"`
   // SnapName
   SnapName string `json:"snap_name,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
}
