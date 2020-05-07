/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ApplicationServer


// ApplicationServer :
type ApplicationServer struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Hostname
   Hostname string `json:"hostname,omitempty"`
   // Port
   Port float64 `json:"port,omitempty"`
   // Username
   Username string `json:"username,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // Password
   Password string `json:"password,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
}
