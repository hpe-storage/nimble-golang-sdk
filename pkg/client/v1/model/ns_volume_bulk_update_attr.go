/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVolumeBulkUpdateAttr


// NsVolumeBulkUpdateAttr :
type NsVolumeBulkUpdateAttr struct {
   // ID
   ID string `json:"id,omitempty"`
   // FolderID
   FolderID string `json:"folder_id,omitempty"`
   // Online
   Online bool `json:"online,omitempty"`
}
