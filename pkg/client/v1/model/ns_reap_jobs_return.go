/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsReapJobsReturn


// NsReapJobsReturn :
type NsReapJobsReturn struct {
   // Reaped
   Reaped float64 `json:"reaped,omitempty"`
   // Remaining
   Remaining float64 `json:"remaining,omitempty"`
}
