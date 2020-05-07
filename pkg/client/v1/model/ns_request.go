/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsRequest


// NsRequest :
type NsRequest struct {
   // Method
   Method float64 `json:"method,omitempty"`
   // Path
   Path string `json:"path,omitempty"`
}
