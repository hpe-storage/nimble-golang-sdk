/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsAlarmCount


// NsAlarmCount :
type NsAlarmCount struct {
   // Critical
   Critical float64 `json:"critical,omitempty"`
   // Warning
   Warning float64 `json:"warning,omitempty"`
}
