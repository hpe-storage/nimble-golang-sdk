/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/DebugLog


// DebugLog :
type DebugLog struct {
   // Tag
   Tag string `json:"tag,omitempty"`
   // Message
   Message string `json:"message,omitempty"`
}
