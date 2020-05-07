/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSnapRetainLimit


// NsSnapRetainLimit :
type NsSnapRetainLimit struct {
   // RetainLimit
   RetainLimit float64 `json:"retain_limit,omitempty"`
   // RetainNum
   RetainNum float64 `json:"retain_num,omitempty"`
}
