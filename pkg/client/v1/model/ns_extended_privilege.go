/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsExtendedPrivilege


// NsExtendedPrivilege :
type NsExtendedPrivilege struct {
   // ObjectType
   ObjectType string `json:"object_type,omitempty"`
   // Operation
   Operation string `json:"operation,omitempty"`
   // Allow
   Allow bool `json:"allow,omitempty"`
}
