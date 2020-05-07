/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArraySummary


// NsArraySummary :
type NsArraySummary struct {
   // ID
   ID string `json:"id,omitempty"`
   // ArrayID
   ArrayID string `json:"array_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
}
