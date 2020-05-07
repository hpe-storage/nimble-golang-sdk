/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/User


// User :
type User struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // RoleID
   RoleID string `json:"role_id,omitempty"`
   // Password
   Password string `json:"password,omitempty"`
   // AuthPassword
   AuthPassword string `json:"auth_password,omitempty"`
   // InactivityTimeout
   InactivityTimeout float64 `json:"inactivity_timeout,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // EmailAddr
   EmailAddr string `json:"email_addr,omitempty"`
   // Disabled
   Disabled bool `json:"disabled,omitempty"`
   // AuthLock
   AuthLock bool `json:"auth_lock,omitempty"`
   // LastLogin
   LastLogin float64 `json:"last_login,omitempty"`
   // LastLogout
   LastLogout float64 `json:"last_logout,omitempty"`
   // LoggedIn
   LoggedIn bool `json:"logged_in,omitempty"`
}
