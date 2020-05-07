/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/UserGroup


// UserGroup :
type UserGroup struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // RoleID
   RoleID string `json:"role_id,omitempty"`
   // InactivityTimeout
   InactivityTimeout float64 `json:"inactivity_timeout,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // Disabled
   Disabled bool `json:"disabled,omitempty"`
   // ExternalID
   ExternalID string `json:"external_id,omitempty"`
   // DomainID
   DomainID string `json:"domain_id,omitempty"`
   // DomainName
   DomainName string `json:"domain_name,omitempty"`
}
