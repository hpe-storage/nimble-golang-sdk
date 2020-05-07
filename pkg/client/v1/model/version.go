/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Version


// Version :
type Version struct {
   // Name
   Name string `json:"name,omitempty"`
   // SoftwareVersion
   SoftwareVersion string `json:"software_version,omitempty"`
}
