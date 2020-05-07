/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/FolderSet


// FolderSet :
type FolderSet struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // AppUuID
   AppUuID string `json:"app_uuid,omitempty"`
   // AppserverID
   AppserverID string `json:"appserver_id,omitempty"`
   // AppserverName
   AppserverName string `json:"appserver_name,omitempty"`
}
