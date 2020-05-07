/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ActiveDirectoryMembership


// ActiveDirectoryMembership :
type ActiveDirectoryMembership struct {
   // ID
   ID string `json:"id,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Netbios
   Netbios string `json:"netbios,omitempty"`
   // ServerList
   ServerList string `json:"server_list,omitempty"`
   // ComputerName
   ComputerName string `json:"computer_name,omitempty"`
   // OrganizationalUnit
   OrganizationalUnit string `json:"organizational_unit,omitempty"`
   // User
   User string `json:"user,omitempty"`
   // Password
   Password string `json:"password,omitempty"`
   // Enabled
   Enabled bool `json:"enabled,omitempty"`
}
